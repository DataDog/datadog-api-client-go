// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfluencePostmortemSettings Settings for a postmortem template stored in Confluence. Required when `location` is `confluence`.
type ConfluencePostmortemSettings struct {
	// The ID of the Confluence integration account.
	AccountId string `json:"account_id"`
	// The ID of the parent Confluence page under which postmortems are created.
	ParentId datadog.NullableString `json:"parent_id,omitempty"`
	// The ID of the Confluence space where postmortems are created.
	SpaceId string `json:"space_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfluencePostmortemSettings instantiates a new ConfluencePostmortemSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfluencePostmortemSettings(accountId string, spaceId string) *ConfluencePostmortemSettings {
	this := ConfluencePostmortemSettings{}
	this.AccountId = accountId
	this.SpaceId = spaceId
	return &this
}

// NewConfluencePostmortemSettingsWithDefaults instantiates a new ConfluencePostmortemSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfluencePostmortemSettingsWithDefaults() *ConfluencePostmortemSettings {
	this := ConfluencePostmortemSettings{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *ConfluencePostmortemSettings) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ConfluencePostmortemSettings) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *ConfluencePostmortemSettings) SetAccountId(v string) {
	o.AccountId = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfluencePostmortemSettings) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ConfluencePostmortemSettings) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ConfluencePostmortemSettings) HasParentId() bool {
	return o != nil && o.ParentId.IsSet()
}

// SetParentId gets a reference to the given datadog.NullableString and assigns it to the ParentId field.
func (o *ConfluencePostmortemSettings) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil.
func (o *ConfluencePostmortemSettings) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil.
func (o *ConfluencePostmortemSettings) UnsetParentId() {
	o.ParentId.Unset()
}

// GetSpaceId returns the SpaceId field value.
func (o *ConfluencePostmortemSettings) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *ConfluencePostmortemSettings) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value.
func (o *ConfluencePostmortemSettings) SetSpaceId(v string) {
	o.SpaceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfluencePostmortemSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	toSerialize["space_id"] = o.SpaceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfluencePostmortemSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string                `json:"account_id"`
		ParentId  datadog.NullableString `json:"parent_id,omitempty"`
		SpaceId   *string                `json:"space_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.SpaceId == nil {
		return fmt.Errorf("required field space_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "parent_id", "space_id"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.ParentId = all.ParentId
	o.SpaceId = *all.SpaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
