package main

import (
	"os"
)

var tinamarURL string = getEnv("TINAMAR_URL", "http://ligatinamar.com/category/once_veteranos_38b_2020")
var mongoURL string = getEnv("MONGO_URL", "mongodb://localhost:27017")
var executionPeriod string = getEnv("EXECUTION_PERIOD", "1")

func getEnv(name string, defValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defValue
}
