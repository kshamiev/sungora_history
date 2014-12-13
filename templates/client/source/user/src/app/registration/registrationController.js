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
