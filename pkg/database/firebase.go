package database

import (
	"context"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google.golang.org/api/option"
)

var FirebaseBucket *storage.BucketHandle

func SetupFirebaseStorage() {
	opt := option.WithCredentialsFile("udrio-firebasesdk.json")
	client, err := storage.NewClient(context.Background(), opt)
	if err != nil {
		log.Error("fail to initialize storage client: ", err.Error())
		panic(err.Error())
	}

	FirebaseBucket = client.Bucket(viper.GetString("FIREBASE_BUCKET"))
}
