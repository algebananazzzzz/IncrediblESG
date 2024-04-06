const express = require('express')
const awsServerlessExpress = require('aws-serverless-express')
const cors = require('cors')
const bodyParser = require('body-parser')

const app = express()
app.use(bodyParser.json())

app.use(cors())

app.get('/', (req, res) => {
    res.send('Hello World!')
})

if ("LAMBDA_TASK_ROOT" in process.env) {
    const server = awsServerlessExpress.createServer(app)
    exports.handler = (event, context) => awsServerlessExpress.proxy(server, event, context)
} else {
    app.listen(3000, () => {
        console.log(`Example app listening on http://localhost:3000/`)
    })
}