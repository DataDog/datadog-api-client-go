// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequest20270101OneOf4457459488 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequest20270101OneOf4457459488 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AuditQuery LogQueryDefinition `json:"audit_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendRegionsRequest20270101OneOf4457459488 instantiates a new GeomapBackendRegionsRequest20270101OneOf4457459488 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendRegionsRequest20270101OneOf4457459488(auditQuery LogQueryDefinition) *GeomapBackendRegionsRequest20270101OneOf4457459488 {
	this := GeomapBackendRegionsRequest20270101OneOf4457459488{}
	this.AuditQuery = auditQuery
	return &this
}

// NewGeomapBackendRegionsRequest20270101OneOf4457459488WithDefaults instantiates a new GeomapBackendRegionsRequest20270101OneOf4457459488 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendRegionsRequest20270101OneOf4457459488WithDefaults() *GeomapBackendRegionsRequest20270101OneOf4457459488 {
	this := GeomapBackendRegionsRequest20270101OneOf4457459488{}
	return &this
}

// GetAuditQuery returns the AuditQuery field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetAuditQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.AuditQuery
}

// GetAuditQueryOk returns a tuple with the AuditQuery field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetAuditQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditQuery, true
}

// SetAuditQuery sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) SetAuditQuery(v LogQueryDefinition) {
	o.AuditQuery = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetStyle() GeomapRequestStyle {
	if o == nil || o.Style == nil {
		var ret GeomapRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) GetStyleOk() (*GeomapRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapRequestStyle and assigns it to the Style field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) SetStyle(v GeomapRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendRegionsRequest20270101OneOf4457459488) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["audit_query"] = o.AuditQuery
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendRegionsRequest20270101OneOf4457459488) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuditQuery         *LogQueryDefinition       `json:"audit_query"`
		ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
		Style              *GeomapRequestStyle       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuditQuery == nil {
		return fmt.Errorf("required field audit_query missing")
	}

	hasInvalidField := false
	if all.AuditQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuditQuery = *all.AuditQuery
	o.ConditionalFormats = all.ConditionalFormats
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
