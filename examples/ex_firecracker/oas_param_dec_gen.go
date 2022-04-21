// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func decodePatchGuestDriveByIDParams(args [1]string, r *http.Request) (PatchGuestDriveByIDParams, error) {
	var (
		params PatchGuestDriveByIDParams
	)
	// Decode path: drive_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: drive_id: not specified")
		}
	}
	return params, nil
}

func decodePatchGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (PatchGuestNetworkInterfaceByIDParams, error) {
	var (
		params PatchGuestNetworkInterfaceByIDParams
	)
	// Decode path: iface_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: iface_id: not specified")
		}
	}
	return params, nil
}

func decodePutGuestDriveByIDParams(args [1]string, r *http.Request) (PutGuestDriveByIDParams, error) {
	var (
		params PutGuestDriveByIDParams
	)
	// Decode path: drive_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: drive_id: not specified")
		}
	}
	return params, nil
}

func decodePutGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (PutGuestNetworkInterfaceByIDParams, error) {
	var (
		params PutGuestNetworkInterfaceByIDParams
	)
	// Decode path: iface_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: iface_id: not specified")
		}
	}
	return params, nil
}
