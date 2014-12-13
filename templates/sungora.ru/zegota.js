/*! zegota - v0.2.0 - 2014-11-24
 * Copyright (c) 2014 ;
 * Licensed 
 */
angular.module('zegotaUsr', [
  'zegota.core',
  'zegota.ui',
  'zegota.recovery',
  'zegota.registration',
  'zegota.profile',
  'templates.user'//,
  //'ngMockE2E'
])
.config(function($interpolateProvider){
    $interpolateProvider.startSymbol('{[{').endSymbol('}]}');
  }
)
//.run(function($httpBackend) {

  //$httpBackend.whenPOST('/api/v1.0/session').respond(function(method,url,data) {
    //var icomingData, userWithCaptcha, userWithoutCaptcha, fakeUser, responseData;

    //userWithCaptcha = {
      //Login: 'alexp',
      //Password: '123123',
      //Captcha: '583191'
    //};

    //userWithoutCaptcha = {
      //Login: 'alexp',
      //Password: '123123',
    //};

    //incomingData = angular.fromJson(data);

    //for (prop in incomingData) {
      //if (incomingData.hasOwnProperty('Captcha')) {
        //fakeUser = userWithCaptcha;
      //} else {
        //fakeUser = userWithoutCaptcha;
      //}
    //}

    //if (JSON.stringify(fakeUser) === JSON.stringify(incomingData)) {
      //responseData = {
        //ErrorCode: 0,
        //ErrorMessage: '',
        //Token: 'a468d9e24f81fe38919d31a34e629d844d899be0470824afdb656b2d3899e5b5'
      //};
      //return [200, responseData];
    //} else {
      //responseData = {
        //ErrorCode: 2,
        //ErrorMessage: 'Ошибка,  логин или пароль указан не верно',
        //Token: ''
      //};
      //return [404, responseData];
    //}
  //});

  //$httpBackend.whenGET('/api/v1.0/session/')
    //.respond({
      //ErrorCode: 0,
      //ErrorMessage: '',
      //Token: 'a468d9e24f81fe38919d31a34e629d844d899be0470824afdb656b2d3899e5b5'
    //});

  //$httpBackend.whenGET('/api/v1.0/session/a468d9e24f81fe38919d31a34e629d844d899be0470824afdb656b2d3899e5b5')
    //.respond({
      //ErrorCode: 0,
      //ErrorMessage: ''
  //});

  //$httpBackend.whenGET('/api/v1.0/currentuser/a468d9e24f81fe38919d31a34e629d844d899be0470824afdb656b2d3899e5b5')
    //.respond({
      //Login: "alexp",
      //LastName: "p",
      //Name: "alex"
  //});

  //$httpBackend.whenDELETE('/api/v1.0/session/a468d9e24f81fe38919d31a34e629d844d899be0470824afdb656b2d3899e5b5')
    //.respond(function () {
      //return [200, {}, {}];
  //});

  //$httpBackend.whenGET('/api/v1.0/captcha/native')
    //.respond(function () {
      //var responseData, captchaImage;
      //captchaImage = "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAPAAAABQCAMAAAAQlwhOAAAAP1BMVEUAAABvQWA1ByakdpUyBCN4SmmqfJtCFDNeME9RI0JgMlGzhaRmOFe1h6aHWXiMXn13SWi2iKd4SmmugJ+NX359HBcRAAAAAXRSTlMAQObYZgAABsVJREFUeJzkW+ly2zgMJpL8qWfaZNP3f8Q92p3JvxZr8QJAAjwkam23cJvEMgniIw4CkPzipuirc3/Ozbg3epoa/TX+P0Dvx6YfJgb4crmcv977rRET4Ev8fy79dX3dlCD/5cF+dIbfqw+/OffP2MhJwHdKb9uPMcRk0h/uYfHOEDuWfgO0jpv0QxP5cM+bfxHAmbrePJlp3T19c+61OWAq05pLy25Bm8F+644YpSven4fEkSsjOFzHLjLdfjS5zijt50K8AO4EvAM0ZaXL8Aa7OgEvdtk+r1+0T2B6kv3JKroFYFW/3sbT60Tqm/TR5cv5Bj90wau332dC7gKGsPf7CGrpwXDgBBbiz7NoKGjt2nNIVooAnJMxmmsYT4zfPcCQfk3qGZLOPACJGGtEEay/jmG1k6gDOMhHQo3KAZvKEPLveBU8Vo0NnSf9k+UQtRGExAWQjxoRJik1TUXGD6FmUlwCP3NgnR3U1jCmF7PBrpIhOGSYlQNv+MTwTjYmvlesoG9fz8/9U7YNGHhQxSRMj2LWiGE0VFC0EC3yTKx3ZsSfnp2aV3wW75qAhTlKPduygByJmDnZAvc0nCLmnlD2WSJuAq7gYZSfJKEMKUmjmCMmbWPlrpmNvCj3FojDPP17fYmlGgSd2JEDU3g51CdJjGpZWF1kc1gEbAnrvD3/6AzpAXatXYUigBsSyR0wNFwHbtoCyOF+QeRumTSIX/XHlBxhzgv1YT1BSx8WRg4szz5OfR824mOIIwlrHKM6JzdWsPQE5Vt+DPqFFuXXDcDAksN6tSAQ5LOKgvFGr5KFMlMuVag4HuSR5RYqlyXYDcDJSlv1C/I/SKRX3jrsyZnCPMtA45pJimxpp1OURS8cyhQfALiG/Tx/jS7XjPIYl3679K4atgjSyCDUstvKHaGMxpRLMwOVk2h9GospLJcDV9j04LbpXRkpU1EmSABxv4ojOXANF0WmpWzMonLiAOByz0swn5z7kOe0BhgZFNogrFKRVRoebNNqIbI+PBl5vO5CabWSV4b0mOfd4SodvJzfrihdO/5oXxqhhoeCpRSJ1z52VoIoc3MqyIYKxjaBKzP0jYYb8UWrhrhiDOVFQiWBaN0qQMVqEJOoRzUMvsFS0dSdB77J/mEcOqrrOyfiiYKqH2kcdXmhwxreIr2ioAkmMm69bw/kxKgTX/3JxECN+uXYI1E6mKOW1M3w4IK8+8ePeBRuSiS7WS28jsoj68wbE7azne3p2ICWsbREAtnP6x0yUK4yCdi+MzmWabFwVULjyUNd239x7rtcB7v6zeOtjsCIxKa5GUGLhxQQxVJaOHeZZD0s48SX+J+CLFq3Wor16wszURrsEtyM0tuSLy/pERCqihCTOv3pgYk5gn3c5mH0uRpOaI9dykPzujNR2nRg166HN7QvmUWWEz3qmIgQb0zGqpU10RqkodSD8ghUOA0ruLkxBmAeLmLuJyt89IdcUBorilFO/u6SD5fsFemjAYD+OMS4itvebvHY4o9vam9tQMy+WcWRinndf8SqDrIPDFYbFkPHD+J9gP0nqe1pxWBjE6ItIDuMXHgGCJv+RcWzo03GMRxcgk+txygbVlJ/pPYAyqtRMHDyhpr3nR/tgFJilqMHn3OCrUyzEdtBC0pHU9qtIdUXNpsUgZQRM930+jSxnPRxI3FJ2Xcjm5BiN6lxLCVJcyBSqptk61rplysGKgKH8iVM/9JgupHTj1rdyPZk7XlYEHNSoYupnUVUzueutQvBoHsfpOBAK8wUSu0Hv0G3EyiismkozOUc2bcxei495Cs4NWiqg/smreUJIRAhO2EN4qdzND0zBwTVK3oUTS0lrk1dV3GnpievFp8us8SufzeI5GF2m+++LCXM4Rq7j66NOHngRfUqKE3SLg8poEF/OPf3MFN9CSx9rRrV1XAYCalsu3KcxStikzOPHrji3TDvIlGQtYyoK/ZTLICyHzZar01xov8E065ZHL1RgrFlk+xaZTeipJfMDqnA3SVRjFbB6GIrXWTCu9gSAaY6Ru+gDR4BIHkuup3B2xqE9W2/D3uu7aR00AuP7ntj9Qg5/r1gL5uYRs/4c763sVXLwTdCnbfmdjY29DOC95IGnkXULVjG0Wg2D1n05f/48h2su5NNDI1lunNP1/AZZFXgw2b0aICpy4fi/bDbPBxg3kuJ7/LfI3T/366riHd+5uPDA2pY5L+xvbm70rkFjeZI9bxW5WTR7U0a3M5tx17lpNJNvoon6IiN7Zh7ew0foR2usMSHj32xeK8P717tMHkrWfhV6lPpsU16B/12gP8LAAD//6Fm9sTJxyrcAAAAAElFTkSuQmCC";
      //responseData = {
        //ErrorCode: 0,
        //ErrorMessage: '',
        //CaptchaImage: captchaImage
      //};
      //return [200, responseData, {}];
  //});
