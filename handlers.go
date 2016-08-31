package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func ItemIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := data[vars["itemId"]]

	if item != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(item); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func RecipeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := data[vars["itemId"]]

	if item != nil {
		var existingItemIds map[string]int

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body, &existingItemIds); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			// if err := json.NewEncoder(w).Encode(err); err != nil {
			// 	panic(err)
			// }
		}

		existingItems := make(Items, 0)
		for itemId, quantity := range existingItemIds {
			for i := 0; i < quantity; i++ {
				existingItems = append(existingItems, data[itemId])
			}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		recipe := Recipe(item, List{leftover: existingItems})
		result := map[string]map[string]int{
			"required": count_items(recipe.required),
			"leftover": count_items(recipe.leftover),
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func ReusableIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	reusables := make(map[string]*Item)

	for itemId, item := range data {
		fmt.Println(itemId)
		if item.Reusable == true {
			reusables[itemId] = item
		}
	}

	if err := json.NewEncoder(w).Encode(reusables); err != nil {
		panic(err)
	}
}
