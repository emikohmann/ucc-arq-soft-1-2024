package main

import "errors"

func main() {

}

func division(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("b no puede ser 0")
	}
	return a / b, nil
}
