package controller

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-services/graphql-files/src"
)

// FilesHandler ...
func FilesHandler(r *mux.Router, bucket string) error {

	r.HandleFunc("/{uid}", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("access_token")
		contentDisposition := r.URL.Query().Get("content-disposition")
		uid := mux.Vars(r)["uid"]

		ctx := context.Background()
		file, err := src.FetchFile(ctx, uid, "Bearer "+token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if file == nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		s3Object, err := src.GetS3Object(src.GetS3ObjectConfig{
			Bucket: bucket,
			Key:    uid,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer s3Object.Body.Close()

		w.Header().Set("content-type", file.ContentType)
		w.Header().Set("content-size", string(file.Size))
		if contentDisposition != "" {
			w.Header().Set("content-disposition", contentDisposition)
		}
		io.Copy(w, s3Object.Body)
	}).Methods("GET")

	return nil
}