//})
;

angular.module('zegota.profile.config', [])

.value('profileUrl', {
  profilePutData: 'http://gocmf.webdesk.ru/api/v1/access/users/t/',
  profileGetCurrentUserData: 'api/users/current-user'
});

"use strict";

angular.module('zegota.profile', [
  'zegota.profile.config',
  'zegota.profile.service',
  'zegota.profile.directive',
  'zegota.profile.controller'
]);

"use strict";

angular.module('zegota.profile.controller', [])

.controller('ProfileController', ['$scope', 'security', 'profile', '$http', 'profileUrl',
  function ($scope, security, profile, $http, profileUrl) {

  $scope.model = {};
  var token = security.getCookieToken();

  function getUserData () {
    var request = $http.get(profileUrl.profileGetCurrentUserData);
    request.then(function (response) {
      if (response.status == 200) {
        angular.extend($scope.model, response.data);
      }
    }, function (error) {
      console.log('error is: ', error);
    });
  }

  getUserData ();

  $scope.submitUserData = profile.profile($scope.model, token);
  $scope.cancelProfile = profile.cancelProfile;
}]);

angular.module('zegota.profile.directive', [])

.directive('profileToolbar', ['security', 'profile', function(security, profile) {
  return {
    templateUrl: 'profile/profile.toolbar.tpl.html',
    restrict: 'E',
    replace: true,
    scope: true,
    link: function ($scope) {
      $scope.isAuthenticated = security.isAuthenticated;
      $scope.profile = profile.showProfile;
    }
  }
}]);

angular.module('zegota.profile.service', [])

.factory('profile', ['$http', '$modal', 'profileUrl', '$window', 'modalWindowControl', function($http, $modal, profileUrl, $window, modalWindowControl){

    var dialogParams = {
      keyboard: false,
      backdrop: 'static',
      windowClass: 'modal-profile',
      templateUrl: 'profile/profile.modal.tpl.html',
      controllerName: 'ProfileController'
    }

    return {
      showProfile: function () {
        modalWindowControl.showDialog(dialogParams);
      },

      cancelProfile: function () {
        modalWindowControl.cancelDialog()
      },

      profile: function(model, token){
        var request = $http.put(profileUrl.profileUriPut+token+'/profile', {LastName: model.LastName, Name: model.Name, MiddleName: model.MiddleName, Password: model.Password, PasswordR: model.PasswordR});
        return request.then(function(success) {
          return success;
        }, function(error) {
          return error;
        })
      }
    }
  }])

"use strict";

angular.module('zegota.recovery.config', [])

.value('recoveryUrl', {
  uriRecoveryStepOne: '/users/recovery',
  uriRecoveryStepTwo: '/users/'
})

.config(['$stateProvider', function($stateProvider) {
  $stateProvider
    .state('recoveryStepTwo', {
      url: '/recovery/:hashCode',
      templateUrl: 'recovery/step2.tpl.html',
      controller: 'RecoveryStepTwoController'
    })
}])
;

"use strict";

angular.module('zegota.recovery.controllerStepOne', [])

  .controller('RecoveryStepOneController', ['localizedMessages', 'recoveryUrl', '$scope', 'recovery',
    function(localizedMessages, recoveryUrl, $scope, recovery){

      var response = null;
      $scope.model = {};

      $scope.cancelRecovery = recovery.cancelRecovery;

      $scope.register = function () {
        $scope.registerError = null;
        $scope.registerSuccess = null;

        response = recovery.recovery($scope.model);
        response.then(function(success) {
          $scope.registerSuccess = localizedMessages.get('user.recovery.step1.success');
        }, function(error) { //TODO  that he would return as a result of errors
          console.log('error', error.message);
        })
      };

      $scope.clearForm = function(){
        $scope.model = {};
      };
  }]);

"use strict";

angular.module('zegota.recovery.controllerStepTwo', [])

.controller('RecoveryStepTwoController',
['localizedMessages', 'recoverynUrl', '$scope', '$stateParams', '$http',
  function(localizedMessages, recoveryUrl, $scope, $stateParams, $http) {
    $scope.registerError = null;
    $scope.registerSuccess = null;

    $http.post(recoveryUrl.uriRecoveryStepOne+$stateParams.hashCode, {}).then(function () {
      $scope.registerSuccess = localizedMessages.get('user.recovery.step2.success');
    }, function (response) {
     if (response.data.errorMessage) {
        $scope.registerError = response.data.errorMessage;
      }
    });
}])
;

"use strict";

angular.module('zegota.recovery.controllers', [
    'zegota.recovery.controllerStepOne', 'zegota.recovery.controllerStepTwo'
]);

"use strict";

angular.module('zegota.recovery', [
  'restangular',
  'zegota.recovery.config',
  'zegota.recovery.directive',
  'zegota.recovery.controllers',
  'zegota.recovery.service'
]);

"use strict";

angular.module('zegota.recovery.directive', [])

.directive('recoveryToolbar', ['recovery', 'security', function(recovery, security) {
  return {
    templateUrl: 'recovery/recovery.toolbar.tpl.html',
    restrict: 'E',
    replace: true,
    scope: true,
    link: function ($scope) {
      $scope.isAuthenticated = security.isAuthenticated;
      $scope.recovery = recovery.showRecovery;
    }
  }
}]);

"use strict";

angular.module('zegota.recovery.service', [])

.factory('recovery', ['$http', '$modal', 'recoveryUrl', '$window', 'modalWindowControl',
  function($http, $modal, recoveryUrl, $window, mWC){

    var dialogParams = {
      keyboard: false,
      backdrop: 'static',
      windowClass: 'modal-recovery',
      templateUrl: 'recovery/step1.modal.tpl.html',
      controllerName: 'RecoveryStepOneController'
    }

    return {
      showRecovery: function () {
        mWC.showDialog(dialogParams);
      },

      cancelRecovery: function () {
        mWC.cancelDialog()
      },

      recovery: function(model){
        var request = $http.post(recoveryUrl.uriRecoveryStepOne, {Email: model.email, Hash: model.hash});
        return request.then(function(success) {
          return success;
        }, function(error) {
          return error;
        })
      }
    }
  }]);

"use strict";

angular.module('zegota.registration.config', [])

.value('registrationUrl', {
    uriRegistration: '/api/v1.0/session/registration'
})
;

"use strict";

angular.module('zegota.registration', [
  'zegota.registration.config',
  'zegota.registration.service',
  'zegota.registration.directive',
  'zegota.registration.controller'
]);





"use strict";

angular.module('zegota.registration.controller', [])

