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
