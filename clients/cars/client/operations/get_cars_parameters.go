// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCarsParams creates a new GetCarsParams object
// with the default values initialized.
func NewGetCarsParams() *GetCarsParams {
	var ()
	return &GetCarsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCarsParamsWithTimeout creates a new GetCarsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCarsParamsWithTimeout(timeout time.Duration) *GetCarsParams {
	var ()
	return &GetCarsParams{

		timeout: timeout,
	}
}

// NewGetCarsParamsWithContext creates a new GetCarsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCarsParamsWithContext(ctx context.Context) *GetCarsParams {
	var ()
	return &GetCarsParams{

		Context: ctx,
	}
}

// NewGetCarsParamsWithHTTPClient creates a new GetCarsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCarsParamsWithHTTPClient(client *http.Client) *GetCarsParams {
	var ()
	return &GetCarsParams{
		HTTPClient: client,
	}
}

/*GetCarsParams contains all the parameters to send to the API endpoint
for the get cars operation typically these are written to a http.Request
*/
type GetCarsParams struct {

	/*Lat
	  Latitude

	*/
	Lat float64
	/*Limit
	  Number of cars requested


	*/
	Limit int64
	/*Lng
	  Longitude

	*/
	Lng float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cars params
func (o *GetCarsParams) WithTimeout(timeout time.Duration) *GetCarsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cars params
func (o *GetCarsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cars params
func (o *GetCarsParams) WithContext(ctx context.Context) *GetCarsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cars params
func (o *GetCarsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cars params
func (o *GetCarsParams) WithHTTPClient(client *http.Client) *GetCarsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cars params
func (o *GetCarsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLat adds the lat to the get cars params
func (o *GetCarsParams) WithLat(lat float64) *GetCarsParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the get cars params
func (o *GetCarsParams) SetLat(lat float64) {
	o.Lat = lat
}

// WithLimit adds the limit to the get cars params
func (o *GetCarsParams) WithLimit(limit int64) *GetCarsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get cars params
func (o *GetCarsParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithLng adds the lng to the get cars params
func (o *GetCarsParams) WithLng(lng float64) *GetCarsParams {
	o.SetLng(lng)
	return o
}

// SetLng adds the lng to the get cars params
func (o *GetCarsParams) SetLng(lng float64) {
	o.Lng = lng
}

// WriteToRequest writes these params to a swagger request
func (o *GetCarsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param lat
	qrLat := o.Lat
	qLat := swag.FormatFloat64(qrLat)
	if qLat != "" {
		if err := r.SetQueryParam("lat", qLat); err != nil {
			return err
		}
	}

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)
	if qLimit != "" {
		if err := r.SetQueryParam("limit", qLimit); err != nil {
			return err
		}
	}

	// query param lng
	qrLng := o.Lng
	qLng := swag.FormatFloat64(qrLng)
	if qLng != "" {
		if err := r.SetQueryParam("lng", qLng); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
