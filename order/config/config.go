package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int64 {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return int64(port)
}

func GetPaymentServiceUrl() string {
	return getEnvironmentValue("PAYMENT_SERVICE_URL")
}

func getEnvironmentValue(key string) string {
	k := os.Getenv(key)
	if k == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return k
}
