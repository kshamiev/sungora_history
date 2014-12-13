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
