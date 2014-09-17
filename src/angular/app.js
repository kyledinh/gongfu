var SERVER_PORT = "";
var SERVER_URL = "192.168.10.10" + SERVER_PORT;
var LAPPA_API_URL = "https://" + SERVER_URL + "/api";

var app = angular.module("app", []);

app.config(function ($routeProvider) {
    $routeProvider.when('/', { controller:"AppCtrl", templateUrl:"app_tpl" })     
        .when('/dashboard', { controller:"DashboardCtrl", templateUrl:"/html/dashboard.tpl" })
        .when('/loggedin/:id', { controller:"LoginCtrl", templateUrl:"/html/login.tpl" })                      
        .when('/login', { controller:"LoginCtrl", templateUrl:"/html/login.tpl" })  
        .when('/logout', { controller:"LogoutCtrl", templateUrl:"/html/login.tpl" })                          
        .when('/signup', { controller:"SignupCtrl", templateUrl:"/html/signup.tpl" })           
        .when('/error', { controller:"ErrorCtrl", templateUrl:"error_tpl", 
            resolve: {
                makeError: ErrorCtrl.makeError,
            } 
        })     
        .otherwise( {template: "404 Not Found!"} );
});

app.run(function ($rootScope, $log) {
    $rootScope.$log = $log;
});

app.controller("AppCtrl", function ($rootScope) {  
    $rootScope.$on("$routeChangeError", function (event, current, previous, rejection) {
        console.log(rejection);
    })   
});

