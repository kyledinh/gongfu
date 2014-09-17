// factories.js

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
