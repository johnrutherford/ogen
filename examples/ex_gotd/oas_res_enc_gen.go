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

func encodeAddStickerToSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAnswerCallbackQueryResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAnswerInlineQueryResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAnswerPreCheckoutQueryResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAnswerShippingQueryResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeApproveChatJoinRequestResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeBanChatMemberResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeBanChatSenderChatResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeCloseResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeCopyMessageResponse(response ResultMessageId, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeCreateChatInviteLinkResponse(response ResultChatInviteLink, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeCreateNewStickerSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeclineChatJoinRequestResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteChatPhotoResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteChatStickerSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteMessageResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteMyCommandsResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteStickerFromSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeDeleteWebhookResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditChatInviteLinkResponse(response ResultChatInviteLink, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditMessageCaptionResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditMessageLiveLocationResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditMessageMediaResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditMessageReplyMarkupResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeEditMessageTextResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeExportChatInviteLinkResponse(response ResultString, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeForwardMessageResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetChatResponse(response ResultChat, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetChatAdministratorsResponse(response ResultArrayOfChatMember, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetChatMemberResponse(response ResultChatMember, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetChatMemberCountResponse(response ResultInt, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetFileResponse(response ResultFile, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetGameHighScoresResponse(response ResultArrayOfGameHighScore, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetMeResponse(response ResultUser, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetMyCommandsResponse(response ResultArrayOfBotCommand, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetStickerSetResponse(response ResultStickerSet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetUpdatesResponse(response ResultArrayOfUpdate, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetUserProfilePhotosResponse(response ResultUserProfilePhotos, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetWebhookInfoResponse(response ResultWebhookInfo, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeLeaveChatResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeLogOutResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePinChatMessageResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePromoteChatMemberResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeRestrictChatMemberResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeRevokeChatInviteLinkResponse(response ResultChatInviteLink, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendAnimationResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendAudioResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendChatActionResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendContactResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendDiceResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendDocumentResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendGameResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendInvoiceResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendLocationResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendMediaGroupResponse(response ResultArrayOfMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendMessageResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendPhotoResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendPollResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendStickerResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendVenueResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendVideoResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendVideoNoteResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSendVoiceResponse(response ResultMessage, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatAdministratorCustomTitleResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatDescriptionResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatPermissionsResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatPhotoResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatStickerSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetChatTitleResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetGameScoreResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetMyCommandsResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetPassportDataErrorsResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetStickerPositionInSetResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetStickerSetThumbResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetWebhookResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeStopMessageLiveLocationResponse(response ResultMessageOrBoolean, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeStopPollResponse(response ResultPoll, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeUnbanChatMemberResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeUnbanChatSenderChatResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeUnpinAllChatMessagesResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeUnpinChatMessageResponse(response Result, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeUploadStickerFileResponse(response ResultFile, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
func encodeErrorResponse(response ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
