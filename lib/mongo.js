const MongoClient = require('mongodb').MongoClient;
const chalk = require('chalk');
const config = require('../config')

const init = () => new Promise((resolve, reject) =>
  MongoClient.connect(config.database.mongoUrl, { 
    useUnifiedTopology: true, 
    useNewUrlParser: true 
  }, (err, client) => {
    console.log(chalk.greenBright(`[mongo] Successfully connected on db ${config.database.dbName}!`))
    const db = client.db(config.database.dbName)
    resolve(db)
  })
)

module.exports = {
  init
}