const http = require("http");
const routes = require("./routes");

// event loop: 実行できるコードがある限り実行しづつける
const server = http.createServer(routes);

server.listen(3000);
