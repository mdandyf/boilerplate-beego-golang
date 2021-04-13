package services

import (
	"boilerplate-golang/models"
)

func GetMainService() (models.MainDataModel, error) {
	return models.MainDataModel{
		ApplicationName:    "Boilerplate-Golang",
		ApplicationVersion: "1.0.0",
	}, nil
}
