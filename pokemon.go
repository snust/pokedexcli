package main

type Pokemon struct {
	name           string
	baseExperience int
	height         int
	weight         int
	stats          map[string]int
	types          []string
}
