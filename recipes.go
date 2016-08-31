/* -*- compile-command: "go build -v" -*- */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type List struct {
	required Items
	leftover Items
}

func ReadData() map[string]*Item {
	dataDir := "./data"

	data := make(map[string]*Item)

	filepath.Walk(dataDir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".json") {

			itemsData, _ := ioutil.ReadFile(path)

			decoder := json.NewDecoder(bytes.NewReader(itemsData))
			decoder.Decode(&data)
		}

		return nil
	})

	hydrated := 1

	for hydrated > 0 {
		hydrated = 0
		for _, item := range data {
			if item.Hydrated == false {
				if len(item.Ingredients) == 0 {
					// Hydrate simple items (no ingredients)
					item.Hydrated = true
					hydrated++

				} else {
					// Hydrate items with ingredients
					for ingredientKey, quantity := range item.Ingredients {
						for i := 0; i < quantity; i++ {
							if ingredient, ok := data[ingredientKey]; ok {
								if ok == true {
									item.Recipe = append(item.Recipe, ingredient)
									hydrated++
								} else {
								}
							} else {
								panic(fmt.Errorf("Could not find recipe for %s", ingredientKey))
							}
						}

						delete(item.Ingredients, ingredientKey)
					}
				}
			}
		}
	}

	return data
}

func Recipe(item *Item, list List) List {
	fmt.Printf("Making: %s\n", item.Name)

	// First check the bag
	for i, element := range list.leftover {
		if item == element {

			// Remove from the bag if not reusable
			if item.Reusable != true {
				list.leftover = append(list.leftover[:i], list.leftover[i+1:]...)
			}

			return list
		}
	}

	// Then make the item
	for _, element := range item.Recipe {
		list = Recipe(element, list)
	}

	// Store the item used (if no ingredients)
	if len(item.Recipe) == 0 {
		list.required = append(list.required, item)
	}

	// Store the leftover items, reusables are always stored
	made := item.Makes
	if item.Reusable != true {
		made--
	}

	fmt.Printf("Using: %s\n", item.Name)
	for i := 0; i < made; i++ {
		list.leftover = append(list.leftover, item)
	}

	return list
}

func count_items(items Items) map[string]int {
	item_counts := make(map[string]int)

	for _, element := range items {
		_, exist := item_counts[element.Name]

		if exist {
			item_counts[element.Name] += 1
		} else {
			item_counts[element.Name] = 1
		}
	}

	return item_counts
}
