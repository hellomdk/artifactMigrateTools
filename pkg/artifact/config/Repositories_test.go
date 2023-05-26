package config

import (
	"fmt"
	"testing"
)

func TestNewRepositories_NewRepositories(t *testing.T) {
	repositories, er := NewRepositories()
	if er != nil {
		fmt.Println(repositories)
	}

	fmt.Println(repositories)
}
