const mongoose = require("mongoose");

const Schema = mongoose.Schema;

mongoose.connect("mongodb://127.0.0.1:27017/", {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

const db = mongoose.connection;

db.on("error", () => console.error("连接错误！"));
db.once("open", () => console.log("连接成功！"));