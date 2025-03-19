// Code generated by goa v3.20.0, DO NOT EDIT.
//
// MailAddresses HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/ynoacamino/alert-go/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	mailaddresses "github.com/ynoacamino/alert-go/gen/mail_addresses"
	mailaddressesviews "github.com/ynoacamino/alert-go/gen/mail_addresses/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListMailAddressesRequest instantiates a HTTP request object with method
// and path set to call the "MailAddresses" service "listMailAddresses" endpoint
func (c *Client) BuildListMailAddressesRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMailAddressesMailAddressesPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("MailAddresses", "listMailAddresses", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListMailAddressesRequest returns an encoder for requests sent to the
// MailAddresses listMailAddresses server.
func EncodeListMailAddressesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*mailaddresses.ListMailAddressesPayload)
		if !ok {
			return goahttp.ErrInvalidType("MailAddresses", "listMailAddresses", "*mailaddresses.ListMailAddressesPayload", v)
		}
		values := req.URL.Query()
		values.Add("page", fmt.Sprintf("%v", p.Page))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListMailAddressesResponse returns a decoder for responses returned by
// the MailAddresses listMailAddresses endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeListMailAddressesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListMailAddressesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "listMailAddresses", err)
			}
			p := NewListMailAddressesMaillistOK(&body)
			view := "default"
			vres := &mailaddressesviews.Maillist{Projected: p, View: view}
			if err = mailaddressesviews.ValidateMaillist(vres); err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "listMailAddresses", err)
			}
			res := mailaddresses.NewMaillist(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("MailAddresses", "listMailAddresses", resp.StatusCode, string(body))
		}
	}
}

// BuildGetMailAddressesRequest instantiates a HTTP request object with method
// and path set to call the "MailAddresses" service "getMailAddresses" endpoint
func (c *Client) BuildGetMailAddressesRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*mailaddresses.GetMailAddressesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("MailAddresses", "getMailAddresses", "*mailaddresses.GetMailAddressesPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetMailAddressesMailAddressesPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("MailAddresses", "getMailAddresses", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetMailAddressesResponse returns a decoder for responses returned by
// the MailAddresses getMailAddresses endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeGetMailAddressesResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetMailAddressesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetMailAddressesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "getMailAddresses", err)
			}
			err = ValidateGetMailAddressesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "getMailAddresses", err)
			}
			res := NewGetMailAddressesMailOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetMailAddressesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "getMailAddresses", err)
			}
			err = ValidateGetMailAddressesNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "getMailAddresses", err)
			}
			return nil, NewGetMailAddressesNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("MailAddresses", "getMailAddresses", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateMailAddressRequest instantiates a HTTP request object with method
// and path set to call the "MailAddresses" service "createMailAddress" endpoint
func (c *Client) BuildCreateMailAddressRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateMailAddressMailAddressesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("MailAddresses", "createMailAddress", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateMailAddressRequest returns an encoder for requests sent to the
// MailAddresses createMailAddress server.
func EncodeCreateMailAddressRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*mailaddresses.MailPayload)
		if !ok {
			return goahttp.ErrInvalidType("MailAddresses", "createMailAddress", "*mailaddresses.MailPayload", v)
		}
		body := NewCreateMailAddressRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("MailAddresses", "createMailAddress", err)
		}
		return nil
	}
}

// DecodeCreateMailAddressResponse returns a decoder for responses returned by
// the MailAddresses createMailAddress endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeCreateMailAddressResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateMailAddressResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "createMailAddress", err)
			}
			err = ValidateCreateMailAddressResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "createMailAddress", err)
			}
			res := NewCreateMailAddressMailCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("MailAddresses", "createMailAddress", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateMailAddressRequest instantiates a HTTP request object with method
// and path set to call the "MailAddresses" service "updateMailAddress" endpoint
func (c *Client) BuildUpdateMailAddressRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*mailaddresses.UpdateMailAddressPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("MailAddresses", "updateMailAddress", "*mailaddresses.UpdateMailAddressPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateMailAddressMailAddressesPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("MailAddresses", "updateMailAddress", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateMailAddressRequest returns an encoder for requests sent to the
// MailAddresses updateMailAddress server.
func EncodeUpdateMailAddressRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*mailaddresses.UpdateMailAddressPayload)
		if !ok {
			return goahttp.ErrInvalidType("MailAddresses", "updateMailAddress", "*mailaddresses.UpdateMailAddressPayload", v)
		}
		body := NewUpdateMailAddressRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("MailAddresses", "updateMailAddress", err)
		}
		return nil
	}
}

// DecodeUpdateMailAddressResponse returns a decoder for responses returned by
// the MailAddresses updateMailAddress endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeUpdateMailAddressResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeUpdateMailAddressResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusNotFound:
			var (
				body UpdateMailAddressNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "updateMailAddress", err)
			}
			err = ValidateUpdateMailAddressNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "updateMailAddress", err)
			}
			return nil, NewUpdateMailAddressNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("MailAddresses", "updateMailAddress", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteMailAddressRequest instantiates a HTTP request object with method
// and path set to call the "MailAddresses" service "deleteMailAddress" endpoint
func (c *Client) BuildDeleteMailAddressRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*mailaddresses.DeleteMailAddressPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("MailAddresses", "deleteMailAddress", "*mailaddresses.DeleteMailAddressPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteMailAddressMailAddressesPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("MailAddresses", "deleteMailAddress", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteMailAddressResponse returns a decoder for responses returned by
// the MailAddresses deleteMailAddress endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeDeleteMailAddressResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteMailAddressResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DeleteMailAddressNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("MailAddresses", "deleteMailAddress", err)
			}
			err = ValidateDeleteMailAddressNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("MailAddresses", "deleteMailAddress", err)
			}
			return nil, NewDeleteMailAddressNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("MailAddresses", "deleteMailAddress", resp.StatusCode, string(body))
		}
	}
}

// unmarshalMailResponseBodyToMailaddressesviewsMailView builds a value of type
// *mailaddressesviews.MailView from a value of type *MailResponseBody.
func unmarshalMailResponseBodyToMailaddressesviewsMailView(v *MailResponseBody) *mailaddressesviews.MailView {
	res := &mailaddressesviews.MailView{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdateAt:  v.UpdateAt,
		Address:   v.Address,
		Active:    v.Active,
	}

	return res
}
