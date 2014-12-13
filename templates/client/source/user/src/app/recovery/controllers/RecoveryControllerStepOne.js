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
