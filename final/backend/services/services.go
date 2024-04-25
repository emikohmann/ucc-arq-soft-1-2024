package services

import "final/models"

func GetResult(input string) models.Response {
	return models.Response{
		Message: "Recibi este mensaje: " + input,
	}
}
