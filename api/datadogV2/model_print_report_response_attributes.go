// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PrintReportResponseAttributes The configuration and download URL for the initiated print-only report.
type PrintReportResponseAttributes struct {
	// The URL from which the rendered PDF report can be downloaded.
	DownloadUrl string `json:"download_url"`
	// The start of the rendered time range, as a Unix timestamp in milliseconds.
	FromTs int64 `json:"from_ts"`
	// The identifier of the dashboard or integration dashboard.
	ResourceId string `json:"resource_id"`
	// The type of dashboard resource the report schedule targets.
	ResourceType ReportScheduleResourceType `json:"resource_type"`
	// The dashboard template variables applied when rendering the report.
	TemplateVariables []ReportScheduleTemplateVariable `json:"template_variables"`
	// The relative time window used, if one was specified in the request.
	Timeframe *string `json:"timeframe,omitempty"`
	// The IANA time zone identifier used when rendering the report.
	Timezone string `json:"timezone"`
	// The end of the rendered time range, as a Unix timestamp in milliseconds.
	ToTs int64 `json:"to_ts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPrintReportResponseAttributes instantiates a new PrintReportResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPrintReportResponseAttributes(downloadUrl string, fromTs int64, resourceId string, resourceType ReportScheduleResourceType, templateVariables []ReportScheduleTemplateVariable, timezone string, toTs int64) *PrintReportResponseAttributes {
	this := PrintReportResponseAttributes{}
	this.DownloadUrl = downloadUrl
	this.FromTs = fromTs
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.TemplateVariables = templateVariables
	this.Timezone = timezone
	this.ToTs = toTs
	return &this
}

// NewPrintReportResponseAttributesWithDefaults instantiates a new PrintReportResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPrintReportResponseAttributesWithDefaults() *PrintReportResponseAttributes {
	this := PrintReportResponseAttributes{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value.
func (o *PrintReportResponseAttributes) GetDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadUrl, true
}

// SetDownloadUrl sets field value.
func (o *PrintReportResponseAttributes) SetDownloadUrl(v string) {
	o.DownloadUrl = v
}

// GetFromTs returns the FromTs field value.
func (o *PrintReportResponseAttributes) GetFromTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetFromTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTs, true
}

// SetFromTs sets field value.
func (o *PrintReportResponseAttributes) SetFromTs(v int64) {
	o.FromTs = v
}

// GetResourceId returns the ResourceId field value.
func (o *PrintReportResponseAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *PrintReportResponseAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
func (o *PrintReportResponseAttributes) GetResourceType() ReportScheduleResourceType {
	if o == nil {
		var ret ReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetResourceTypeOk() (*ReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *PrintReportResponseAttributes) SetResourceType(v ReportScheduleResourceType) {
	o.ResourceType = v
}

// GetTemplateVariables returns the TemplateVariables field value.
func (o *PrintReportResponseAttributes) GetTemplateVariables() []ReportScheduleTemplateVariable {
	if o == nil {
		var ret []ReportScheduleTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetTemplateVariablesOk() (*[]ReportScheduleTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// SetTemplateVariables sets field value.
func (o *PrintReportResponseAttributes) SetTemplateVariables(v []ReportScheduleTemplateVariable) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *PrintReportResponseAttributes) GetTimeframe() string {
	if o == nil || o.Timeframe == nil {
		var ret string
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetTimeframeOk() (*string, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *PrintReportResponseAttributes) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given string and assigns it to the Timeframe field.
func (o *PrintReportResponseAttributes) SetTimeframe(v string) {
	o.Timeframe = &v
}

// GetTimezone returns the Timezone field value.
func (o *PrintReportResponseAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *PrintReportResponseAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetToTs returns the ToTs field value.
func (o *PrintReportResponseAttributes) GetToTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value
// and a boolean to check if the value has been set.
func (o *PrintReportResponseAttributes) GetToTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTs, true
}

// SetToTs sets field value.
func (o *PrintReportResponseAttributes) SetToTs(v int64) {
	o.ToTs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PrintReportResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["download_url"] = o.DownloadUrl
	toSerialize["from_ts"] = o.FromTs
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["template_variables"] = o.TemplateVariables
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	toSerialize["timezone"] = o.Timezone
	toSerialize["to_ts"] = o.ToTs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PrintReportResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DownloadUrl       *string                           `json:"download_url"`
		FromTs            *int64                            `json:"from_ts"`
		ResourceId        *string                           `json:"resource_id"`
		ResourceType      *ReportScheduleResourceType       `json:"resource_type"`
		TemplateVariables *[]ReportScheduleTemplateVariable `json:"template_variables"`
		Timeframe         *string                           `json:"timeframe,omitempty"`
		Timezone          *string                           `json:"timezone"`
		ToTs              *int64                            `json:"to_ts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DownloadUrl == nil {
		return fmt.Errorf("required field download_url missing")
	}
	if all.FromTs == nil {
		return fmt.Errorf("required field from_ts missing")
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
	if all.ToTs == nil {
		return fmt.Errorf("required field to_ts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"download_url", "from_ts", "resource_id", "resource_type", "template_variables", "timeframe", "timezone", "to_ts"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DownloadUrl = *all.DownloadUrl
	o.FromTs = *all.FromTs
	o.ResourceId = *all.ResourceId
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.TemplateVariables = *all.TemplateVariables
	o.Timeframe = all.Timeframe
	o.Timezone = *all.Timezone
	o.ToTs = *all.ToTs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
