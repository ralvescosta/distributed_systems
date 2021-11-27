const pino = require('pino');
const pinoInspector = require('pino-inspector');
const { createContainer, InjectionMode, asValue, asClass, Lifetime } = require('awilix')

const { Telemetry } = require('./infra/telemetry/telemetry')
const { MessagingBroker,  } = require('./infra/message_broker/message_broker')
const { DbConnection } = require('./infra/database/connection')
const { InventoryClient } = require('./infra/grpc_client/inventory_client')
const { PaymentClient } = require('./infra/http_client/payment_client')
const { PubClient } = require('./infra/pub_client/pub_client')
const { PurchaseUseCase } = require('./application/usecases/purchase_usecase')
const { PurchaseSubscriber } = require('./interface/amqp/subscribers/purchase_subscriber')
const { PurchaseController } = require('./interface/amqp/controllers/purchase_controller')

const container = createContainer({
  injectionMode: InjectionMode.CLASSIC
})

const registerInjections = () => {
  container.register({
    logger: asValue(createLoggerInstance()),
    telemetry: asClass(Telemetry, { lifetime:  Lifetime.SINGLETON}),
    messageBroker: asClass(MessagingBroker, { lifetime:  Lifetime.SINGLETON }),
    dbConnection: asClass(DbConnection, { lifetime:  Lifetime.SINGLETON }),
    inventoryClient: asClass(InventoryClient, { lifetime:  Lifetime.SCOPED }),
    paymentClient: asClass(PaymentClient, { lifetime:  Lifetime.SCOPED }),
    pubClient: asClass(PubClient, { lifetime: Lifetime.SCOPED }),
    purchaseUseCase: asClass(PurchaseUseCase, { lifetime:  Lifetime.SCOPED }),
    purchaseController: asClass(PurchaseController, { lifetime:  Lifetime.SINGLETON }),
    purchaseSubscriber: asClass(PurchaseSubscriber, { lifetime:  Lifetime.SINGLETON }),
  })

  return container
}

const createLoggerInstance = () => {
  const debug = process.env.LOG_LEVEL === 'trace' ||  process.env.LOG_LEVEL === 'debug'
  
  let logger;
  if (debug) {
    logger = pino({
      enabled: process.env.ENABLE_LOG === 'true',
      level: process.env.LOG_LEVEL || 'trace',
      prettyPrint: true,
      prettifier: pinoInspector,
    })
  } else {
    logger = pino({
      enabled: process.env.ENABLE_LOG === 'true',
      level: process.env.LOG_LEVEL || 'warn'
    })
  }

  return logger;
}

module.exports = {
  container,
  registerInjections
}