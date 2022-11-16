// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardReportCreateAttributes Attributes for the dashboard report schema used to create a dashboard report.
type DashboardReportCreateAttributes struct {
	// Description of the report in both the user interface and as part of the email body for each report sent. Must be a UTF-8 encoded string less than 4 KiB in size.
	Description datadog.NullableString `json:"description,omitempty"`
	// Report destination-specific configuration. This determines how reports are sent. Only email destinations are supported.
	Destinations DashboardReportDestination `json:"destinations"`
	// Report schedule-specific configuration that defines when and how often a report is sent.
	Schedule DashboardReportSchedule `json:"schedule"`
	// Template variables used to parameterize the dashboard when generating a report.
	TemplateVariables map[string]string `json:"template_variables,omitempty"`
	// Time period covered by the widgets in the dashboard report, at the time of report generation.
	Timeframe DashboardReportTimeframe `json:"timeframe"`
	// Title of the report in both the user interface, and as part of the email subject for each report sent.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportCreateAttributes instantiates a new DashboardReportCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportCreateAttributes(destinations DashboardReportDestination, schedule DashboardReportSchedule, timeframe DashboardReportTimeframe, title string) *DashboardReportCreateAttributes {
	this := DashboardReportCreateAttributes{}
	this.Destinations = destinations
	this.Schedule = schedule
	this.Timeframe = timeframe
	this.Title = title
	return &this
}

// NewDashboardReportCreateAttributesWithDefaults instantiates a new DashboardReportCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportCreateAttributesWithDefaults() *DashboardReportCreateAttributes {
	this := DashboardReportCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardReportCreateAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardReportCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardReportCreateAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *DashboardReportCreateAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *DashboardReportCreateAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *DashboardReportCreateAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDestinations returns the Destinations field value.
func (o *DashboardReportCreateAttributes) GetDestinations() DashboardReportDestination {
	if o == nil {
		var ret DashboardReportDestination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *DashboardReportCreateAttributes) GetDestinationsOk() (*DashboardReportDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// SetDestinations sets field value.
func (o *DashboardReportCreateAttributes) SetDestinations(v DashboardReportDestination) {
	o.Destinations = v
}

// GetSchedule returns the Schedule field value.
func (o *DashboardReportCreateAttributes) GetSchedule() DashboardReportSchedule {
	if o == nil {
		var ret DashboardReportSchedule
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *DashboardReportCreateAttributes) GetScheduleOk() (*DashboardReportSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value.
func (o *DashboardReportCreateAttributes) SetSchedule(v DashboardReportSchedule) {
	o.Schedule = v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *DashboardReportCreateAttributes) GetTemplateVariables() map[string]string {
	if o == nil || o.TemplateVariables == nil {
		var ret map[string]string
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportCreateAttributes) GetTemplateVariablesOk() (*map[string]string, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *DashboardReportCreateAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given map[string]string and assigns it to the TemplateVariables field.
func (o *DashboardReportCreateAttributes) SetTemplateVariables(v map[string]string) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value.
func (o *DashboardReportCreateAttributes) GetTimeframe() DashboardReportTimeframe {
	if o == nil {
		var ret DashboardReportTimeframe
		return ret
	}
	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *DashboardReportCreateAttributes) GetTimeframeOk() (*DashboardReportTimeframe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value.
func (o *DashboardReportCreateAttributes) SetTimeframe(v DashboardReportTimeframe) {
	o.Timeframe = v
}

// GetTitle returns the Title field value.
func (o *DashboardReportCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardReportCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *DashboardReportCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["destinations"] = o.Destinations
	toSerialize["schedule"] = o.Schedule
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Destinations *DashboardReportDestination `json:"destinations"`
		Schedule     *DashboardReportSchedule    `json:"schedule"`
		Timeframe    *DashboardReportTimeframe   `json:"timeframe"`
		Title        *string                     `json:"title"`
	}{}
	all := struct {
		Description       datadog.NullableString     `json:"description,omitempty"`
		Destinations      DashboardReportDestination `json:"destinations"`
		Schedule          DashboardReportSchedule    `json:"schedule"`
		TemplateVariables map[string]string          `json:"template_variables,omitempty"`
		Timeframe         DashboardReportTimeframe   `json:"timeframe"`
		Title             string                     `json:"title"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Destinations == nil {
		return fmt.Errorf("required field destinations missing")
	}
	if required.Schedule == nil {
		return fmt.Errorf("required field schedule missing")
	}
	if required.Timeframe == nil {
		return fmt.Errorf("required field timeframe missing")
	}
	if required.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Timeframe; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Description = all.Description
	if all.Destinations.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Destinations = all.Destinations
	if all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
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
