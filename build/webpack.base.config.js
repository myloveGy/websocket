const htmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
    entry: './assets/index.ts',
    output: {
        filename: 'app.js'
    },
    resolve: {
        extensions: ['.js', '.ts', '.tsx']
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/i,
                use: [{
                    loader: 'ts-loader',
                }],
                exclude: /node_modules/
            }  
        ]
    },
    plugins: [
        new htmlWebpackPlugin({
            template: './assets/index.html'
        })
    ]
}