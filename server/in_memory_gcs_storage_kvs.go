package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/minedia/orca-graphql-server/service"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	m := map[string]interface{}{
		"a": "world",
	}

	r.Get("/read/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")

		if m[key] != nil {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]interface{}{
				"status": "success",
				"value":  m[key],
			})
			return
		}

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]interface{}{
			"status": "notFound",
			"value":  "null",
		})
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
			})
			return
		}
		render.JSON(w, r, params)
	})

	r.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		d, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("failed to marshal: %v\n", err)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		err = service.UploadToGcs(r.Context(),
			"orca-cache", // GCSのバケット名
			"cache.json", // Storageに保存されるファイル名
			d,
		)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "success")
	})

	r.Get("/load", func(w http.ResponseWriter, r *http.Request) {
		
	})

	fmt.Printf("server is running on port %v\n", 8080)
	http.ListenAndServe(":8080", r)
}
