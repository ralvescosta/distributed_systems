const Environments = require('./infra/environment/environments')
const { registerInjections }  = require('./ioc')

;(async () => {
  Environments.registerEnvironments()
  const container = registerInjections()

  const { logger, telemetry, messageBroker, dbConnection, purchaseSubscriber } = container.cradle;
  telemetry.start()

  const isConnected = await messageBroker.connectToBroker()
  if (isConnected.isLeft()) {
    logger.error("Exiting...")
    process.exit(1)
  }

  const isDbConnected = await dbConnection.connect()
  if (isDbConnected.isLeft()) {
    logger.error("Exiting...")
    process.exit(1)
  }

  purchaseSubscriber.subscribe()
})()