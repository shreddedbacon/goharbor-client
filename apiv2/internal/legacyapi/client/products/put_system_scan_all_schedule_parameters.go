// Code generated by go-swagger; DO NOT EDIT.

package products

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

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// NewPutSystemScanAllScheduleParams creates a new PutSystemScanAllScheduleParams object
// with the default values initialized.
func NewPutSystemScanAllScheduleParams() *PutSystemScanAllScheduleParams {
	var ()
	return &PutSystemScanAllScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemScanAllScheduleParamsWithTimeout creates a new PutSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSystemScanAllScheduleParamsWithTimeout(timeout time.Duration) *PutSystemScanAllScheduleParams {
	var ()
	return &PutSystemScanAllScheduleParams{

		timeout: timeout,
	}
}

// NewPutSystemScanAllScheduleParamsWithContext creates a new PutSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSystemScanAllScheduleParamsWithContext(ctx context.Context) *PutSystemScanAllScheduleParams {
	var ()
	return &PutSystemScanAllScheduleParams{

		Context: ctx,
	}
}

// NewPutSystemScanAllScheduleParamsWithHTTPClient creates a new PutSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSystemScanAllScheduleParamsWithHTTPClient(client *http.Client) *PutSystemScanAllScheduleParams {
	var ()
	return &PutSystemScanAllScheduleParams{
		HTTPClient: client,
	}
}

/*PutSystemScanAllScheduleParams contains all the parameters to send to the API endpoint
for the put system scan all schedule operation typically these are written to a http.Request
*/
type PutSystemScanAllScheduleParams struct {

	/*Schedule
	  Updates the schedule of scan all job, which scans all of images in Harbor.

	*/
	Schedule *legacy.AdminJobSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) WithTimeout(timeout time.Duration) *PutSystemScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) WithContext(ctx context.Context) *PutSystemScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) WithHTTPClient(client *http.Client) *PutSystemScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchedule adds the schedule to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) WithSchedule(schedule *legacy.AdminJobSchedule) *PutSystemScanAllScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the put system scan all schedule params
func (o *PutSystemScanAllScheduleParams) SetSchedule(schedule *legacy.AdminJobSchedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
