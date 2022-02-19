// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Encode implements json.Marshaler.
func (s Book) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.ID.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.ID.Set {
			e.RawStr("\"id\"" + ":")
			s.ID.Encode(e)
		}
	}
	{
		if s.MediaID.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.MediaID.Set {
			e.RawStr("\"media_id\"" + ":")
			s.MediaID.Encode(e)
		}
	}
	{
		if s.Images.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Images.Set {
			e.RawStr("\"images\"" + ":")
			s.Images.Encode(e)
		}
	}
	{
		if s.Title.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Title.Set {
			e.RawStr("\"title\"" + ":")
			s.Title.Encode(e)
		}
	}
	{
		if s.Tags != nil {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Tags != nil {
			e.RawStr("\"tags\"" + ":")
			e.ArrStart()
			if len(s.Tags) >= 1 {
				// Encode first element without comma.
				{
					elem := s.Tags[0]
					elem.Encode(e)
				}
				for _, elem := range s.Tags[1:] {
					e.Comma()
					elem.Encode(e)
				}
			}
			e.ArrEnd()
		}
	}
	{
		if s.Scanlator.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Scanlator.Set {
			e.RawStr("\"scanlator\"" + ":")
			s.Scanlator.Encode(e)
		}
	}
	{
		if s.UploadDate.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.UploadDate.Set {
			e.RawStr("\"upload_date\"" + ":")
			s.UploadDate.Encode(e)
		}
	}
	{
		if s.NumPages.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.NumPages.Set {
			e.RawStr("\"num_pages\"" + ":")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.NumFavorites.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.NumFavorites.Set {
			e.RawStr("\"num_favorites\"" + ":")
			s.NumFavorites.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfBook = [9]string{
	0: "id",
	1: "media_id",
	2: "images",
	3: "title",
	4: "tags",
	5: "scanlator",
	6: "upload_date",
	7: "num_pages",
	8: "num_favorites",
}

// Decode decodes Book from json.
func (s *Book) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Book to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "media_id":
			if err := func() error {
				s.MediaID.Reset()
				if err := s.MediaID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"media_id\"")
			}
		case "images":
			if err := func() error {
				s.Images.Reset()
				if err := s.Images.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"images\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "tags":
			if err := func() error {
				s.Tags = make([]Tag, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Tag
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Tags = append(s.Tags, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tags\"")
			}
		case "scanlator":
			if err := func() error {
				s.Scanlator.Reset()
				if err := s.Scanlator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"scanlator\"")
			}
		case "upload_date":
			if err := func() error {
				s.UploadDate.Reset()
				if err := s.UploadDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"upload_date\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "num_favorites":
			if err := func() error {
				s.NumFavorites.Reset()
				if err := s.NumFavorites.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_favorites\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Book")
	}

	return nil
}

// Encode implements json.Marshaler.
func (s Image) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.T.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.T.Set {
			e.RawStr("\"t\"" + ":")
			s.T.Encode(e)
		}
	}
	{
		if s.W.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.W.Set {
			e.RawStr("\"w\"" + ":")
			s.W.Encode(e)
		}
	}
	{
		if s.H.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.H.Set {
			e.RawStr("\"h\"" + ":")
			s.H.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfImage = [3]string{
	0: "t",
	1: "w",
	2: "h",
}

// Decode decodes Image from json.
func (s *Image) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Image to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "t":
			if err := func() error {
				s.T.Reset()
				if err := s.T.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"t\"")
			}
		case "w":
			if err := func() error {
				s.W.Reset()
				if err := s.W.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"w\"")
			}
		case "h":
			if err := func() error {
				s.H.Reset()
				if err := s.H.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"h\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Image")
	}

	return nil
}

// Encode implements json.Marshaler.
func (s Images) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.Pages != nil {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Pages != nil {
			e.RawStr("\"pages\"" + ":")
			e.ArrStart()
			if len(s.Pages) >= 1 {
				// Encode first element without comma.
				{
					elem := s.Pages[0]
					elem.Encode(e)
				}
				for _, elem := range s.Pages[1:] {
					e.Comma()
					elem.Encode(e)
				}
			}
			e.ArrEnd()
		}
	}
	{
		if s.Cover.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Cover.Set {
			e.RawStr("\"cover\"" + ":")
			s.Cover.Encode(e)
		}
	}
	{
		if s.Thumbnail.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Thumbnail.Set {
			e.RawStr("\"thumbnail\"" + ":")
			s.Thumbnail.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfImages = [3]string{
	0: "pages",
	1: "cover",
	2: "thumbnail",
}

// Decode decodes Images from json.
func (s *Images) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Images to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pages":
			if err := func() error {
				s.Pages = make([]Image, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Image
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Pages = append(s.Pages, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pages\"")
			}
		case "cover":
			if err := func() error {
				s.Cover.Reset()
				if err := s.Cover.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover\"")
			}
		case "thumbnail":
			if err := func() error {
				s.Thumbnail.Reset()
				if err := s.Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thumbnail\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Images")
	}

	return nil
}

// Encode encodes Image as json.
func (o OptImage) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Image from json.
func (o *OptImage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptImage to nil")
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptImage", d.Next())
	}
}

// Encode encodes Images as json.
func (o OptImages) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Images from json.
func (o *OptImages) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptImages to nil")
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptImages", d.Next())
	}
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	switch d.Next() {
	case jx.Number:
		o.Set = true
		v, err := d.Int()
		if err != nil {
			return err
		}
		o.Value = int(v)
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptInt", d.Next())
	}
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptString", d.Next())
	}
}

// Encode encodes TagType as json.
func (o OptTagType) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes TagType from json.
func (o *OptTagType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTagType to nil")
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptTagType", d.Next())
	}
}

