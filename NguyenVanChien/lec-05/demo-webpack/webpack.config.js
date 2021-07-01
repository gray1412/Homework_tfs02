const path = require('path');
const htmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    mode: 'development',
    entry: path.resolve(__dirname, 'src/index.js'),
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "main.js"
    },
    module: {
        rules: [
            {
                test: /\.js$/, use: 'babel-loader',
            },
            {
                test: /\.css$/i,  use: ["style-loader", "css-loader"],
            },
        ],
    },
    devServer: {
        port: 8080,
        contentBase: path.resolve(__dirname, "public"),
    },
    plugins: [
        new htmlWebpackPlugin({
            title: "Hello",
            template: path.resolve(__dirname, "public/index.html"),
        })
    ]
}