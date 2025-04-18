// Code generated by goa v3.20.0, DO NOT EDIT.
//
// MailAddresses HTTP client CLI support package
//
// Command:
// $ goa gen github.com/ynoacamino/alert-go/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode/utf8"

	mailaddresses "github.com/ynoacamino/alert-go/gen/mail_addresses"
	goa "goa.design/goa/v3/pkg"
)

// BuildListMailAddressesPayload builds the payload for the MailAddresses
// listMailAddresses endpoint from CLI flags.
func BuildListMailAddressesPayload(mailAddressesListMailAddressesPage string, mailAddressesListMailAddressesLimit string) (*mailaddresses.ListMailAddressesPayload, error) {
	var err error
	var page int
	{
		if mailAddressesListMailAddressesPage != "" {
			var v int64
			v, err = strconv.ParseInt(mailAddressesListMailAddressesPage, 10, strconv.IntSize)
			page = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for page, must be INT")
			}
			if page < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("page", page, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if mailAddressesListMailAddressesLimit != "" {
			var v int64
			v, err = strconv.ParseInt(mailAddressesListMailAddressesLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
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
		}
	}
	v := &mailaddresses.ListMailAddressesPayload{}
	v.Page = page
	v.Limit = limit

	return v, nil
}

// BuildGetMailAddressesPayload builds the payload for the MailAddresses
// getMailAddresses endpoint from CLI flags.
func BuildGetMailAddressesPayload(mailAddressesGetMailAddressesID string) (*mailaddresses.GetMailAddressesPayload, error) {
	var id string
	{
		id = mailAddressesGetMailAddressesID
	}
	v := &mailaddresses.GetMailAddressesPayload{}
	v.ID = id

	return v, nil
}

// BuildCreateMailAddressPayload builds the payload for the MailAddresses
// createMailAddress endpoint from CLI flags.
func BuildCreateMailAddressPayload(mailAddressesCreateMailAddressBody string) (*mailaddresses.MailPayload, error) {
	var err error
	var body CreateMailAddressRequestBody
	{
		err = json.Unmarshal([]byte(mailAddressesCreateMailAddressBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"active\": true,\n      \"address\": \"c\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.address", body.Address, goa.FormatEmail))
		if utf8.RuneCountInString(body.Address) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", body.Address, utf8.RuneCountInString(body.Address), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &mailaddresses.MailPayload{
		Address: body.Address,
		Active:  body.Active,
	}
	{
		var zero bool
		if v.Active == zero {
			v.Active = true
		}
	}

	return v, nil
}

// BuildUpdateMailAddressPayload builds the payload for the MailAddresses
// updateMailAddress endpoint from CLI flags.
func BuildUpdateMailAddressPayload(mailAddressesUpdateMailAddressBody string, mailAddressesUpdateMailAddressID string) (*mailaddresses.UpdateMailAddressPayload, error) {
	var err error
	var body UpdateMailAddressRequestBody
	{
		err = json.Unmarshal([]byte(mailAddressesUpdateMailAddressBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"active\": true,\n      \"address\": \"iy\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.address", body.Address, goa.FormatEmail))
		if utf8.RuneCountInString(body.Address) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", body.Address, utf8.RuneCountInString(body.Address), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = mailAddressesUpdateMailAddressID
	}
	v := &mailaddresses.UpdateMailAddressPayload{
		Address: body.Address,
		Active:  body.Active,
	}
	v.ID = id

	return v, nil
}

// BuildDeleteMailAddressPayload builds the payload for the MailAddresses
// deleteMailAddress endpoint from CLI flags.
func BuildDeleteMailAddressPayload(mailAddressesDeleteMailAddressID string) (*mailaddresses.DeleteMailAddressPayload, error) {
	var id string
	{
		id = mailAddressesDeleteMailAddressID
	}
	v := &mailaddresses.DeleteMailAddressPayload{}
	v.ID = id

	return v, nil
}
