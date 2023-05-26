package config

import (
	"fmt"
	"testing"
)

func TestNewArti_NewArti(t *testing.T) {
	artiList, er := NewArti()
	if er != nil {
		fmt.Println(artiList)
	}

	fmt.Println(artiList)
}
