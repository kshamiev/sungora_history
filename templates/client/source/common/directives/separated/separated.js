angular.module('directives.separated', [
    'zegota.filters'
    ])

    // Форматированный вывод целого числа
    // Пример: 1234567 = 1'234'567
    .directive('currencyInput', function($filter, $browser) {
        return {
            require: 'ngModel',
            link: function($scope, $element, $attrs, ngModelCtrl) {
                var listener = function() {
                    var value = $element.val().replace(/'/g, '');
                    $element.val($filter('separatedNumber')(value));
                };

                // Вызывается при обновлении input
                ngModelCtrl.$parsers.push(function(viewValue) {
                    return viewValue.replace(/'/g, '');
                });

                // Вызывается при обновлении модели
                ngModelCtrl.$render = function() {
                    $element.val($filter('separatedNumber')(ngModelCtrl.$viewValue));
                };

                // Вызывается при нажатии клавиш
                $element.bind('change', listener);
                $element.bind('keydown', function(event) {
                    var key = event.keyCode;
                    if (key == 91 || (15 < key && key < 19) || (37 <= key && key <= 40))
                        return
                    $browser.defer(listener);
                });

                // Смотрим события вырезания и вставки
                $element.bind('paste cut', function() {
                    $browser.defer(listener);
                })
            }
        }
    });
