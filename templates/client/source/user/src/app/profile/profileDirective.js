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
