package configuration

import "flag"

type Configuration struct {
	DatabaseConfiguration *DatabaseConfiguration
	SessionConfiguration  *SessionConfiguration
}

func NewConfiguration() (*Configuration, error) {
	var err error

	configuration := &Configuration{
		DatabaseConfiguration: GetDatabaseConfiguration(),
		SessionConfiguration:  GetSessionConfiguration(),
	}

	flag.Parse()

	if err = configuration.DatabaseConfiguration.Validate(); err != nil {
		return nil, err
	}

	if err = configuration.SessionConfiguration.Validate(); err != nil {
		return nil, err
	}

	return configuration, nil
}

func (c *Configuration) GetMongoDBURI() string {
	return *c.DatabaseConfiguration.ConnectionURI
}

func (c *Configuration) GetMongoDBName() string {
	return *c.SessionConfiguration.MongoDBName
}

func (c *Configuration) GetMongoDBCollectionName() string {
	return *c.SessionConfiguration.CollectionName
}
