package database

import (
	"context"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"

	"google.golang.org/api/option"
)

var FirebaseBucket *storage.BucketHandle

func SetupFirebaseStorage() {
	opt := option.WithCredentialsFile("udrio-firebasesdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Error("fail to initialize firabase app: ", err)
		return
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Error("fail to initialize firebase storage client: ", err)
		return
	}

	FirebaseBucket, err = client.DefaultBucket()
	if err != nil {
		log.Error("fail to initialize firebase storage bucket: ", err)
		return
	}
}
