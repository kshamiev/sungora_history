"use strict";

angular.module('zegota.registration.service', [])

.factory('registration', ['$http', '$modal', 'registrationUrl', '$window', 'modalWindowControl', 'security',
  function($http, $modal, registrationUrl, $window, mWC, security){

    var dialogParams = {
          keyboard: false,
          backdrop:'static',
          windowClass: "modal-registration",
          templateUrl: 'registration/registration.modal.tpl.html',
          controllerName: 'RegistrationController'
    }

    var registrationError = null;

    var factory = {
      showRegistration: function () {
        mWC.showDialog(dialogParams);
      },

      cancelRegistration: function () {
        mWC.cancelDialog();
      },

      registration: function(model){
        var request = $http.post(registrationUrl.uriRegistration, {Email: model.Email, CaptchaHash: model.captchaHash, CaptchaValue: model.captchaValue });
        console.log("reg request is: ", request);
        return request.then(function (response) {
          if ( typeof response.data === 'object' ) {
            if ( typeof response.data.ErrorCode === 'number' ) {
              if ( response.data.ErrorCode === 0 ) {
                security.userToken = response.data.Token;
                security.setCookieToken();
                security.currentUser = security.requestCurrentUser();
              } else {
                factory.registrationError = response.data;
              }
              if (security.isAuthenticated()) {
                mWC.cancelDialog(true); //check this function
              } else {
                mWC.cancelDialog(false); //check this function
              }
              return security.isAuthenticated();
            }
          }
        });
      }
    }
    return factory;
  }]);
