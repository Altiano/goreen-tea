package shared

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// static
var (
	config Config
	once   sync.Once
)

// Config .
type Config struct {
	AppPort     string `envconfig:"APP_PORT" required:"true"`
	MongoURI    string `envconfig:"MONGO_URI" required:"true"`
	MongoDBName string `envconfig:"MONGO_DB_NAME" required:"true"`
	JeagerUrl   string `envconfig:"JEAGER_URL"`
}

// NewConfig .
func NewConfig() Config {
	once.Do(func() {
		filename := EnvFilename

		file, err := os.Stat(filename)

		if err != nil && os.IsNotExist(err) == false {
			panic(err)
		}

		if file != nil {
			if err := godotenv.Load(filename); err != nil {
				panic(err)
			}
		}

		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})

	return config
}
