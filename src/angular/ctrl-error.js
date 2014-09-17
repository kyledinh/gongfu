// ctrl-error.js

var ErrorCtrl = app.controller("ErrorCtrl", function ($scope, $http) {
    $scope.toggleAccordion = Util.toggleAccordion;
    $scope.result = "OUTPUT";
});

// Error Controller
ErrorCtrl.loadData = function ($q, $timeout) {
    var defer = $q.defer();
    $timeout(function () {
        defer.resolve();
        console.log("loadData");
    }, 500);
    return defer.promise;
};

ErrorCtrl.makeError = function ($q, $timeout) {
    var defer = $q.defer();
    $timeout(function () {
        defer.reject();
    }, 500);
    return defer.promise;
};

