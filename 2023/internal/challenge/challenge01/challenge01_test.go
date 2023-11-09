package challenge01_test

import (
	"strings"
	"testing"

	"codember/internal/challenge/challenge01"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestChallenge01_SliceToLower(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it converts every item in a string slice to lower case", func(t *testing.T) {
		input := "cat dog dog car Cat doG sun"
		expected := []string{"cat", "dog", "dog", "car", "cat", "dog", "sun"}

		assert.Equal(expected, challenge01.SliceToLower(strings.Split(input, " ")))
	})
}

func TestChallenge01_InsertionOrderedMap(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it is possible to insert and get key-value pairs", func(t *testing.T) {
		orderedMap := challenge01.NewInsertionOrderedMap()
		orderedMap.Set("foo", 1)
		orderedMap.Set("bar", 2)

		assert.Equal([]string{"foo", "bar"}, orderedMap.Keys())
		assert.Equal(1, orderedMap.Get("foo"))
		assert.Equal(2, orderedMap.Get("bar"))
	})

	t.Run("when inserting or getting, it ignores the case of the key", func(t *testing.T) {
		orderedMap := challenge01.NewInsertionOrderedMap()
		orderedMap.Set("foo", 1)
		orderedMap.Set("FOO", 2)

		assert.Equal([]string{"foo"}, orderedMap.Keys())
		assert.Equal(2, orderedMap.Get("foo"))
		assert.Equal(2, orderedMap.Get("FOO"))
	})

	t.Run("it is possible to get the raw map of key-value pairs", func(t *testing.T) {
		orderedMap := challenge01.NewInsertionOrderedMap()
		orderedMap.Set("foo", 1)
		orderedMap.Set("bar", 2)

		assert.Equal(map[string]int{"foo": 1, "bar": 2}, orderedMap.ToMap())
	})

	t.Run("it is possible to get a string of all key-value pairs in order", func(t *testing.T) {
		orderedMap := challenge01.NewInsertionOrderedMap()
		orderedMap.Set("foo", 1)
		orderedMap.Set("bar", 2)
		orderedMap.Set("baz", 3)

		assert.Equal("foo1bar2baz3", orderedMap.ToString())
	})
}

func TestChallenge01_CountOccurrences(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it counts the number of words in a string", func(t *testing.T) {
		input := []string{"keys", "house", "house", "house", "keys"}
		expected := map[string]int{
			"keys":  2,
			"house": 3,
		}

		assert.Equal(expected, challenge01.CountOccurrences(input).ToMap())
	})
}

func TestChallenge01_Run(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it gives the correct output based on the challenge input", func(t *testing.T) {
		input := "cat dog dog car Cat doG sun"
		expected := "cat2dog3car1sun1"

		assert.Equal(expected, challenge01.Run(input))
	})
}
