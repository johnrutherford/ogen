// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func (s *AnswerShippingQueryPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AnswerShippingQueryPostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AnswerShippingQueryPostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CallbackQuery) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Chat) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EncryptedPassportElement) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Game) Validate() error {
	var failures []validate.FieldError
	if s.Photo == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "photo",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetGameHighScoresPostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if s.Result == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "result",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetMyCommandsPostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if s.Result == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "result",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetStickerSetPostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUpdatesPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUpdatesPostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUpdatesPostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUpdatesPostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if s.Result == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "result",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUserProfilePhotosPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUserProfilePhotosPostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUserProfilePhotosPostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUserProfilePhotosPostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *InlineKeyboardMarkup) Validate() error {
	var failures []validate.FieldError
	if s.InlineKeyboard == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "inline_keyboard",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MaskPosition) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Message) Validate() error {
	var failures []validate.FieldError
	if s.Chat == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "chat",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MessageEntity) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PassportData) Validate() error {
	var failures []validate.FieldError
	if s.Data == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "data",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Poll) Validate() error {
	var failures []validate.FieldError
	if s.Options == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "options",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PollAnswer) Validate() error {
	var failures []validate.FieldError
	if s.OptionIds == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "option_ids",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendGamePostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendGamePostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendGamePostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendGamePostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoicePostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if s.Prices == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "prices",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoicePostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if s.Prices == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "prices",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoicePostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if s.Prices == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "prices",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoicePostResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetMyCommandsPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if s.Commands == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "commands",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetMyCommandsPostReqApplicationXWwwFormUrlencoded) Validate() error {
	var failures []validate.FieldError
	if s.Commands == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "commands",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetMyCommandsPostReqMultipartFormData) Validate() error {
	var failures []validate.FieldError
	if s.Commands == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "commands",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ShippingOption) Validate() error {
	var failures []validate.FieldError
	if s.Prices == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "prices",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Sticker) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *StickerSet) Validate() error {
	var failures []validate.FieldError
	if s.Stickers == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "stickers",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Update) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *UserProfilePhotos) Validate() error {
	var failures []validate.FieldError
	if s.Photos == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "photos",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
