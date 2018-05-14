const path = require('path');
const fs = require('fs');
const lessToJs = require('less-vars-to-js');
const rewireLess = require('react-app-rewire-less');
const { injectBabelPlugin } = require('react-app-rewired');
const rewireCssModules = require('react-app-rewire-css-modules');

const themeVars = lessToJs(fs.readFileSync(path.join(__dirname, './src/utils/styles/antTheme.less'), 'utf8'));

/* config-overrides.js */
module.exports = function override(config, env) {
  config = injectBabelPlugin(
    ['import', { libraryName: 'antd', style: true, libraryDirectory: 'es' }],
    config,
  );

  config = rewireLess.withLoaderOptions({
    modifyVars: themeVars,
    javascriptEnabled: true,
  })(config, env);

  config = rewireCssModules(config, env);

  return config;
};
