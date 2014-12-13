angular.module('templates.bootstrap', ['template/accordion/accordion-group.html', 'template/accordion/accordion.html', 'template/alert/alert.html', 'template/carousel/carousel.html', 'template/carousel/slide.html', 'template/datepicker/datepicker.html', 'template/datepicker/day.html', 'template/datepicker/month.html', 'template/datepicker/popup.html', 'template/datepicker/year.html', 'template/dialog/message.html', 'template/modal/backdrop.html', 'template/modal/window.html', 'template/pagination/pager.html', 'template/pagination/pagination.html', 'template/popover/popover.html', 'template/progressbar/bar.html', 'template/progressbar/progress.html', 'template/progressbar/progressbar.html', 'template/rating/rating.html', 'template/tabs/tab.html', 'template/tabs/tabset-titles.html', 'template/tabs/tabset.html', 'template/timepicker/timepicker.html', 'template/tooltip/tooltip-html-unsafe-popup.html', 'template/tooltip/tooltip-popup.html', 'template/typeahead/typeahead-match.html', 'template/typeahead/typeahead-popup.html']);

angular.module("template/accordion/accordion-group.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/accordion/accordion-group.html",
    "<div class=\"accordion-group\">\n" +
    "  <div class=\"accordion-heading\" ><a class=\"accordion-toggle\" ng-click=\"isOpen = !isOpen\" accordion-transclude=\"heading\">{{heading}}</a></div>\n" +
    "  <div class=\"accordion-body\" collapse=\"!isOpen\">\n" +
    "    <div class=\"accordion-inner\" ng-transclude></div>  </div>\n" +
    "</div>");
}]);

angular.module("template/accordion/accordion.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/accordion/accordion.html",
    "<div class=\"accordion\" ng-transclude></div>");
}]);

angular.module("template/alert/alert.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/alert/alert.html",
    "<div class='alert' ng-class='type && \"alert-\" + type'>\n" +
    "    <button ng-show='closeable' type='button' class='close' ng-click='close()'>&times;</button>\n" +
    "    <div ng-transclude></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/carousel/carousel.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/carousel/carousel.html",
    "<div ng-mouseenter=\"pause()\" ng-mouseleave=\"play()\" class=\"carousel\">\n" +
    "    <ol class=\"carousel-indicators\" ng-show=\"slides().length > 1\">\n" +
    "        <li ng-repeat=\"slide in slides()\" ng-class=\"{active: isActive(slide)}\" ng-click=\"select(slide)\"></li>\n" +
    "    </ol>\n" +
    "    <div class=\"carousel-inner\" ng-transclude></div>\n" +
    "    <a ng-click=\"prev()\" class=\"carousel-control left\" ng-show=\"slides().length > 1\">&lsaquo;</a>\n" +
    "    <a ng-click=\"next()\" class=\"carousel-control right\" ng-show=\"slides().length > 1\">&rsaquo;</a>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/carousel/slide.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/carousel/slide.html",
    "<div ng-class=\"{\n" +
    "    'active': leaving || (active && !entering),\n" +
    "    'prev': (next || active) && direction=='prev',\n" +
    "    'next': (next || active) && direction=='next',\n" +
    "    'right': direction=='prev',\n" +
    "    'left': direction=='next'\n" +
    "  }\" class=\"item\" ng-transclude></div>\n" +
    "");
}]);

