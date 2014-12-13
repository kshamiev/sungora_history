"use strict";

angular.module('zegota.profile.controller', [])

.controller('ProfileController', ['$scope', 'security', 'profile', '$http', 'profileUrl',
  function ($scope, security, profile, $http, profileUrl) {

  $scope.model = {};
  var token = security.getCookieToken();

  function getUserData () {
    var request = $http.get(profileUrl.profileGetCurrentUserData);
    request.then(function (response) {
      if (response.status == 200) {
        angular.extend($scope.model, response.data);
      }
    }, function (error) {
      console.log('error is: ', error);
    });
  }

  getUserData ();

  $scope.submitUserData = profile.profile($scope.model, token);
  $scope.cancelProfile = profile.cancelProfile;
}]);
