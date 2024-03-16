
const Dotenv = require('dotenv-webpack');

module.exports = {
    chainWebpack: config => {
        config.plugin('define').tap(definitions => {
            definitions[0]['process.env']['MODE_ENV'] = JSON.stringify(process.env.MODE_ENV || 'development');
            return definitions;
        });
    },
    configureWebpack: config => {
        let envPath = '.env.debug'; // 默认加载开发环境配置

        if (process.env.MODE_ENV === 'release') {
            envPath = '.env.release';
        } else if (process.env.MODE_ENV === 'test') {
            envPath = '.env.test';
        }

        config.plugins.push(new Dotenv({ path: envPath }));
    }
};