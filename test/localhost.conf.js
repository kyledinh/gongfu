exports.config = {

  // Capabilities to be passed to the webdriver instance.
  baseUrl: 'https://192.168.10.10/',
  capabilities: {
    'browserName': 'chrome',
    'chromeOptions': {
      args: ['--test-type']
    }
  },

  // Spec patterns are relative to the current working directly when
  // Active tests has a '-' dash in this format: 01-homepage.js
  // Inactive test can be left in diretory but remove the '##-' number prefix
  specs: [ 'e2e/*-*.js' ]
};
