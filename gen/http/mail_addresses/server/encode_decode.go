// Code generated by goa v3.20.0, DO NOT EDIT.
//
// MailAddresses HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/ynoacamino/alert-go/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	mailaddresses "github.com/ynoacamino/alert-go/gen/mail_addresses"
	mailaddressesviews "github.com/ynoacamino/alert-go/gen/mail_addresses/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListMailAddressesResponse returns an encoder for responses returned by
// the MailAddresses listMailAddresses endpoint.
func EncodeListMailAddressesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*mailaddressesviews.Maillist)
		enc := encoder(ctx, w)
		body := NewListMailAddressesResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListMailAddressesRequest returns a decoder for requests sent to the
// MailAddresses listMailAddresses endpoint.
func DecodeListMailAddressesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			page  int
			limit int
			err   error
		)
		qp := r.URL.Query()
		{
			pageRaw := qp.Get("page")
			if pageRaw == "" {
				page = 1
			} else {
				v, err2 := strconv.ParseInt(pageRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("page", pageRaw, "integer"))
				}
				page = int(v)
			}
		}
		if page < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("page", page, 0, true))
		}
		{
			limitRaw := qp.Get("limit")
			if limitRaw == "" {
				limit = 10
			} else {
				v, err2 := strconv.ParseInt(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
				}
				limit = int(v)
			}
		}
		if limit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 0, true))
		}
		if limit > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListMailAddressesPayload(page, limit)

		return payload, nil
	}
}

// EncodeGetMailAddressesResponse returns an encoder for responses returned by
// the MailAddresses getMailAddresses endpoint.
func EncodeGetMailAddressesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*mailaddresses.Mail)
		enc := encoder(ctx, w)
		body := NewGetMailAddressesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetMailAddressesRequest returns a decoder for requests sent to the
// MailAddresses getMailAddresses endpoint.
func DecodeGetMailAddressesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetMailAddressesPayload(id)

		return payload, nil
	}
}

// EncodeGetMailAddressesError returns an encoder for errors returned by the
// getMailAddresses MailAddresses endpoint.
func EncodeGetMailAddressesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetMailAddressesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateMailAddressResponse returns an encoder for responses returned by
// the MailAddresses createMailAddress endpoint.
func EncodeCreateMailAddressResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*mailaddresses.Mail)
		enc := encoder(ctx, w)
		body := NewCreateMailAddressResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateMailAddressRequest returns a decoder for requests sent to the
// MailAddresses createMailAddress endpoint.
func DecodeCreateMailAddressRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateMailAddressRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateMailAddressRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateMailAddressMailPayload(&body)

		return payload, nil
	}
}

// EncodeUpdateMailAddressResponse returns an encoder for responses returned by
// the MailAddresses updateMailAddress endpoint.
func EncodeUpdateMailAddressResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateMailAddressRequest returns a decoder for requests sent to the
// MailAddresses updateMailAddress endpoint.
func DecodeUpdateMailAddressRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateMailAddressRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateMailAddressRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewUpdateMailAddressPayload(&body, id)

		return payload, nil
	}
}

// EncodeUpdateMailAddressError returns an encoder for errors returned by the
// updateMailAddress MailAddresses endpoint.
func EncodeUpdateMailAddressError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateMailAddressNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteMailAddressResponse returns an encoder for responses returned by
// the MailAddresses deleteMailAddress endpoint.
func EncodeDeleteMailAddressResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteMailAddressRequest returns a decoder for requests sent to the
// MailAddresses deleteMailAddress endpoint.
func DecodeDeleteMailAddressRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewDeleteMailAddressPayload(id)

		return payload, nil
	}
}

// EncodeDeleteMailAddressError returns an encoder for errors returned by the
// deleteMailAddress MailAddresses endpoint.
func EncodeDeleteMailAddressError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteMailAddressNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalMailaddressesviewsMailViewToMailResponseBody builds a value of type
// *MailResponseBody from a value of type *mailaddressesviews.MailView.
func marshalMailaddressesviewsMailViewToMailResponseBody(v *mailaddressesviews.MailView) *MailResponseBody {
	res := &MailResponseBody{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		UpdateAt:  *v.UpdateAt,
		Address:   *v.Address,
		Active:    *v.Active,
	}

	return res
}
