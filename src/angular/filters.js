// filters.js

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
