// Code generated by goa v3.20.0, DO NOT EDIT.
//
// MailAddresses endpoints
//
// Command:
// $ goa gen github.com/ynoacamino/alert-go/design

package mailaddresses

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "MailAddresses" service endpoints.
type Endpoints struct {
	ListMailAddresses goa.Endpoint
	GetMailAddresses  goa.Endpoint
	CreateMailAddress goa.Endpoint
	UpdateMailAddress goa.Endpoint
	DeleteMailAddress goa.Endpoint
}

// NewEndpoints wraps the methods of the "MailAddresses" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ListMailAddresses: NewListMailAddressesEndpoint(s),
		GetMailAddresses:  NewGetMailAddressesEndpoint(s),
		CreateMailAddress: NewCreateMailAddressEndpoint(s),
		UpdateMailAddress: NewUpdateMailAddressEndpoint(s),
		DeleteMailAddress: NewDeleteMailAddressEndpoint(s),
	}
}

// Use applies the given middleware to all the "MailAddresses" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListMailAddresses = m(e.ListMailAddresses)
	e.GetMailAddresses = m(e.GetMailAddresses)
	e.CreateMailAddress = m(e.CreateMailAddress)
	e.UpdateMailAddress = m(e.UpdateMailAddress)
	e.DeleteMailAddress = m(e.DeleteMailAddress)
}

// NewListMailAddressesEndpoint returns an endpoint function that calls the
// method "listMailAddresses" of service "MailAddresses".
func NewListMailAddressesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListMailAddressesPayload)
		res, err := s.ListMailAddresses(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMaillist(res, "default")
		return vres, nil
	}
}

// NewGetMailAddressesEndpoint returns an endpoint function that calls the
// method "getMailAddresses" of service "MailAddresses".
func NewGetMailAddressesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetMailAddressesPayload)
		return s.GetMailAddresses(ctx, p)
	}
}

// NewCreateMailAddressEndpoint returns an endpoint function that calls the
// method "createMailAddress" of service "MailAddresses".
func NewCreateMailAddressEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MailPayload)
		return s.CreateMailAddress(ctx, p)
	}
}

// NewUpdateMailAddressEndpoint returns an endpoint function that calls the
// method "updateMailAddress" of service "MailAddresses".
func NewUpdateMailAddressEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateMailAddressPayload)
		return s.UpdateMailAddress(ctx, p)
	}
}

// NewDeleteMailAddressEndpoint returns an endpoint function that calls the
// method "deleteMailAddress" of service "MailAddresses".
func NewDeleteMailAddressEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteMailAddressPayload)
		return nil, s.DeleteMailAddress(ctx, p)
	}
}
