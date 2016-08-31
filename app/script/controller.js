angular.module('RecipesApp')
  .controller('ItemIndexController', ['$scope', 'item', function($scope, itemService) {
    itemService.getAllItems().then(function(data) { $scope.items = data; });
  }])

  .controller('ItemShowController', ['$scope', '$routeParams', '$cookies', 'item', function($scope, $routeParams, $cookies, itemService) {
    var keptItems = function() {
      return $cookies.getObject('keptItems') || {};
    };

    itemService.getItem($routeParams.itemId).then(function(data) { $scope.item = data; });
    itemService.getRecipe($routeParams.itemId, keptItems()).then(function(data) { $scope.recipe = data; });
  }])

  .controller('ReusableListController', ['$scope', '$cookies', 'item', function($scope, $cookies, itemService) {
    var keptItemsKey = 'keptItems',
        keptItems = $cookies.getObject(keptItemsKey) || {};

    itemService.getReusables().then(function(data) { $scope.items = data; });

    $scope.keep = function(itemId) {
      console.log($scope.kept(itemId));
      keptItems[itemId] = 1;
      $cookies.putObject(keptItemsKey, keptItems);
    };

    $scope.drop = function(itemId) {
      delete keptItems[itemId];
      $cookies.putObject(keptItemsKey, keptItems);
    };

    $scope.kept = function(itemId) {
      return (keptItems[itemId] != undefined);
    };
  }]);
