angular.module('zegotaAdmin', [
  'zegota.core',
  'zegota.ui',
  'templates.admin'
])
.config(function ($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise('/admin');
  $stateProvider
    .state('/admin', {
      abstract: true,
      url: '/admin',
      tempate: '<div ui-view></div>'
    })
    .state('admin.dashboard', {
      url: '/dashboard',
      tempate: '<h1>admin dashboard</div>'
    })
    .state('admin.users', {
      url: '/users',
      tempate: '<h1>users</h1>'
    })
    .state('admin.groups', {
      url: '/groups',
      template: '<h1>groups</h1>'
    })
    .state('admin.routes', {
      url: '/routes',
      template: '<h1>routes</h1>'
    })
    .state('admin.controlles', {
      url: '/controllers',
      temlate: '<h1>controllers</h1>'
    })
})
;
