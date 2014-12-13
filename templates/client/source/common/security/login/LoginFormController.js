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