.controller('RegistrationController', ['localizedMessages', 'registrationUrl', '$scope', 'registration', 'security',
  function(localizedMessages, registrationUrl, $scope, registration, security){

  var response = null;
  var registrationError = null;
  $scope.model = {};
  $scope.showCaptcha = false;
  $scope.captcha = null;
  $scope.captchaHash = null;
  $scope.authReason = null;
  $scope.registerError = null;
  $scope.registerSuccess = null;

  function getCaptcha() {
    security.getCaptcha().then(
      function (captchaObj) {
        $scope.captcha = captchaObj.captchaImage;
        $scope.captchaHash = captchaObj.captchaHash;
      },
      function (exception) {
        console.log("captcha error");
      }
    );
  }

  getCaptcha();

  $scope.reloadCaptcha = function () {
    getCaptcha();
  };

  $scope.cancelRegistration = registration.cancelRegistration;

  $scope.register = function () {
    registration.registration($scope.model).then(
      function (success) {
        if (!success) {
          $scope.registerError= localizedMessages.get('user.register.step1.fail');
        }
      },
      function (error) {
        registrationError = error.data;
        $scope.reloadCaptcha();
        $scope.showCaptcha = true;
        $scope.model.captcha = {};
        $scope.registerError = (registrationError)?(registrationError):(localizedMessages.get('user.register.step1.fail'));
      }
    )
  };

  $scope.clearForm = function(){
    $scope.model = {};
  };
}]);

"use strict";

angular.module('zegota.registration.directive', [])

.directive('registrationToolbar', ['registration', 'security', function(registration, security) {
  return {
    templateUrl: 'registration/registration.toolbar.tpl.html',
    restrict: 'E',
    replace: true,
    scope: true,
    link: function ($scope) {
      $scope.isAuthenticated = security.isAuthenticated;
      $scope.registration = registration.showRegistration;
    }
  }
}]);

"use strict";

angular.module('zegota.registration.service', [])

.factory('registration', ['$http', '$modal', 'registrationUrl', '$window', 'modalWindowControl', 'security',
  function($http, $modal, registrationUrl, $window, mWC, security){

    var dialogParams = {
          keyboard: false,
          backdrop:'static',
          windowClass: "modal-registration",
          templateUrl: 'registration/registration.modal.tpl.html',
          controllerName: 'RegistrationController'
    }

    var registrationError = null;

    var factory = {
      showRegistration: function () {
        mWC.showDialog(dialogParams);
      },

      cancelRegistration: function () {
        mWC.cancelDialog();
      },

      registration: function(model){
        var request = $http.post(registrationUrl.uriRegistration, {Email: model.Email, CaptchaHash: model.captchaHash, CaptchaValue: model.captchaValue });
        console.log("reg request is: ", request);
        return request.then(function (response) {
          if ( typeof response.data === 'object' ) {
            if ( typeof response.data.ErrorCode === 'number' ) {
              if ( response.data.ErrorCode === 0 ) {
                security.userToken = response.data.Token;
                security.setCookieToken();
                security.currentUser = security.requestCurrentUser();
              } else {
                factory.registrationError = response.data;
              }
              if (security.isAuthenticated()) {
                mWC.cancelDialog(true); //check this function
              } else {
                mWC.cancelDialog(false); //check this function
              }
              return security.isAuthenticated();
            }
          }
        });
      }
    }
    return factory;
  }]);

angular.module('zegota.bootstrap', [

	'ui.bootstrap',
	'templates.bootstrap' //customize bootstrap templates

]);
angular.module('zegota.common.config', [])
  .constant('ZEGOTA_CONFIG', {
    api: {
      baseUrl: '/api/v1'
    },

    ngTableParams: {
      page: 1,             // show first page
      count: 50,           // count per page
      sorting: {
        Position: 'asc'  // initial sorting
      }
    }
  })
  .constant('I18N.MESSAGES', {
    'captcha.error.serverError': 'Ошибка севрера. Не могу получить Captcha.',
    'login.error.serverError': "Ошибка авторизации",
    'login.reason.notAuthorized': "Ошибка авторизации",
    'login.reason.notAuthenticated': "Ошибка аутентификации",
    'login.error.invalidCredentials': "Неправильные логин/пароль",
    'user.register.step1.success': "Регистрация прошла успешно. Проверьте почтовый ящик.",
    'user.register.step1.fail': "Во время регистрации произошла ошибка.",
    'user.register.step2.success': "Вы успешно зарегистрированы.",
    'user.recovery.step1.success': "Вам выслано письмо. Проверьте почтовый ящик.",
    'user.recovery.step2.success': "Вы успешно восстановили пароль.",
    'modal.error': "modal.error",
    'modal.notify': "modal.notify"
  })
  ;

angular.module('zegota.core', [
      'ui.router',
      'zegota.common.config',
      'templates.common',
      'common.services',
      'common.security'
    ])

    .config(['ZEGOTA_CONFIG', '$logProvider', '$locationProvider',
        function (ZEGOTA_CONFIG, $logProvider, $locationProvider) {
            $locationProvider.html5Mode(true).hashPrefix('!');
            $logProvider.debugEnabled(true);
        }])

    .run(['$rootScope', '$state', '$stateParams', '$location', 'authUrl', 'security',
        function ($rootScope, $state, $stateParams, $location, authUrl, security) {

            security.getCookieToken().then(
              function(data) {
                if (data[0] !== undefined) {
                  security.userToken = data;
                  security.checkUser();
                };
              }
            );

            $rootScope.$state = $state;
            $rootScope.$stateParams = $stateParams;

            $rootScope.baseUrl = $location.absUrl();

            // @see https://github.com/angular-ui/ui-router/issues/110#issuecomment-18348811
            // scrollTop
            $rootScope.$on('$stateChangeSuccess', function (event, toState, toParams, fromState, fromParams) {
                if (toState.scrollTop !== false) {
                    $("html, body").animate({ scrollTop: 0 }, 200);
                }
            });
        }])

    .controller('AppCtrl', ['$state', '$scope', '$rootScope', 'security', 'i18nNotifications',
        function ($state, $scope, $rootScope, security, i18nNotifications) {
            $scope.notifications = i18nNotifications;
            $scope.removeNotification = function (notification) {
                i18nNotifications.remove(notification);
            };

            $scope.isAuthenticated = security.isAuthenticated;
            $scope.isAdmin = security.isAdmin;
        }])
;

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

angular.module('security.authorization', ['security.service'])

	.provider('securityAuthorization', {
		requireAdminUser: ['securityAuthorization', function (securityAuthorization) {
			return securityAuthorization.requireAdminUser();
		}],

		requireAuthenticatedUser: ['securityAuthorization', function (securityAuthorization) {
			return securityAuthorization.requireAuthenticatedUser();
		}],

		$get: ['security', 'securityRetryQueue', function (security, queue) {
			var service = {

				requireAuthenticatedUser: function () {
					var promise = security.requestCurrentUser().then(function (userInfo) {
						if (!security.isAuthenticated()) {
							return queue.pushRetryFn('unauthenticated-client', service.requireAuthenticatedUser);
						}
					});
					return promise;
				},

				requireAdminUser: function () {
					var promise = security.requestCurrentUser().then(function (userInfo) {
						if (!security.isAdmin()) {
							return queue.pushRetryFn('unauthorized-client', service.requireAdminUser);
						}
					});
					return promise;
				}

			};

			return service;
		}]
	})
;

angular.module('common.security', [
  'security.retryQueue',
  'security.service',
  'security.interceptor',
  'security.login',
  'security.authorization'
]);

angular.module('security.interceptor', ['security.retryQueue'])

	.factory('securityInterceptor', ['$injector', 'securityRetryQueue', function ($injector, queue) {
		return function (promise) {
			return promise.then(null, function (originalResponse) {
				if (originalResponse.status === 401) {
					promise = queue.pushRetryFn('unauthorized-server', function retryRequest() {
						return $injector.get('$http')(originalResponse.config);
					});
				}
				return promise;
			});
		};
	}])

	.config(['$httpProvider', function ($httpProvider) {
		$httpProvider.responseInterceptors.push('securityInterceptor');
	}])
