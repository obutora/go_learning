package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learn/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const (
	Bucket = "kvs-test"
	Path   = "cache.json"
)

// In-memory data store.
var m = map[string]interface{}{}

// Create channel to listen for signals.
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func init() {
	ctx := context.Background()
	d, err := service.ReadFromGcs(ctx, Bucket, Path)
	if err != nil {
		fmt.Printf("failed to load from GCS: %v\n", err)
	}
	m = d
}

//NOTE: CloudRunのライフサイクル
// https://cloud.google.com/blog/ja/products/serverless/lifecycle-container-cloud-run

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/read/{key}", func(w http.ResponseWriter, r *http.Request) {
		// 計測開始
		s := time.Now()

		key := chi.URLParam(r, "key")

		if m[key] != nil {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]interface{}{
				"status": "success",
				"value":  m[key],
			})
			// 経過時間を出力
			fmt.Printf("process time: %s\n", time.Since(s))
			return
		}

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]interface{}{
			"status": "notFound",
			"value":  "null",
		})
		// 経過時間を出力
		fmt.Printf("process time: %s\n", time.Since(s))

	})

	// NOTE: POST bodyに{"key": "hoge", "value": "fuga"}　を入れて実行する
	r.Post("/write", func(w http.ResponseWriter, r *http.Request) {
		params := map[string]interface{}{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&params); err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
		}
		defer r.Body.Close()

		key := params["key"].(string)
		value := params["value"]

		if key != "" && value != nil {
			m[key] = value
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]interface{}{
				"status": "success",
				"value":  value,
			})
		}
	})

	go func() {
		fmt.Printf("server is running on port %v\n", 8080)
		http.ListenAndServe(":8080", r)
	}()

	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Receive output from signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught", sig)

	// Timeout if waiting for connections to return idle.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := saveGcs(ctx)
	if err != nil {
		fmt.Printf("failed to save GCS: %v\n", err)
	}

}

func saveGcs(ctx context.Context) error {
	s := time.Now()

	// GCSのデータとインメモリのデータをマージする
	gcsData, err := service.ReadFromGcs(ctx, Bucket, Path)
	if err != nil {
		fmt.Printf("failed to load from GCS: %v\n", err)
		return err
	}

	// GCSのデータをインメモリのデータで上書きする
	for k, v := range m {
		gcsData[k] = v
	}

	j, err := json.Marshal(gcsData)
	if err != nil {
		fmt.Printf("failed to marshal: %v\n", err)
	}

	err = service.UploadToGcs(ctx, Bucket, Path, j)
	if err != nil {
		fmt.Printf("failed to upload to GCS: %v\n", err)
		return err
	}
	fmt.Printf("save gcs execution time: %v\n", time.Since(s))
	return nil
}
