const express = require("express");


const app = express()
app.use(express.json());
app.get("/", (req, res) => {
  res.status(200).json({message: "welcome my test api"})
});

const port = process.env.PORT || 3500;
app.listen(port, () => {console.log("server listening on port " + port)});
