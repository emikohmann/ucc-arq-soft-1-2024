package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
	var item map[string]interface{}
	var err error

	item, err = GetItem("MLA904895322")
	assert.NotNil(t, item)
	assert.Nil(t, err)
	assert.Equal(t, "Camisa Manga Larga Lisa Hombre Variedad Colores Con Bolsillo", item["title"])

	item, err = GetItem("Invalid")
	assert.Nil(t, item)
	assert.NotNil(t, err)
	assert.Equal(t, "unexpected 404 in http.Get", err.Error())
}
