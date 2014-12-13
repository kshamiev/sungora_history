angular.module('security.service', [
    'ngCookies',
		'security.retryQueue',
		'security.login',
    'zegota.bootstrap',
    'ui.bootstrap.modal',
    'dialogs.controllers',
    'common.services.modalWindowControl'
	])

	.value('authUrl', {
      cookieName:        'ZEGOTACOOKIE',
      uriCheckCookie:    '/api/v1.0/session/authorization',
      uriGetCurrentUser: '/api/v1.0/session',
      uriLogIn:          '/api/v1.0/session/authorization',
      uriLogOut:         '/api/v1.0/session/authorization',
      uriCaptcha:        '/api/v1.0/session/captcha/native'
	})

	.factory('security', ['$window', '$state', '$http', '$q', '$location', 'securityRetryQueue', '$modal', 'authUrl', '$cookieStore', 'modalWindowControl',
		function ($window, $state, $http, $q, $location, queue, $modal, authUrl, $cookieStore, mWC) {


      var dialogParams = {
        keyboard: false,
        backdrop: 'static',
        windowClass: "modal-login",
        templateUrl: 'security/login/form.tpl.html',
        controllerName: 'LoginFormController'
      };

			queue.onItemAddedCallbacks.push(function (retryItem) {
				if (queue.hasMore()) {
					service.showLogin();
				}
			});

			var service = {
        // Токен авторизации пользователя
        // Если пусто - пользователь не авторизован
        // Если не пусто, проверяется на сервере
        userToken: '',

        // Объект с информацией о текущем пользователе
        currentUser: null,

        // Причина ошибки авторизации
        loginError: null,

        captcha: null,
        captchaHash: null,

        getLoginReason: function () {
					return queue.retryReason();
				},

				showLogin: function () {
					mWC.showDialog(dialogParams);
				},

				cancelLogin: function (redirectTo) { //check this param
					mWC.cancelDialog();
				},

//      Процесс входа в систему
				login: function (login, password, remember, captcha) {
  				var request = $http.post(authUrl.uriLogIn, {Login: login, Password: password, Remember: remember, Captcha: captcha});
					return request.then(function (response) {
            if ( typeof response.data === 'object' ) {
              if ( typeof response.data.ErrorCode === 'number' ) {
                if ( response.data.ErrorCode === 0 ) {
                  service.userToken = response.data.Token;
                  service.setCookieToken();
                  service.currentUser = service.requestCurrentUser();
                  console.log("current user is: ",service.currentUser);
                } else {
                  service.loginError = response.data;
                }
                if (service.isAuthenticated()) {
                  mWC.cancelDialog(true); //check this function
                } else {
                  mWC.cancelDialog(false); //check this function
                }
                return service.isAuthenticated();
              }
            }
					});
				},

				logout: function (redirectTo) {
          var uri = authUrl.uriLogOut + '/' + service.userToken;
          $http.delete(uri).then(function () {
						service.currentUser = null;
            service.userToken = '';
            $cookieStore.remove(authUrl.cookieName);
					});
				},

				requestCurrentUser: function () {
          var resp;
					if (service.isAuthenticated()) {
            resp = $q.when(service.currentUser);
					} else {
            resp = service.getCookieToken().then(function(token){
              service.userToken = token;
              service.setCookieToken();
              //service.checkUser();
            });
					}
          return resp;
				},

  //    Проверка авторизован ли пользователь
				isAuthenticated: function () {
          if ( service.userToken === '' || service.currentUser == null ) {
            return false;
          } else {
            return true
          }
				},

				isAdmin: function () {
					var isAdmin = false;
					angular.forEach(service.currentAccessGroups, function (accessGroup) {
						if (accessGroup.Id === 2 || accessGroup.Id === 3) {
							isAdmin = true;
						}
					});
					return isAdmin;
				},

//      Получение проверенного токена из Cookie
//      Если Cookie нет, или токен не верный вернётся пустая строка
				getCookieToken: function() {
          var cookieValue = $cookieStore.get(authUrl.cookieName);
          // Если в куках токена нет, пользователь не авторизован
          if ( cookieValue === undefined ) {
            $cookieStore.remove(authUrl.cookieName);
            cookieValue = '';
          }
          var uri = authUrl.uriCheckCookie + '/' + cookieValue;
          // Если в куках есть токен, его надо проверить на актуальность
          var promise = $http.get(uri).then(function (response) {
            var data = response.data;
            var resp = '';
            if ( typeof data === 'object' ) {
              if ( typeof data.ErrorCode === 'number' ) {
                if ( data.ErrorCode === 0 )
                    resp = cookieValue;
              }
            }
            return resp;
          });
          return promise;
				},

        setCookieToken: function() {
          if ( service.userToken !== '' ) {
            $cookieStore.put(authUrl.cookieName, service.userToken);
          }
        },

        checkUser: function() {
          var uri = authUrl.uriGetCurrentUser + '/' + service.userToken;
          return $http.get(uri).then(function (response) {
            if ( typeof response.data === 'object' ) {
              if ( response.data.ErrorCode !== 0 ) {
                service.currentUser = response.data;
              }
              return service.currentUser;
            }
          });
        },

        getCaptcha: function() {
          var uri = authUrl.uriCaptcha;
          var request = $http.get(uri);
          return request.then(
            function (response) {
              if ( typeof response === 'object' ) {
                if ( response.ErrorCode === 0 ) {
                  console.log(response.data);
                  service.captcha = response.CaptchaImage;
                  servive.captchaHash = response.CaptchaHash;
                  return { captchaImage: service.captcha, captchaHash: service.captchaHash };
                } else {
                  return response;
                }
              }
            }
          );
        }
			};

			return service;
		}]);
