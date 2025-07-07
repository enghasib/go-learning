import express from 'express'
import morgan from 'morgan'

const app = express()

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(morgan("dev"))


app.get('/',(_,res)=>{
    res.status(200).send("This response is form the root url!")
})

app.get('/get',(_,res)=>{
    res.status(200).json({
        message: "This is a json response from  the get url",
    })
})

app.post('/post',(req,res)=>{
    res.status(200).json(req.body)
})

app.post('/postForm',(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

const port = 5000
app.listen(port,()=>{
    console.log(`Server is running on  http://localhost:${port}`)
})