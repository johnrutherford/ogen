package uri

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathEncoder(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		tests := []struct {
			Param   string
			Input   string
			Expect  string
			Style   PathStyle
			Explode bool
		}{
			{
				Param:   "id",
				Input:   "foo",
				Expect:  "foo",
				Style:   PathStyleSimple,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   "foo",
				Expect:  "foo",
				Style:   PathStyleSimple,
				Explode: true,
			},
			{
				Param:   "id",
				Input:   "foo",
				Expect:  ".foo",
				Style:   PathStyleLabel,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   "foo",
				Expect:  ".foo",
				Style:   PathStyleLabel,
				Explode: true,
			},
			{
				Param:   "id",
				Input:   "foo",
				Expect:  ";id=foo",
				Style:   PathStyleMatrix,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   "foo",
				Expect:  ";id=foo",
				Style:   PathStyleMatrix,
				Explode: true,
			},
		}

		for i, test := range tests {
			result := NewPathEncoder(PathEncoderConfig{
				Param:   test.Param,
				Style:   test.Style,
				Explode: test.Explode,
			}).EncodeString(test.Input)
			require.Equal(t, test.Expect, result, fmt.Sprintf("Test %d", i+1))
		}
	})

	t.Run("Strings", func(t *testing.T) {
		tests := []struct {
			Param   string
			Input   []string
			Expect  string
			Style   PathStyle
			Explode bool
		}{
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  "foo,bar",
				Style:   PathStyleSimple,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  "foo,bar",
				Style:   PathStyleSimple,
				Explode: true,
			},
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  ".foo,bar",
				Style:   PathStyleLabel,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  ".foo.bar",
				Style:   PathStyleLabel,
				Explode: true,
			},
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  ";id=foo,bar",
				Style:   PathStyleMatrix,
				Explode: false,
			},
			{
				Param:   "id",
				Input:   []string{"foo", "bar"},
				Expect:  ";id=foo;id=bar",
				Style:   PathStyleMatrix,
				Explode: true,
			},
		}

		for i, test := range tests {
			result := NewPathEncoder(PathEncoderConfig{
				Param:   test.Param,
				Style:   test.Style,
				Explode: test.Explode,
			}).EncodeStrings(test.Input)
			require.Equal(t, test.Expect, result, fmt.Sprintf("Test %d", i+1))
		}
	})

	t.Run("Object", func(t *testing.T) {
		tests := []struct {
			Param   string
			Input   []PathObjectField
			Expect  string
			Style   PathStyle
			Explode bool
		}{
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleSimple,
				Explode: false,
				Expect:  "role,admin,firstName,Alex",
			},
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleSimple,
				Explode: true,
				Expect:  "role=admin,firstName=Alex",
			},
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleLabel,
				Explode: false,
				Expect:  ".role,admin,firstName,Alex",
			},
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleLabel,
				Explode: true,
				Expect:  ".role=admin.firstName=Alex",
			},
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleMatrix,
				Explode: false,
				Expect:  ";id=role,admin,firstName,Alex",
			},
			{
				Param: "id",
				Input: []PathObjectField{
					{"role", "admin"},
					{"firstName", "Alex"},
				},
				Style:   PathStyleMatrix,
				Explode: true,
				Expect:  ";role=admin;firstName=Alex",
			},
		}

		for i, test := range tests {
			result := NewPathEncoder(PathEncoderConfig{
				Param:   test.Param,
				Style:   test.Style,
				Explode: test.Explode,
			}).EncodeObject(test.Input)
			require.Equal(t, test.Expect, result, fmt.Sprintf("Test %d", i+1))
		}
	})
}
