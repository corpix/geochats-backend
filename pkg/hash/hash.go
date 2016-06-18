package hash

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
)

// Create returns a hex hash value for arbitrary data
func Create(v interface{}) (string, error) {
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash), nil
}
