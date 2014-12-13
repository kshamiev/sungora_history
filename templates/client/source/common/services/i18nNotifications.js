angular.module('common.services.i18nNotifications', [
  'common.services.localizedMessages',
  'common.services.notifications',
  'dialogs'
])
.factory('i18nNotifications', ['localizedMessages', 'notifications', '$dialogs', function (localizedMessages, notifications, $dialogs) {

	var prepareNotification = function (msgKey, type, interpolateParams, otherProperties) {
		return angular.extend({
			message: localizedMessages.get(msgKey, interpolateParams),
			type: type
		}, otherProperties);
	};

	var I18nNotifications = {
		modal: function (msgKey, type, interpolateParams, otherProperties) {
			var notify = prepareNotification(msgKey, type, interpolateParams, otherProperties);
			switch (notify.type) {
				case 'error':
					$dialogs.error(localizedMessages.get("modal.error"), notify.message);
					break;
				case 'notify':
					$dialogs.notify(localizedMessages.get("modal.notify"), notify.message);
					break;
			}
		},
		pushSticky: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushSticky(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		pushForCurrentRoute: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushForCurrentRoute(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		pushForNextRoute: function (msgKey, type, interpolateParams, otherProperties) {
			return notifications.pushForNextRoute(prepareNotification(msgKey, type, interpolateParams, otherProperties));
		},
		getCurrent: function () {
			return notifications.getCurrent();
		},
		remove: function (notification) {
			return notifications.remove(notification);
		}
	};

	return I18nNotifications;
}]);
