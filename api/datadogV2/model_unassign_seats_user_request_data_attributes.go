// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnassignSeatsUserRequestDataAttributes
type UnassignSeatsUserRequestDataAttributes struct {
	//
	ProductCode *string `json:"product_code,omitempty"`
	//
	UserUuids []string `json:"user_uuids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUnassignSeatsUserRequestDataAttributes instantiates a new UnassignSeatsUserRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnassignSeatsUserRequestDataAttributes() *UnassignSeatsUserRequestDataAttributes {
	this := UnassignSeatsUserRequestDataAttributes{}
	return &this
}

// NewUnassignSeatsUserRequestDataAttributesWithDefaults instantiates a new UnassignSeatsUserRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnassignSeatsUserRequestDataAttributesWithDefaults() *UnassignSeatsUserRequestDataAttributes {
	this := UnassignSeatsUserRequestDataAttributes{}
	return &this
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *UnassignSeatsUserRequestDataAttributes) GetProductCode() string {
	if o == nil || o.ProductCode == nil {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnassignSeatsUserRequestDataAttributes) GetProductCodeOk() (*string, bool) {
	if o == nil || o.ProductCode == nil {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *UnassignSeatsUserRequestDataAttributes) HasProductCode() bool {
	return o != nil && o.ProductCode != nil
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *UnassignSeatsUserRequestDataAttributes) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetUserUuids returns the UserUuids field value if set, zero value otherwise.
func (o *UnassignSeatsUserRequestDataAttributes) GetUserUuids() []string {
	if o == nil || o.UserUuids == nil {
		var ret []string
		return ret
	}
	return o.UserUuids
}

// GetUserUuidsOk returns a tuple with the UserUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnassignSeatsUserRequestDataAttributes) GetUserUuidsOk() (*[]string, bool) {
	if o == nil || o.UserUuids == nil {
		return nil, false
	}
	return &o.UserUuids, true
}

// HasUserUuids returns a boolean if a field has been set.
func (o *UnassignSeatsUserRequestDataAttributes) HasUserUuids() bool {
	return o != nil && o.UserUuids != nil
}

// SetUserUuids gets a reference to the given []string and assigns it to the UserUuids field.
func (o *UnassignSeatsUserRequestDataAttributes) SetUserUuids(v []string) {
	o.UserUuids = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UnassignSeatsUserRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ProductCode != nil {
		toSerialize["product_code"] = o.ProductCode
	}
	if o.UserUuids != nil {
		toSerialize["user_uuids"] = o.UserUuids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UnassignSeatsUserRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProductCode *string  `json:"product_code,omitempty"`
		UserUuids   []string `json:"user_uuids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"product_code", "user_uuids"})
	} else {
		return err
	}
	o.ProductCode = all.ProductCode
	o.UserUuids = all.UserUuids

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
