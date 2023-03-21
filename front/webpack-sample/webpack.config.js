const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const RemoveEmptyScriptsPlugin = require("webpack-remove-empty-scripts");
const CopyPlugin = require("copy-webpack-plugin");
const ImageMinimizerPlugin = require('image-minimizer-webpack-plugin');
const path = require('path');

module.exports = {
  mode: 'development',
  entry: './assets/index.js',
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'js/bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.scss$/i,
        use: [
          MiniCssExtractPlugin.loader,
          "css-loader",
          "sass-loader",
        ],
      },
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "css/styles.css",
    }),
    new RemoveEmptyScriptsPlugin(),
    new CopyPlugin({
      patterns: [
        { from: 'assets/img', to: 'img' },
        { from: 'assets/fonts', to: 'fonts' }
      ]
    }),
    new ImageMinimizerPlugin({
      test: /\.(png|jpe?g)$/i,
      minimizer: {
        implementation: ImageMinimizerPlugin.sharpMinify,
        options: {
          png: {
            quality: 80,
          },
        },
      },
    }),
  ],
}
