angular.module('zegota.filters', [])

    // find by id
    .filter('getById', function () {
        return function (input, Id) {
            var i = 0, len = input.length;
            for (; i < len; i++) {
                if (+input[i].Id === +Id) {
                    return input[i];
                }
            }
            return null;
        };
    })

	// поиск по Code в объекте
	.filter('getByCode', function () {
		return function (input, Code) {
			var i = 0, len = input.length;
			for (; i < len; i++) {
				if (input[i].Code === Code) {
					return input[i];
				}
			}
			return null;
		};
	})

	// upper first letter
	.filter('capitalize', function() {
		return function(input) {
			return input.substring(0,1).toUpperCase()+input.substring(1);
		};
	})

    .filter('currency', function(){
        return function(input) {
            var decimal   = 2;
            var separator = "'";
            var inp = parseFloat(input)
            if ( input === 0 ) {
                return "-";
            }

            var exp10=Math.pow(10,decimal);
            inp = Math.round(inp*exp10)/exp10;
            var out = Number(inp).toFixed(decimal).toString().split('.');
            var b = out[0].replace(/(\d{1,3}(?=(\d{3})+(?:\.\d|\b)))/g, '\$1'+separator);
            var ret = (out[1]?b+'.'+out[1]:b) + " руб.";
            return ret;
        };
    })

    .filter('separatedNumber', function(){
        return function(input) {
            var decimal   = 2;
            var separator = "'";
            var inp = parseInt(input||0)
            if ( input === 0 ) {
                return 0;
            }

            var exp10=Math.pow(10,decimal);
            inp = Math.round(inp*exp10)/exp10;
            var out = Number(inp).toFixed(decimal).toString().split('.');
            var b = out[0].replace(/(\d{1,3}(?=(\d{3})+(?:\.\d|\b)))/g, '\$1'+separator);
            return b;
        };
    });
;
