// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListActionConnectionsResponseMeta Metadata for the list connections response.
type ListActionConnectionsResponseMeta struct {
	// Count of integrations by type.
	IntegrationCounts []IntegrationCounts `json:"integration_counts,omitempty"`
	// The total number of connections.
	Total int64 `json:"total"`
	// The total number of connections that match the specified filters.
	TotalFiltered int64 `json:"total_filtered"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListActionConnectionsResponseMeta instantiates a new ListActionConnectionsResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListActionConnectionsResponseMeta(total int64, totalFiltered int64) *ListActionConnectionsResponseMeta {
	this := ListActionConnectionsResponseMeta{}
	this.Total = total
	this.TotalFiltered = totalFiltered
	return &this
}

// NewListActionConnectionsResponseMetaWithDefaults instantiates a new ListActionConnectionsResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListActionConnectionsResponseMetaWithDefaults() *ListActionConnectionsResponseMeta {
	this := ListActionConnectionsResponseMeta{}
	return &this
}

// GetIntegrationCounts returns the IntegrationCounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListActionConnectionsResponseMeta) GetIntegrationCounts() []IntegrationCounts {
	if o == nil {
		var ret []IntegrationCounts
		return ret
	}
	return o.IntegrationCounts
}

// GetIntegrationCountsOk returns a tuple with the IntegrationCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ListActionConnectionsResponseMeta) GetIntegrationCountsOk() (*[]IntegrationCounts, bool) {
	if o == nil || o.IntegrationCounts == nil {
		return nil, false
	}
	return &o.IntegrationCounts, true
}

// HasIntegrationCounts returns a boolean if a field has been set.
func (o *ListActionConnectionsResponseMeta) HasIntegrationCounts() bool {
	return o != nil && o.IntegrationCounts != nil
}

// SetIntegrationCounts gets a reference to the given []IntegrationCounts and assigns it to the IntegrationCounts field.
func (o *ListActionConnectionsResponseMeta) SetIntegrationCounts(v []IntegrationCounts) {
	o.IntegrationCounts = v
}

// GetTotal returns the Total field value.
func (o *ListActionConnectionsResponseMeta) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListActionConnectionsResponseMeta) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ListActionConnectionsResponseMeta) SetTotal(v int64) {
	o.Total = v
}

// GetTotalFiltered returns the TotalFiltered field value.
func (o *ListActionConnectionsResponseMeta) GetTotalFiltered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFiltered
}

// GetTotalFilteredOk returns a tuple with the TotalFiltered field value
// and a boolean to check if the value has been set.
func (o *ListActionConnectionsResponseMeta) GetTotalFilteredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFiltered, true
}

// SetTotalFiltered sets field value.
func (o *ListActionConnectionsResponseMeta) SetTotalFiltered(v int64) {
	o.TotalFiltered = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListActionConnectionsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IntegrationCounts != nil {
		toSerialize["integration_counts"] = o.IntegrationCounts
	}
	toSerialize["total"] = o.Total
	toSerialize["total_filtered"] = o.TotalFiltered

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListActionConnectionsResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IntegrationCounts []IntegrationCounts `json:"integration_counts,omitempty"`
		Total             *int64              `json:"total"`
		TotalFiltered     *int64              `json:"total_filtered"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.TotalFiltered == nil {
		return fmt.Errorf("required field total_filtered missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"integration_counts", "total", "total_filtered"})
	} else {
		return err
	}
	o.IntegrationCounts = all.IntegrationCounts
	o.Total = *all.Total
	o.TotalFiltered = *all.TotalFiltered

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
