package database

import (
	"context"

	"cloud.google.com/go/storage"

	"google.golang.org/api/option"
)

func SetupFirebaseStorage(sdkPath, bucket string) (*storage.BucketHandle, error) {
	opt := option.WithCredentialsFile(sdkPath)
	client, err := storage.NewClient(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	b := client.Bucket(bucket)
	return b, nil
}
