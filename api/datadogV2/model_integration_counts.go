// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationCounts Count of integrations by type.
type IntegrationCounts struct {
	// The filtered count of integrations.
	FilteredCount *int64 `json:"filtered_count,omitempty"`
	// The integration type.
	IntegrationType *string `json:"integration_type,omitempty"`
	// The total count of integrations.
	TotalCount *int64 `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationCounts instantiates a new IntegrationCounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationCounts() *IntegrationCounts {
	this := IntegrationCounts{}
	return &this
}

// NewIntegrationCountsWithDefaults instantiates a new IntegrationCounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationCountsWithDefaults() *IntegrationCounts {
	this := IntegrationCounts{}
	return &this
}

// GetFilteredCount returns the FilteredCount field value if set, zero value otherwise.
func (o *IntegrationCounts) GetFilteredCount() int64 {
	if o == nil || o.FilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.FilteredCount
}

// GetFilteredCountOk returns a tuple with the FilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCounts) GetFilteredCountOk() (*int64, bool) {
	if o == nil || o.FilteredCount == nil {
		return nil, false
	}
	return o.FilteredCount, true
}

// HasFilteredCount returns a boolean if a field has been set.
func (o *IntegrationCounts) HasFilteredCount() bool {
	return o != nil && o.FilteredCount != nil
}

// SetFilteredCount gets a reference to the given int64 and assigns it to the FilteredCount field.
func (o *IntegrationCounts) SetFilteredCount(v int64) {
	o.FilteredCount = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *IntegrationCounts) GetIntegrationType() string {
	if o == nil || o.IntegrationType == nil {
		var ret string
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCounts) GetIntegrationTypeOk() (*string, bool) {
	if o == nil || o.IntegrationType == nil {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *IntegrationCounts) HasIntegrationType() bool {
	return o != nil && o.IntegrationType != nil
}

// SetIntegrationType gets a reference to the given string and assigns it to the IntegrationType field.
func (o *IntegrationCounts) SetIntegrationType(v string) {
	o.IntegrationType = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *IntegrationCounts) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCounts) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *IntegrationCounts) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *IntegrationCounts) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FilteredCount != nil {
		toSerialize["filtered_count"] = o.FilteredCount
	}
	if o.IntegrationType != nil {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationCounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilteredCount   *int64  `json:"filtered_count,omitempty"`
		IntegrationType *string `json:"integration_type,omitempty"`
		TotalCount      *int64  `json:"total_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filtered_count", "integration_type", "total_count"})
	} else {
		return err
	}
	o.FilteredCount = all.FilteredCount
	o.IntegrationType = all.IntegrationType
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
