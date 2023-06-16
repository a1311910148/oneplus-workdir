const mongoose = require('mongoose')
const dbUrl = 'mongodb://101.43.18.212:32769/danmu'

mongoose.connect(dbUrl, {useNewUrlParser: true, useUnifiedTopology: true})
const db = mongoose.connection

db.on('error', console.error.bind(console, 'MongoDB connection error:'))
db.once('open', () => {
  console.log('MongoDB connected successfully')
})


const Schema = mongoose.Schema;

const MyDataSchema = new Schema({
  text: String,
  date: String
});

const MyDataModel = mongoose.model('danmu', MyDataSchema);

module.exports = MyDataModel;