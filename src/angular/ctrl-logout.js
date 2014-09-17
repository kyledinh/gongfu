// ctrl-logout.js

function LogoutCtrl($scope, $http, $routeParams, $location, SessionFactory) {
    SessionFactory.dump();
    var logout = Util.mkJsonApiGet($scope, $http, "https://" + SERVER_URL + "/logout", "", "");
    logout();
}
   
