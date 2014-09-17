(function () {

var fs = require('fs');
var Util = {};

Util.snapshot = function (filename) {
    browser.sleep(2000);
    browser.takeScreenshot().then(function(data) {
        var buf = new Buffer(data, 'base64'),
            stream = fs.createWriteStream(filename);
        stream.write(buf);
        stream.end();
    });
};

Util.makeToken = function () {
    var prefix = 'bot-';
    return prefix + Math.floor(Math.random()*999 +1).toString();
};

// Nodejs export
if (typeof exports !== 'undefined') {
    exports.snapshot = Util.snapshot;
    exports.loginWithTestAccount = Util.loginWithTestAccount;
    exports.makeToken = Util.makeToken;
} else {
    this.Util = Util;
}

})();