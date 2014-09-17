// system.js
// places util into the global namespace

var Util = {};

Util.toggleAccordion = function (id) {
    console.log("boot style.height: " + document.getElementById(id).style.height);
    document.getElementById(id).classList.add('in');   
    if (document.getElementById(id).style.height === "auto") { 
        document.getElementById(id).style.height = "0px";
        document.getElementById(id).style.display = "none";
    } else {
        document.getElementById(id).style.display = "block";
        document.getElementById(id).style.height = "auto";
    }
};

Util.openModal = function (id) {
    document.getElementById("body").classList.add('modal-open');           
    document.getElementById(id).classList.add('in');
    document.getElementById(id).setAttribute("aria-hidden","true");
    document.getElementById(id).style.display = "block";
    document.getElementById("backdrop").classList.add('modal-backdrop');
};

Util.closeModal = function (id) { 
    document.getElementById("body").classList.remove('modal-open');
    document.getElementById(id).classList.remove('in');
    document.getElementById(id).setAttribute("aria-hidden","false");      
    document.getElementById(id).style.display = "none";
    document.getElementById("backdrop").classList.remove('modal-backdrop');      
};

// target is the $scope[obj]
Util.mkJsonApiCaller = function ($scope, $http, _method, _url, _data, _params, target, callback) {
    return function () {
        $http({
            method: _method,
            url: _url,
            data: _data,
            params: _params,
            headers: {"Content-Type": "application/json"}
        }).success(function (data, status, headers, config) {   
            $scope[target] = data.payload;
            if (callback != null) {
                callback();
            }                
        }).error(function (data, status, headers, config) {
            //$scope.message = status + " " + config;
            alert(status + " error, can not process request.");
        });
    };
};

Util.mkJsonApiGet = function ($scope, $http, _url, _params, target, callback) {
    return function () {
        $http({
            method: "GET",
            url: _url,
            params: _params,
            headers: {"Content-Type": "application/json"}
        }).success(function (data, status, headers, config) {
            $scope[target] = data.payload;
            if (callback != null) {
                callback();
            }
        }).error(function (data, status, headers, config) {
            //$scope.message = status + " " + config;
            alert(status + " error, can not process request.");
        });
    };
};

Util.mkJsonApiDelete = function ($scope, $http, _url, _params, target, callback) {
    return function () {
        $http({
            method: "DELETE",
            url: _url,
            params: _params,
            headers: {"Content-Type": "application/json"}
        }).success(function (data, status, headers, config) {   
            $scope[target] = data.payload;
            if (callback != null) {
                callback();
            }                
        }).error(function (data, status, headers, config) {
            //$scope.message = status + " " + config;
            alert(status + " error, request not updated.");
        });
    };
};

Util.mkJsonApiPost = function ($scope, $http, _url, _data, target, callback) {
    return function () {
        $http({
            method: "POST",
            url: _url,
            data: JSON.stringify(_data),
            headers: {"Content-Type": "application/json"}
        }).success(function (data, status, headers, config) {
            $scope[target] = data.payload;
            if (callback != null) {
                callback();
            }                
        }).error(function (data, status, headers, config) {
            //$scope.message = status + " " + config;
            alert(status + " error, request not updated.");
        });
    };
};

Util.mkJsonApiPatch = function ($scope, $http, _url, _data, target, callback) {
    return function () {
        $http({
            method: "PATCH",
            url: _url,
            data: JSON.stringify(_data),
            headers: {"Content-Type": "application/json"}
        }).success(function (data, status, headers, config) {
            $scope[target] = data.payload;
            if (callback != null) {
                callback();
            }                
        }).error(function (data, status, headers, config) {
            //$scope.message = status + " " + config;
            alert(status + " error, request not updated.");
        });
    };
};
