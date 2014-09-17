// ctrl-dashboard.js

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
   
