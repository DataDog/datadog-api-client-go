// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDependencyAttributesResponse Attributes for dataset dependency.
type SecurityMonitoringDatasetDependencyAttributesResponse struct {
	// The count of resources that depend on the dataset.
	Count int64 `json:"count"`
	// Array of IDs of resources that depend on the dataset.
	Ids []string `json:"ids"`
	// The type of resource that depends on the dataset.
	ResourceType string `json:"resource_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDependencyAttributesResponse instantiates a new SecurityMonitoringDatasetDependencyAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDependencyAttributesResponse(count int64, ids []string, resourceType string) *SecurityMonitoringDatasetDependencyAttributesResponse {
	this := SecurityMonitoringDatasetDependencyAttributesResponse{}
	this.Count = count
	this.Ids = ids
	this.ResourceType = resourceType
	return &this
}

// NewSecurityMonitoringDatasetDependencyAttributesResponseWithDefaults instantiates a new SecurityMonitoringDatasetDependencyAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDependencyAttributesResponseWithDefaults() *SecurityMonitoringDatasetDependencyAttributesResponse {
	this := SecurityMonitoringDatasetDependencyAttributesResponse{}
	return &this
}

// GetCount returns the Count field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) SetCount(v int64) {
	o.Count = v
}

// GetIds returns the Ids field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) SetIds(v []string) {
	o.Ids = v
}

// GetResourceType returns the ResourceType field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) SetResourceType(v string) {
	o.ResourceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDependencyAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["ids"] = o.Ids
	toSerialize["resource_type"] = o.ResourceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetDependencyAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count        *int64    `json:"count"`
		Ids          *[]string `json:"ids"`
		ResourceType *string   `json:"resource_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Ids == nil {
		return fmt.Errorf("required field ids missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "ids", "resource_type"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Ids = *all.Ids
	o.ResourceType = *all.ResourceType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
