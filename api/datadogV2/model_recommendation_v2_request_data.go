// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationV2RequestData JSON:API resource object for the SPA v2 recommendation request.
type RecommendationV2RequestData struct {
	// Attributes for requesting SPA recommendations by forwarding a Spark job's raw arguments
	// instead of a pre-computed shard.
	Attributes RecommendationV2RequestAttributes `json:"attributes"`
	// JSON:API resource type for the SPA v2 recommendation request.
	Type RecommendationV2RequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationV2RequestData instantiates a new RecommendationV2RequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationV2RequestData(attributes RecommendationV2RequestAttributes, typeVar RecommendationV2RequestType) *RecommendationV2RequestData {
	this := RecommendationV2RequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewRecommendationV2RequestDataWithDefaults instantiates a new RecommendationV2RequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationV2RequestDataWithDefaults() *RecommendationV2RequestData {
	this := RecommendationV2RequestData{}
	var typeVar RecommendationV2RequestType = RECOMMENDATIONV2REQUESTTYPE_RECOMMENDATION_V2_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RecommendationV2RequestData) GetAttributes() RecommendationV2RequestAttributes {
	if o == nil {
		var ret RecommendationV2RequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RecommendationV2RequestData) GetAttributesOk() (*RecommendationV2RequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RecommendationV2RequestData) SetAttributes(v RecommendationV2RequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *RecommendationV2RequestData) GetType() RecommendationV2RequestType {
	if o == nil {
		var ret RecommendationV2RequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecommendationV2RequestData) GetTypeOk() (*RecommendationV2RequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RecommendationV2RequestData) SetType(v RecommendationV2RequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationV2RequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationV2RequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RecommendationV2RequestAttributes `json:"attributes"`
		Type       *RecommendationV2RequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
