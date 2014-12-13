angular.module('templates.user', ['src/app/profile/profile.modal.tpl.html', 'src/app/profile/profile.toolbar.tpl.html', 'src/app/recovery/recovery.toolbar.tpl.html', 'src/app/recovery/step1.modal.tpl.html', 'src/app/recovery/step2.tpl.html', 'src/app/registration/registration.toolbar.tpl.html', 'src/app/registration/step1.modal.tpl.html', 'src/app/registration/step2.tpl.html']);

angular.module("src/app/profile/profile.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/profile/profile.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelProfile()\">&times;</button>\n" +
    "        <h3>Профиль</h3>\n" +
    "      </div>\n" +
    "      <form novalidate name=\"profileForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"profileError\">\n" +
    "            {[{profileError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"profileSuccess\">\n" +
    "            {[{profileSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <fieldset>\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfName\">Имя</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfName\" name=\"inputProfName\" class=\"form-control\" type=\"text\" ng-model=\"model.Name\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfMiddleName\">Отчество</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfMiddleName\" name=\"inputProfMiddleName\" class=\"form-control\" type=\"text\" ng-model=\"model.MiddleName\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfLastName\">Фамилия</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfLastName\" name=\"inputProfLastName\" class=\"form-control\" type=\"text\" ng-model=\"model.LastName\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputProfPassword\">Пароль</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfPassword\" name=\"inputProfPassword\" class=\"form-control\" type=\"password\" ng-model=\"model.Password\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <label class=\"control-label\" for=\"inputRegLastName\">Подтверждение</label>\n" +
    "              <input autocomplete=\"off\" id=\"inputProfPasswordR\" name=\"inputProfPasswordR\" class=\"form-control\" type=\"password\" ng-model=\"model.PasswordR\">\n" +
    "            </div>\n" +
    "\n" +
    "            <div class=\"form-group\">\n" +
    "              <button class=\"btn btn-primary btn-block\" ng-click=\"submitUserData()\" ng-disabled=\"profileFormForm.$invalid\" type=\"submit\">Соxранить</button>\n" +
    "            </div>\n" +
    "          </fieldset>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("src/app/profile/profile.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/profile/profile.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-show=\"isAuthenticated()\" class=\"profile\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"profile()\">Профиль</a>\n" +
    "  </li>\n" +
    "</ul>\n" +
    "");
}]);

angular.module("src/app/recovery/recovery.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/recovery/recovery.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-hide=\"isAuthenticated()\" class=\"recovery\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"recovery()\">Восстановление</a>\n" +
    "  </li>\n" +
    "</ul>\n" +
    "");
}]);

angular.module("src/app/recovery/step1.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/recovery/step1.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelRecovery()\">&times;</button>\n" +
    "        <h3>Восстановление пароля</h3>\n" +
    "      </div>\n" +
    "      <form novalidate=\"true\" name=\"recoveryForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"registerError\">\n" +
    "            {[{registerError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"registerSuccess\">\n" +
    "            {[{registerSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div ng-hide=\"registerSuccess\">\n" +
    "              <fieldset>\n" +
    "                <div class=\"form-group\">\n" +
    "                  <label class=\"control-label\" for=\"inputRegEmail\">Email</label>\n" +
    "                  <input autocomplete=\"off\" id=\"inputRegEmail\" name=\"inputRegEmail\" class=\"form-control\" type=\"email\" ng-model=\"model.Email\" required=\"true\" autofocus />\n" +
    "                </div>\n" +
    "\n" +
    "                <div class=\"form-group\">\n" +
    "                  <button class=\"btn btn-primary btn-block\" ng-click=\"recovery()\" ng-disabled=\"recoveryForm.$invalid\" type=\"submit\">Восстановить пароль</button>\n" +
    "                  <button class=\"btn btn-default btn-block\" ng-click=\"clearForm()\" type=\"button\">Очистить форму</button>\n" +
    "                </div>\n" +
    "            </fieldset>\n" +
    "          </div>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("src/app/recovery/step2.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/recovery/step2.tpl.html",
    "<div ng-if=\"registerSuccess\">\n" +
    "  <p>{[{ registerSuccess }]}</p>\n" +
    "</div>\n" +
    "<div ng-if=\"registerError\">\n" +
    "  <p>{[{ registerError }]}</p>\n" +
    "</div>\n" +
    "");
}]);

angular.module("src/app/registration/registration.toolbar.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/registration/registration.toolbar.tpl.html",
    "<ul class=\"navbar-nav nav\">\n" +
    "  <li ng-hide=\"isAuthenticated()\" class=\"registration\">\n" +
    "    <a class=\"dropdown-toggle block-support__link block-support__link_user\" ng-click=\"registration()\">Регистрация</a>\n" +
    "  </li>\n" +
    "</ul>");
}]);

angular.module("src/app/registration/step1.modal.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/registration/step1.modal.tpl.html",
    "<div class=\"modal\">\n" +
    "  <div class=\"modal-dialog\">\n" +
    "    <div class=\"modal-content\">\n" +
    "      <div class=\"modal-header\">\n" +
    "        <button type=\"button\" class=\"close\" ng-click=\"cancelRegistration()\">&times;</button>\n" +
    "        <h3>Регистрация</h3>\n" +
    "      </div>\n" +
    "      <form novalidate name=\"registrationForm\">\n" +
    "        <div class=\"modal-body\">\n" +
    "\n" +
    "          <div class=\"alert alert-error\" ng-show=\"registerError\">\n" +
    "            {[{registerError}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div class=\"alert alert-success\" ng-show=\"registerSuccess\">\n" +
    "            {[{registerSuccess}]}\n" +
    "          </div>\n" +
    "\n" +
    "          <div ng-hide=\"registerSuccess\">\n" +
    "            <fieldset>\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegEmail\">Email</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegEmail\" name=\"inputRegEmail\" class=\"form-control\" type=\"email\" ng-model=\"model.Email\" required=\"true\" autofocus>\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegPhone\">Логин</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegPhone\" name=\"inputRegPhone\" class=\"form-control\" type=\"text\" ng-model=\"model.Login\" required=\"true\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegName\">Имя</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegName\" name=\"inputRegName\" class=\"form-control\" type=\"text\" ng-model=\"model.Name\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegMiddleName\">Отчество</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegMiddleName\" name=\"inputRegMiddleName\" class=\"form-control\" type=\"text\" ng-model=\"model.MiddleName\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <label class=\"control-label\" for=\"inputRegLastName\">Фамилия</label>\n" +
    "                <input autocomplete=\"off\" id=\"inputRegLastName\" name=\"inputRegLastName\" class=\"form-control\" type=\"text\" ng-model=\"model.LastName\">\n" +
    "              </div>\n" +
    "\n" +
    "              <div class=\"form-group\">\n" +
    "                <button class=\"btn btn-primary btn-block\" ng-click=\"register()\" ng-disabled=\"registrationForm.$invalid\" type=\"submit\">Зарегестрироваться</button>\n" +
    "                <button class=\"btn btn-default btn-block\" ng-click=\"clearForm()\" type=\"button\">Очистить форму</button>\n" +
    "              </div>\n" +
    "            </fieldset>\n" +
    "          </div>\n" +
    "        </div>\n" +
    "      </form>\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("src/app/registration/step2.tpl.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("src/app/registration/step2.tpl.html",
    "<div ng-if=\"registerSuccess\">\n" +
    "  <p>{[{ registerSuccess }]}</p>\n" +
    "</div>\n" +
    "<div ng-if=\"registerError\">\n" +
    "  <p>{[{ registerError }]}</p>\n" +
    "</div>\n" +
    "");
}]);
