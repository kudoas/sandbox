// ws-server.js
const WebSocket = require('ws');
const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(426, { 'Content-Type': 'text/plain' });
  res.end('WebSocket Required');
});

const wss = new WebSocket.Server({ server });

wss.on('connection', (ws, req) => {
  const headers = req.headers;

  // 接続直後にヘッダを送信
  ws.send(JSON.stringify(headers, null, 2));

  // http version
  ws.send(`HTTP Version: ${req.httpVersion}`);

  // echo
  ws.on('message', (msg) => {
    ws.send(`echo: ${msg}`);
  });
});

server.listen(3000, () => {
  console.log('WebSocket server listening on :3000');
});
