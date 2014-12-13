angular.module('zegota.common.config', [])
  .constant('ZEGOTA_CONFIG', {
    api: {
      baseUrl: '/api/v1'
    },

    ngTableParams: {
      page: 1,             // show first page
      count: 50,           // count per page
      sorting: {
        Position: 'asc'  // initial sorting
      }
    }
  })
  .constant('I18N.MESSAGES', {
    'captcha.error.serverError': 'Ошибка севрера. Не могу получить Captcha.',
    'login.error.serverError': "Ошибка авторизации",
    'login.reason.notAuthorized': "Ошибка авторизации",
    'login.reason.notAuthenticated': "Ошибка аутентификации",
    'login.error.invalidCredentials': "Неправильные логин/пароль",
    'user.register.step1.success': "Регистрация прошла успешно. Проверьте почтовый ящик.",
    'user.register.step1.fail': "Во время регистрации произошла ошибка.",
    'user.register.step2.success': "Вы успешно зарегистрированы.",
    'user.recovery.step1.success': "Вам выслано письмо. Проверьте почтовый ящик.",
    'user.recovery.step2.success': "Вы успешно восстановили пароль.",
    'modal.error': "modal.error",
    'modal.notify': "modal.notify"
  })
  ;
