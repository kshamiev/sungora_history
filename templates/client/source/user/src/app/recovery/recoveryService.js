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
