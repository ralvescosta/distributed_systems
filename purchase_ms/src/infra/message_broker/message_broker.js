const amqp = require('amqplib/callback_api')

const { left, right } = require('../../domain/entities/either')
const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class MessagingBroker {
  constructor (logger) {
    this._brokerChannel
    this._connection
    this.logger = logger
  }

  getBrokerChannel () {
    return this._brokerChannel
  }

  getConnection () {
    return this._connection
  }

  async connectToBroker () {
    // Creates connection
    const connection = await this._amqpConnectCallbackToPromise()
    if (connection.isLeft()) {
      return left(connection.value)
    }
    this._connection = connection.value

    // Creates channel
    const channel = await this._amqpCreateChannelCallbackToPromise()
    if (channel.isLeft()) {
      return left(channel.value)
    }
    this._brokerChannel = channel.value

    return right(true)
  }

  async pub (message, exchange, routingKey = '') {
    if (!this._brokerChannel) {
      throw new Error('Channel not established')
    }

    this._brokerChannel.publish(exchange, routingKey, Buffer.from(message))

    await this.closeConnection()
  }

  sub (queueName, controller, options = { noAck: true }) {
    this.logger.info(`Register Subscribe in Queue: ${queueName}`)
    if (!this._brokerChannel) {
      throw new Error('Channel not established')
    }

    if (!options.noAck) {
      this._noAckMessages(queueName, controller, options)
      return
    }

    this._brokerChannel.consume(queueName, 
      async (message) => {
        await controller.handle(message)
      }, options)
  }

  _noAckMessages(queueName, controller, options) {
    this._brokerChannel.consume(queueName, async (message) => {
      // received on message per time
      this._brokerChannel.prefetch(1)

      const result = await controller.handle(message)
      if (result.isRight()) {
        this._brokerChannel.ack(message)
        return
      }

      if(result.value.code >= ErrorCodeEnum.InternalError) {
        this._brokerChannel.nack(message)
        return
      }

      const DEAD_LETTER_EXCHANGE = process.env.DEAD_LETTER_EXCHANGE
      const DEAD_LETTER_ROUTING_KEY = process.env.DEAD_LETTER_ROUTING_KEY
      // remove queue message
      this._brokerChannel.ack(message)
      // publish message into dead letter
      this._brokerChannel.publish(DEAD_LETTER_EXCHANGE, DEAD_LETTER_ROUTING_KEY, message.content)
    }, options)
  }

  async closeConnection (){
    if (!this._brokerChannel) {
      throw new Error('Channel not established')
    }

    this._brokerChannel.close((err) => {
      if (err) {
        throw new Error('Error trying to close channel')
      }

      this._connection.close()
    })
  }

  _amqpConnectCallbackToPromise () {
    const AMQP_URI = process.env.AMQP_URI
    this.logger.info(`Connecting to RabbitMQ Broker: ${AMQP_URI} ...`)
    return new Promise((resolve, rejects) => {
      amqp.connect(AMQP_URI, (connError, connection) => {
        if (connError) {
          this.logger.error(connError)
          return rejects(connError)
        }
        this.logger.info('RabbitMQ Broker Connected!')
        return resolve(right(connection))
      })
    })
  }

  _amqpCreateChannelCallbackToPromise () {
    const AMQP_EXCHANGE = process.env.AMQP_EXCHANGE
    const AMQP_EXCHANGE_KIND = process.env.AMQP_EXCHANGE_KIND
    const AMQP_QUEUE = process.env.AMQP_QUEUE
    const AMQP_ROUTING_KEY = process.env.AMQP_ROUTING_KEY
    const DEAD_LETTER_EXCHANGE = process.env.DEAD_LETTER_EXCHANGE
    const DEAD_LETTER_QUEUE = process.env.DEAD_LETTER_QUEUE
    const DEAD_LETTER_ROUTING_KEY = process.env.DEAD_LETTER_ROUTING_KEY


    return new Promise((resolve, rejects) => {
      this._connection.createChannel((channelError, channel) => {
        if (channelError) return rejects(channelError)

        channel.assertExchange(AMQP_EXCHANGE, AMQP_EXCHANGE_KIND, { durable: true }, (err) => {
          if (err) return rejects(err)
        })

        // Dead Letter
        channel.assertExchange(DEAD_LETTER_EXCHANGE, AMQP_EXCHANGE_KIND, { durable: true }, (err) => {
          if (err) return rejects(err)
        })
        channel.assertQueue(DEAD_LETTER_QUEUE, { durable: true }, (err) => {
          if (err) return rejects(err)
        })
        channel.bindQueue(DEAD_LETTER_QUEUE, DEAD_LETTER_EXCHANGE, DEAD_LETTER_ROUTING_KEY, {}, (err) => {
          if (err) return rejects(err)
        })

        channel.assertQueue(AMQP_QUEUE, { 
          durable: true,  
          deadLetterExchange: DEAD_LETTER_EXCHANGE,
          deadLetterRoutingKey: DEAD_LETTER_ROUTING_KEY,
        }, (err) => {
          if (err) return rejects(err)
        })
        channel.bindQueue(AMQP_QUEUE, AMQP_EXCHANGE, AMQP_ROUTING_KEY, {}, (err) => {
          if (err) return rejects(err)
        })

        return resolve(right(channel))
      })
    })
  }
}

module.exports = { MessagingBroker }
