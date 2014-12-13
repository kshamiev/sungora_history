angular.module('directives.datetimepicker', [
		'ui.bootstrap',
		'templates.bootstrap',
		'directives/datetimepicker/datetimepicker.tpl.html',
        'mgcrea.ngStrap.datepicker',
        'mgcrea.ngStrap.timepicker',
	])

    .config(function($datepickerProvider) {
        angular.extend($datepickerProvider.defaults, {
            dateFormat: 'dd.MM.yyyy',
            startWeek: 1,
            delay: { show: 0, hide: 100 },
            autoclose: true
        });
    })

    .config(function($timepickerProvider) {
        angular.extend($timepickerProvider.defaults, {
            timeFormat: 'HH:mm',
            length: 7,
            hourStep: 1,
            minuteStep: 1,
            delay: { show: 0, hide: 100 }
        });
    })

	.directive('zegotaDatepicker', [function () {
		return {
			restrict: 'E',
			scope: {
				datetime: "=ngModel"
			},
			require: 'ngModel',
			templateUrl: 'directives/datetimepicker/datetimepicker.tpl.html',
			link: function ($scope, $element, $attrs, $controller) {

			},
			controller: ['$scope',
				function ($scope) {
					var today = function () {
                        $scope.datetime = new Date();
					};
                    // Разделение на Дату и время
                    var Parse = function() {
                        $scope.date = new Date($scope.datetime.getFullYear(), $scope.datetime.getMonth(), $scope.datetime.getDate(), 0, 0, 0, 0);
                        $scope.time = new Date(0,0,0, $scope.datetime.getHours(), $scope.datetime.getMinutes(), $scope.datetime.getSeconds(), 0);
                    };

                    // Исключаем возможные ошибки
                    if ( typeof $scope.datetime === 'undefined')
                        today();
                    else
                        $scope.datetime = new Date($scope.datetime);

                    // Первоначальное...
                    Parse();

                    // Если модель поменялась, заново разделяем.
                    $scope.$watch("datetime", function(nV, oV){
                        Parse();
                    });

                    // Собираем куски в в модель
					$scope.$watch("[date,time]", function (newValues, oldValues) {
                        if (typeof newValues[0] === 'object' && typeof newValues[1] === 'object') {
                            var h = newValues[1].getHours() * 60 * 60;
                            var m = newValues[1].getMinutes() * 60;
//                            var s = newValues[1].getSeconds();
                            $scope.datetime = new Date(newValues[0].getTime() + (h+m)*1000 );
                        }
					}, true);
				}]
		};
	}]);
