// Paquete
package main

// Dependencias
import (
	"encoding/json"
	"errors"
	"fmt"
)

// Estructuras
type Producto struct {
	Nombre     string
	Precio     float64
	Categoria  string
	Disponible bool
	Color      string
}

// Funciones
func main() {
	tv := Producto{
		Nombre:     "TV 50 pulgadas",
		Precio:     100000,
		Categoria:  "Tecnologia",
		Disponible: true,
		Color:      "Negro",
	}

	fmt.Println("Intro Go")

	var err error
	err = mostrar(tv)
	if err != nil {
		fmt.Println(err.Error())
	}

	aplicarDescuento(&tv)

	err = mostrar(tv)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func mostrar(prod Producto) error {
	if prod.Nombre == "" {
		return errors.New("El titulo no puede estar vacio")
	}

	prodJson, err := json.Marshal(prod)
	if err != nil {
		return err
	}
	prodString := string(prodJson)
	fmt.Println(prodString)
	return nil
}

// Implementar una función aplicarDescuento()
// Recibe un producto y retorna uno nuevo
// a la mitad de precio
// Devuelve un error si el precio es 0 antes de aplicar el descuento
func aplicarDescuento(product *Producto) {
	product.Precio = product.Precio * 0.5
}
