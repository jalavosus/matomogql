package matomo

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

const (
	envApiKey   = "MATOMO_API_KEY"
	envEndpoint = "MATOMO_ENDPOINT"
)

type envData struct {
	apiKey   string
	endpoint string
}

var envOnce = sync.OnceValue(func() envData {
	data := envData{}

	if apiKey, ok := os.LookupEnv(envApiKey); !ok {
		panic(ErrorMissingAPIKey)
	} else {
		data.apiKey = apiKey
	}

	if endpoint, ok := os.LookupEnv(envEndpoint); !ok {
		panic(ErrorMissingEndpoint)
	} else {
		data.endpoint = endpoint
	}

	return data
})

func GetEnv() (apiKey, endpoint string) {
	data := envOnce()
	return data.apiKey, data.endpoint
}
