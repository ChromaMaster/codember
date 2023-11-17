package challenge02_test

import (
	"codember/internal/challenge/challenge02"
	"testing"

	assertpkg "github.com/stretchr/testify/assert"
)

func TestMiniCompiler_New(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it creates a new minicompiler with initial value 0", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		assert.Equal(0, compiler.Value)
	})

	t.Run("it creates a new minicompiler with initial empty output buffer", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		assert.Equal("", compiler.Output)
	})
}

func TestMiniCompiler_ParseInput(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("when received '#', it increases the value by 1", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		err := compiler.ParseInput('#')
		assert.Nil(err)

		assert.Equal(1, compiler.Value)
	})

	t.Run("when received '@', it decreases the value by 1", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		err := compiler.ParseInput('@')
		assert.Nil(err)

		assert.Equal(-1, compiler.Value)
	})

	t.Run("when received '*', it multiplies the value by itself", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		compiler.Value = 3

		err := compiler.ParseInput('*')
		assert.Nil(err)

		assert.Equal(9, compiler.Value)
	})

	t.Run("when received '&', it appends the value to the output buffer", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		compiler.Value = 21

		err := compiler.ParseInput('&')
		assert.Nil(err)

		assert.Equal("21", compiler.Output)
	})

	t.Run("when the input is not supported, it returns an error", func(t *testing.T) {
		compiler := challenge02.NewMiniCompiler()

		err := compiler.ParseInput('g')
		assert.ErrorIs(err, challenge02.ErrInputNotSupported)
	})
}

func TestChallenge02_Run(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it gives the correct output based on the challenge input", func(t *testing.T) {
		input := "&##&*&@&"
		expected := "0243"

		assert.Equal(expected, challenge02.Run(input))
	})
}
