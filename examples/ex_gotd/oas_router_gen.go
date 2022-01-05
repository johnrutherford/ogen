// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
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
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	args := map[string]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleBanChatMemberRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if prefix := "a"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleAnswerCallbackQueryRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'd': // Prefix: "ddStickerToSet"
					if prefix := "ddStickerToSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: AddStickerToSet
					s.handleAddStickerToSetRequest(args, w, r)
					return
				case 'n': // Prefix: "nswer"
					if prefix := "nswer"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleAnswerInlineQueryRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "CallbackQuery"
						if prefix := "CallbackQuery"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: AnswerCallbackQuery
						s.handleAnswerCallbackQueryRequest(args, w, r)
						return
					case 'I': // Prefix: "InlineQuery"
						if prefix := "InlineQuery"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: AnswerInlineQuery
						s.handleAnswerInlineQueryRequest(args, w, r)
						return
					case 'P': // Prefix: "PreCheckoutQuery"
						if prefix := "PreCheckoutQuery"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: AnswerPreCheckoutQuery
						s.handleAnswerPreCheckoutQueryRequest(args, w, r)
						return
					case 'S': // Prefix: "ShippingQuery"
						if prefix := "ShippingQuery"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: AnswerShippingQuery
						s.handleAnswerShippingQueryRequest(args, w, r)
						return
					}
				case 'p': // Prefix: "pproveChatJoinRequest"
					if prefix := "pproveChatJoinRequest"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: ApproveChatJoinRequest
					s.handleApproveChatJoinRequestRequest(args, w, r)
					return
				}
			case 'b': // Prefix: "banChatMember"
				if prefix := "banChatMember"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: BanChatMember
				s.handleBanChatMemberRequest(args, w, r)
				return
			case 'c': // Prefix: "c"
				if prefix := "c"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreateChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'o': // Prefix: "opyMessage"
					if prefix := "opyMessage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: CopyMessage
					s.handleCopyMessageRequest(args, w, r)
					return
				case 'r': // Prefix: "reate"
					if prefix := "reate"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleCreateNewStickerSetRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "ChatInviteLink"
						if prefix := "ChatInviteLink"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: CreateChatInviteLink
						s.handleCreateChatInviteLinkRequest(args, w, r)
						return
					case 'N': // Prefix: "NewStickerSet"
						if prefix := "NewStickerSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: CreateNewStickerSet
						s.handleCreateNewStickerSetRequest(args, w, r)
						return
					}
				}
			case 'd': // Prefix: "de"
				if prefix := "de"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleDeleteChatPhotoRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'c': // Prefix: "clineChatJoinRequest"
					if prefix := "clineChatJoinRequest"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: DeclineChatJoinRequest
					s.handleDeclineChatJoinRequestRequest(args, w, r)
					return
				case 'l': // Prefix: "lete"
					if prefix := "lete"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleDeleteMessageRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "Chat"
						if prefix := "Chat"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleDeleteChatStickerSetRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'P': // Prefix: "Photo"
							if prefix := "Photo"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: DeleteChatPhoto
							s.handleDeleteChatPhotoRequest(args, w, r)
							return
						case 'S': // Prefix: "StickerSet"
							if prefix := "StickerSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: DeleteChatStickerSet
							s.handleDeleteChatStickerSetRequest(args, w, r)
							return
						}
					case 'M': // Prefix: "M"
						if prefix := "M"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleDeleteMyCommandsRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'e': // Prefix: "essage"
							if prefix := "essage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: DeleteMessage
							s.handleDeleteMessageRequest(args, w, r)
							return
						case 'y': // Prefix: "yCommands"
							if prefix := "yCommands"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: DeleteMyCommands
							s.handleDeleteMyCommandsRequest(args, w, r)
							return
						}
					case 'S': // Prefix: "StickerFromSet"
						if prefix := "StickerFromSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: DeleteStickerFromSet
						s.handleDeleteStickerFromSetRequest(args, w, r)
						return
					case 'W': // Prefix: "Webhook"
						if prefix := "Webhook"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: DeleteWebhook
						s.handleDeleteWebhookRequest(args, w, r)
						return
					}
				}
			case 'e': // Prefix: "e"
				if prefix := "e"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleExportChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'd': // Prefix: "dit"
					if prefix := "dit"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleEditMessageCaptionRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "ChatInviteLink"
						if prefix := "ChatInviteLink"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: EditChatInviteLink
						s.handleEditChatInviteLinkRequest(args, w, r)
						return
					case 'M': // Prefix: "Message"
						if prefix := "Message"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleEditMessageLiveLocationRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Caption"
							if prefix := "Caption"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: EditMessageCaption
							s.handleEditMessageCaptionRequest(args, w, r)
							return
						case 'L': // Prefix: "LiveLocation"
							if prefix := "LiveLocation"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: EditMessageLiveLocation
							s.handleEditMessageLiveLocationRequest(args, w, r)
							return
						case 'M': // Prefix: "Media"
							if prefix := "Media"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: EditMessageMedia
							s.handleEditMessageMediaRequest(args, w, r)
							return
						case 'R': // Prefix: "ReplyMarkup"
							if prefix := "ReplyMarkup"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: EditMessageReplyMarkup
							s.handleEditMessageReplyMarkupRequest(args, w, r)
							return
						case 'T': // Prefix: "Text"
							if prefix := "Text"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: EditMessageText
							s.handleEditMessageTextRequest(args, w, r)
							return
						}
					}
				case 'x': // Prefix: "xportChatInviteLink"
					if prefix := "xportChatInviteLink"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: ExportChatInviteLink
					s.handleExportChatInviteLinkRequest(args, w, r)
					return
				}
			case 'f': // Prefix: "forwardMessage"
				if prefix := "forwardMessage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: ForwardMessage
				s.handleForwardMessageRequest(args, w, r)
				return
			case 'g': // Prefix: "get"
				if prefix := "get"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleGetFileRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'C': // Prefix: "Chat"
					if prefix := "Chat"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetChatRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'A': // Prefix: "Administrators"
						if prefix := "Administrators"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: GetChatAdministrators
						s.handleGetChatAdministratorsRequest(args, w, r)
						return
					case 'M': // Prefix: "Member"
						if prefix := "Member"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleGetChatMemberRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Count"
							if prefix := "Count"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: GetChatMemberCount
							s.handleGetChatMemberCountRequest(args, w, r)
							return
						}
					}
				case 'F': // Prefix: "File"
					if prefix := "File"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: GetFile
					s.handleGetFileRequest(args, w, r)
					return
				case 'G': // Prefix: "GameHighScores"
					if prefix := "GameHighScores"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: GetGameHighScores
					s.handleGetGameHighScoresRequest(args, w, r)
					return
				case 'M': // Prefix: "M"
					if prefix := "M"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetMyCommandsRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'e': // Prefix: "e"
						if prefix := "e"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: GetMe
						s.handleGetMeRequest(args, w, r)
						return
					case 'y': // Prefix: "yCommands"
						if prefix := "yCommands"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: GetMyCommands
						s.handleGetMyCommandsRequest(args, w, r)
						return
					}
				case 'S': // Prefix: "StickerSet"
					if prefix := "StickerSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: GetStickerSet
					s.handleGetStickerSetRequest(args, w, r)
					return
				case 'U': // Prefix: "U"
					if prefix := "U"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetUserProfilePhotosRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'p': // Prefix: "pdates"
						if prefix := "pdates"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: GetUpdates
						s.handleGetUpdatesRequest(args, w, r)
						return
					case 's': // Prefix: "serProfilePhotos"
						if prefix := "serProfilePhotos"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: GetUserProfilePhotos
						s.handleGetUserProfilePhotosRequest(args, w, r)
						return
					}
				}
			case 'l': // Prefix: "leaveChat"
				if prefix := "leaveChat"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: LeaveChat
				s.handleLeaveChatRequest(args, w, r)
				return
			case 'p': // Prefix: "p"
				if prefix := "p"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePromoteChatMemberRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'i': // Prefix: "inChatMessage"
					if prefix := "inChatMessage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: PinChatMessage
					s.handlePinChatMessageRequest(args, w, r)
					return
				case 'r': // Prefix: "romoteChatMember"
					if prefix := "romoteChatMember"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: PromoteChatMember
					s.handlePromoteChatMemberRequest(args, w, r)
					return
				}
			case 'r': // Prefix: "re"
				if prefix := "re"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleRevokeChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 's': // Prefix: "strictChatMember"
					if prefix := "strictChatMember"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: RestrictChatMember
					s.handleRestrictChatMemberRequest(args, w, r)
					return
				case 'v': // Prefix: "vokeChatInviteLink"
					if prefix := "vokeChatInviteLink"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: RevokeChatInviteLink
					s.handleRevokeChatInviteLinkRequest(args, w, r)
					return
				}
			case 's': // Prefix: "s"
				if prefix := "s"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleStopMessageLiveLocationRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'e': // Prefix: "e"
					if prefix := "e"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSetChatAdministratorCustomTitleRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'n': // Prefix: "nd"
						if prefix := "nd"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleSendChatActionRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'A': // Prefix: "A"
							if prefix := "A"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendAudioRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'n': // Prefix: "nimation"
								if prefix := "nimation"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendAnimation
								s.handleSendAnimationRequest(args, w, r)
								return
							case 'u': // Prefix: "udio"
								if prefix := "udio"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendAudio
								s.handleSendAudioRequest(args, w, r)
								return
							}
						case 'C': // Prefix: "C"
							if prefix := "C"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendContactRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'h': // Prefix: "hatAction"
								if prefix := "hatAction"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendChatAction
								s.handleSendChatActionRequest(args, w, r)
								return
							case 'o': // Prefix: "ontact"
								if prefix := "ontact"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendContact
								s.handleSendContactRequest(args, w, r)
								return
							}
						case 'D': // Prefix: "D"
							if prefix := "D"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendDocumentRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'i': // Prefix: "ice"
								if prefix := "ice"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendDice
								s.handleSendDiceRequest(args, w, r)
								return
							case 'o': // Prefix: "ocument"
								if prefix := "ocument"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendDocument
								s.handleSendDocumentRequest(args, w, r)
								return
							}
						case 'G': // Prefix: "Game"
							if prefix := "Game"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SendGame
							s.handleSendGameRequest(args, w, r)
							return
						case 'I': // Prefix: "Invoice"
							if prefix := "Invoice"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SendInvoice
							s.handleSendInvoiceRequest(args, w, r)
							return
						case 'L': // Prefix: "Location"
							if prefix := "Location"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SendLocation
							s.handleSendLocationRequest(args, w, r)
							return
						case 'M': // Prefix: "Me"
							if prefix := "Me"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendMessageRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'd': // Prefix: "diaGroup"
								if prefix := "diaGroup"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendMediaGroup
								s.handleSendMediaGroupRequest(args, w, r)
								return
							case 's': // Prefix: "ssage"
								if prefix := "ssage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendMessage
								s.handleSendMessageRequest(args, w, r)
								return
							}
						case 'P': // Prefix: "P"
							if prefix := "P"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendPollRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'h': // Prefix: "hoto"
								if prefix := "hoto"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendPhoto
								s.handleSendPhotoRequest(args, w, r)
								return
							case 'o': // Prefix: "oll"
								if prefix := "oll"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendPoll
								s.handleSendPollRequest(args, w, r)
								return
							}
						case 'S': // Prefix: "Sticker"
							if prefix := "Sticker"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SendSticker
							s.handleSendStickerRequest(args, w, r)
							return
						case 'V': // Prefix: "V"
							if prefix := "V"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendVideoRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'e': // Prefix: "enue"
								if prefix := "enue"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendVenue
								s.handleSendVenueRequest(args, w, r)
								return
							case 'i': // Prefix: "ideo"
								if prefix := "ideo"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								if len(elem) == 0 {
									s.handleSendVideoRequest(args, w, r)
									return
								}
								switch elem[0] {
								case 'N': // Prefix: "Note"
									if prefix := "Note"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
										elem = elem[len(prefix):]
									} else {
										break
									}

									// Leaf: SendVideoNote
									s.handleSendVideoNoteRequest(args, w, r)
									return
								}
							case 'o': // Prefix: "oice"
								if prefix := "oice"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SendVoice
								s.handleSendVoiceRequest(args, w, r)
								return
							}
						}
					case 't': // Prefix: "t"
						if prefix := "t"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleSetGameScoreRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Chat"
							if prefix := "Chat"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSetChatDescriptionRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'A': // Prefix: "AdministratorCustomTitle"
								if prefix := "AdministratorCustomTitle"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetChatAdministratorCustomTitle
								s.handleSetChatAdministratorCustomTitleRequest(args, w, r)
								return
							case 'D': // Prefix: "Description"
								if prefix := "Description"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetChatDescription
								s.handleSetChatDescriptionRequest(args, w, r)
								return
							case 'P': // Prefix: "P"
								if prefix := "P"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								if len(elem) == 0 {
									s.handleSetChatPhotoRequest(args, w, r)
									return
								}
								switch elem[0] {
								case 'e': // Prefix: "ermissions"
									if prefix := "ermissions"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
										elem = elem[len(prefix):]
									} else {
										break
									}

									// Leaf: SetChatPermissions
									s.handleSetChatPermissionsRequest(args, w, r)
									return
								case 'h': // Prefix: "hoto"
									if prefix := "hoto"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
										elem = elem[len(prefix):]
									} else {
										break
									}

									// Leaf: SetChatPhoto
									s.handleSetChatPhotoRequest(args, w, r)
									return
								}
							case 'S': // Prefix: "StickerSet"
								if prefix := "StickerSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetChatStickerSet
								s.handleSetChatStickerSetRequest(args, w, r)
								return
							case 'T': // Prefix: "Title"
								if prefix := "Title"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetChatTitle
								s.handleSetChatTitleRequest(args, w, r)
								return
							}
						case 'G': // Prefix: "GameScore"
							if prefix := "GameScore"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SetGameScore
							s.handleSetGameScoreRequest(args, w, r)
							return
						case 'M': // Prefix: "MyCommands"
							if prefix := "MyCommands"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SetMyCommands
							s.handleSetMyCommandsRequest(args, w, r)
							return
						case 'P': // Prefix: "PassportDataErrors"
							if prefix := "PassportDataErrors"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SetPassportDataErrors
							s.handleSetPassportDataErrorsRequest(args, w, r)
							return
						case 'S': // Prefix: "Sticker"
							if prefix := "Sticker"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSetStickerSetThumbRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'P': // Prefix: "PositionInSet"
								if prefix := "PositionInSet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetStickerPositionInSet
								s.handleSetStickerPositionInSetRequest(args, w, r)
								return
							case 'S': // Prefix: "SetThumb"
								if prefix := "SetThumb"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Leaf: SetStickerSetThumb
								s.handleSetStickerSetThumbRequest(args, w, r)
								return
							}
						case 'W': // Prefix: "Webhook"
							if prefix := "Webhook"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: SetWebhook
							s.handleSetWebhookRequest(args, w, r)
							return
						}
					}
				case 't': // Prefix: "top"
					if prefix := "top"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleStopPollRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'M': // Prefix: "MessageLiveLocation"
						if prefix := "MessageLiveLocation"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: StopMessageLiveLocation
						s.handleStopMessageLiveLocationRequest(args, w, r)
						return
					case 'P': // Prefix: "Poll"
						if prefix := "Poll"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: StopPoll
						s.handleStopPollRequest(args, w, r)
						return
					}
				}
			case 'u': // Prefix: "u"
				if prefix := "u"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleUploadStickerFileRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'n': // Prefix: "n"
					if prefix := "n"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleUnpinAllChatMessagesRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'b': // Prefix: "banChatMember"
						if prefix := "banChatMember"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: UnbanChatMember
						s.handleUnbanChatMemberRequest(args, w, r)
						return
					case 'p': // Prefix: "pin"
						if prefix := "pin"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleUnpinChatMessageRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'A': // Prefix: "AllChatMessages"
							if prefix := "AllChatMessages"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: UnpinAllChatMessages
							s.handleUnpinAllChatMessagesRequest(args, w, r)
							return
						case 'C': // Prefix: "ChatMessage"
							if prefix := "ChatMessage"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: UnpinChatMessage
							s.handleUnpinChatMessageRequest(args, w, r)
							return
						}
					}
				case 'p': // Prefix: "ploadStickerFile"
					if prefix := "ploadStickerFile"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					// Leaf: UploadStickerFile
					s.handleUploadStickerFileRequest(args, w, r)
					return
				}
			}
		}
	}
	s.notFound(w, r)
}
