package database

import (
	"fmt"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IntializeMongoDB() {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		panic(err)
	}
}
