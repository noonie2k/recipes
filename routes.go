package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"ItemIndex", "GET", "/api/items", ItemIndex},
	Route{"ItemShow", "GET", "/api/item/{itemId}", ItemShow},
	Route{"RecipeShow", "GET", "/api/recipe/{itemId}", RecipeShow},
	Route{"RecipeShow", "POST", "/api/recipe/{itemId}", RecipeShow},
	Route{"ReusableIndex", "GET", "/api/items/reusable", ReusableIndex},
	Route{"BuildShow", "POST", "/api/build", BuildShow},
}
