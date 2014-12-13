angular.module('security.interceptor', ['security.retryQueue'])

	.factory('securityInterceptor', ['$injector', 'securityRetryQueue', function ($injector, queue) {
		return function (promise) {
			return promise.then(null, function (originalResponse) {
				if (originalResponse.status === 401) {
					promise = queue.pushRetryFn('unauthorized-server', function retryRequest() {
						return $injector.get('$http')(originalResponse.config);
					});
				}
				return promise;
			});
		};
	}])

	.config(['$httpProvider', function ($httpProvider) {
		$httpProvider.responseInterceptors.push('securityInterceptor');
	}])
;
