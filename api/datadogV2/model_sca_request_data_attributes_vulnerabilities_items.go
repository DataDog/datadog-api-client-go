// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesVulnerabilitiesItems
type ScaRequestDataAttributesVulnerabilitiesItems struct {
	//
	Affects []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems `json:"affects,omitempty"`
	//
	BomRef *string `json:"bom_ref,omitempty"`
	//
	Id *string `json:"id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesVulnerabilitiesItems instantiates a new ScaRequestDataAttributesVulnerabilitiesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesVulnerabilitiesItems() *ScaRequestDataAttributesVulnerabilitiesItems {
	this := ScaRequestDataAttributesVulnerabilitiesItems{}
	return &this
}

// NewScaRequestDataAttributesVulnerabilitiesItemsWithDefaults instantiates a new ScaRequestDataAttributesVulnerabilitiesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesVulnerabilitiesItemsWithDefaults() *ScaRequestDataAttributesVulnerabilitiesItems {
	this := ScaRequestDataAttributesVulnerabilitiesItems{}
	return &this
}

// GetAffects returns the Affects field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetAffects() []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems {
	if o == nil || o.Affects == nil {
		var ret []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems
		return ret
	}
	return o.Affects
}

// GetAffectsOk returns a tuple with the Affects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetAffectsOk() (*[]ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems, bool) {
	if o == nil || o.Affects == nil {
		return nil, false
	}
	return &o.Affects, true
}

// HasAffects returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) HasAffects() bool {
	return o != nil && o.Affects != nil
}

// SetAffects gets a reference to the given []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems and assigns it to the Affects field.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) SetAffects(v []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) {
	o.Affects = v
}

// GetBomRef returns the BomRef field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetBomRef() string {
	if o == nil || o.BomRef == nil {
		var ret string
		return ret
	}
	return *o.BomRef
}

// GetBomRefOk returns a tuple with the BomRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetBomRefOk() (*string, bool) {
	if o == nil || o.BomRef == nil {
		return nil, false
	}
	return o.BomRef, true
}

// HasBomRef returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) HasBomRef() bool {
	return o != nil && o.BomRef != nil
}

// SetBomRef gets a reference to the given string and assigns it to the BomRef field.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) SetBomRef(v string) {
	o.BomRef = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) SetId(v string) {
	o.Id = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesVulnerabilitiesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Affects != nil {
		toSerialize["affects"] = o.Affects
	}
	if o.BomRef != nil {
		toSerialize["bom_ref"] = o.BomRef
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesVulnerabilitiesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Affects []ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems `json:"affects,omitempty"`
		BomRef  *string                                                    `json:"bom_ref,omitempty"`
		Id      *string                                                    `json:"id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"affects", "bom_ref", "id"})
	} else {
		return err
	}
	o.Affects = all.Affects
	o.BomRef = all.BomRef
	o.Id = all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
