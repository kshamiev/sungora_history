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
