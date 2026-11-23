// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationExtraField A recommendation field selected by the request's `extra_fields` patterns.
type CostRecommendationExtraField struct {
	// Dotted path of the matched field in the recommendation event's custom attributes.
	Path string `json:"path"`
	// Matched JSON value. Paths through arrays return the matched values as an array.
	Value interface{} `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationExtraField instantiates a new CostRecommendationExtraField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationExtraField(path string, value interface{}) *CostRecommendationExtraField {
	this := CostRecommendationExtraField{}
	this.Path = path
	this.Value = value
	return &this
}

// NewCostRecommendationExtraFieldWithDefaults instantiates a new CostRecommendationExtraField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationExtraFieldWithDefaults() *CostRecommendationExtraField {
	this := CostRecommendationExtraField{}
	return &this
}

// GetPath returns the Path field value.
func (o *CostRecommendationExtraField) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationExtraField) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *CostRecommendationExtraField) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value.
func (o *CostRecommendationExtraField) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationExtraField) GetValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CostRecommendationExtraField) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationExtraField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["path"] = o.Path
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationExtraField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Path  *string      `json:"path"`
		Value *interface{} `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"path", "value"})
	} else {
		return err
	}
	o.Path = *all.Path
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
