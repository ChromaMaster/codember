package challenge03

import (
	"fmt"
	"strconv"
	"strings"
)

type EncryptionPolicy struct {
	Min    int
	Max    int
	Letter string
}

func NewEncryptionPolicy(policyStr string) (*EncryptionPolicy, error) {
	options := strings.Split(policyStr, " ")
	minMax := strings.Split(options[0], "-")

	minValue, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, fmt.Errorf("cannot parse minValue: %w", err)
	}

	maxValue, err := strconv.Atoi(minMax[1])
	if err != nil {
		return nil, fmt.Errorf("cannot parse maxValue: %w", err)
	}

	return &EncryptionPolicy{
		minValue,
		maxValue,
		options[1],
	}, nil
}

func (e *EncryptionPolicy) Evaluate(password string) bool {
	count := strings.Count(password, e.Letter)

	if count >= e.Min && count <= e.Max {
		return true
	}

	return false
}

func Run(input []string, nth int) string {
	var wrongPasswords []string

	for _, line := range input {
		parts := strings.Split(line, ":")
		policyStr := strings.TrimSpace(parts[0])
		password := strings.TrimSpace(parts[1])

		policy, _ := NewEncryptionPolicy(policyStr)
		if !policy.Evaluate(password) {
			wrongPasswords = append(wrongPasswords, password)
			continue
		}
	}

	return wrongPasswords[nth-1]
}
