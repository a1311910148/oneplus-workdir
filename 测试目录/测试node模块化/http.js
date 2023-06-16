// var HTTPParser = process.binding("http_parser").HTTPParser;
// var parser = new HTTPParser(HTTPParser.REQUEST);
// const kOnHeaders = HTTPParser.kOnHeaders;
// const kOnHeadersComplete = HTTPParser.kOnHeadersComplete;
// const kOnBody = HTTPParser.kOnBody;
// const kOnMessageComplete = HTTPParser.kOnMessageComplete;
// const kOnExecute = HTTPParser.kOnExecute;

// parser[kOnHeaders] = function (headers, url) {
//   console.log("kOnHeaders", headers.length, url);
// };
// parser[kOnHeadersComplete] = function (
//   versionMajor,
//   versionMinor,
//   headers,
//   method,
//   url,
//   statusCode,
//   statusMessage,
//   upgrade,
//   shouldKeepAlive
// ) {
//   console.log("kOnHeadersComplete", headers);
// };

// parser[kOnBody] = function (b, start, len) {
//   console.log("kOnBody", b.slice(start).toString("utf-8"));
// };
// parser[kOnMessageComplete] = function () {
//   console.log("kOnMessageComplete");
// };
// parser[kOnExecute] = function () {
//   console.log("kOnExecute");
// };
// parser.execute(
//   Buffer.from(
//     "GET / HTTP/1.1\r\n" +
//       "Host: http://localhost\r\n" +
//       "a: b\r\n" +
//       // 很多'a: b\r\n'+
//       "content-length: 1\r\n\r\n" +
//       "我的世界"
//   )
// );

// const { ServerResponse } = require("http");
// const http = require("http");

// const option = {
//   hostname: "101.43.18.212",
//   port: 80,
//   path: "/",
//   method: "GET",
// };

// http
//   .request(option, (res) => {
//     console.log(res.headers);
//     res.on("data", (chunk) => {
//       console.log(`响应主体: ${chunk}`);
//     });
//     res.on("end", () => {
//       console.log("响应中已无数据。");
//     });
//   })
//   .end();

// const { Server, ClientRequest, ServerResponse, request } = require("http");

// console.log(new ServerResponse(new ClientRequest("http://101.43.18.212:5244")));

// let s = new Server((req, res) => {
//   console.log(1);
//   res.writeHead(200, { "Content-Type": "text/plain" });
//   res.end("okay");
// }).listen(5244);

// setTimeout(() => {
//   s.emit(
//     "request",
//     new ClientRequest("http://101.43.18.212:5244"),
//     new ServerResponse(new ClientRequest("http://101.43.18.212:5244"))
//   );
// }, 1000);
const fs = require("fs");
const {
  Server,
  ClientRequest,
  ServerResponse,
  request,
  IncomingMessage,
} = require("http");
const { Module } = require("module");
const path = require("path");

let c = fs.createWriteStream("./test.png");
// https: github.com/theanarkh/understand-nodejs.git
// request(
//   {
//     hostname: "101.43.18.212",
//     port: 8883,
//     path: "/1682000287864.png",
//     method: "HEAD",
//   },
//   (res) => {
//     console.log(res.headers);
//     res.on("data", (chunk) => {
//       console.log(`响应主体: ${chunk}`);
//       c.write(chunk);
//     });
//     res.on("end", () => {
//       console.log("响应中已无数据。");
//     });
//   }
// ).end();

// 创建一个 path与处理函数的映射表
const routes = {
  "/": [],
  "/about": [],
  get: function (path, handler) {
    this[path].push(handler);
  },
  listen: function (port, callback) {
    if (!port) throw new Error("port is required");
    if (typeof port != "number") throw new Error("port is not a number");
    if (this.server) throw new Error("server is already running");

    this.server = new Server((req, res) => {
      // console.log(req.headers, req.url);
      let isStatic = false;
      console.log(req.url.slice(1));
      try {
        fs.accessSync(req.url.slice(1), fs.constants.F_OK);
        isStatic = true;
      } catch (error) {
        isStatic = false;
      }
      if (isStatic) {
        this.static(req.url.slice(1), req, res);
        return;
      }
      if (!this[req.url]) {
        res.writeHead(404, { "Content-Type": "text/plain" });
        res.end("Not Found");
        return;
      }

      this[req.url].forEach((handler) => handler(req, res));
    }).listen(port, callback);
  },
  use: function (handler) {
    if (!typeof handler === "function")
      throw new Error("handler must be a function");
    // push 到所有的路由中
    Object.keys(this).forEach((key) => {
      if (typeof this[key] === "function") return;
      this[key].push(handler);
    });
  },
};

// 请求日志
routes.use(function (req, res) {
  let record = `${new Date().toLocaleString()} ${req.method} ${req.url} ${
    req.headers["user-agent"]
  }`;
  fs.appendFile("./log.txt", record + "\n", (err) => {
    if (err) throw err;
  });
});

// 静态文件处理
routes.static = function (path, req, res) {
  console.log("static");
  fs.readFile(path, (err, data) => {
    console.log(data);
    if (err) throw err;
    res.writeHead(200, { "Content-Type": "text/plain" });
    res.end(data);
  });
};

routes.get("/", function (req, res) {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("index");
});

routes.get("/about", function (req, res) {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("about");
});

routes.listen(8883, () => console.log("在http://0.0.0.0:8883监听成功"));

// 判断文件是否存在
routes.server.on("error", (e) => {
  if (e.code === "EADDRINUSE") {
    console.log("Address in use, retrying...");
    setTimeout(() => {
      server.close();
      server.listen(PORT, HOST);
    }, 1000);
  }
});

const path = require("path");
const { Module } = require("module");

const filenmae = "./package.json";
const dirname = path.dirname(filenmae);
console.log(dirname);
console.log(path.resolve(dirname, filenmae));
let mod = new Module(filenmae, null);
console.log(mod);
mod.load(filenmae);
console.log(mod);
console.log(process.binding("os").getHostname());
