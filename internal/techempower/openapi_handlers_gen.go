// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

func NewDBHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.DB(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := encodeDBResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewJSONHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.JSON(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := encodeJSONResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewQueriesHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeQueriesParams(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.Queries(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := encodeQueriesResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func respondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
}
