module.exports = {
  presets: ['@babel/preset-env', '@vue/app'],
  plugins: [
    [
      'component',
      {
        'libraryName': 'element-ui',
        'styleLibraryName': 'theme-chalk'
      }
    ]
  ]
}
