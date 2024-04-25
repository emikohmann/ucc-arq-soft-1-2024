package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	resultado, err := division(8, 4)

	if resultado != 2 {
		fmt.Println("El resultado deberia ser 2")
		t.Fail()
	}

	if err != nil {
		fmt.Println("No deberia haber error")
		t.Fail()
	}
}

func TestDivisionEmi(t *testing.T) {
	resultado, err := division(8, 4)
	assert.Equal(t, resultado, 2)
	assert.Equal(t, err, nil)
}
