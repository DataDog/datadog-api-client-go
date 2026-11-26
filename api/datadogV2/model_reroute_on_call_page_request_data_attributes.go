// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RerouteOnCallPageRequestDataAttributes Details about where to reroute the On-Call Page.
type RerouteOnCallPageRequestDataAttributes struct {
	// An optional comment explaining why the Page is being rerouted.
	Comment *string `json:"comment,omitempty"`
	// Information about the target to notify (such as a team or user).
	Target CreatePageRequestDataAttributesTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRerouteOnCallPageRequestDataAttributes instantiates a new RerouteOnCallPageRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRerouteOnCallPageRequestDataAttributes(target CreatePageRequestDataAttributesTarget) *RerouteOnCallPageRequestDataAttributes {
	this := RerouteOnCallPageRequestDataAttributes{}
	this.Target = target
	return &this
}

// NewRerouteOnCallPageRequestDataAttributesWithDefaults instantiates a new RerouteOnCallPageRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRerouteOnCallPageRequestDataAttributesWithDefaults() *RerouteOnCallPageRequestDataAttributes {
	this := RerouteOnCallPageRequestDataAttributes{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RerouteOnCallPageRequestDataAttributes) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RerouteOnCallPageRequestDataAttributes) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RerouteOnCallPageRequestDataAttributes) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RerouteOnCallPageRequestDataAttributes) SetComment(v string) {
	o.Comment = &v
}

// GetTarget returns the Target field value.
func (o *RerouteOnCallPageRequestDataAttributes) GetTarget() CreatePageRequestDataAttributesTarget {
	if o == nil {
		var ret CreatePageRequestDataAttributesTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *RerouteOnCallPageRequestDataAttributes) GetTargetOk() (*CreatePageRequestDataAttributesTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *RerouteOnCallPageRequestDataAttributes) SetTarget(v CreatePageRequestDataAttributesTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RerouteOnCallPageRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RerouteOnCallPageRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comment *string                                `json:"comment,omitempty"`
		Target  *CreatePageRequestDataAttributesTarget `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comment", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Comment = all.Comment
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
