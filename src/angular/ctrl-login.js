// ctrl-login.js

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
   
