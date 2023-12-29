package challenge03_test

import (
	"codember/internal/challenge/challenge03"
	assertpkg "github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptionPolicy_New(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it creates as new encryption policy with the correct settings", func(t *testing.T) {
		policyStr := "2-4 f"
		policy, err := challenge03.NewEncryptionPolicy(policyStr)
		assert.NoError(err)

		assert.Equal(2, policy.Min)
		assert.Equal(4, policy.Max)
		assert.Equal("f", policy.Letter)
	})

	t.Run("it returns error when min is not an int value", func(t *testing.T) {
		policyStr := "a-4 f"
		_, err := challenge03.NewEncryptionPolicy(policyStr)
		assert.Error(err)
	})

	t.Run("it returns error when max is not an int value", func(t *testing.T) {
		policyStr := "2-b f"
		_, err := challenge03.NewEncryptionPolicy(policyStr)
		assert.Error(err)
	})
}

func TestEncryptionPolicy_Evaluate(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it identifies a correct password", func(t *testing.T) {
		policyStr := "3-4 f"
		policy, err := challenge03.NewEncryptionPolicy(policyStr)
		assert.NoError(err)

		assert.True(policy.Evaluate("fgff"))
		assert.True(policy.Evaluate("fgsdfasdff"))
	})

	t.Run("it identifies a wrong password", func(t *testing.T) {
		policyStr := "3-4 f"
		policy, err := challenge03.NewEncryptionPolicy(policyStr)
		assert.NoError(err)

		assert.False(policy.Evaluate("fgf"))
		assert.False(policy.Evaluate("fgsdfasfdff"))
	})
}

func TestChallenge02_Run(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("it returns the nth invalid password", func(t *testing.T) {
		input := []string{
			"2-4 f: fgff",
			"4-6 z: zzzsg",
			"1-6 h: hhhhhh",
			"4-8 d: qdwdedrdtdydudidod",
		}
		password := challenge03.Run(input, 1)
		assert.Equal("zzzsg", password)
	})
}