;

angular.module('security.login.form', [])

.controller('LoginFormController', ['$scope', 'security', 'localizedMessages', '$window', '$location', '$http', 'authUrl',
  function ($scope, security, localizedMessages, $window, $location, $http, authUrl) {

  var loginError = null;
  $scope.user = {};
  $scope.authError = null;
  $scope.showCaptcha = false;
  $scope.captcha = undefined;
  $scope.authReason = null;
  $scope.authError = null;

  function getCaptcha() {
    security.getCaptcha().then(function (response) {
      $scope.captcha = "data:image/png;base64,";
      $scope.captcha += response.data.CaptchaImage;
    }, function (exception) {
      console.log("captcha error : ", exception);
    });
  }

  getCaptcha();

  $scope.reloadCaptcha = function () {
    getCaptcha();
  };

  if (security.getLoginReason()) {
    $scope.authReason = ( security.isAuthenticated() ) ?
      localizedMessages.get('login.reason.notAuthorized') :
      localizedMessages.get('login.reason.notAuthenticated');
  }

  $scope.login = function () {
    security.login($scope.user.login, $scope.user.password, $scope.user.remember, $scope.user.captcha).then(
      function (loggedIn) {
        if (!loggedIn) {
          $scope.authError = localizedMessages.get('login.error.serverError');
        }
      },
      function (exception) {
        loginError = exception.data;
        $scope.reloadCaptcha();
        $scope.showCaptcha = true;
        $scope.user.captcha = null;
        $scope.authError = (loginError.errorMessage) ? (loginError.errorMessage) : localizedMessages.get('login.error.invalidCredentials');
      }
    );
  };

  $scope.clearForm = function () {
    $scope.user = {};
  };

  $scope.cancelLogin = security.cancelLogin;
}])
;

angular.module('security.login', ['security.login.form', 'security.login.toolbar']);

angular.module('security.login.toolbar', [])

	.directive('loginToolbar', ['security', function (security) {
		var directive = {
			templateUrl: 'security/login/toolbar.tpl.html',
			restrict: 'E',
			replace: true,
			scope: true,
			link: function ($scope, $element, $attrs, $controller) {
				$scope.isAuthenticated = security.isAuthenticated;
				$scope.login = security.showLogin;
				$scope.logout = security.logout;
				$scope.$watch(function () {
					return security.currentUser;
				}, function (currentUser) {
					$scope.currentUser = currentUser;
				});
			}
		};
		return directive;
	}])
;

angular.module('security.retryQueue', [])

	.factory('securityRetryQueue', ['$q', '$log', function ($q, $log) {
		var retryQueue = [];
		var service = {
			onItemAddedCallbacks: [],

			hasMore: function () {
				return retryQueue.length > 0;
			},
			push: function (retryItem) {
				retryQueue.push(retryItem);
				angular.forEach(service.onItemAddedCallbacks, function (cb) {
					try {
						cb(retryItem);
					} catch (e) {
						$log.error('securityRetryQueue.push(retryItem): callback threw an error' + e);
					}
				});
			},
			pushRetryFn: function (reason, retryFn) {
				if (arguments.length === 1) {
					retryFn = reason;
					reason = undefined;
				}

				var deferred = $q.defer();
				var retryItem = {
					reason: reason,
					retry: function () {
						$q.when(retryFn()).then(function (value) {
							deferred.resolve(value);
						}, function (value) {
							deferred.reject(value);
						});
					},
					cancel: function () {
						deferred.reject();
					}
				};
				service.push(retryItem);
				return deferred.promise;
			},
			retryReason: function () {
				return service.hasMore() && retryQueue[0].reason;
			},
			cancelAll: function () {
				while (service.hasMore()) {
					retryQueue.shift().cancel();
				}
			},
			retryAll: function () {
				while (service.hasMore()) {
					retryQueue.shift().retry();
				}
			}
		};
		return service;
	}])
;

angular.module('security.service', [
    'ngCookies',
		'security.retryQueue',
		'security.login',
    'zegota.bootstrap',
    'ui.bootstrap.modal',
    'dialogs.controllers',
    'common.services.modalWindowControl'
	])

	.value('authUrl', {
      cookieName:        'ZEGOTACOOKIE',
      uriCheckCookie:    '/api/v1.0/session/authorization',
      uriGetCurrentUser: '/api/v1.0/session',
      uriLogIn:          '/api/v1.0/session/authorization',
      uriLogOut:         '/api/v1.0/session/authorization',
      uriCaptcha:        '/api/v1.0/session/captcha/native'
	})

	.factory('security', ['$window', '$state', '$http', '$q', '$location', 'securityRetryQueue', '$modal', 'authUrl', '$cookieStore', 'modalWindowControl',
		function ($window, $state, $http, $q, $location, queue, $modal, authUrl, $cookieStore, mWC) {


      var dialogParams = {
        keyboard: false,
        backdrop: 'static',
        windowClass: "modal-login",
        templateUrl: 'security/login/form.tpl.html',
        controllerName: 'LoginFormController'
      };

			queue.onItemAddedCallbacks.push(function (retryItem) {
				if (queue.hasMore()) {
					service.showLogin();
				}
			});

			var service = {
        // Токен авторизации пользователя
        // Если пусто - пользователь не авторизован
        // Если не пусто, проверяется на сервере
        userToken: '',

        // Объект с информацией о текущем пользователе
        currentUser: null,

        // Причина ошибки авторизации
        loginError: null,

        captcha: null,
        captchaHash: null,

        getLoginReason: function () {
					return queue.retryReason();
				},

				showLogin: function () {
					mWC.showDialog(dialogParams);
				},

				cancelLogin: function (redirectTo) { //check this param
					mWC.cancelDialog();
				},

//      Процесс входа в систему
				login: function (login, password, remember, captcha) {
  				var request = $http.post(authUrl.uriLogIn, {Login: login, Password: password, Remember: remember, Captcha: captcha});
					return request.then(function (response) {
            if ( typeof response.data === 'object' ) {
              if ( typeof response.data.ErrorCode === 'number' ) {
                if ( response.data.ErrorCode === 0 ) {
                  service.userToken = response.data.Token;
                  service.setCookieToken();
                  service.currentUser = service.requestCurrentUser();
                  console.log("current user is: ",service.currentUser);
                } else {
                  service.loginError = response.data;
                }
                if (service.isAuthenticated()) {
                  mWC.cancelDialog(true); //check this function
                } else {
                  mWC.cancelDialog(false); //check this function
                }
                return service.isAuthenticated();
              }
            }
					});
				},

				logout: function (redirectTo) {
          var uri = authUrl.uriLogOut + '/' + service.userToken;
          $http.delete(uri).then(function () {
						service.currentUser = null;
            service.userToken = '';
            $cookieStore.remove(authUrl.cookieName);
					});
				},

				requestCurrentUser: function () {
          var resp;
					if (service.isAuthenticated()) {
            resp = $q.when(service.currentUser);
					} else {
            resp = service.getCookieToken().then(function(token){
              service.userToken = token;
              service.setCookieToken();
              //service.checkUser();
            });
					}
          return resp;
				},

  //    Проверка авторизован ли пользователь
				isAuthenticated: function () {
          if ( service.userToken === '' || service.currentUser == null ) {
            return false;
          } else {
            return true
          }
				},

				isAdmin: function () {
					var isAdmin = false;
					angular.forEach(service.currentAccessGroups, function (accessGroup) {
						if (accessGroup.Id === 2 || accessGroup.Id === 3) {
							isAdmin = true;
						}
					});
					return isAdmin;
				},

//      Получение проверенного токена из Cookie
//      Если Cookie нет, или токен не верный вернётся пустая строка
				getCookieToken: function() {
          var cookieValue = $cookieStore.get(authUrl.cookieName);
          // Если в куках токена нет, пользователь не авторизован
          if ( cookieValue === undefined ) {
            $cookieStore.remove(authUrl.cookieName);
            cookieValue = '';
          }
          var uri = authUrl.uriCheckCookie + '/' + cookieValue;
          // Если в куках есть токен, его надо проверить на актуальность
          var promise = $http.get(uri).then(function (response) {
            var data = response.data;
            var resp = '';
            if ( typeof data === 'object' ) {
              if ( typeof data.ErrorCode === 'number' ) {
                if ( data.ErrorCode === 0 )
                    resp = cookieValue;
              }
            }
            return resp;
          });
          return promise;
				},

        setCookieToken: function() {
          if ( service.userToken !== '' ) {
            $cookieStore.put(authUrl.cookieName, service.userToken);
          }
        },

        checkUser: function() {
          var uri = authUrl.uriGetCurrentUser + '/' + service.userToken;
          return $http.get(uri).then(function (response) {
            if ( typeof response.data === 'object' ) {
              if ( response.data.ErrorCode !== 0 ) {
                service.currentUser = response.data;
              }
              return service.currentUser;
            }
          });
        },

        getCaptcha: function() {
          var uri = authUrl.uriCaptcha;
          var request = $http.get(uri);
          return request.then(
            function (response) {
              if ( typeof response === 'object' ) {
                if ( response.ErrorCode === 0 ) {
                  console.log(response.data);
                  service.captcha = response.CaptchaImage;
                  servive.captchaHash = response.CaptchaHash;
                  return { captchaImage: service.captcha, captchaHash: service.captchaHash };
                } else {
                  return response;
                }
              }
            }
          );
        }
			};

			return service;
		}]);

