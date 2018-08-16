'use strict'

const webpack = require('webpack')
const { VueLoaderPlugin } = require('vue-loader')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const WorkboxPlugin = require('workbox-webpack-plugin');

const path = require('path')

function resolve(dir) {
    return path.join(__dirname, '..', dir)
}

module.exports = {
    watch: true,
    mode: 'development',

    entry: [
        './src/app.js'
    ],

    module: {
        rules: [{
                test: /\.vue$/,
                use: 'vue-loader'
            },
            {
                test: /\.css$/,
                use: [
                    'vue-style-loader',
                    'css-loader'
                ]
            },
            {
                test: /\.js$/,
                use: 'babel-loader'
            }
        ]
    },

    plugins: [
        new VueLoaderPlugin(),
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: 'src/index.html',
            inject: true
        }),
        new CopyWebpackPlugin([{
            from: resolve('static/img'),
            to: resolve('dist/static/img'),
            toType: 'dir'
        }]),
        new WorkboxPlugin.GenerateSW({
            // these options encourage the ServiceWorkers to get in there fast 
            // and not allow any straggling "old" SWs to hang around
            clientsClaim: true,
            skipWaiting: true
        })
    ]
}
