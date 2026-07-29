// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumConfigCreateAttributes Attributes of the RUM configuration to create.
type RumConfigCreateAttributes struct {
	// Whether application tags are enforced for the RUM applications in the organization.
	EnforcedApplicationTags bool `json:"enforced_application_tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumConfigCreateAttributes instantiates a new RumConfigCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumConfigCreateAttributes(enforcedApplicationTags bool) *RumConfigCreateAttributes {
	this := RumConfigCreateAttributes{}
	this.EnforcedApplicationTags = enforcedApplicationTags
	return &this
}

// NewRumConfigCreateAttributesWithDefaults instantiates a new RumConfigCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumConfigCreateAttributesWithDefaults() *RumConfigCreateAttributes {
	this := RumConfigCreateAttributes{}
	return &this
}

// GetEnforcedApplicationTags returns the EnforcedApplicationTags field value.
func (o *RumConfigCreateAttributes) GetEnforcedApplicationTags() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnforcedApplicationTags
}

// GetEnforcedApplicationTagsOk returns a tuple with the EnforcedApplicationTags field value
// and a boolean to check if the value has been set.
func (o *RumConfigCreateAttributes) GetEnforcedApplicationTagsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforcedApplicationTags, true
}

// SetEnforcedApplicationTags sets field value.
func (o *RumConfigCreateAttributes) SetEnforcedApplicationTags(v bool) {
	o.EnforcedApplicationTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumConfigCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enforced_application_tags"] = o.EnforcedApplicationTags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumConfigCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnforcedApplicationTags *bool `json:"enforced_application_tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnforcedApplicationTags == nil {
		return fmt.Errorf("required field enforced_application_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced_application_tags"})
	} else {
		return err
	}
	o.EnforcedApplicationTags = *all.EnforcedApplicationTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
