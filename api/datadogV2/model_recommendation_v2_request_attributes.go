// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationV2RequestAttributes Attributes for requesting SPA recommendations by forwarding a Spark job's raw arguments
// instead of a precomputed shard.
type RecommendationV2RequestAttributes struct {
	// Raw, unfiltered Spark job arguments as submitted (for example, `--org_id=2`).
	// SPA determines which arguments are relevant for the given service.
	Arguments []string `json:"arguments"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationV2RequestAttributes instantiates a new RecommendationV2RequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationV2RequestAttributes(arguments []string) *RecommendationV2RequestAttributes {
	this := RecommendationV2RequestAttributes{}
	this.Arguments = arguments
	return &this
}

// NewRecommendationV2RequestAttributesWithDefaults instantiates a new RecommendationV2RequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationV2RequestAttributesWithDefaults() *RecommendationV2RequestAttributes {
	this := RecommendationV2RequestAttributes{}
	return &this
}

// GetArguments returns the Arguments field value.
func (o *RecommendationV2RequestAttributes) GetArguments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *RecommendationV2RequestAttributes) GetArgumentsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value.
func (o *RecommendationV2RequestAttributes) SetArguments(v []string) {
	o.Arguments = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationV2RequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arguments"] = o.Arguments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationV2RequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments *[]string `json:"arguments"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Arguments == nil {
		return fmt.Errorf("required field arguments missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments"})
	} else {
		return err
	}
	o.Arguments = *all.Arguments

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
