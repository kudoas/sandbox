// mysql2での直接接続する場合
// const mysql = require("mysql2");

// const pool = mysql.createPool({
//   host: "127.0.0.1",
//   user: "root",
//   database: "sample_db",
//   password: "root"
// });

// module.exports = pool.promise();

// ORM
const Sequelize = require("sequelize").Sequelize;

const sequelize = new Sequelize("sample_db", "root", "root", {
  dialect: "mysql",
  host: "127.0.0.1"
});

module.exports = sequelize;
