// directives.js

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
