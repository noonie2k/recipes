angular.module('RecipesApp')
  .factory('item', ['$http', function($http) {
    return {
      getAllItems: function() {
        var request = {
          method: 'GET',
          url: '/api/items'
        };

        return $http(request)
          .then(function(response) {
            return response.data;
          });
      },

      getItem: function(itemId) {
        var request = {
          method: 'GET',
          url: '/api/item/' + itemId
        };

        return $http(request)
          .then(function(response) {
            console.log(response.data);
            return response.data;
          });
      },

      getRecipe: function getRecipes(itemId, existingItems = {}) {
        var request = {
          method: 'POST',
          url: '/api/recipe/' + itemId,
          data: existingItems
        };

        return $http(request)
          .then(function(response) {
            return response.data;
          });
      },

      getReusables: function() {
        var request = {
          method: 'GET',
          url: '/api/items/reusable'
        };

        return $http(request)
          .then(function(response) {
            return response.data;
          });
      }
    };
  }]);