angular.module("template/datepicker/datepicker.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/datepicker.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"pull-left\" ng-click=\"move(-1)\"><span class=\"nav-arrow nav-arrow_left\"></span></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"{{rows[0].length - 2 + showWeekNumbers}}\">\n" +
    "			<button type=\"button\" class=\"\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"pull-right\" ng-click=\"move(1)\"><span class=\"nav-arrow nav-arrow_right\"></span></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	<tr ng-show=\"labels.length > 0\" class=\"h6\">\n" +
    "		<th ng-show=\"showWeekNumbers\" class=\"text-center\">#</th>\n" +
    "		<th ng-repeat=\"label in labels\" class=\"text-center\">{{label}}</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows\">\n" +
    "		<td ng-show=\"showWeekNumbers\" class=\"text-center\"><em>{{ getWeekNumber(row) }}</em></td>\n" +
    "		<td ng-repeat=\"dt in row\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"\"\n" +
    "			        ng-class=\"{'btn-info': dt.selected}\" ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span\n" +
    "					ng-class=\"{'text-muted': dt.secondary}\">{{dt.label}}</span></button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/day.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/day.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"{{5 + showWeeks}}\">\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	<tr>\n" +
    "		<th ng-show=\"showWeeks\" class=\"text-center\"></th>\n" +
    "		<th ng-repeat=\"label in labels track by $index\" class=\"text-center\">\n" +
    "			<small>{{label}}</small>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-show=\"showWeeks\" class=\"text-center\">\n" +
    "			<small><em>{{ weekNumbers[$index] }}</em></small>\n" +
    "		</td>\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default btn-sm\"\n" +
    "			        ng-class=\"{'btn-info': dt.selected}\" ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span\n" +
    "					ng-class=\"{'text-muted': dt.secondary, 'text-info': dt.current}\">{{dt.label}}</span></button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/month.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/month.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default\" ng-class=\"{'btn-info': dt.selected}\"\n" +
    "			        ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span ng-class=\"{'text-info': dt.current}\">{{dt.label}}</span>\n" +
    "			</button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/datepicker/popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/popup.html",
    "<ul class=\"dropdown-menu dropdown-datepicker\"\n" +
    "    ng-style=\"{display: (isOpen && 'block') || 'none', top: position.top+'px', left: position.left+'px'}\">\n" +
    "	<li ng-transclude></li>\n" +
    "\n" +
    "	<!--<li ng-show=\"showButtonBar\" style=\"padding:10px 9px 2px\">-->\n" +
    "		<!--<span class=\"btn-group\">-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-info\" ng-click=\"today()\">{{currentText}}</button>-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-default\" ng-click=\"showWeeks = ! showWeeks\"-->\n" +
    "			        <!--ng-class=\"{active: showWeeks}\">{{toggleWeeksText}}-->\n" +
    "			<!--</button>-->\n" +
    "			<!--<button type=\"button\" class=\"btn btn-sm btn-danger\" ng-click=\"clear()\">{{clearText}}</button>-->\n" +
    "		<!--</span>-->\n" +
    "		<!--<button type=\"button\" class=\"btn btn-sm btn-success pull-right\" ng-click=\"isOpen = false\">{{closeText}}</button>-->\n" +
    "	<!--</li>-->\n" +
    "</ul>\n" +
    "");
}]);

angular.module("template/datepicker/year.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/datepicker/year.html",
    "<table>\n" +
    "	<thead>\n" +
    "	<tr>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-left\" ng-click=\"move(-1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-left\"></i></button>\n" +
    "		</th>\n" +
    "		<th colspan=\"3\">\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm btn-block\" ng-click=\"toggleMode()\">\n" +
    "				<strong>{{title}}</strong></button>\n" +
    "		</th>\n" +
    "		<th>\n" +
    "			<button type=\"button\" class=\"btn btn-default btn-sm pull-right\" ng-click=\"move(1)\"><i\n" +
    "					class=\"glyphicon glyphicon-chevron-right\"></i></button>\n" +
    "		</th>\n" +
    "	</tr>\n" +
    "	</thead>\n" +
    "	<tbody>\n" +
    "	<tr ng-repeat=\"row in rows track by $index\">\n" +
    "		<td ng-repeat=\"dt in row track by dt.date\" class=\"text-center\">\n" +
    "			<button type=\"button\" style=\"width:100%;\" class=\"btn btn-default\" ng-class=\"{'btn-info': dt.selected}\"\n" +
    "			        ng-click=\"select(dt.date)\" ng-disabled=\"dt.disabled\"><span ng-class=\"{'text-info': dt.current}\">{{dt.label}}</span>\n" +
    "			</button>\n" +
    "		</td>\n" +
    "	</tr>\n" +
    "	</tbody>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/dialog/message.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/dialog/message.html",
    "<div class=\"modal-header\">\n" +
    "	<h3>{{ title }}</h3>\n" +
    "</div>\n" +
    "<div class=\"modal-body\">\n" +
    "	<p>{{ message }}</p>\n" +
    "</div>\n" +
    "<div class=\"modal-footer\">\n" +
    "	<button ng-repeat=\"btn in buttons\" ng-click=\"close(btn.result)\" class=\"btn\" ng-class=\"btn.cssClass\">{{ btn.label }}</button>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/modal/backdrop.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/modal/backdrop.html",
    "<div class=\"modal-backdrop fade\" ng-class=\"{in: animate}\" ng-style=\"{'z-index': 1040 + index*10}\" ng-click=\"close($event)\"></div>");
}]);

