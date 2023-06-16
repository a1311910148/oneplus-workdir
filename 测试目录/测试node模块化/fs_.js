const fs = require("fs");

// let data = fs.readFileSync("./mod.js", "utf-8");
// let data = fs.readFile("./mod.js", "utf-8", (err, data) => {
//   if (err) console.log(err);
//   console.log(data);
// });

// fs.createReadStream("./mod.js", { start: 10, end: 29 }).on("data", (chunk) => {
//   console.log(chunk.toString());
// });
// console.log(data);

let writeStreamInstans = fs.createWriteStream("./bundle.js");

writeStreamInstans.on("open", async () => {
  writeStreamInstans._write(Buffer.from("abc", "utf-8"), "utf-8", (err) => {
    if (err) console.log(err);
    if (!err) console.log("写入成功");
  });

  writeStreamInstans._writev(
    [
      {
        chunk: Buffer.from("实现可写流批量写钩子\n", "utf-8"),
        encoding: "utf-8",
      },
      {
        chunk: Buffer.from(
          "计算待写入的出总大小，并且把数据保存到chunk数组中，准备写入  ",
          "utf-8"
        ),
        encoding: "utf-8",
      },
    ],
    (err) => {
      if (err) console.log(err);
      if (!err) console.log("写入成功");
    }
  );
});

const net = require("net");

net
  .createServer((socket) => {
    console.log("客户端已连接");
    socket.on("data", (data) => {
      console.log(data.toString());
    });
  })
  .listen(3000, () => {});