angular.module('common.services.i18nNotifications', [
  'common.services.localizedMessages',
  'common.services.notifications',
  'dialogs'
])
.factory('i18nNotifications', ['localizedMessages', 'notifications', '$dialogs', function (localizedMessages, notifications, $dialogs) {

	var prepareNotification = function (msgKey, type, interpolateParams, otherProperties) {
		return angular.extend({
			message: localizedMessages.get(msgKey, interpolateParams),
			type: type
		}, otherProperties);
	};

	var I18nNotifications = {
		modal: function (msgKey, type, interpolateParams, otherProperties) {
			var notify = prepareNotification(msgKey, type, interpolateParams, otherProperties);
			switch (notify.type) {
				case 'error':
					$dialogs.error(localizedMessages.get("modal.error"), notify.message);
					break;
				case 'notify':
					$dialogs.notify(localizedMessages.get("modal.notify"), notify.message);
					break;
			}
		},
		pushSticky: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushSticky(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		pushForCurrentRoute: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushForCurrentRoute(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		pushForNextRoute: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushForNextRoute(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		getCurrent: function () {
			return notifications.getCurrent();
		},
		remove: function (notification) {
			return notifications.remove(notification);
		}
	};

	return I18nNotifications;
}]);

angular.module('common.services', [
  'common.services.modalWindowControl',
  'common.services.notifications',
  'common.services.localizedMessages',
  'common.services.i18nNotifications',
]);

angular.module('common.services.localizedMessages', [])
.factory('localizedMessages', ['$interpolate', 'I18N.MESSAGES', function ($interpolate, i18nmessages) {

  var handleNotFound = function (msg, msgKey) {
    return msg || '?' + msgKey + '?';
  };

  return {
    get: function (msgKey, interpolateParams) {
      var msg = i18nmessages[msgKey];
      if (msg) {
        return $interpolate(msg)(interpolateParams);
      } else {
        return handleNotFound(msg, msgKey);
      }
    }
  };
}]);

angular.module('common.services.modalWindowControl', [])
.factory('modalWindowControl', ['$modal', '$window', '$location', function($modal, $window, $location) {

  var dialog = null;

  function redirect(url) {
    url = url || '/';
    if ($window.location.pathname !== url) {
      $window.location = url;
    }
  }

  function openDialog(dialogParams) {
    if (dialog) {
      throw new Error('Trying to open a dialog that is already open!');
    }
    if (!dialog) {
      dialog = $modal.open({
        keyboard: dialogParams.keyboard,
        backdrop: dialogParams.backdrop,
        windowClass: dialogParams.windowClass,
        templateUrl: dialogParams.templateUrl,
        controller: dialogParams.controllerName
      });
      dialog.result.then(onDialogClose);
    }
  }

  function closeDialog(success, redirectTo) { //check input params
    if (dialog) {
      dialog.close({success: success, redirectTo: redirectTo });
    }
  }

  function onDialogClose(result) {
    dialog = null;
    if (typeof queue !== "undefined") {
      if (result.success) {
        queue.retryAll();
      } else {
        queue.cancelAll();
        redirect(result.redirectTo);
      }
    }
  }

  return {
    showDialog: function (dialogParams) {
      openDialog(dialogParams);
    },

    cancelDialog: function () {
      closeDialog();
    }
  }

}]);

angular.module('common.services.notifications', [])
.factory('notifications', ['$rootScope', function ($rootScope) {

	var notifications = {
		'STICKY': [],
		'ROUTE_CURRENT': [],
		'ROUTE_NEXT': []
	};
	var notificationsService = {};

	var addNotification = function (notificationsArray, notificationObj) {
		if (!angular.isObject(notificationObj)) {
			throw new Error("Only object can be added to the notification service");
		}
		notificationsArray.push(notificationObj);
		return notificationObj;
	};

	$rootScope.$on('$routeChangeSuccess', function () {
		notifications.ROUTE_CURRENT.length = 0;

		notifications.ROUTE_CURRENT = angular.copy(notifications.ROUTE_NEXT);
		notifications.ROUTE_NEXT.length = 0;
	});

	notificationsService.getCurrent = function () {
		return [].concat(notifications.STICKY, notifications.ROUTE_CURRENT);
	};

	notificationsService.pushSticky = function (notification) {
		return addNotification(notifications.STICKY, notification);
	};

	notificationsService.pushForCurrentRoute = function (notification) {
		return addNotification(notifications.ROUTE_CURRENT, notification);
	};

	notificationsService.pushForNextRoute = function (notification) {
		return addNotification(notifications.ROUTE_NEXT, notification);
	};

	notificationsService.remove = function (notification) {
		angular.forEach(notifications, function (notificationsByType) {
			var idx = notificationsByType.indexOf(notification);
			if (idx > -1) {
				notificationsByType.splice(idx, 1);
			}
		});
	};

	notificationsService.removeAll = function () {
		angular.forEach(notifications, function (notificationsByType) {
			notificationsByType.length = 0;
		});
	};

	return notificationsService;
}]);

angular.module('zegota.ui', [
    'zegota.bootstrap',
    //'ngTable',
    //'ui.select2',
    //'angular-redactor',
    //'ui.sortable',
    //'angularFileUpload',
    //'dialogs',
    //'ngAnimate',
    //'ui.slider',
    //'xeditable',
    //'ui.mask'
]);

angular.module('templates.user', ['profile/profile.modal.tpl.html', 'profile/profile.toolbar.tpl.html', 'recovery/recovery.toolbar.tpl.html', 'recovery/step1.modal.tpl.html', 'recovery/step2.tpl.html', 'registration/registration.modal.tpl.html', 'registration/registration.toolbar.tpl.html']);

angular.module("profile/profile.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("profile/profile.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelProfile()\">&times;</button>\n" +
    "        <h3>Профиль</h3>\n" +
    "      </div>\n" +
    "      <form novalidate name=\"profileForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"profileError\">\n" +
    "            {[{profileError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"profileSuccess\">\n" +
    "            {[{profileSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <fieldset>\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfName\">Имя</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfName\" name=\"inputProfName\" class=\"form-control\" type=\"text\" ng-model=\"model.Name\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfMiddleName\">Отчество</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfMiddleName\" name=\"inputProfMiddleName\" class=\"form-control\" type=\"text\" ng-model=\"model.MiddleName\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfLastName\">Фамилия</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfLastName\" name=\"inputProfLastName\" class=\"form-control\" type=\"text\" ng-model=\"model.LastName\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfPassword\">Пароль</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfPassword\" name=\"inputProfPassword\" class=\"form-control\" type=\"password\" ng-model=\"model.Password\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputRegLastName\">Подтверждение</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfPasswordR\" name=\"inputProfPasswordR\" class=\"form-control\" type=\"password\" ng-model=\"model.PasswordR\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <button class=\"btn btn-primary btn-block\" ng-click=\"submitUserData()\" ng-disabled=\"profileFormForm.$invalid\" type=\"submit\">Соxранить</button>\n" +
    "            </div>\n" +
    "          </fieldset>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("profile/profile.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("profile/profile.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-show=\"isAuthenticated()\" class=\"profile\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"profile()\">Профиль</a>\n" +
    "  </li>\n" +
    "</ul>\n" +
    "");
}]);

angular.module("recovery/recovery.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("recovery/recovery.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-hide=\"isAuthenticated()\" class=\"recovery\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"recovery()\">Восстановление</a>\n" +
    "  </li>\n" +
    "</ul>\n" +
    "");
}]);

