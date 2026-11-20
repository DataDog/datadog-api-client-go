// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchFilters Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchFilters struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AudienceFilters *AudienceFilters `json:"audience_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	StringFilter *string `json:"string_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearchFilters instantiates a new RetentionSearchFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearchFilters() *RetentionSearchFilters {
	this := RetentionSearchFilters{}
	return &this
}

// NewRetentionSearchFiltersWithDefaults instantiates a new RetentionSearchFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchFiltersWithDefaults() *RetentionSearchFilters {
	this := RetentionSearchFilters{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *RetentionSearchFilters) GetAudienceFilters() AudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret AudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchFilters) GetAudienceFiltersOk() (*AudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *RetentionSearchFilters) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given AudienceFilters and assigns it to the AudienceFilters field.
func (o *RetentionSearchFilters) SetAudienceFilters(v AudienceFilters) {
	o.AudienceFilters = &v
}

// GetStringFilter returns the StringFilter field value if set, zero value otherwise.
func (o *RetentionSearchFilters) GetStringFilter() string {
	if o == nil || o.StringFilter == nil {
		var ret string
		return ret
	}
	return *o.StringFilter
}

// GetStringFilterOk returns a tuple with the StringFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchFilters) GetStringFilterOk() (*string, bool) {
	if o == nil || o.StringFilter == nil {
		return nil, false
	}
	return o.StringFilter, true
}

// HasStringFilter returns a boolean if a field has been set.
func (o *RetentionSearchFilters) HasStringFilter() bool {
	return o != nil && o.StringFilter != nil
}

// SetStringFilter gets a reference to the given string and assigns it to the StringFilter field.
func (o *RetentionSearchFilters) SetStringFilter(v string) {
	o.StringFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearchFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.StringFilter != nil {
		toSerialize["string_filter"] = o.StringFilter
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearchFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *AudienceFilters `json:"audience_filters,omitempty"`
		StringFilter    *string          `json:"string_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.StringFilter = all.StringFilter

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
