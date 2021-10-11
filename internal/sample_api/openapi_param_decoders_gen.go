// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/uri"
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
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
)

func decodeFoobarGetParams(r *http.Request) (FoobarGetParams, error) {
	var params FoobarGetParams
	// Decode param 'inlinedParam' located in 'Query'.
	if err := func() error {
		values, ok := r.URL.Query()["inlinedParam"]
		if !ok {
			return fmt.Errorf("query parameter 'inlinedParam' not specified")
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt64()
		if err != nil {
			return err
		}
		params.InlinedParam = int64(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param 'skip' located in 'Query'.
	if err := func() error {
		values, ok := r.URL.Query()["skip"]
		if !ok {
			return fmt.Errorf("query parameter 'skip' not specified")
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt32()
		if err != nil {
			return err
		}
		params.Skip = int32(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodePetGetParams(r *http.Request) (PetGetParams, error) {
	var params PetGetParams
	// Decode param 'petID' located in 'Query'.
	if err := func() error {
		values, ok := r.URL.Query()["petID"]
		if !ok {
			return fmt.Errorf("query parameter 'petID' not specified")
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt64()
		if err != nil {
			return err
		}
		params.PetID = int64(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param 'x-tags' located in 'Header'.
	if err := func() error {
		param := r.Header.Values("x-tags")

		if len(param) == 0 {
			return nil
		}

		v, err := conv.ToUUIDArray(param)
		if err != nil {
			return fmt.Errorf("parse header param 'x-tags': %w", err)
		}

		params.XTags = []uuid.UUID(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param 'x-scope' located in 'Header'.
	if err := func() error {
		param := r.Header.Values("x-scope")

		if len(param) == 0 {
			return fmt.Errorf("header parameter 'x-scope' not specified")
		}

		v, err := conv.ToStringArray(param)
		if err != nil {
			return fmt.Errorf("parse header param 'x-scope': %w", err)
		}

		params.XScope = []string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param 'token' located in 'Cookie'.
	if err := func() error {
		c, err := r.Cookie("token")
		if err != nil {
			return fmt.Errorf("get cookie 'token': %w", err)
		}

		param := c.Value
		if len(param) == 0 {
			return fmt.Errorf("cookie parameter 'token' not specified")
		}

		v, err := conv.ToString(param)
		if err != nil {
			return fmt.Errorf("parse cookie param 'token': %w", err)
		}

		params.Token = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodePetGetByNameParams(r *http.Request) (PetGetByNameParams, error) {
	var params PetGetByNameParams
	// Decode param 'name' located in 'Path'.
	if err := func() error {
		param := chi.URLParam(r, "name")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'name' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "name",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeString()
		if err != nil {
			return err
		}
		params.Name = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}
