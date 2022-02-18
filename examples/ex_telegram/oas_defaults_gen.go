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
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// setDefaults set default value of fields.
func (s *BotCommandScopeAllChatAdministrators) setDefaults() {
	{
		val := string("all_chat_administrators")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeAllGroupChats) setDefaults() {
	{
		val := string("all_group_chats")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeAllPrivateChats) setDefaults() {
	{
		val := string("all_private_chats")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeChat) setDefaults() {
	{
		val := string("chat")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeChatAdministrators) setDefaults() {
	{
		val := string("chat_administrators")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeChatMember) setDefaults() {
	{
		val := string("chat_member")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *BotCommandScopeDefault) setDefaults() {
	{
		val := string("default")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *Error) setDefaults() {
	{
		val := bool(false)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *GetUpdates) setDefaults() {
	{
		val := int(0)

		s.Offset.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GetUserProfilePhotos) setDefaults() {
	{
		val := int(0)

		s.Offset.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultArticle) setDefaults() {
	{
		val := string("article")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultAudio) setDefaults() {
	{
		val := string("audio")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedAudio) setDefaults() {
	{
		val := string("audio")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedDocument) setDefaults() {
	{
		val := string("document")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedGif) setDefaults() {
	{
		val := string("gif")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedMpeg4Gif) setDefaults() {
	{
		val := string("mpeg4_gif")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedPhoto) setDefaults() {
	{
		val := string("photo")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedSticker) setDefaults() {
	{
		val := string("sticker")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedVideo) setDefaults() {
	{
		val := string("video")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedVoice) setDefaults() {
	{
		val := string("voice")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultContact) setDefaults() {
	{
		val := string("contact")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultDocument) setDefaults() {
	{
		val := string("document")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultGame) setDefaults() {
	{
		val := string("game")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultGif) setDefaults() {
	{
		val := string("gif")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultLocation) setDefaults() {
	{
		val := string("location")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultMpeg4Gif) setDefaults() {
	{
		val := string("mpeg4_gif")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultPhoto) setDefaults() {
	{
		val := string("photo")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultVenue) setDefaults() {
	{
		val := string("venue")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultVideo) setDefaults() {
	{
		val := string("video")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultVoice) setDefaults() {
	{
		val := string("voice")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InputMediaAnimation) setDefaults() {
	{
		val := string("animation")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InputMediaAudio) setDefaults() {
	{
		val := string("audio")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InputMediaDocument) setDefaults() {
	{
		val := string("document")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InputMediaPhoto) setDefaults() {
	{
		val := string("photo")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InputMediaVideo) setDefaults() {
	{
		val := string("video")

		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *MessageEntity) setDefaults() {
	{
		val := int(0)

		s.Offset = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorDataField) setDefaults() {
	{
		val := string("data")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFile) setDefaults() {
	{
		val := string("file")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFiles) setDefaults() {
	{
		val := string("files")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFrontSide) setDefaults() {
	{
		val := string("front_side")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorReverseSide) setDefaults() {
	{
		val := string("reverse_side")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorSelfie) setDefaults() {
	{
		val := string("selfie")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorTranslationFile) setDefaults() {
	{
		val := string("translation_file")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorTranslationFiles) setDefaults() {
	{
		val := string("translation_files")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorUnspecified) setDefaults() {
	{
		val := string("unspecified")

		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *Result) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfBotCommand) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfChatMember) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfGameHighScore) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfMessage) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfUpdate) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChat) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChatInviteLink) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChatMember) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultFile) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultInt) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultMessage) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultMessageId) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultPoll) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultString) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultUser) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultUserProfilePhotos) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultWebhookInfo) setDefaults() {
	{
		val := bool(true)

		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *UploadStickerFile) setDefaults() {
	{
		val := string("up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »")

		s.PNGSticker = val
	}
}
