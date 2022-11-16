// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardReportUpdateAttributes Attributes of a dashboard report that can be updated. Fields that are not to be updated can be omitted, with some exceptions for `repeat_on_day_of_week` and `repeat_on_day_of_month`, as noted in their respective sections.
type DashboardReportUpdateAttributes struct {
	// Description of the report in both the user interface and as part of the email body for each report sent. Must be a UTF-8 encoded string less than 4 KiB in size.
	Description datadog.NullableString `json:"description,omitempty"`
	// Report destination-specific configuration. This determines how reports are sent. Only email destinations are supported.
	Destinations *DashboardReportDestination `json:"destinations,omitempty"`
	// Report schedule-specific configuration that defines when and how often a report is sent.
	Schedule *DashboardReportSchedule `json:"schedule,omitempty"`
	// Template variables used to parameterize the dashboard when generating a report.
	TemplateVariables map[string]string `json:"template_variables,omitempty"`
	// Time period covered by the widgets in the dashboard report, at the time of report generation.
	Timeframe *DashboardReportTimeframe `json:"timeframe,omitempty"`
	// Title of the report in both the user interface, and as part of the email subject for each report sent.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportUpdateAttributes instantiates a new DashboardReportUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportUpdateAttributes() *DashboardReportUpdateAttributes {
	this := DashboardReportUpdateAttributes{}
	return &this
}

// NewDashboardReportUpdateAttributesWithDefaults instantiates a new DashboardReportUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportUpdateAttributesWithDefaults() *DashboardReportUpdateAttributes {
	this := DashboardReportUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardReportUpdateAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardReportUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *DashboardReportUpdateAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *DashboardReportUpdateAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *DashboardReportUpdateAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *DashboardReportUpdateAttributes) GetDestinations() DashboardReportDestination {
	if o == nil || o.Destinations == nil {
		var ret DashboardReportDestination
		return ret
	}
	return *o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportUpdateAttributes) GetDestinationsOk() (*DashboardReportDestination, bool) {
	if o == nil || o.Destinations == nil {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasDestinations() bool {
	return o != nil && o.Destinations != nil
}

// SetDestinations gets a reference to the given DashboardReportDestination and assigns it to the Destinations field.
func (o *DashboardReportUpdateAttributes) SetDestinations(v DashboardReportDestination) {
	o.Destinations = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DashboardReportUpdateAttributes) GetSchedule() DashboardReportSchedule {
	if o == nil || o.Schedule == nil {
		var ret DashboardReportSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportUpdateAttributes) GetScheduleOk() (*DashboardReportSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given DashboardReportSchedule and assigns it to the Schedule field.
func (o *DashboardReportUpdateAttributes) SetSchedule(v DashboardReportSchedule) {
	o.Schedule = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *DashboardReportUpdateAttributes) GetTemplateVariables() map[string]string {
	if o == nil || o.TemplateVariables == nil {
		var ret map[string]string
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportUpdateAttributes) GetTemplateVariablesOk() (*map[string]string, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given map[string]string and assigns it to the TemplateVariables field.
func (o *DashboardReportUpdateAttributes) SetTemplateVariables(v map[string]string) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *DashboardReportUpdateAttributes) GetTimeframe() DashboardReportTimeframe {
	if o == nil || o.Timeframe == nil {
		var ret DashboardReportTimeframe
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportUpdateAttributes) GetTimeframeOk() (*DashboardReportTimeframe, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given DashboardReportTimeframe and assigns it to the Timeframe field.
func (o *DashboardReportUpdateAttributes) SetTimeframe(v DashboardReportTimeframe) {
	o.Timeframe = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DashboardReportUpdateAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportUpdateAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DashboardReportUpdateAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DashboardReportUpdateAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Destinations != nil {
		toSerialize["destinations"] = o.Destinations
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Description       datadog.NullableString      `json:"description,omitempty"`
		Destinations      *DashboardReportDestination `json:"destinations,omitempty"`
		Schedule          *DashboardReportSchedule    `json:"schedule,omitempty"`
		TemplateVariables map[string]string           `json:"template_variables,omitempty"`
		Timeframe         *DashboardReportTimeframe   `json:"timeframe,omitempty"`
		Title             *string                     `json:"title,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Timeframe; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Description = all.Description
	if all.Destinations != nil && all.Destinations.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Destinations = all.Destinations
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Schedule = all.Schedule
	o.TemplateVariables = all.TemplateVariables
	o.Timeframe = all.Timeframe
	o.Title = all.Title
	return nil
}
