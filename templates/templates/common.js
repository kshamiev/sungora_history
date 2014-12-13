angular.module('templates.common', ['directives/datetimepicker/datetimepicker.tpl.html', 'security/login/form.tpl.html', 'security/login/toolbar.tpl.html', 'services/notifications.tpl.html']);

angular.module("directives/datetimepicker/datetimepicker.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("directives/datetimepicker/datetimepicker.tpl.html",
    "<span class=\"block-date-selector\">\n" +
    "    <span class=\"input-append\">\n" +
    "        <input class=\"\" ng-model=\"date\" ng-value=\"date\" name=\"date\" bs-datepicker type=\"text\">\n" +
    "    </span>\n" +
    "    <input type=\"text\" class=\"input-timepicker\" bs-timepicker ng-model=\"time\" ng-value=\"time\" data-autoclose=\"0\" placeholder=\"Время\" >\n" +
    "</span>\n" +
    "");
}]);

angular.module("security/login/form.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("security/login/form.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <form name=\"loginForm\">\n" +
    "        <div class=\"modal-header\">\n" +
    "            <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-hidden=\"true\" ng-click=\"cancelLogin()\">×</button>\n" +
    "            <h4 class=\"modal-title\">Авторизация</h4>\n" +
    "        </div>\n" +
    "        <div class=\"modal-body\">\n" +
    "          <div class=\"alert alert-danger\" ng-show=\"authError\">\n" +
    "            {[{authError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <fieldset>\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputLogin\">Логин</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputLogin\" name=\"inputLogin\" type=\"text\" ng-model=\"user.login\" required=\"true\" placeholder=\"Введите ваш логин\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputPassword\">Пароль</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputPassword\" name=\"inputPassword\" type=\"password\" ng-model=\"user.password\" required=\"true\" placeholder=\"Введите ваш пароль\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputRememberMe\">Запомнить?</label>\n" +
    "              <input id=\"inputRememberMe\" name=\"inputRememberMe\" type=\"checkbox\" ng-model=\"user.remember\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\" ng-show=\"showCaptcha && captcha\">\n" +
    "              <img class=\"imgCaptcha\" ng-src=\"{[{captcha}]}\" ng-click=\"reloadCaptcha()\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div  class=\"form-group\" ng-show=\"showCaptcha && captcha\">\n" +
    "              <label  class=\"control-label\" for=\"inputCaptcha\">Введите код с картинки</label>\n" +
    "              <input autocomplete=\"off\" class=\"form-control\" id=\"inputCaptcha\" name=\"inputCaptcha\" type=\"text\" ng-model=\"user.captcha\" placeholder=\"Ведите код с картинки\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <button type=\"button\" class=\"btn btn-primary btn-block\" ng-disabled=\"loginForm.$invalid\" ng-click=\"login()\">Войти</button>\n" +
    "              <button type=\"button\" class=\"btn btn-default btn-block\" ng-click=\"clearForm()\">Очистить форму</button>\n" +
    "            </div>\n" +
    "          </fieldset>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("security/login/toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("security/login/toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "    <li ng-show=\"isAuthenticated()\">\n" +
    "        <a class=\"block-support__link block-support__link_tech-support\" href>Поддержка</a>\n" +
    "    </li>\n" +
    "\n" +
    "    <li ng-hide=\"isAuthenticated()\" class=\"login\">\n" +
    "        <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"login()\">Войти</a>\n" +
    "    </li>\n" +
    "\n" +
    "    <li ng-show=\"isAuthenticated()\" class=\"dropdown\">\n" +
    "        <a href=\"#\" class=\"dropdown-toggle block-support__link block-support__link_user\" data-toggle=\"dropdown\">\n" +
    "            {{currentUser.LastName}} {{currentUser.Name.substring(0,1)}}.\n" +
    "        </a>\n" +
    "        <ul class=\"dropdown-menu\">\n" +
    "            <li><a ng-click=\"logout()\">Выход</a></li>\n" +
    "        </ul>\n" +
    "    </li>\n" +
    "</ul>");
}]);

angular.module("services/notifications.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("services/notifications.tpl.html",
    "<div class=\"alert-messages\">\n" +
    "  <div class=\"container\">\n" +
    "    <div ng-class=\"['alert', 'alert-'+notification.type]\" ng-repeat=\"notification in notifications.getCurrent()\">\n" +
    "      <button class=\"close\" ng-click=\"removeNotification(notification)\">x</button>\n" +
    "      {{notification.message}}\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);
