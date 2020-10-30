module.exports = {
    outputDir: "public",
    publicPath: "/public",
    pages: {
        index: {
            // page 的入口
            entry: 'assets/main.ts',
            // 模板来源
            template: 'index.html',
            // 在 dist/index.html 的输出
            filename: 'index.html',
            // 当使用 title 选项时，
            // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
            title: 'Vue WebSocket TypeScript Demo',
            // 在这个页面中包含的块，默认情况下会包含
            // 提取出来的通用 chunk 和 vendor chunk。
            chunks: ['chunk-vendors', 'chunk-common', 'index']
        }
    }
}