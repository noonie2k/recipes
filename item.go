package main

type Item struct {
	Name        string         `json:"name"`
	Makes       int            `json:"makes"`
	Reusable    bool           `json:"reusable"`
	Recipe      Items          `json:"recipe"`
	Ingredients map[string]int `json:"ingredients"`
	Hydrated    bool           `json:"-"`
}

type Items []*Item
