// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyPathTarget A reference to the range of steps between two nodes of the journey.
type ProductAnalyticsJourneyPathTarget struct {
	// Alias of the node the path ends at.
	End string `json:"end"`
	// Alias of the node the path starts at.
	Start string `json:"start"`
	// The discriminator identifying a target that references a range of steps.
	Type ProductAnalyticsJourneyPathTargetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyPathTarget instantiates a new ProductAnalyticsJourneyPathTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyPathTarget(end string, start string, typeVar ProductAnalyticsJourneyPathTargetType) *ProductAnalyticsJourneyPathTarget {
	this := ProductAnalyticsJourneyPathTarget{}
	this.End = end
	this.Start = start
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsJourneyPathTargetWithDefaults instantiates a new ProductAnalyticsJourneyPathTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyPathTargetWithDefaults() *ProductAnalyticsJourneyPathTarget {
	this := ProductAnalyticsJourneyPathTarget{}
	return &this
}

// GetEnd returns the End field value.
func (o *ProductAnalyticsJourneyPathTarget) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyPathTarget) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *ProductAnalyticsJourneyPathTarget) SetEnd(v string) {
	o.End = v
}

// GetStart returns the Start field value.
func (o *ProductAnalyticsJourneyPathTarget) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyPathTarget) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *ProductAnalyticsJourneyPathTarget) SetStart(v string) {
	o.Start = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsJourneyPathTarget) GetType() ProductAnalyticsJourneyPathTargetType {
	if o == nil {
		var ret ProductAnalyticsJourneyPathTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyPathTarget) GetTypeOk() (*ProductAnalyticsJourneyPathTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsJourneyPathTarget) SetType(v ProductAnalyticsJourneyPathTargetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyPathTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end"] = o.End
	toSerialize["start"] = o.Start
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyPathTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *string                                `json:"end"`
		Start *string                                `json:"start"`
		Type  *ProductAnalyticsJourneyPathTargetType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.End = *all.End
	o.Start = *all.Start
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
