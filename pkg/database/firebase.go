package database

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google.golang.org/api/option"
)

var FirebaseBucket *storage.BucketHandle

func CreateConfig() {
	conf := []byte(viper.GetString("FIREBASE_CONFIG"))
	if err := os.WriteFile("udrio-firebasesdk.json", conf, os.ModePerm); err != nil {
		panic(err)
	}
}

func SetupFirebaseStorage() {
	CreateConfig()
	opt := option.WithCredentialsFile("udrio-firebasesdk.json")
	client, err := storage.NewClient(context.Background(), opt)
	if err != nil {
		log.Error("fail to initialize storage client: ", err.Error())
		panic(err.Error())
	}

	FirebaseBucket = client.Bucket(viper.GetString("FIREBASE_BUCKET"))
}
