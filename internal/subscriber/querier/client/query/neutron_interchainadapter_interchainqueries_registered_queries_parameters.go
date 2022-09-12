// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParams creates a new NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParams() *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	return &NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithTimeout creates a new NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams object
// with the ability to set a timeout on a request.
func NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithTimeout(timeout time.Duration) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	return &NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams{
		timeout: timeout,
	}
}

// NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithContext creates a new NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams object
// with the ability to set a context for a request.
func NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithContext(ctx context.Context) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	return &NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams{
		Context: ctx,
	}
}

// NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithHTTPClient creates a new NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewNeutronInterchainadapterInterchainqueriesRegisteredQueriesParamsWithHTTPClient(client *http.Client) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	return &NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams{
		HTTPClient: client,
	}
}

/* NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams contains all the parameters to send to the API endpoint
   for the neutron interchainadapter interchainqueries registered queries operation.

   Typically these are written to a http.Request.
*/
type NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams struct {

	// Owner.
	Owner *string

	/* PaginationCountTotal.

	     count_total is set to true  to indicate that the result set should include
	a count of the total number of items available for pagination in UIs.
	count_total is only respected when offset is used. It is ignored when key
	is set.
	*/
	PaginationCountTotal *bool

	/* PaginationKey.

	     key is a value returned in PageResponse.next_key to begin
	querying the next page most efficiently. Only one of offset or key
	should be set.

	     Format: byte
	*/
	PaginationKey *strfmt.Base64

	/* PaginationLimit.

	     limit is the total number of results to be returned in the result page.
	If left empty it will default to a value to be set by each app.

	     Format: uint64
	*/
	PaginationLimit *string

	/* PaginationOffset.

	     offset is a numeric offset that can be used when key is unavailable.
	It is less efficient than using key. Only one of offset or key should
	be set.

	     Format: uint64
	*/
	PaginationOffset *string

	/* PaginationReverse.

	     reverse is set to true if results are to be returned in the descending order.

	Since: cosmos-sdk 0.43
	*/
	PaginationReverse *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the neutron interchainadapter interchainqueries registered queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithDefaults() *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the neutron interchainadapter interchainqueries registered queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithTimeout(timeout time.Duration) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithContext(ctx context.Context) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithHTTPClient(client *http.Client) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithOwner(owner *string) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WithPaginationCountTotal adds the paginationCountTotal to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithPaginationCountTotal(paginationCountTotal *bool) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithPaginationKey(paginationKey *strfmt.Base64) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithPaginationLimit(paginationLimit *string) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithPaginationOffset(paginationOffset *string) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WithPaginationReverse(paginationReverse *bool) *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the neutron interchainadapter interchainqueries registered queries params
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WriteToRequest writes these params to a swagger request
func (o *NeutronInterchainadapterInterchainqueriesRegisteredQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Owner != nil {

		// query param owner
		var qrOwner string

		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}
	}

	if o.PaginationCountTotal != nil {

		// query param pagination.count_total
		var qrPaginationCountTotal bool

		if o.PaginationCountTotal != nil {
			qrPaginationCountTotal = *o.PaginationCountTotal
		}
		qPaginationCountTotal := swag.FormatBool(qrPaginationCountTotal)
		if qPaginationCountTotal != "" {

			if err := r.SetQueryParam("pagination.count_total", qPaginationCountTotal); err != nil {
				return err
			}
		}
	}

	if o.PaginationKey != nil {

		// query param pagination.key
		var qrPaginationKey strfmt.Base64

		if o.PaginationKey != nil {
			qrPaginationKey = *o.PaginationKey
		}
		qPaginationKey := qrPaginationKey.String()
		if qPaginationKey != "" {

			if err := r.SetQueryParam("pagination.key", qPaginationKey); err != nil {
				return err
			}
		}
	}

	if o.PaginationLimit != nil {

		// query param pagination.limit
		var qrPaginationLimit string

		if o.PaginationLimit != nil {
			qrPaginationLimit = *o.PaginationLimit
		}
		qPaginationLimit := qrPaginationLimit
		if qPaginationLimit != "" {

			if err := r.SetQueryParam("pagination.limit", qPaginationLimit); err != nil {
				return err
			}
		}
	}

	if o.PaginationOffset != nil {

		// query param pagination.offset
		var qrPaginationOffset string

		if o.PaginationOffset != nil {
			qrPaginationOffset = *o.PaginationOffset
		}
		qPaginationOffset := qrPaginationOffset
		if qPaginationOffset != "" {

			if err := r.SetQueryParam("pagination.offset", qPaginationOffset); err != nil {
				return err
			}
		}
	}

	if o.PaginationReverse != nil {

		// query param pagination.reverse
		var qrPaginationReverse bool

		if o.PaginationReverse != nil {
			qrPaginationReverse = *o.PaginationReverse
		}
		qPaginationReverse := swag.FormatBool(qrPaginationReverse)
		if qPaginationReverse != "" {

			if err := r.SetQueryParam("pagination.reverse", qPaginationReverse); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
