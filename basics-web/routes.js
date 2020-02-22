const fs = require("fs");

const requestHandler = (req, res) => {
  // urlによるルーティングの設定
  const url = req.url;
  const method = req.method;
  if (url === "/") {
    // console.log(req.url, req.method, req.headers);
    // process終了後にループを抜ける
    // process.exit( );
    res.write("<html>");
    res.write("<header><title>Enter Message</title></header>");
    res.write(
      "<body><form action='/message' method='POST'><input type='text' name='message'><button type='submit'>Send</button></form></body>"
    );
    res.write("</html>");
    return res.end();
  }

  // formでPOSTした際のりダイレク先を指定
  if (url === "/message" && method === "POST") {
    // Parsing Request Bodies(stream, buffer)
    const body = [];
    req.on("data", chuck => {
      console.log("chuck", chuck);
      body.push(chuck);
    });
    req.on("end", () => {
      // bufferに読み込まれたchuckを文字列に戻す
      const parsedBody = Buffer.concat(body).toString();
      console.log("parsedBody", parsedBody);

      const message = parsedBody.split("=")[1];
      // writeFileSync: 同期的にファイル書き込みを行う
      // このコードが実行されるまで次のコードをブロックする
      // fs.writeFileSync("messagge.txt", message);

      // callbackとしてもブロックコードを実行できる
      fs.writeFile("messagge.txt", message, () => {
        // 302(リダイレクト)を返す
        res.statusCode = 302;
        res.setHeader("Location", "/");
        return res.end();
      });
    });
  }
  // headerを設定しすると設定に合った書き込みが可能
  res.setHeader("Content-Type", "text/html");
  res.write("<html>");
  res.write("<header><title>First Page</title></header>");
  res.write("<body><h1>Hello from my Node.js Server!</h1></body>");
  res.write("</html>");
  res.end();
};

module.exports = requestHandler;
