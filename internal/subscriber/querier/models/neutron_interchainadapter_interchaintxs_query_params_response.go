// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NeutronInterchainadapterInterchaintxsQueryParamsResponse QueryParamsResponse is response type for the Query/Params RPC method.
//
// swagger:model neutron.interchainadapter.interchaintxs.QueryParamsResponse
type NeutronInterchainadapterInterchaintxsQueryParamsResponse struct {

	// params holds all the parameters of this module.
	Params interface{} `json:"params,omitempty"`
}

// Validate validates this neutron interchainadapter interchaintxs query params response
func (m *NeutronInterchainadapterInterchaintxsQueryParamsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchaintxs query params response based on context it is used
func (m *NeutronInterchainadapterInterchaintxsQueryParamsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchaintxsQueryParamsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchaintxsQueryParamsResponse) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchaintxsQueryParamsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
