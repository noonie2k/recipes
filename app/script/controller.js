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
  }])

  .controller('BuildCreateController', ['$scope', '$cookies', 'item', function($scope, $cookies, itemService) {
    itemService.getAllItems().then(function(data) { $scope.items = data; });

    $scope.getBuild = function() {
      return $cookies.getObject('build');
    };

    $scope.querySearch = function(query) {
      return query ? $scope.items.filter(createFilterFor(query)) : $scope.items;
    };

    $scope.addItem = function() {
      modifyBuild(1);
    };

    $scope.removeItem = function() {
      modifyBuild(-1);
    };

    $scope.clearBuild = function() {
      $cookies.remove('build');
      delete $scope.buildItems;
    };

    $scope.calculateBuild = function() {
      var build = $cookies.getObject('build') || {};

      itemService.calculateBuild({
        "requiredItems": build, "existingItems": $cookies.getObject('keptItems')
      }).then(function(data) { $scope.buildItems = data; });
    };

    function modifyBuild(change) {
      if ($scope.selectedItem !== null) {

        var build = $cookies.getObject('build') || {};
        console.log(build);
        if (!build[$scope.selectedItem.id]) build[$scope.selectedItem.id] = 0;

        build[$scope.selectedItem.id] += change;

        if (build[$scope.selectedItem.id] <= 0) {
          delete build[$scope.selectedItem.id];
        }

        console.log(build);

        $cookies.putObject('build', build);
      }
    }

    function createFilterFor(query) {
      var lowercaseQuery = angular.lowercase(query);

      return function(state) {
        return (
          state.id.indexOf(lowercaseQuery) !== -1
            || angular.lowercase(state.name).indexOf(lowercaseQuery) !== -1
        );
      };
    }
  }]);
