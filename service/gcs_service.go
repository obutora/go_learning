package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
)

// func credentialsOption() option.ClientOption {
// 	return option.WithCredentialsFile("kvs-test-account.json")
// }

func UploadToGcs(ctx context.Context, bucket, path string, data []byte) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("[GCS upload] failed to create client: %v", err)
		return err
	}

	b := client.Bucket(bucket)
	obj := b.Object(path)

	wc := obj.NewWriter(ctx)
	_, err = wc.Write(data)
	if err != nil {
		fmt.Printf("[GCS upload] failed to write data: %v", err)
		return err
	}

	if err := wc.Close(); err != nil {
		fmt.Printf("[GCS upload] failed to close writer: %v", err)
		return err
	}

	log.Println("[GCS upload] file upload success")
	return nil
}

func ReadFromGcs(ctx context.Context, bucket, path string) (map[string]interface{}, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("[GCS upload] failed to create client: %v", err)
		return nil, err
	}
	b := client.Bucket(bucket)
	obj := b.Object(path)

	reader, err := obj.NewReader(ctx)
	if err != nil {
		fmt.Printf("Failed to create object reader: %v", err)
		return nil, err
	}
	defer reader.Close()

	d, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Failed to read object: %v", err)
		return nil, err
	}

	var m map[string]interface{}
	json.Unmarshal(d, &m)

	fmt.Printf("Object contents: %s\n", d)
	return m, nil
}
