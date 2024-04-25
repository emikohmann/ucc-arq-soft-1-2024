// Package
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Response struct {
	Results []Result
}

type Result struct {
	Title              string  `json:"title"`
	Price              float64 `json:"price"`
	AvailableQuantity  int     `json:"available_quantity "`
	Condition          string  `json:"condition"`
	AcceptsMercadopago bool    `json:"accepts_mercadopago"`
	CategoryID         string  `json:"category_id"`
}

type Category struct {
	Name string `json:"name"`
}

// Ejercicio: Implementar la funcion getCategory que
// Recibe un category id (string) y
// Retorna el nombre de la categoria (string)

func main() {
	fmt.Print("Ingresa una busqueda: ")
	input := readInput()
	results := search(input)
	for i, result := range results {
		// Celular Iphone 13 - MLA1234
		category := getCategory(result.CategoryID)
		product := fmt.Sprintf(" [%s] [%d] %s - $%.2f", category.Name, i+1, result.Title, result.Price)
		fmt.Println(product)
		time.Sleep(200 * time.Millisecond)
	}
}

func getCategory(categoryID string) Category {
	// Construyo URL
	categoryID = url.QueryEscape(categoryID)
	urlML := "http://api.mercadolibre.com/categories/" + categoryID

	// Ejecuto request
	resp, err := http.Get(urlML)
	if err != nil {
		fmt.Println(err.Error())
		return Category{}
	}

	// Valido Status Code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received error status code")
		return Category{}
	}

	// Leer respuesta
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return Category{}
	}

	// Asignar los datos en la estructura
	var response Category
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		fmt.Println(err.Error())
		return Category{}
	}

	return response
}

// Input de busqueda
func readInput() string {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

func search(query string) []Result {
	// Construyo URL
	query = url.QueryEscape(query)
	urlML := "http://api.mercadolibre.com/sites/MLA/search?q=" + query

	// Ejecuto request
	resp, err := http.Get(urlML)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Valido Status Code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received error status code")
		return nil
	}

	// Leer respuesta
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Asignar los datos en la estructura
	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return response.Results
}

// Request API Mercado Libre Search
// Leer resultados
// Mostrar resultados
