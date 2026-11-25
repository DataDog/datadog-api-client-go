// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneratedMetric Object for a single generated (derived) metric. Only returned for organizations with
// generated metrics enabled.
type GeneratedMetric struct {
	// The metric name for this resource.
	Id string `json:"id"`
	// Relationships for a generated metric.
	Relationships *GeneratedMetricRelationships `json:"relationships,omitempty"`
	// The generated metric resource type.
	Type GeneratedMetricType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGeneratedMetric instantiates a new GeneratedMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeneratedMetric(id string, typeVar GeneratedMetricType) *GeneratedMetric {
	this := GeneratedMetric{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewGeneratedMetricWithDefaults instantiates a new GeneratedMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeneratedMetricWithDefaults() *GeneratedMetric {
	this := GeneratedMetric{}
	var typeVar GeneratedMetricType = GENERATEDMETRICTYPE_GENERATED_METRIC
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *GeneratedMetric) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GeneratedMetric) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GeneratedMetric) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GeneratedMetric) GetRelationships() GeneratedMetricRelationships {
	if o == nil || o.Relationships == nil {
		var ret GeneratedMetricRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratedMetric) GetRelationshipsOk() (*GeneratedMetricRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GeneratedMetric) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given GeneratedMetricRelationships and assigns it to the Relationships field.
func (o *GeneratedMetric) SetRelationships(v GeneratedMetricRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *GeneratedMetric) GetType() GeneratedMetricType {
	if o == nil {
		var ret GeneratedMetricType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeneratedMetric) GetTypeOk() (*GeneratedMetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GeneratedMetric) SetType(v GeneratedMetricType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeneratedMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeneratedMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                       `json:"id"`
		Relationships *GeneratedMetricRelationships `json:"relationships,omitempty"`
		Type          *GeneratedMetricType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
