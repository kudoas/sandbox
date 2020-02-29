// const products = [];

// file system
const fs = require("fs");
const path = require("path");

const p = path.join(path.dirname(process.mainModule.filename), "data", "products.json");

const getProductsFromFile = cb => {
  fs.readFile(p, (err, fileContent) => {
    if (err) {
      cb([]);
    } else {
      cb(JSON.parse(fileContent));
    }
  });
};

module.exports = class Product {
  constructor(t) {
    this.title = t;
  }

  save() {
    // products.push(this);

    // file system
    getProductsFromFile(products => {
      products.push(this);
      fs.writeFile(p, JSON.stringify(products), err => {
        console.log(err);
      });
    });
  }

  // static: class instanceを作らず呼び出せるmethod
  // fetchAll()でcallbackを呼び出す
  static fetchAll(cb) {
    getProductsFromFile(cb);

    // fetchall自体は何も返していないのでエラー
    // jsは非同期なので値を渡したい場合はreturnではなくcallbackとして渡す必要あり？
    // fs.readFile(p, (err, fileContent) => {
    //   if (err) {
    //     return [];
    //   }
    //   return JSON.parse(fileContent);
    // });
  }
};