// Encode encodes Title as json.
func (o OptTitle) Encode(e *jx.Writer) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Title from json.
func (o *OptTitle) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTitle to nil")
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("unexpected type %q while reading OptTitle", d.Next())
	}
}

// Encode encodes SearchByTagIDOKApplicationJSON as json.
func (s SearchByTagIDOKApplicationJSON) Encode(e *jx.Writer) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	if len(unwrapped) >= 1 {
		// Encode first element without comma.
		{
			elem := unwrapped[0]
			elem.Encode(e)
		}
		for _, elem := range unwrapped[1:] {
			e.Comma()
			elem.Encode(e)
		}
	}
	e.ArrEnd()
}

// Decode decodes SearchByTagIDOKApplicationJSON from json.
func (s *SearchByTagIDOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchByTagIDOKApplicationJSON to nil")
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = make([]SearchResponse, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchByTagIDOKApplicationJSON(unwrapped)
	return nil
}

// Encode encodes SearchOKApplicationJSON as json.
func (s SearchOKApplicationJSON) Encode(e *jx.Writer) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	if len(unwrapped) >= 1 {
		// Encode first element without comma.
		{
			elem := unwrapped[0]
			elem.Encode(e)
		}
		for _, elem := range unwrapped[1:] {
			e.Comma()
			elem.Encode(e)
		}
	}
	e.ArrEnd()
}

// Decode decodes SearchOKApplicationJSON from json.
func (s *SearchOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchOKApplicationJSON to nil")
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = make([]SearchResponse, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchOKApplicationJSON(unwrapped)
	return nil
}

// Encode implements json.Marshaler.
func (s SearchResponse) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.Result != nil {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Result != nil {
			e.RawStr("\"result\"" + ":")
			e.ArrStart()
			if len(s.Result) >= 1 {
				// Encode first element without comma.
				{
					elem := s.Result[0]
					elem.Encode(e)
				}
				for _, elem := range s.Result[1:] {
					e.Comma()
					elem.Encode(e)
				}
			}
			e.ArrEnd()
		}
	}
	{
		if s.NumPages.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.NumPages.Set {
			e.RawStr("\"num_pages\"" + ":")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.PerPage.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.PerPage.Set {
			e.RawStr("\"per_page\"" + ":")
			s.PerPage.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfSearchResponse = [3]string{
	0: "result",
	1: "num_pages",
	2: "per_page",
}

// Decode decodes SearchResponse from json.
func (s *SearchResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			if err := func() error {
				s.Result = make([]Book, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Book
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "per_page":
			if err := func() error {
				s.PerPage.Reset()
				if err := s.PerPage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"per_page\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchResponse")
	}

	return nil
}

// Encode implements json.Marshaler.
func (s Tag) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.ID.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.ID.Set {
			e.RawStr("\"id\"" + ":")
			s.ID.Encode(e)
		}
	}
	{
		if s.Type.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Type.Set {
			e.RawStr("\"type\"" + ":")
			s.Type.Encode(e)
		}
	}
	{
		if s.Name.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Name.Set {
			e.RawStr("\"name\"" + ":")
			s.Name.Encode(e)
		}
	}
	{
		if s.URL.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.URL.Set {
			e.RawStr("\"url\"" + ":")
			s.URL.Encode(e)
		}
	}
	{
		if s.Count.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Count.Set {
			e.RawStr("\"count\"" + ":")
			s.Count.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTag = [5]string{
	0: "id",
	1: "type",
	2: "name",
	3: "url",
	4: "count",
}

// Decode decodes Tag from json.
func (s *Tag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Tag to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "count":
			if err := func() error {
				s.Count.Reset()
				if err := s.Count.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Tag")
	}

	return nil
}

// Encode encodes TagType as json.
func (s TagType) Encode(e *jx.Writer) {
	e.Str(string(s))
}

// Decode decodes TagType from json.
func (s *TagType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TagType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch TagType(v) {
	case TagTypeParody:
		*s = TagTypeParody
	case TagTypeCharacter:
		*s = TagTypeCharacter
	case TagTypeTag:
		*s = TagTypeTag
	case TagTypeArtist:
		*s = TagTypeArtist
	case TagTypeGroup:
		*s = TagTypeGroup
	case TagTypeCategory:
		*s = TagTypeCategory
	case TagTypeLanguage:
		*s = TagTypeLanguage
	default:
		*s = TagType(v)
	}

	return nil
}

// Encode implements json.Marshaler.
func (s Title) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if s.English.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.English.Set {
			e.RawStr("\"english\"" + ":")
			s.English.Encode(e)
		}
	}
	{
		if s.Japanese.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Japanese.Set {
			e.RawStr("\"japanese\"" + ":")
			s.Japanese.Encode(e)
		}
	}
	{
		if s.Pretty.Set {
			if !first {
				e.Comma()
			}
			first = false
		}
		if s.Pretty.Set {
			e.RawStr("\"pretty\"" + ":")
			s.Pretty.Encode(e)
		}
	}
	e.ObjEnd()
}

var jsonFieldsNameOfTitle = [3]string{
	0: "english",
	1: "japanese",
	2: "pretty",
}

// Decode decodes Title from json.
func (s *Title) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Title to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "english":
			if err := func() error {
				s.English.Reset()
				if err := s.English.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"english\"")
			}
		case "japanese":
			if err := func() error {
				s.Japanese.Reset()
				if err := s.Japanese.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"japanese\"")
			}
		case "pretty":
			if err := func() error {
				s.Pretty.Reset()
				if err := s.Pretty.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pretty\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Title")
	}

	return nil
}