angular.module("recovery/step1.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("recovery/step1.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelRecovery()\">&times;</button>\n" +
    "        <h3>Восстановление пароля</h3>\n" +
    "      </div>\n" +
    "      <form novalidate=\"true\" name=\"recoveryForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"registerError\">\n" +
    "            {[{registerError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"registerSuccess\">\n" +
    "            {[{registerSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div ng-hide=\"registerSuccess\">\n" +
    "              <fieldset>\n" +
    "                <div class=\"form-group\">\n" +
    "                  <label class=\"control-label\" for=\"inputRegEmail\">Email</label>\n" +
    "                  <input autocomplete=\"off\" id=\"inputRegEmail\" name=\"inputRegEmail\" class=\"form-control\" type=\"email\" ng-model=\"model.Email\" required=\"true\" autofocus />\n" +
    "                </div>\n" +
    "\n" +
    "                <div class=\"form-group\">\n" +
    "                  <button class=\"btn btn-primary btn-block\" ng-click=\"recovery()\" ng-disabled=\"recoveryForm.$invalid\" type=\"submit\">Восстановить пароль</button>\n" +
    "                  <button class=\"btn btn-default btn-block\" ng-click=\"clearForm()\" type=\"button\">Очистить форму</button>\n" +
    "                </div>\n" +
    "            </fieldset>\n" +
    "          </div>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("recovery/step2.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("recovery/step2.tpl.html",
    "<div ng-if=\"registerSuccess\">\n" +
    "  <p>{[{ registerSuccess }]}</p>\n" +
    "</div>\n" +
    "<div ng-if=\"registerError\">\n" +
    "  <p>{[{ registerError }]}</p>\n" +
    "</div>\n" +
    "");
}]);

angular.module("registration/registration.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("registration/registration.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelRegistration()\">&times;</button>\n" +
    "        <h3>Регистрация</h3>\n" +
    "      </div>\n" +
    "      <form novalidate name=\"registrationForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"registerError\">\n" +
    "            {[{registerError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"registerSuccess\">\n" +
    "            {[{registerSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div ng-hide=\"registerSuccess\">\n" +
    "            <fieldset>\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegEmail\">Email</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegEmail\" name=\"inputRegEmail\" class=\"form-control\" type=\"email\" ng-model=\"model.Email\" required=\"true\" autofocus>\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\" ng-show=\"showCaptcha && captcha\">\n" +
    "                <img class=\"imgCaptcha\" ng-src=\"{[{captcha}]}\" ng-click=\"reloadCaptcha()\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div  class=\"form-group\" ng-if=\"showCaptcha && captcha\">\n" +
    "                <label  class=\"control-label\" for=\"inputCaptcha\">Введите код с картинки</label>\n" +
    "                <input autocomplete=\"off\" class=\"form-control\" id=\"inputCaptcha\" name=\"inputCaptcha\" type=\"text\" ng-model=\"model.captcha\" placeholder=\"Ведите код с картинки\" required=\"true\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <button class=\"btn btn-primary btn-block\" ng-click=\"register()\" ng-disabled=\"registrationForm.$invalid\" type=\"submit\">Зарегестрироваться</button>\n" +
    "                <button class=\"btn btn-default btn-block\" ng-click=\"clearForm()\" type=\"button\">Очистить форму</button>\n" +
    "              </div>\n" +
    "            </fieldset>\n" +
    "          </div>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("registration/registration.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("registration/registration.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-hide=\"isAuthenticated()\" class=\"registration\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"registration()\">Регистрация</a>\n" +
    "  </li>\n" +
    "</ul>");
}]);

angular.module('templates.bootstrap', ['template/accordion/accordion-group.html', 'template/accordion/accordion.html', 'template/alert/alert.html', 'template/carousel/carousel.html', 'template/carousel/slide.html', 'template/datepicker/datepicker.html', 'template/datepicker/day.html', 'template/datepicker/month.html', 'template/datepicker/popup.html', 'template/datepicker/year.html', 'template/dialog/message.html', 'template/modal/backdrop.html', 'template/modal/window.html', 'template/pagination/pager.html', 'template/pagination/pagination.html', 'template/popover/popover.html', 'template/progressbar/bar.html', 'template/progressbar/progress.html', 'template/progressbar/progressbar.html', 'template/rating/rating.html', 'template/tabs/tab.html', 'template/tabs/tabset-titles.html', 'template/tabs/tabset.html', 'template/timepicker/timepicker.html', 'template/tooltip/tooltip-html-unsafe-popup.html', 'template/tooltip/tooltip-popup.html', 'template/typeahead/typeahead-match.html', 'template/typeahead/typeahead-popup.html']);

angular.module("template/accordion/accordion-group.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/accordion/accordion-group.html",
    "<div class=\"accordion-group\">\n" +
    "  <div class=\"accordion-heading\" ><a class=\"accordion-toggle\" ng-click=\"isOpen = !isOpen\" accordion-transclude=\"heading\">{{heading}}</a></div>\n" +
    "  <div class=\"accordion-body\" collapse=\"!isOpen\">\n" +
    "    <div class=\"accordion-inner\" ng-transclude></div>  </div>\n" +
    "</div>");
}]);

angular.module("template/accordion/accordion.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/accordion/accordion.html",
    "<div class=\"accordion\" ng-transclude></div>");
}]);

angular.module("template/alert/alert.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/alert/alert.html",
    "<div class='alert' ng-class='type && \"alert-\" + type'>\n" +
    "    <button ng-show='closeable' type='button' class='close' ng-click='close()'>&times;</button>\n" +
    "    <div ng-transclude></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/carousel/carousel.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/carousel/carousel.html",
    "<div ng-mouseenter=\"pause()\" ng-mouseleave=\"play()\" class=\"carousel\">\n" +
    "    <ol class=\"carousel-indicators\" ng-show=\"slides().length > 1\">\n" +
    "        <li ng-repeat=\"slide in slides()\" ng-class=\"{active: isActive(slide)}\" ng-click=\"select(slide)\"></li>\n" +
    "    </ol>\n" +
    "    <div class=\"carousel-inner\" ng-transclude></div>\n" +
    "    <a ng-click=\"prev()\" class=\"carousel-control left\" ng-show=\"slides().length > 1\">&lsaquo;</a>\n" +
    "    <a ng-click=\"next()\" class=\"carousel-control right\" ng-show=\"slides().length > 1\">&rsaquo;</a>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/carousel/slide.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/carousel/slide.html",
    "<div ng-class=\"{\n" +
    "    'active': leaving || (active && !entering),\n" +
    "    'prev': (next || active) && direction=='prev',\n" +
    "    'next': (next || active) && direction=='next',\n" +
    "    'right': direction=='prev',\n" +
    "    'left': direction=='next'\n" +
    "  }\" class=\"item\" ng-transclude></div>\n" +
    "");
}]);

