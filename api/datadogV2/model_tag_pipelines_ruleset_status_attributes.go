// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPipelinesRulesetStatusAttributes Attributes for a tag pipeline ruleset status.
type TagPipelinesRulesetStatusAttributes struct {
	// The processing status of the ruleset.
	ProcessingStatus string `json:"processing_status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPipelinesRulesetStatusAttributes instantiates a new TagPipelinesRulesetStatusAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPipelinesRulesetStatusAttributes(processingStatus string) *TagPipelinesRulesetStatusAttributes {
	this := TagPipelinesRulesetStatusAttributes{}
	this.ProcessingStatus = processingStatus
	return &this
}

// NewTagPipelinesRulesetStatusAttributesWithDefaults instantiates a new TagPipelinesRulesetStatusAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPipelinesRulesetStatusAttributesWithDefaults() *TagPipelinesRulesetStatusAttributes {
	this := TagPipelinesRulesetStatusAttributes{}
	return &this
}

// GetProcessingStatus returns the ProcessingStatus field value.
func (o *TagPipelinesRulesetStatusAttributes) GetProcessingStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value
// and a boolean to check if the value has been set.
func (o *TagPipelinesRulesetStatusAttributes) GetProcessingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingStatus, true
}

// SetProcessingStatus sets field value.
func (o *TagPipelinesRulesetStatusAttributes) SetProcessingStatus(v string) {
	o.ProcessingStatus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPipelinesRulesetStatusAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["processing_status"] = o.ProcessingStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPipelinesRulesetStatusAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProcessingStatus *string `json:"processing_status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProcessingStatus == nil {
		return fmt.Errorf("required field processing_status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"processing_status"})
	} else {
		return err
	}
	o.ProcessingStatus = *all.ProcessingStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
