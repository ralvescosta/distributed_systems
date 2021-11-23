const Environments = require('./infra/environment/environments')
const { registerInjections }  = require('./ioc')

;(async () => {
  Environments.registerEnvironments()
  const container = registerInjections()

  const { messageBroker, logger, purchaseSubscriber } = container.cradle;
  
  const isConnected = await messageBroker.connectToBroker()
  if (isConnected.isLeft()) {
    logger.error("Exiting...")
    process.exit(1)
  }

  purchaseSubscriber.subscribe()
})()