angular.module("template/datepicker/datepicker.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/datepicker.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"pull-left\" ng-click=\"move(-1)\"><span class=\"nav-arrow nav-arrow_left\"></span></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"{{rows[0].length - 2 + showWeekNumbers}}\">\n" +
    "			<button type=\"button\" class=\"\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"pull-right\" ng-click=\"move(1)\"><span class=\"nav-arrow nav-arrow_right\"></span></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	<tr ng-show=\"labels.length > 0\" class=\"h6\">\n" +
    "		<th ng-show=\"showWeekNumbers\" class=\"text-center\">#</th>\n" +
    "		<th ng-repeat=\"label in labels\" class=\"text-center\">{{label}}</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows\">\n" +
    "		<td ng-show=\"showWeekNumbers\" class=\"text-center\"><em>{{ getWeekNumber(row) }}</em></td>\n" +
    "		<td ng-repeat=\"dt in row\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"\"\n" +
    "			        ng-class=\"{'btn-info': dt.selected}\" ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span\n" +
    "					ng-class=\"{'text-muted': dt.secondary}\">{{dt.label}}</span></button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/day.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/day.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"{{5 + showWeeks}}\">\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	<tr>\n" +
    "		<th ng-show=\"showWeeks\" class=\"text-center\"></th>\n" +
    "		<th ng-repeat=\"label in labels track by $index\" class=\"text-center\">\n" +
    "			<small>{{label}}</small>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-show=\"showWeeks\" class=\"text-center\">\n" +
    "			<small><em>{{ weekNumbers[$index] }}</em></small>\n" +
    "		</td>\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default btn-sm\"\n" +
    "			        ng-class=\"{'btn-info': dt.selected}\" ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span\n" +
    "					ng-class=\"{'text-muted': dt.secondary, 'text-info': dt.current}\">{{dt.label}}</span></button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/month.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/month.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default\" ng-class=\"{'btn-info': dt.selected}\"\n" +
    "			        ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span ng-class=\"{'text-info': dt.current}\">{{dt.label}}</span>\n" +
    "			</button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/popup.html",
    "<ul class=\"dropdown-menu dropdown-datepicker\"\n" +
    "    ng-style=\"{display: (isOpen && 'block') || 'none', top: position.top+'px', left: position.left+'px'}\">\n" +
    "	<li ng-transclude></li>\n" +
    "\n" +
    "	<!--<li ng-show=\"showButtonBar\" style=\"padding:10px 9px 2px\">-->\n" +
    "		<!--<span class=\"btn-group\">-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-info\" ng-click=\"today()\">{{currentText}}</button>-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-default\" ng-click=\"showWeeks = ! showWeeks\"-->\n" +
    "			        <!--ng-class=\"{active: showWeeks}\">{{toggleWeeksText}}-->\n" +
    "			<!--</button>-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-danger\" ng-click=\"clear()\">{{clearText}}</button>-->\n" +
    "		<!--</span>-->\n" +
    "		<!--<button type=\"button\" class=\"btn btn-sm btn-success pull-right\" ng-click=\"isOpen = false\">{{closeText}}</button>-->\n" +
    "	<!--</li>-->\n" +
    "</ul>\n" +
    "");
}]);

angular.module("template/datepicker/year.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/year.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"3\">\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default\" ng-class=\"{'btn-info': dt.selected}\"\n" +
    "			        ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span ng-class=\"{'text-info': dt.current}\">{{dt.label}}</span>\n" +
    "			</button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/dialog/message.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/dialog/message.html",
    "<div class=\"modal-header\">\n" +
    "	<h3>{{ title }}</h3>\n" +
    "</div>\n" +
    "<div class=\"modal-body\">\n" +
    "	<p>{{ message }}</p>\n" +
    "</div>\n" +
    "<div class=\"modal-footer\">\n" +
    "	<button ng-repeat=\"btn in buttons\" ng-click=\"close(btn.result)\" class=\"btn\" ng-class=\"btn.cssClass\">{{ btn.label }}</button>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/modal/backdrop.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/modal/backdrop.html",
    "<div class=\"modal-backdrop fade\" ng-class=\"{in: animate}\" ng-style=\"{'z-index': 1040 + index*10}\" ng-click=\"close($event)\"></div>");
}]);

angular.module("template/modal/window.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/modal/window.html",
    "<div tabindex=\"-1\" class=\"modal fade {{ windowClass }}\" ng-class=\"{in: animate}\" ng-style=\"{'z-index': 1050 + index*10}\" ng-transclude></div>");
}]);

angular.module("template/pagination/pager.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/pagination/pager.html",
    "<div class=\"pager\">\n" +
    "  <ul>\n" +
    "    <li ng-repeat=\"page in pages\" ng-class=\"{disabled: page.disabled, previous: page.previous, next: page.next}\"><a ng-click=\"selectPage(page.number)\">{{page.text}}</a></li>\n" +
    "  </ul>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/pagination/pagination.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/pagination/pagination.html",
    "<div class=\"pagination\"><ul>\n" +
    "  <li ng-repeat=\"page in pages\" ng-class=\"{active: page.active, disabled: page.disabled}\"><a ng-click=\"selectPage(page.number)\">{{page.text}}</a></li>\n" +
    "  </ul>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/popover/popover.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/popover/popover.html",
    "<div class=\"popover {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"arrow\"></div>\n" +
    "\n" +
    "  <div class=\"popover-inner\">\n" +
    "      <h3 class=\"popover-title\" ng-bind=\"title\" ng-show=\"title\"></h3>\n" +
    "      <div class=\"popover-content\" ng-bind=\"content\"></div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/progressbar/bar.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/bar.html",
    "<div class=\"bar\" ng-class=\"type && 'bar-' + type\" ng-transclude></div>");
}]);

angular.module("template/progressbar/progress.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/progress.html",
    "<div class=\"progress\" ng-transclude></div>");
}]);

angular.module("template/progressbar/progressbar.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/progressbar.html",
    "<div class=\"progress\"><div class=\"bar\" ng-class=\"type && 'bar-' + type\" ng-transclude></div></div>");
}]);

angular.module("template/rating/rating.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/rating/rating.html",
    "<span ng-mouseleave=\"reset()\">\n" +
    "	<i ng-repeat=\"r in range\" ng-mouseenter=\"enter($index + 1)\" ng-click=\"rate($index + 1)\" ng-class=\"$index < val && (r.stateOn || 'icon-star') || (r.stateOff || 'icon-star-empty')\"></i>\n" +
    "</span>");
}]);

angular.module("template/tabs/tab.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tab.html",
    "<li ng-class=\"{active: active, disabled: disabled}\">\n" +
    "  <a ng-click=\"select()\" tab-heading-transclude>{{heading}}</a>\n" +
    "</li>\n" +
    "");
}]);

angular.module("template/tabs/tabset-titles.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tabset-titles.html",
    "<ul class=\"nav {{type && 'nav-' + type}}\" ng-class=\"{'nav-stacked': vertical}\">\n" +
    "</ul>\n" +
    "");
}]);

