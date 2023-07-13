const HtmlWebpackPlugin = require("html-webpack-plugin");
const copyWebpackPlugin = require("copy-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const fileListTxtWebpackPlugin = require("./build/fileListTxtWebpackPlugin");
const webpack = require("webpack");
const path = require("path");

module.exports = {
  mode: "production",
  entry: "./src/index.js",
  module: {
    rules: [
      { test: /\.css$/, use: [MiniCssExtractPlugin.loader, "css-loader"] },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin(),
    new webpack.ProgressPlugin(),
    // new webpack.BannerPlugin("banner"),
    new HtmlWebpackPlugin({ template: "./src/index.html" }),
    new copyWebpackPlugin({
      patterns: [
        {
          from: "src/**/*.json",
          to: ({ context }) => {
            console.log(context);
            console.log(path.parse(context).base);
            return Promise.resolve("src/[name][ext]");
          },
        },
      ],
    }),
  ],
  optimization: {
    splitChunks: {
      chunks: "all",
      //   name: "vendor",
      minSize: 2000,
      cacheGroups: {
        vendor: {
          name: "vendor",
          priority: 10,
          test: /[\\/]node_modules[\\/]/,
        },
        // lodash: {
        //   name: "chunk-lodash",
        //   priority: 20,
        //   test: /[\\/]node_modules[\\/](.*)lodash(.*)/,
        // },
        // css: {
        //   name: "chunk-css",
        //   priority: 10,
        //   test: /.*css/,
        // },
      },
    },
  },
};
