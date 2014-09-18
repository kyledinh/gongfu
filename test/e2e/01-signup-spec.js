describe('gongfu page', function() {

    describe('Gongfu Signup', function() {
        var rows, target, token;
        var util = require('./Util');

        token = util.makeToken();

        beforeEach(function() {
        });

        it('signup for new account', function() {
            browser.get('webapp/app.html#/signup');
            element(by.id('inputName')).clear();
            element(by.id('inputName')).sendKeys(token);
            element(by.id('inputEmail')).sendKeys(token + "@gongfu.com");
            element(by.id('inputPassword1')).sendKeys(token);   
            element(by.id('inputPassword2')).sendKeys(token);                 
            element(by.id('btn_create_account')).click();

        });

        it('login', function() {
            browser.get('webapp/app.html#/login');
            element(by.id('inputEmail')).sendKeys(token + "@gongfu.com");
            element(by.id('inputPassword')).sendKeys(token);                   
            element(by.css('[class="btn btn-success"]')).click();     

            target = element(by.css('h1'));
            expect(target.getText()).toEqual('Chirpboard');

        });

        it('chirps', function() {
            element(by.css('[class="btn btn-xs btn-success glyphicon glyphicon-plus"]')).click();     
            element(by.id('neu-type')).sendKeys(token);                   
            element(by.id('neu-message')).sendKeys(token);                   
            element(by.id('neu-chirp')).click();     
        });

    });
});
