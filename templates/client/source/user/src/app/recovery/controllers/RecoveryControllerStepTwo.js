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
