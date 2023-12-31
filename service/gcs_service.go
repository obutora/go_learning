package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/minedia/orca-graphql-server/logger"
)

func UploadToGcs(ctx context.Context, bucket, path string, data []byte) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.Error("[GCS upload] failed to create client", err)
		return err
	}

	// bucketオブジェクトの作成
	b := client.Bucket(bucket)

	// bucket内のオブジェクトの作成
	obj := b.Object(path)

	// file upload
	wc := obj.NewWriter(ctx)
	_, err = wc.Write(data)
	if err != nil {
		logger.Error("[GCS upload] failed to write data", err)
		return err
	}

	if err := wc.Close(); err != nil {
		logger.Error("[GCS upload] failed to close writer", err)
		return err
	}

	log.Println("[GCS upload] file upload success")
	return nil
}

func ReadFromGcs(ctx context.Context, bucket, path string) (map[string]interface{}, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.Error("[GCS upload] failed to create client", err)
		return nil, err
	}
	// bucketオブジェクトの作成
	b := client.Bucket(bucket)
	// bucket内のオブジェクトの作成
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
