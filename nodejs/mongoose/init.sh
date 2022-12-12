use shop
db.createUser({user:"testuser", pwd:"password", roles:[{role:"readWrite", db: "shop"}]})
