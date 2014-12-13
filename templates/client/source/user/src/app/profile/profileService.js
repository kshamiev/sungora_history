angular.module('zegota.profile.service', [])

.factory('profile', ['$http', '$modal', 'profileUrl', '$window', 'modalWindowControl', function($http, $modal, profileUrl, $window, modalWindowControl){

    var dialogParams = {
      keyboard: false,
      backdrop: 'static',
      windowClass: 'modal-profile',
      templateUrl: 'profile/profile.modal.tpl.html',
      controllerName: 'ProfileController'
    }

    return {
      showProfile: function () {
        modalWindowControl.showDialog(dialogParams);
      },

      cancelProfile: function () {
        modalWindowControl.cancelDialog()
      },

      profile: function(model, token){
        var request = $http.put(profileUrl.profileUriPut+token+'/profile', {LastName: model.LastName, Name: model.Name, MiddleName: model.MiddleName, Password: model.Password, PasswordR: model.PasswordR});
        return request.then(function(success) {
          return success;
        }, function(error) {
          return error;
        })
      }
    }
  }])
