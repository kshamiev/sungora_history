angular.module('zegota.profile.config', [])

.value('profileUrl', {
  profilePutData: 'http://gocmf.webdesk.ru/api/v1/access/users/t/',
  profileGetCurrentUserData: 'api/users/current-user'
});
