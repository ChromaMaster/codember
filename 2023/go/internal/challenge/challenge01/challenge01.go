package challenge01

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type InsertionOrderedMap struct {
	keys   []string
	values map[string]int
}

func NewInsertionOrderedMap() *InsertionOrderedMap {
	return &InsertionOrderedMap{
		keys:   make([]string, 0),
		values: make(map[string]int),
	}
}

func (i *InsertionOrderedMap) Set(key string, value int) {
	key = strings.ToLower(key)

	if !slices.Contains(i.keys, key) {
		i.keys = append(i.keys, key)
	}

	i.values[key] = value
}

func (i *InsertionOrderedMap) Get(key string) int {
	key = strings.ToLower(key)

	return i.values[key]
}

func (i *InsertionOrderedMap) Keys() []string {
	return slices.Clone(i.keys)
}

func (i *InsertionOrderedMap) ToMap() map[string]int {
	return maps.Clone(i.values)
}

func (i *InsertionOrderedMap) ToString() string {
	output := ""
	for _, key := range i.keys {
		output = fmt.Sprintf("%s%s%d", output, key, i.values[key])
	}
	return output
}

func Run(input string) string {
	return CountOccurrences(SliceToLower(strings.Split(input, " "))).ToString()
}

func SliceToLower(input []string) []string {
	output := make([]string, 0, len(input))

	for _, word := range input {
		output = append(output, strings.ToLower(word))
	}

	return output
}

func CountOccurrences(input []string) *InsertionOrderedMap {
	output := NewInsertionOrderedMap()

	for _, word := range input {
		output.Set(word, output.Get(word)+1)
	}

	return output
}
