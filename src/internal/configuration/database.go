package configuration

import (
	"flag"

	validation "github.com/go-ozzo/ozzo-validation"
)

type DatabaseConfiguration struct {
	ConnectionURI *string `yaml:"connectionURI" json:"connectionURI"`
}

func (dc DatabaseConfiguration) Validate() error {
	return validation.ValidateStruct(&dc,
		validation.Field(&dc.ConnectionURI, validation.NilOrNotEmpty),
	)
}

func newDatabaseConfiguration(connectionURI *string) *DatabaseConfiguration {
	return &DatabaseConfiguration{
		ConnectionURI: connectionURI,
	}
}

func GetDatabaseConfiguration() *DatabaseConfiguration {
	connectionURI := flag.String("db-connection-uri", "mongodb+srv://kukaNku:kuka0404@cluster0.nkjmlah.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", "MongoDB connection URI")

	return newDatabaseConfiguration(connectionURI)
}
