// Package swagger
// Date: 2022/9/23 13:28
// Author: Amu
// Description:
package swagger

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/allenai/go-swaggerui"
)

type Swagger struct {
	FilePath string
}

func NewSwagger(options ...Option) *Swagger {
	swagger := &Swagger{}
	for _, option := range options {
		option(swagger)
	}
	return swagger
}

func (s *Swagger) Run() error {
	if s.FilePath == "" {
		return errors.New("swagger filepath 不能为空")
	}
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}
	// Redirect /api to /api/ when serving UI.
	http.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusFound)
	})

	// Serve Swagger UI using a relative path to the schema exposed below.
	http.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler("v1/openapi.json")))

	// Serve the JSON-encoded schema.
	http.HandleFunc("/swagger/v1/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}
