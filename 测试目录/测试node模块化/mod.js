function test() {
    console.log(1);
}

function server(port, callback) {
    const express = require('express');
    let router = express.Router()
    router.get('/', (req, res) => res.send('router'))

    let app = express()
    app.use('/api/home', router)
    app.use('/', express.static('./'))
    app.get('/api', callback)
    app.listen(port, () => console.log('成功监听在' + port + '端口'))
    return app
}
module.exports = { test, server }
