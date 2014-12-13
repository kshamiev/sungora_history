angular.module('common.services.modalWindowControl', [])
.factory('modalWindowControl', ['$modal', '$window', '$location', function($modal, $window, $location) {

  var dialog = null;

  function redirect(url) {
    url = url || '/';
    if ($window.location.pathname !== url) {
      $window.location = url;
    }
  }

  function openDialog(dialogParams) {
    if (dialog) {
      throw new Error('Trying to open a dialog that is already open!');
    }
    if (!dialog) {
      dialog = $modal.open({
        keyboard: dialogParams.keyboard,
        backdrop: dialogParams.backdrop,
        windowClass: dialogParams.windowClass,
        templateUrl: dialogParams.templateUrl,
        controller: dialogParams.controllerName
      });
      dialog.result.then(onDialogClose);
    }
  }

  function closeDialog(success, redirectTo) { //check input params
    if (dialog) {
      dialog.close({success: success, redirectTo: redirectTo });
    }
  }

  function onDialogClose(result) {
    dialog = null;
    if (typeof queue !== "undefined") {
      if (result.success) {
        queue.retryAll();
      } else {
        queue.cancelAll();
        redirect(result.redirectTo);
      }
    }
  }

  return {
    showDialog: function (dialogParams) {
      openDialog(dialogParams);
    },

    cancelDialog: function () {
      closeDialog();
    }
  }

}]);