angular.module("template/tabs/tabset.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tabset.html",
    "<div class=\"tabbable\">\n" +
    "  <ul class=\"nav {{type && 'nav-' + type}}\" ng-class=\"{'nav-stacked': vertical}\" ng-transclude>\n" +
    "  </ul>\n" +
    "  <div class=\"tab-content\">\n" +
    "    <div class=\"tab-pane\" \n" +
    "         ng-repeat=\"tab in tabs\" \n" +
    "         ng-class=\"{active: tab.active}\"\n" +
    "         tab-content-transclude=\"tab\">\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/timepicker/timepicker.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/timepicker/timepicker.html",
    "<table class=\"form-inline\">\n" +
    "	<tr class=\"text-center\">\n" +
    "		<td><a ng-click=\"incrementHours()\" class=\"btn btn-link\"><i class=\"icon-chevron-up\"></i></a></td>\n" +
    "		<td>&nbsp;</td>\n" +
    "		<td><a ng-click=\"incrementMinutes()\" class=\"btn btn-link\"><i class=\"icon-chevron-up\"></i></a></td>\n" +
    "		<td ng-show=\"showMeridian\"></td>\n" +
    "	</tr>\n" +
    "	<tr>\n" +
    "		<td class=\"control-group\" ng-class=\"{'error': invalidHours}\"><input type=\"text\" ng-model=\"hours\" ng-change=\"updateHours()\" class=\"span1 text-center\" ng-mousewheel=\"incrementHours()\" ng-readonly=\"readonlyInput\" maxlength=\"2\"></td>\n" +
    "		<td>:</td>\n" +
    "		<td class=\"control-group\" ng-class=\"{'error': invalidMinutes}\"><input type=\"text\" ng-model=\"minutes\" ng-change=\"updateMinutes()\" class=\"span1 text-center\" ng-readonly=\"readonlyInput\" maxlength=\"2\"></td>\n" +
    "		<td ng-show=\"showMeridian\"><button type=\"button\" ng-click=\"toggleMeridian()\" class=\"btn text-center\">{{meridian}}</button></td>\n" +
    "	</tr>\n" +
    "	<tr class=\"text-center\">\n" +
    "		<td><a ng-click=\"decrementHours()\" class=\"btn btn-link\"><i class=\"icon-chevron-down\"></i></a></td>\n" +
    "		<td>&nbsp;</td>\n" +
    "		<td><a ng-click=\"decrementMinutes()\" class=\"btn btn-link\"><i class=\"icon-chevron-down\"></i></a></td>\n" +
    "		<td ng-show=\"showMeridian\"></td>\n" +
    "	</tr>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/tooltip/tooltip-html-unsafe-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tooltip/tooltip-html-unsafe-popup.html",
    "<div class=\"tooltip {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"tooltip-arrow\"></div>\n" +
    "  <div class=\"tooltip-inner\" bind-html-unsafe=\"content\"></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/tooltip/tooltip-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tooltip/tooltip-popup.html",
    "<div class=\"tooltip {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"tooltip-arrow\"></div>\n" +
    "  <div class=\"tooltip-inner\" ng-bind=\"content\"></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/typeahead/typeahead-match.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/typeahead/typeahead-match.html",
    "<a tabindex=\"-1\" bind-html-unsafe=\"match.label | typeaheadHighlight:query\"></a>");
}]);

angular.module("template/typeahead/typeahead-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/typeahead/typeahead-popup.html",
    "<ul class=\"typeahead dropdown-menu\" ng-style=\"{display: isOpen()&&'block' || 'none', top: position.top+'px', left: position.left+'px'}\">\n" +
    "    <li ng-repeat=\"match in matches\" ng-class=\"{active: isActive($index) }\" ng-mouseenter=\"selectActive($index)\" ng-click=\"selectMatch($index)\">\n" +
    "        <div typeahead-match index=\"$index\" match=\"match\" query=\"query\" template-url=\"templateUrl\"></div>\n" +
    "    </li>\n" +
    "</ul>");
}]);

angular.module('templates.common', ['directives/datetimepicker/datetimepicker.tpl.html', 'security/login/form.tpl.html', 'security/login/toolbar.tpl.html', 'services/notifications.tpl.html']);

angular.module("directives/datetimepicker/datetimepicker.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("directives/datetimepicker/datetimepicker.tpl.html",
    "<span class=\"block-date-selector\">\n" +
    "    <span class=\"input-append\">\n" +
    "        <input class=\"\" ng-model=\"date\" ng-value=\"date\" name=\"date\" bs-datepicker type=\"text\">\n" +
    "    </span>\n" +
    "    <input type=\"text\" class=\"input-timepicker\" bs-timepicker ng-model=\"time\" ng-value=\"time\" data-autoclose=\"0\" placeholder=\"Время\" >\n" +
    "</span>\n" +
    "");
}]);

angular.module("security/login/form.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("security/login/form.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <form name=\"loginForm\">\n" +
    "        <div class=\"modal-header\">\n" +
    "            <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-hidden=\"true\" ng-click=\"cancelLogin()\">×</button>\n" +
    "            <h4 class=\"modal-title\">Авторизация</h4>\n" +
    "        </div>\n" +
    "        <div class=\"modal-body\">\n" +
    "          <div class=\"alert alert-danger\" ng-show=\"authError\">\n" +
    "            {[{authError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <fieldset>\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputLogin\">Логин</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputLogin\" name=\"inputLogin\" type=\"text\" ng-model=\"user.login\" required=\"true\" placeholder=\"Введите ваш логин\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputPassword\">Пароль</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputPassword\" name=\"inputPassword\" type=\"password\" ng-model=\"user.password\" required=\"true\" placeholder=\"Введите ваш пароль\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputRememberMe\">Запомнить?</label>\n" +
    "              <input id=\"inputRememberMe\" name=\"inputRememberMe\" type=\"checkbox\" ng-model=\"user.remember\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\" ng-show=\"showCaptcha && captcha\">\n" +
    "              <img class=\"imgCaptcha\" src=\"{[{captcha}]}\" ng-click=\"reloadCaptcha()\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div  class=\"form-group\" ng-show=\"showCaptcha && captcha\">\n" +
    "              <label  class=\"control-label\" for=\"inputCaptcha\">Введите код с картинки</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputCaptcha\" name=\"inputCaptcha\" type=\"text\" ng-model=\"user.captcha\" placeholder=\"Ведите код с картинки\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <button type=\"button\" class=\"btn btn-primary btn-block\" ng-disabled=\"loginForm.$invalid\" ng-click=\"login()\">Войти</button>\n" +
    "              <button type=\"button\" class=\"btn btn-default btn-block\" ng-click=\"clearForm()\">Очистить форму</button>\n" +
    "            </div>\n" +
    "          </fieldset>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("security/login/toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("security/login/toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "    <li ng-show=\"isAuthenticated()\">\n" +
    "        <a class=\"block-support__link block-support__link_tech-support\" href>Поддержка</a>\n" +
    "    </li>\n" +
    "\n" +
    "    <li ng-hide=\"isAuthenticated()\" class=\"login\">\n" +
    "        <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"login()\">Войти</a>\n" +
    "    </li>\n" +
    "\n" +
    "    <li ng-show=\"isAuthenticated()\" class=\"dropdown\">\n" +
    "        <a href=\"#\" class=\"dropdown-toggle block-support__link block-support__link_user\" data-toggle=\"dropdown\">\n" +
    "            {{currentUser.LastName}} {{currentUser.Name.substring(0,1)}}.\n" +
    "        </a>\n" +
    "        <ul class=\"dropdown-menu\">\n" +
    "            <li><a ng-click=\"logout()\">Выход</a></li>\n" +
    "        </ul>\n" +
    "    </li>\n" +
    "</ul>");
}]);

angular.module("services/notifications.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("services/notifications.tpl.html",
    "<div class=\"alert-messages\">\n" +
    "  <div class=\"container\">\n" +
    "    <div ng-class=\"['alert', 'alert-'+notification.type]\" ng-repeat=\"notification in notifications.getCurrent()\">\n" +
    "      <button class=\"close\" ng-click=\"removeNotification(notification)\">x</button>\n" +
    "      {{notification.message}}\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);