angular.module("template/modal/window.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/modal/window.html",
    "<div tabindex=\"-1\" class=\"modal fade {{ windowClass }}\" ng-class=\"{in: animate}\" ng-style=\"{'z-index': 1050 + index*10}\" ng-transclude></div>");
}]);

angular.module("template/pagination/pager.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/pagination/pager.html",
    "<div class=\"pager\">\n" +
    "  <ul>\n" +
    "    <li ng-repeat=\"page in pages\" ng-class=\"{disabled: page.disabled, previous: page.previous, next: page.next}\"><a ng-click=\"selectPage(page.number)\">{{page.text}}</a></li>\n" +
    "  </ul>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/pagination/pagination.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/pagination/pagination.html",
    "<div class=\"pagination\"><ul>\n" +
    "  <li ng-repeat=\"page in pages\" ng-class=\"{active: page.active, disabled: page.disabled}\"><a ng-click=\"selectPage(page.number)\">{{page.text}}</a></li>\n" +
    "  </ul>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/popover/popover.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/popover/popover.html",
    "<div class=\"popover {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"arrow\"></div>\n" +
    "\n" +
    "  <div class=\"popover-inner\">\n" +
    "      <h3 class=\"popover-title\" ng-bind=\"title\" ng-show=\"title\"></h3>\n" +
    "      <div class=\"popover-content\" ng-bind=\"content\"></div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/progressbar/bar.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/bar.html",
    "<div class=\"bar\" ng-class=\"type && 'bar-' + type\" ng-transclude></div>");
}]);

angular.module("template/progressbar/progress.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/progress.html",
    "<div class=\"progress\" ng-transclude></div>");
}]);

angular.module("template/progressbar/progressbar.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/progressbar/progressbar.html",
    "<div class=\"progress\"><div class=\"bar\" ng-class=\"type && 'bar-' + type\" ng-transclude></div></div>");
}]);

angular.module("template/rating/rating.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/rating/rating.html",
    "<span ng-mouseleave=\"reset()\">\n" +
    "	<i ng-repeat=\"r in range\" ng-mouseenter=\"enter($index + 1)\" ng-click=\"rate($index + 1)\" ng-class=\"$index < val && (r.stateOn || 'icon-star') || (r.stateOff || 'icon-star-empty')\"></i>\n" +
    "</span>");
}]);

angular.module("template/tabs/tab.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tab.html",
    "<li ng-class=\"{active: active, disabled: disabled}\">\n" +
    "  <a ng-click=\"select()\" tab-heading-transclude>{{heading}}</a>\n" +
    "</li>\n" +
    "");
}]);

angular.module("template/tabs/tabset-titles.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tabset-titles.html",
    "<ul class=\"nav {{type && 'nav-' + type}}\" ng-class=\"{'nav-stacked': vertical}\">\n" +
    "</ul>\n" +
    "");
}]);

