/* eslint-disable camelcase */
const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const VueLoaderPlugin = require('vue-loader/lib/plugin')
// const WorkboxPlugin = require('workbox-webpack-plugin')
// const WebpackPwaManifest = require('webpack-pwa-manifest')
// const PrerenderSPAPlugin = require('prerender-spa-plugin')
const CopyPlugin = require('copy-webpack-plugin')

const {
  meta, 
  port,
} = require('./config')

const plugins = [
  new HtmlWebpackPlugin({
    inject: true,
    template: './src/index.pug',
  }),

  new CopyPlugin({
    patterns: [
      {
        from: 'public',
        to: '', 
      },
    ],

    options: {
      concurrency: 100,
    },
  }),

  new VueLoaderPlugin(),
]

module.exports = () => ({
  plugins,

  module: {
    rules: [
      {
        test: /\.pug$/,
        oneOf: [
          {
            resourceQuery: /^\?vue/,
            use: ['pug-plain-loader'],
          },
          {
            use: 'pug-loader',
          },
        ],
      },
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader',
        ],
      },
      {
        test: /\.css$/,
        use: [
          'css-loader',
        ],
      },
      {
        test: /\.(mp4|webm|webp|png|jpg|gif|woff|woff2|eot|ttf|otf)$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              esModule: false,
            },
          },
        ],
      },
      {
        test: /\.svg$/,
        use: 'vue-svg-loader',
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
      },
    ],
  },

  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      'config': path.resolve(__dirname, 'config.js'),
      'css': path.resolve(__dirname, 'src/css'),
      'sass': path.resolve(__dirname, 'src/sass'),
      'assets': path.resolve(__dirname, 'src/assets'),
    },
    extensions: ['.js', '.vue'],
  },

  entry: {
    app: './src/index.js',
  },

  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: "[name].bundle.js",
  },

  devtool: 'source-map',

  devServer: {
    port: port,
    host: '0.0.0.0',
    historyApiFallback: true,
  },
})
