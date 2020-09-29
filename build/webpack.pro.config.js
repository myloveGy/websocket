const {CleanWebpackPlugin} = require('clean-webpack-plugin')
const webpack = require('webpack')
const UglifyJsPlugin=require('uglifyjs-webpack-plugin')

module.exports = {
    devtool: '#source-map',
    optimization:{
        minimizer:[
            new UglifyJsPlugin({
                uglifyOptions: {
                    output: {
                        comments: false
                    },
                    compress: {
                        warnings: false,
                        drop_debugger: true,
                        drop_console: true
                    }
                }
            })
        ]
    },
    plugins: [
        new CleanWebpackPlugin(),
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: '"production"'
            }
        }),
        new webpack.LoaderOptionsPlugin({
            minimize: true
        })
    ]
}