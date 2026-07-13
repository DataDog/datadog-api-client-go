// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PrintReportRequestAttributes The configuration for a print-only report. Specify exactly one of `timeframe` (for a
// relative time window) or both `from_ts` and `to_ts` (for an absolute time range).
type PrintReportRequestAttributes struct {
	// The start of an absolute time range, as a Unix timestamp in milliseconds.
	// Required when `timeframe` is omitted.
	FromTs *int64 `json:"from_ts,omitempty"`
	// The identifier of the dashboard or integration dashboard to render.
	ResourceId string `json:"resource_id"`
	// The type of dashboard resource the report schedule targets.
	ResourceType ReportScheduleResourceType `json:"resource_type"`
	// The dashboard template variables applied when rendering the report.
	TemplateVariables []ReportScheduleTemplateVariable `json:"template_variables"`
	// A relative time window (for example `1w` or `calendar_month`). Mutually
	// exclusive with `from_ts` and `to_ts`.
	Timeframe *string `json:"timeframe,omitempty"`
	// The IANA time zone identifier used to evaluate the time window.
	Timezone string `json:"timezone"`
	// The end of an absolute time range, as a Unix timestamp in milliseconds.
	// Required when `timeframe` is omitted.
	ToTs *int64 `json:"to_ts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPrintReportRequestAttributes instantiates a new PrintReportRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPrintReportRequestAttributes(resourceId string, resourceType ReportScheduleResourceType, templateVariables []ReportScheduleTemplateVariable, timezone string) *PrintReportRequestAttributes {
	this := PrintReportRequestAttributes{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.TemplateVariables = templateVariables
	this.Timezone = timezone
	return &this
}

// NewPrintReportRequestAttributesWithDefaults instantiates a new PrintReportRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPrintReportRequestAttributesWithDefaults() *PrintReportRequestAttributes {
	this := PrintReportRequestAttributes{}
	return &this
}

// GetFromTs returns the FromTs field value if set, zero value otherwise.
func (o *PrintReportRequestAttributes) GetFromTs() int64 {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret
	}
	return *o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetFromTsOk() (*int64, bool) {
	if o == nil || o.FromTs == nil {
		return nil, false
	}
	return o.FromTs, true
}

// HasFromTs returns a boolean if a field has been set.
func (o *PrintReportRequestAttributes) HasFromTs() bool {
	return o != nil && o.FromTs != nil
}

// SetFromTs gets a reference to the given int64 and assigns it to the FromTs field.
func (o *PrintReportRequestAttributes) SetFromTs(v int64) {
	o.FromTs = &v
}

// GetResourceId returns the ResourceId field value.
func (o *PrintReportRequestAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *PrintReportRequestAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
func (o *PrintReportRequestAttributes) GetResourceType() ReportScheduleResourceType {
	if o == nil {
		var ret ReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetResourceTypeOk() (*ReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *PrintReportRequestAttributes) SetResourceType(v ReportScheduleResourceType) {
	o.ResourceType = v
}

// GetTemplateVariables returns the TemplateVariables field value.
func (o *PrintReportRequestAttributes) GetTemplateVariables() []ReportScheduleTemplateVariable {
	if o == nil {
		var ret []ReportScheduleTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetTemplateVariablesOk() (*[]ReportScheduleTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// SetTemplateVariables sets field value.
func (o *PrintReportRequestAttributes) SetTemplateVariables(v []ReportScheduleTemplateVariable) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *PrintReportRequestAttributes) GetTimeframe() string {
	if o == nil || o.Timeframe == nil {
		var ret string
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetTimeframeOk() (*string, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *PrintReportRequestAttributes) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given string and assigns it to the Timeframe field.
func (o *PrintReportRequestAttributes) SetTimeframe(v string) {
	o.Timeframe = &v
}

// GetTimezone returns the Timezone field value.
func (o *PrintReportRequestAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *PrintReportRequestAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetToTs returns the ToTs field value if set, zero value otherwise.
func (o *PrintReportRequestAttributes) GetToTs() int64 {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret
	}
	return *o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintReportRequestAttributes) GetToTsOk() (*int64, bool) {
	if o == nil || o.ToTs == nil {
		return nil, false
	}
	return o.ToTs, true
}

// HasToTs returns a boolean if a field has been set.
func (o *PrintReportRequestAttributes) HasToTs() bool {
	return o != nil && o.ToTs != nil
}

// SetToTs gets a reference to the given int64 and assigns it to the ToTs field.
func (o *PrintReportRequestAttributes) SetToTs(v int64) {
	o.ToTs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PrintReportRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FromTs != nil {
		toSerialize["from_ts"] = o.FromTs
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["template_variables"] = o.TemplateVariables
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	toSerialize["timezone"] = o.Timezone
	if o.ToTs != nil {
		toSerialize["to_ts"] = o.ToTs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PrintReportRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromTs            *int64                            `json:"from_ts,omitempty"`
		ResourceId        *string                           `json:"resource_id"`
		ResourceType      *ReportScheduleResourceType       `json:"resource_type"`
		TemplateVariables *[]ReportScheduleTemplateVariable `json:"template_variables"`
		Timeframe         *string                           `json:"timeframe,omitempty"`
		Timezone          *string                           `json:"timezone"`
		ToTs              *int64                            `json:"to_ts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.TemplateVariables == nil {
		return fmt.Errorf("required field template_variables missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_ts", "resource_id", "resource_type", "template_variables", "timeframe", "timezone", "to_ts"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FromTs = all.FromTs
	o.ResourceId = *all.ResourceId
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.TemplateVariables = *all.TemplateVariables
	o.Timeframe = all.Timeframe
	o.Timezone = *all.Timezone
	o.ToTs = all.ToTs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
