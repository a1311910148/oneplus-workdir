const express = require('express')
const MyDataModel = require('./mog.js')
const app = express()

// 解析 POST 请求的中间件
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.use((req, res, next) => {     
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Methods', 'PUT,POST,GET,DELETE,OPTIONS'); 
    res.header('Access-Control-Allow-Headers', 'Content-Type,Content-Length, Authorization, Accept,X-Requested-With');     
    next();
})
// POST 请求的处理
app.post('/api', async (req, res) => {
  const { text } = req.body
  const date = new Date()
  try {
    // 将数据存储到 MongoDB 中
    const newMyData = new MyDataModel({ text, date })
    await newMyData.save()

    res.status(200).json({ message: 'Data saved successfully' })
  } catch (err) {
    console.error(err)
    res.status(500).json({ error: 'Server error' })
  }
})

// 服务器监听
app.listen(3000, () => {
  console.log('Server listening on http://localhost:3000')
})