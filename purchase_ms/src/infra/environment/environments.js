const dotEnv = require('dotenv')

module.exports = {
  registerEnvironments: () => {
    const nodeEnv = process.env.NODE_ENV || 'development'
    dotEnv.config({ path: `.env.${nodeEnv}` })
  } 
}