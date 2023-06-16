const path = require("path");
const webpack = require("webpack");
const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
  entry: "./src/pages/index/index.js",
  module: {
    rules: [
      {
        test: /\.worker\.js$/,
        use: { loader: "worker-loader" },
      },
      {
        test: /\.(js|jsx)$/i,
        use: [
          {
            loader: "babel-loader",
            options: {
              presets: ["@babel/preset-env"],
            },
          },
        ],
      },
    ],
  },
  plugins: [
    new CopyWebpackPlugin([
      {
        from: "src/pages/doc/transcode.js",
        to: "data/transcode.js",
      },
    ]),
    // new BundleAnalyzerPlugin()
  ],
  // devServer: conf.devServer || {},
  resolve: {
    extensions: [".js", ".scss", ".less", ".css"],
    modules: [
      "node_modules",
      path.join(process.cwd(), "./src"),
      path.join(process.cwd(), "./src/scss"),
      path.join(process.cwd(), "./src/templates"),
      path.join(process.cwd(), "./src/assets"),
    ],
  },
};
