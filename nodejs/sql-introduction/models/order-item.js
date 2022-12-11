const Squelize = require("sequelize");

const sequelize = require("../util/database");

const OrderItem = sequelize.define("orderItem", {
  id: {
    type: Squelize.INTEGER,
    autoIncrement: true,
    allowNull: false,
    primaryKey: true,
  },
  quantity: Squelize.INTEGER,
});

module.exports = OrderItem;
