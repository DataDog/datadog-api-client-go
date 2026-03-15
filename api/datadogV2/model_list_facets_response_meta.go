// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListFacetsResponseMeta Metadata for facets response.
type ListFacetsResponseMeta struct {
	// Total number of entities.
	TotalEntities int64 `json:"total_entities"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListFacetsResponseMeta instantiates a new ListFacetsResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListFacetsResponseMeta(totalEntities int64) *ListFacetsResponseMeta {
	this := ListFacetsResponseMeta{}
	this.TotalEntities = totalEntities
	return &this
}

// NewListFacetsResponseMetaWithDefaults instantiates a new ListFacetsResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListFacetsResponseMetaWithDefaults() *ListFacetsResponseMeta {
	this := ListFacetsResponseMeta{}
	return &this
}

// GetTotalEntities returns the TotalEntities field value.
func (o *ListFacetsResponseMeta) GetTotalEntities() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalEntities
}

// GetTotalEntitiesOk returns a tuple with the TotalEntities field value
// and a boolean to check if the value has been set.
func (o *ListFacetsResponseMeta) GetTotalEntitiesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalEntities, true
}

// SetTotalEntities sets field value.
func (o *ListFacetsResponseMeta) SetTotalEntities(v int64) {
	o.TotalEntities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListFacetsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["total_entities"] = o.TotalEntities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListFacetsResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalEntities *int64 `json:"total_entities"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalEntities == nil {
		return fmt.Errorf("required field total_entities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_entities"})
	} else {
		return err
	}
	o.TotalEntities = *all.TotalEntities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
