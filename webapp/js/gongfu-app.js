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

;// ctrl-dashboard.js

function DashboardCtrl($scope, $http, $location, SessionFactory) {

    // Redirect if not logged in
    var sess = SessionFactory.email();    
    if (sess.length < 1) { 
        $location.path('/login');
        return;
    }

    $scope.chirps = [];
    $scope.user = {};
    $scope.chirp = {};
    $scope.closeMod = Util.closeModal;
    $scope.openMod = Util.openModal;    
    $scope.toggleAccordion = Util.toggleAccordion;

    $scope.deleteChirp = function (id) {
        // Current error reading the arg in this function, resorted to pulling id from $scope
        pid = $scope.chirp.id;
        console.log("deletePayment: " + pid);
        var xhr = Util.mkJsonApiDelete($scope, $http, LAPPA_API_URL + "/chirp/" + pid, "", "message", xhrChirps);
        xhr();
        Util.closeModal('mod-chirp');             
    };

    $scope.editUser = function (id) {
        //var xhr = Util.mkJsonApiGet($scope, $http, LAPPA_API_URL + "/tenant/" + id, "", "tenant");
        //xhr();
        Util.openModal('mod-user');
    };

    $scope.editChirp = function (id) {     
        var xhr = Util.mkJsonApiGet($scope, $http, LAPPA_API_URL + "/chirp/" + id, "", "chirp");
        xhr();
        Util.openModal('mod-chirp');
    };

    $scope.patchUser = function () {
        var obj = $scope.user;
        var xhr = Util.mkJsonApiPatch($scope, $http, LAPPA_API_URL + "/user", obj, "message");
        xhr();
        Util.closeModal('mod-user');        
    };

    $scope.patchChirp = function () {
        var obj = $scope.chirp;
        var xhr = Util.mkJsonApiPatch($scope, $http, LAPPA_API_URL + "/chirp", obj, "message", xhrChirps);
        xhr();
        Util.closeModal('mod-chirp');                
    };

    $scope.postChirp = function () {
        var obj = $scope.neu;
        var xhr = Util.mkJsonApiPost($scope, $http, LAPPA_API_URL + "/chirp", obj, "message", xhrChirps);
        xhr();
        $scope.neu = {};
        Util.closeModal('mod-new-chirp');                
    };

    var xhrChirps = Util.mkJsonApiGet($scope, $http, LAPPA_API_URL + "/chirps", "", "chirps");
    xhrChirps();

    var xhrUser = Util.mkJsonApiGet($scope, $http, LAPPA_API_URL + "/user", "", "user");
    xhrUser();


}
   
;// ctrl-error.js

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

;// ctrl-login.js

function LoginCtrl($scope, $http, $routeParams, $location, SessionFactory) {
    var id = $routeParams.id || "";

    if (id === "") { 
        // no new email address in id param
        // show login form
    } else {
    	var sess = { email : id };
        SessionFactory.save(sess);
        $location.path('/dashboard');
    }
}
   
;// ctrl-logout.js

function LogoutCtrl($scope, $http, $routeParams, $location, SessionFactory) {
    SessionFactory.dump();
    var logout = Util.mkJsonApiGet($scope, $http, "https://" + SERVER_URL + "/logout", "", "");
    logout();
}
   
;// ctrl-signup.js

function SignupCtrl($scope, $http, DataService) {
    $scope.data = DataService;
}

;// directives.js

app.directive("error", function ($rootScope) {
    return {
        restrict:"E",
        template:'<div class="alert-box alert" ng-show="isError">Error!!!</div>',
        link: function (scope) {
            $rootScope.$on("$routeChangeError", function (event, current, previous, rejection) {
                scope.isError = true;
            })
        }
    }
})

app.directive("areaChart", function () {
    return {
        restrict:"E",
        template:'<areaChart class="alert-box alert"></areaChart>',
        scope: { 'ngModel': '=' },
        link: function (scope, elements, attrs, ctrl) {
            var width = attrs.width || 600;
            var height = attrs.height || 400;
            var margin = {top: 20, right: 20, bottom: 30, left: 50};

            width = width - margin.left - margin.right;
            height = height - margin.top - margin.bottom;

            // https://github.com/mbostock/d3/wiki/Time-Formatting
            var parseDate = d3.time.format("%Y-%m-%d").parse; //2013-05-27

            var x = d3.time.scale().range([0, width]);
            var y = d3.scale.linear().range([height, 0]);

            var xAxis = d3.svg.axis().scale(x).orient("bottom");
            var yAxis = d3.svg.axis().scale(y).orient("left");

            var area = d3.svg.area()
                .x(function(d) { return x(d.date); })
                .y0(height).y1(function(d) { return y(d.close); });

            d3.select("svg").remove();
            var svg = d3.select("#chartdata").append("svg")
                .attr("width", width + margin.left + margin.right)
                .attr("height", height + margin.top + margin.bottom)
                .append("g")
                .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

            //d3.tsv("/data/mtgox.tsv", function(error, data) {
            d3.tsv("/chart/" + daysago, function(error, data) {
                data.forEach(function(d) {
                d.date = parseDate(d.date);
                d.close = +d.close;
            });

            x.domain(d3.extent(data, function(d) { return d.date; }));
            y.domain([0, d3.max(data, function(d) { return d.close; })]);

            svg.append("path").datum(data).attr("class", "area").attr("d", area);

            svg.append("g").attr("class", "x axis").attr("transform", "translate(0," + height + ")").call(xAxis);

            svg.append("g").attr("class", "y axis").call(yAxis)
                .append("text").attr("transform", "rotate(-90)").attr("y", 6)
                .attr("dy", ".71em").style("text-anchor", "end").text("BitStamp Price (USD)");
            });   
            
        }
    }
})
;// factories.js

app.factory('SessionFactory', function () {
    return {
        tenantOptions: function () {
            return [ {name: "", id: 0}, {name: "Annalise", id: 1}, {name: "Bob", id: 2}, {name: "Carson", id: 3} ];
        },
        all: function () {
            var dataStr = window.localStorage['lib-pmt.session'];
            if (dataStr) { return angular.fromJson(dataStr); }
            return [];
        },
        email: function () {
            var dataStr = window.localStorage['gongfu'];
            if (dataStr) { return angular.fromJson(dataStr); }
            return [];
        },
        save: function (data) {
            window.localStorage['gongfu'] = angular.toJson(data);
        },
        last: function (num) {
            var dataStr = window.localStorage['lib-pmt.session'];
            if (dataStr) { return angular.fromJson(dataStr).slice(-num); }
            return [];
        },
        dump: function () {
            var empty = [];
            window.localStorage['gongfu'] = angular.toJson(empty);
        },
        newAct: function (grade, desc, type) {
            var grade = grade || "=";
            var date = Date.now();
            var desc = desc || "";
            var type = type || "";
            return {
                grade: grade,
                date: date,
                desc: desc,
                type: type
            };
        }
    };
});
;// filters.js

app.filter('matchPropertyid', function () {
    return function (items, id) {
        var filtered = [];
        var item, i, n = 0;
        if (Object.prototype.toString.call(items) === '[object Array]') {
            n = items.length;
        }
 
        for (i = 0; i < n; i++) {
            item = items[i];
            if (item.propertyid === id) {
                filtered.push(item);
            }
        }
        return filtered;
    };
});
;// system.js
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
