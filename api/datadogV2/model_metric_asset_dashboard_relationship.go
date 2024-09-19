// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// MetricAssetDashboardRelationship An object of type `dashboard` that can be referenced in the `included` data.
type MetricAssetDashboardRelationship struct {
	// The related dashboard's ID.
	Id *string `json:"id,omitempty"`
	// Dashboard resource type.
	Type *MetricDashboardType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewMetricAssetDashboardRelationship instantiates a new MetricAssetDashboardRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricAssetDashboardRelationship() *MetricAssetDashboardRelationship {
	this := MetricAssetDashboardRelationship{}
	return &this
}

// NewMetricAssetDashboardRelationshipWithDefaults instantiates a new MetricAssetDashboardRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricAssetDashboardRelationshipWithDefaults() *MetricAssetDashboardRelationship {
	this := MetricAssetDashboardRelationship{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricAssetDashboardRelationship) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetDashboardRelationship) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricAssetDashboardRelationship) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricAssetDashboardRelationship) SetId(v string) {
	o.Id = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricAssetDashboardRelationship) GetType() MetricDashboardType {
	if o == nil || o.Type == nil {
		var ret MetricDashboardType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetDashboardRelationship) GetTypeOk() (*MetricDashboardType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricAssetDashboardRelationship) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MetricDashboardType and assigns it to the Type field.
func (o *MetricAssetDashboardRelationship) SetType(v MetricDashboardType) {
	o.Type = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o MetricAssetDashboardRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricAssetDashboardRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id,omitempty"`
		Type *MetricDashboardType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "id", "type",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.Type != nil &&!all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