angular.module("template/tabs/tabset.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tabs/tabset.html",
    "<div class=\"tabbable\">\n" +
    "  <ul class=\"nav {{type && 'nav-' + type}}\" ng-class=\"{'nav-stacked': vertical}\" ng-transclude>\n" +
    "  </ul>\n" +
    "  <div class=\"tab-content\">\n" +
    "    <div class=\"tab-pane\" \n" +
    "         ng-repeat=\"tab in tabs\" \n" +
    "         ng-class=\"{active: tab.active}\"\n" +
    "         tab-content-transclude=\"tab\">\n" +
    "    </div>\n" +
    "  </div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/timepicker/timepicker.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/timepicker/timepicker.html",
    "<table class=\"form-inline\">\n" +
    "	<tr class=\"text-center\">\n" +
    "		<td><a ng-click=\"incrementHours()\" class=\"btn btn-link\"><i class=\"icon-chevron-up\"></i></a></td>\n" +
    "		<td>&nbsp;</td>\n" +
    "		<td><a ng-click=\"incrementMinutes()\" class=\"btn btn-link\"><i class=\"icon-chevron-up\"></i></a></td>\n" +
    "		<td ng-show=\"showMeridian\"></td>\n" +
    "	</tr>\n" +
    "	<tr>\n" +
    "		<td class=\"control-group\" ng-class=\"{'error': invalidHours}\"><input type=\"text\" ng-model=\"hours\" ng-change=\"updateHours()\" class=\"span1 text-center\" ng-mousewheel=\"incrementHours()\" ng-readonly=\"readonlyInput\" maxlength=\"2\"></td>\n" +
    "		<td>:</td>\n" +
    "		<td class=\"control-group\" ng-class=\"{'error': invalidMinutes}\"><input type=\"text\" ng-model=\"minutes\" ng-change=\"updateMinutes()\" class=\"span1 text-center\" ng-readonly=\"readonlyInput\" maxlength=\"2\"></td>\n" +
    "		<td ng-show=\"showMeridian\"><button type=\"button\" ng-click=\"toggleMeridian()\" class=\"btn text-center\">{{meridian}}</button></td>\n" +
    "	</tr>\n" +
    "	<tr class=\"text-center\">\n" +
    "		<td><a ng-click=\"decrementHours()\" class=\"btn btn-link\"><i class=\"icon-chevron-down\"></i></a></td>\n" +
    "		<td>&nbsp;</td>\n" +
    "		<td><a ng-click=\"decrementMinutes()\" class=\"btn btn-link\"><i class=\"icon-chevron-down\"></i></a></td>\n" +
    "		<td ng-show=\"showMeridian\"></td>\n" +
    "	</tr>\n" +
    "</table>\n" +
    "");
}]);

angular.module("template/tooltip/tooltip-html-unsafe-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tooltip/tooltip-html-unsafe-popup.html",
    "<div class=\"tooltip {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"tooltip-arrow\"></div>\n" +
    "  <div class=\"tooltip-inner\" bind-html-unsafe=\"content\"></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/tooltip/tooltip-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/tooltip/tooltip-popup.html",
    "<div class=\"tooltip {{placement}}\" ng-class=\"{ in: isOpen(), fade: animation() }\">\n" +
    "  <div class=\"tooltip-arrow\"></div>\n" +
    "  <div class=\"tooltip-inner\" ng-bind=\"content\"></div>\n" +
    "</div>\n" +
    "");
}]);

angular.module("template/typeahead/typeahead-match.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/typeahead/typeahead-match.html",
    "<a tabindex=\"-1\" bind-html-unsafe=\"match.label | typeaheadHighlight:query\"></a>");
}]);

angular.module("template/typeahead/typeahead-popup.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("template/typeahead/typeahead-popup.html",
    "<ul class=\"typeahead dropdown-menu\" ng-style=\"{display: isOpen()&&'block' || 'none', top: position.top+'px', left: position.left+'px'}\">\n" +
    "    <li ng-repeat=\"match in matches\" ng-class=\"{active: isActive($index) }\" ng-mouseenter=\"selectActive($index)\" ng-click=\"selectMatch($index)\">\n" +
    "        <div typeahead-match index=\"$index\" match=\"match\" query=\"query\" template-url=\"templateUrl\"></div>\n" +
    "    </li>\n" +
    "</ul>");
}]);
