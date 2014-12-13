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
