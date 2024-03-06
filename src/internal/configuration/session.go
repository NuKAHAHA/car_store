package configuration

import (
	"flag"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	SessionKeyUser = "user"
)

type SessionConfiguration struct {
	Secret         *string `yaml:"secret" json:"secret"`
	MongoDBURI     *string `yaml:"mongoDBURI" json:"mongoDBURI"`
	MongoDBName    *string `yaml:"mongoDBName" json:"mongoDBName"`
	CollectionName *string `yaml:"collectionName" json:"collectionName"`
}

func (sc SessionConfiguration) Validate() error {
	return validation.ValidateStruct(&sc,
		validation.Field(&sc.Secret, validation.NilOrNotEmpty),
		validation.Field(&sc.MongoDBURI, validation.NilOrNotEmpty),
		validation.Field(&sc.MongoDBName, validation.NilOrNotEmpty),
		validation.Field(&sc.CollectionName, validation.NilOrNotEmpty),
	)
}

func NewSessionConfiguration(secret, mongoDBURI, mongoDBName, collectionName *string) *SessionConfiguration {
	return &SessionConfiguration{
		Secret:         secret,
		MongoDBURI:     mongoDBURI,
		MongoDBName:    mongoDBName,
		CollectionName: collectionName,
	}
}

func GetSessionConfiguration() *SessionConfiguration {
	sessionSecret := flag.String("session-secret", "secret", "session secret")
	mongoDBURI := flag.String("mongo-db-uri", "mongodb+srv://kukaNku:kuka0404@cluster0.nkjmlah.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", "MongoDB connection URI")
	mongoDBName := flag.String("mongo-db-name", "your_database_name", "MongoDB database name")
	collectionName := flag.String("mongo-collection-name", "your_collection_name", "MongoDB collection name")

	return NewSessionConfiguration(sessionSecret, mongoDBURI, mongoDBName, collectionName)
}

func IsAuthenticated(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	user := session.Get(SessionKeyUser)
	if user == nil {
		return false
	}

	return true
}
