const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const { VueLoaderPlugin } = require("vue-loader");

module.exports = {
	mode: "development",
	entry: "./src/index.js",
	output: {
		filename: "main.js",
		path: path.resolve(__dirname, "dist"),
	},
	module: {
		rules: [
			{
				test: /\.vue$/,
				loader: "vue-loader",
			},
			{
				test: /\.js$/,
				loader: "babel-loader",
			},
			{
				test: /\.css$/,
				use: ["vue-style-loader", "css-loader"],
			},
		],
	},
	plugins: [
		new VueLoaderPlugin(),
		new HtmlWebpackPlugin({
			title: "Cart Demo with Webpack",
			template: path.resolve(__dirname, "public/index.html"),
		}),
	],
	devServer: {
		contentBase: path.join(__dirname, "public"),
		compress: true,
		port: 8000,
	},
};
