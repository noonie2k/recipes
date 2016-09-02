angular.module('RecipesApp')
  .config(function($routeProvider, $locationProvider) {
    $routeProvider
      .when('/items', {
        templateUrl: 'items.html',
        controller: 'ItemIndexController'
      })
      .when('/item/:itemId', {
        templateUrl: 'item.html',
        controller: 'ItemShowController'
      })
      .when('/items/reusable', {
        templateUrl: 'items.html',
        controller: 'ReusableListController'
      })
      .when('/build', {
        templateUrl: 'build.html',
        controller: 'BuildCreateController',
        controllerAs: 'buildCtrl'
      });
  });
