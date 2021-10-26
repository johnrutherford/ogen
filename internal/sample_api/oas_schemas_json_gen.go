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

// WriteJSON implements json.Marshaler.
func (s FoobarGetResNotFound) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarGetResNotFound json value to io.Writer.
func (s FoobarGetResNotFound) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarGetResNotFound json value from io.Reader.
func (s *FoobarGetResNotFound) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarGetResNotFound from json stream.
func (s *FoobarGetResNotFound) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostResDefaultApplicationJSON) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("code")
	j.WriteInt64(s.Code)

	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPostResDefaultApplicationJSON json value to io.Writer.
func (s FoobarPostResDefaultApplicationJSON) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPostResDefaultApplicationJSON json value from io.Reader.
func (s *FoobarPostResDefaultApplicationJSON) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPostResDefaultApplicationJSON from json stream.
func (s *FoobarPostResDefaultApplicationJSON) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "code":
			if err := func() error {
				s.Code = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "message":
			if err := func() error {
				s.Message = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostResDefaultApplicationJSONStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPostResDefaultApplicationJSONStatusCode json value to io.Writer.
func (s FoobarPostResDefaultApplicationJSONStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPostResDefaultApplicationJSONStatusCode json value from io.Reader.
func (s *FoobarPostResDefaultApplicationJSONStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPostResDefaultApplicationJSONStatusCode from json stream.
func (s *FoobarPostResDefaultApplicationJSONStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostResNotFound) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPostResNotFound json value to io.Writer.
func (s FoobarPostResNotFound) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPostResNotFound json value from io.Reader.
func (s *FoobarPostResNotFound) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPostResNotFound from json stream.
func (s *FoobarPostResNotFound) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutResDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPutResDefault json value to io.Writer.
func (s FoobarPutResDefault) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPutResDefault json value from io.Reader.
func (s *FoobarPutResDefault) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPutResDefault from json stream.
func (s *FoobarPutResDefault) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutResDefaultStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPutResDefaultStatusCode json value to io.Writer.
func (s FoobarPutResDefaultStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPutResDefaultStatusCode json value from io.Reader.
func (s *FoobarPutResDefaultStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPutResDefaultStatusCode from json stream.
func (s *FoobarPutResDefaultStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON writes json value of string to json stream.
func (o OptNilString) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptNilString) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		o.Value = string(i.ReadString())
		return i.Error
	case json.NilValue:
		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptNilString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON implements json.Marshaler.
func (s PetGetResDefaultApplicationJSON) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetResDefaultApplicationJSON json value to io.Writer.
func (s PetGetResDefaultApplicationJSON) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetResDefaultApplicationJSON json value from io.Reader.
func (s *PetGetResDefaultApplicationJSON) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetResDefaultApplicationJSON from json stream.
func (s *PetGetResDefaultApplicationJSON) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			if err := func() error {
				s.Message = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetResDefaultApplicationJSONStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetResDefaultApplicationJSONStatusCode json value to io.Writer.
func (s PetGetResDefaultApplicationJSONStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetResDefaultApplicationJSONStatusCode json value from io.Reader.
func (s *PetGetResDefaultApplicationJSONStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetResDefaultApplicationJSONStatusCode from json stream.
func (s *PetGetResDefaultApplicationJSONStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetResOKApplicationJSON) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("birthday")
	json.WriteDate(j, s.Birthday)

	more.More()
	j.WriteObjectField("id")
	j.WriteInt64(s.ID)

	more.More()
	j.WriteObjectField("ip")
	json.WriteIP(j, s.IP)

	more.More()
	j.WriteObjectField("ip_v4")
	json.WriteIP(j, s.IPV4)

	more.More()
	j.WriteObjectField("ip_v6")
	json.WriteIP(j, s.IPV6)

	more.More()
	j.WriteObjectField("kind")
	s.Kind.WriteJSON(j)

	more.More()
	j.WriteObjectField("name")
	j.WriteString(s.Name)

	more.More()
	j.WriteObjectField("next")
	s.Next.WriteJSON(j)

	if s.Nickname.Set {
		more.More()
		j.WriteObjectField("nickname")
		s.Nickname.WriteJSON(j)
	}

	if s.NullStr.Set {
		more.More()
		j.WriteObjectField("nullStr")
		s.NullStr.WriteJSON(j)
	}

	if s.Primary != nil {
		more.More()
		j.WriteObjectField("primary")
		s.Primary.WriteJSON(j)
	}

	more.More()
	j.WriteObjectField("rate")
	json.WriteDuration(j, s.Rate)

	more.More()
	j.WriteObjectField("tag")
	json.WriteUUID(j, s.Tag)

	more.More()
	j.WriteObjectField("testDate")
	json.WriteDate(j, s.TestDate)

	more.More()
	j.WriteObjectField("testDateTime")
	json.WriteDateTime(j, s.TestDateTime)

	more.More()
	j.WriteObjectField("testDuration")
	json.WriteDuration(j, s.TestDuration)

	more.More()
	j.WriteObjectField("testFloat1")
	j.WriteFloat64(s.TestFloat1)

	more.More()
	j.WriteObjectField("testInteger1")
	j.WriteInt(s.TestInteger1)

	more.More()
	j.WriteObjectField("testTime")
	json.WriteTime(j, s.TestTime)

	more.More()
	j.WriteObjectField("type")
	s.Type.WriteJSON(j)

	more.More()
	j.WriteObjectField("unique_id")
	json.WriteUUID(j, s.UniqueID)

	more.More()
	j.WriteObjectField("uri")
	json.WriteURI(j, s.URI)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetResOKApplicationJSON json value to io.Writer.
func (s PetGetResOKApplicationJSON) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetResOKApplicationJSON json value from io.Reader.
func (s *PetGetResOKApplicationJSON) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetResOKApplicationJSON from json stream.
func (s *PetGetResOKApplicationJSON) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "birthday":
			if err := func() error {
				v, err := json.ReadDate(i)
				s.Birthday = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "friends":
			if err := func() error {
				s.Friends = s.Friends[:0]
				var retErr error
				i.ReadArrayCB(func(i *json.Iterator) bool {
					var elem PetGetResOKApplicationJSON
					if err := func() error {
						// Struct or enum.
						if err := elem.ReadJSON(i); err != nil {
							return err
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.Friends = append(s.Friends, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "id":
			if err := func() error {
				s.ID = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IP = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip_v4":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IPV4 = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip_v6":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IPV6 = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "kind":
			if err := func() error {
				s.Kind = PetGetResOKApplicationJSONKind(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "name":
			if err := func() error {
				s.Name = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "next":
			if err := func() error {
				// Struct or enum.
				if err := s.Next.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "nullStr":
			if err := func() error {
				s.NullStr.Reset()
				if err := s.NullStr.ReadJSON(i); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "primary":
			if err := func() error {
				s.Primary = nil
				var elem PetGetResOKApplicationJSON
				if err := func() error {
					// Struct or enum.
					if err := elem.ReadJSON(i); err != nil {
						return err
					}
					return i.Error
				}(); err != nil {
					return err
				}
				s.Primary = &elem
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "rate":
			if err := func() error {
				v, err := json.ReadDuration(i)
				s.Rate = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "tag":
			if err := func() error {
				v, err := json.ReadUUID(i)
				s.Tag = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testArray1":
			if err := func() error {
				s.TestArray1 = s.TestArray1[:0]
				var retErr error
				i.ReadArrayCB(func(i *json.Iterator) bool {
					var elem []string
					if err := func() error {
						elem = elem[:0]
						var retErr error
						i.ReadArrayCB(func(i *json.Iterator) bool {
							var elemElem string
							if err := func() error {
								elemElem = string(i.ReadString())
								return i.Error
							}(); err != nil {
								retErr = err
								return false
							}
							elem = append(elem, elemElem)
							return true
						})
						if retErr != nil {
							return retErr
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.TestArray1 = append(s.TestArray1, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDate":
			if err := func() error {
				v, err := json.ReadDate(i)
				s.TestDate = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDateTime":
			if err := func() error {
				v, err := json.ReadDateTime(i)
				s.TestDateTime = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDuration":
			if err := func() error {
				v, err := json.ReadDuration(i)
				s.TestDuration = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testFloat1":
			if err := func() error {
				s.TestFloat1 = float64(i.ReadFloat64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testInteger1":
			if err := func() error {
				s.TestInteger1 = int(i.ReadInt())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testTime":
			if err := func() error {
				v, err := json.ReadTime(i)
				s.TestTime = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "type":
			if err := func() error {
				s.Type = PetGetResOKApplicationJSONType(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "unique_id":
			if err := func() error {
				v, err := json.ReadUUID(i)
				s.UniqueID = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "uri":
			if err := func() error {
				v, err := json.ReadURI(i)
				s.URI = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetResOKApplicationJSONKind) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads PetGetResOKApplicationJSONKind from json stream.
func (s *PetGetResOKApplicationJSONKind) ReadJSON(i *json.Iterator) error {
	*s = PetGetResOKApplicationJSONKind(i.ReadString())
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetResOKApplicationJSONNext) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("description")
	j.WriteString(s.Description)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetResOKApplicationJSONNext json value to io.Writer.
func (s PetGetResOKApplicationJSONNext) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetResOKApplicationJSONNext json value from io.Reader.
func (s *PetGetResOKApplicationJSONNext) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetResOKApplicationJSONNext from json stream.
func (s *PetGetResOKApplicationJSONNext) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "description":
			if err := func() error {
				s.Description = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetResOKApplicationJSONType) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads PetGetResOKApplicationJSONType from json stream.
func (s *PetGetResOKApplicationJSONType) ReadJSON(i *json.Iterator) error {
	*s = PetGetResOKApplicationJSONType(i.ReadString())
	return i.Error
}
