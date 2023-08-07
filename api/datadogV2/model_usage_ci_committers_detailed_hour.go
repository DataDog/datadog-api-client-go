// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCICommittersDetailedHour The response containing attributions for the specified org.
type UsageCICommittersDetailedHour struct {
	// The response containing CI Committers Detailed attributes for specified organization.
	Attributes *UsageCICommittersDetailedAttributes `json:"attributes,omitempty"`
	// A unique ID for the JSON API resource.
	Id *string `json:"id,omitempty"`
	// The type of shape of the data.
	Type *UsageDataPointType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageCICommittersDetailedHour instantiates a new UsageCICommittersDetailedHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCICommittersDetailedHour() *UsageCICommittersDetailedHour {
	this := UsageCICommittersDetailedHour{}
	var typeVar UsageDataPointType = USAGEDATAPOINTTYPE_USAGE_DATA_POINT
	this.Type = &typeVar
	return &this
}

// NewUsageCICommittersDetailedHourWithDefaults instantiates a new UsageCICommittersDetailedHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCICommittersDetailedHourWithDefaults() *UsageCICommittersDetailedHour {
	this := UsageCICommittersDetailedHour{}
	var typeVar UsageDataPointType = USAGEDATAPOINTTYPE_USAGE_DATA_POINT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedHour) GetAttributes() UsageCICommittersDetailedAttributes {
	if o == nil || o.Attributes == nil {
		var ret UsageCICommittersDetailedAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedHour) GetAttributesOk() (*UsageCICommittersDetailedAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedHour) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UsageCICommittersDetailedAttributes and assigns it to the Attributes field.
func (o *UsageCICommittersDetailedHour) SetAttributes(v UsageCICommittersDetailedAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedHour) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedHour) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedHour) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UsageCICommittersDetailedHour) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedHour) GetType() UsageDataPointType {
	if o == nil || o.Type == nil {
		var ret UsageDataPointType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedHour) GetTypeOk() (*UsageDataPointType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedHour) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UsageDataPointType and assigns it to the Type field.
func (o *UsageCICommittersDetailedHour) SetType(v UsageDataPointType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCICommittersDetailedHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageCICommittersDetailedHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UsageCICommittersDetailedAttributes `json:"attributes,omitempty"`
		Id         *string                              `json:"id,omitempty"`
		Type       *UsageDataPointType                  `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
