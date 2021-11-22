const createLoggerInstance = () => {
  const debug = process.env.DEBUG === 'true'
  
  let logger;
  if (debug) {
    logger = pino({
      enabled: process.env.ENABLE_LOG === 'true',
      level: process.env.LOG_LEVEL || 'warn',
      prettyPrint: true,
      prettifier: pinoInspector
    })
  } else {
    logger = pino({
      enabled: process.env.ENABLE_LOG === 'true',
      level: process.env.LOG_LEVEL || 'warn'
    })
  }

  return logger;
}