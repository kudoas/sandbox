const Products = require("../models/product");

exports.getAddProduct = (req, res, next) => {
  res.render("add-product", {
    pageTitle: "Add Product",
    path: "/admin/add-product",
    formsCSS: true,
    productCSS: true,
    activeAddProduct: true
  });
};

exports.postAddProduct = (req, res, next) => {
  const product = new Products(req.body.title);
  product.save();
  res.redirect("/");
};

exports.getProducts = (req, res, next) => {
  // fetchAllで呼び出してある
  Products.fetchAll(products => {
    res.render("shop", {
      prods: products,
      pageTitle: "Shop",
      path: "/",
      hasProducts: products.length > 0,
      activeShop: true,
      productCSS: true
    });
  });
};

// async, await
// const Product = require("../models/product");

// exports.getAddProduct = (req, res) => {
//   res.render('add-product', { docTitle: 'Add Product', path: '/admin/add-product' });
// }

// exports.postAddProduct = (req, res) => {
//   const price = Math.random() * 20;
//   const product = new Product(req.body.title, price.toFixed(2));
//   product.save();
//   res.redirect('/');
// }

// exports.getProducts = async (req, res) => {
//   const products = await Product.fetchAll();
//   res.render('shop', { prods: products, docTitle: 'Shop', path: '/' });

// }
