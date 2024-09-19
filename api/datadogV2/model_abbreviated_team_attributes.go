// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// AbbreviatedTeamAttributes The definition of `AbbreviatedTeamAttributes` object.
type AbbreviatedTeamAttributes struct {
	// Unicode representation of the avatar for the team, limited to a single grapheme
	Avatar datadog.NullableString `json:"avatar,omitempty"`
	// Banner selection for the team
	Banner *int64 `json:"banner,omitempty"`
	// The team's identifier
	Handle string `json:"handle"`
	// The `AbbreviatedTeamAttributes` `handles`.
	Handles *string `json:"handles,omitempty"`
	// The `AbbreviatedTeamAttributes` `is_open_membership`.
	IsOpenMembership *bool `json:"is_open_membership,omitempty"`
	// The name of the team
	Name string `json:"name"`
	// A brief summary of the team
	Summary *string `json:"summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewAbbreviatedTeamAttributes instantiates a new AbbreviatedTeamAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAbbreviatedTeamAttributes(handle string, name string) *AbbreviatedTeamAttributes {
	this := AbbreviatedTeamAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewAbbreviatedTeamAttributesWithDefaults instantiates a new AbbreviatedTeamAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAbbreviatedTeamAttributesWithDefaults() *AbbreviatedTeamAttributes {
	this := AbbreviatedTeamAttributes{}
	return &this
}
// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbbreviatedTeamAttributes) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AbbreviatedTeamAttributes) GetAvatarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *AbbreviatedTeamAttributes) HasAvatar() bool {
	return o != nil && o.Avatar.IsSet()
}

// SetAvatar gets a reference to the given datadog.NullableString and assigns it to the Avatar field.
func (o *AbbreviatedTeamAttributes) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil.
func (o *AbbreviatedTeamAttributes) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil.
func (o *AbbreviatedTeamAttributes) UnsetAvatar() {
	o.Avatar.Unset()
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *AbbreviatedTeamAttributes) GetBanner() int64 {
	if o == nil || o.Banner == nil {
		var ret int64
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetBannerOk() (*int64, bool) {
	if o == nil || o.Banner == nil {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *AbbreviatedTeamAttributes) HasBanner() bool {
	return o != nil && o.Banner != nil
}

// SetBanner gets a reference to the given int64 and assigns it to the Banner field.
func (o *AbbreviatedTeamAttributes) SetBanner(v int64) {
	o.Banner = &v
}


// GetHandle returns the Handle field value.
func (o *AbbreviatedTeamAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *AbbreviatedTeamAttributes) SetHandle(v string) {
	o.Handle = v
}


// GetHandles returns the Handles field value if set, zero value otherwise.
func (o *AbbreviatedTeamAttributes) GetHandles() string {
	if o == nil || o.Handles == nil {
		var ret string
		return ret
	}
	return *o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetHandlesOk() (*string, bool) {
	if o == nil || o.Handles == nil {
		return nil, false
	}
	return o.Handles, true
}

// HasHandles returns a boolean if a field has been set.
func (o *AbbreviatedTeamAttributes) HasHandles() bool {
	return o != nil && o.Handles != nil
}

// SetHandles gets a reference to the given string and assigns it to the Handles field.
func (o *AbbreviatedTeamAttributes) SetHandles(v string) {
	o.Handles = &v
}


// GetIsOpenMembership returns the IsOpenMembership field value if set, zero value otherwise.
func (o *AbbreviatedTeamAttributes) GetIsOpenMembership() bool {
	if o == nil || o.IsOpenMembership == nil {
		var ret bool
		return ret
	}
	return *o.IsOpenMembership
}

// GetIsOpenMembershipOk returns a tuple with the IsOpenMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetIsOpenMembershipOk() (*bool, bool) {
	if o == nil || o.IsOpenMembership == nil {
		return nil, false
	}
	return o.IsOpenMembership, true
}

// HasIsOpenMembership returns a boolean if a field has been set.
func (o *AbbreviatedTeamAttributes) HasIsOpenMembership() bool {
	return o != nil && o.IsOpenMembership != nil
}

// SetIsOpenMembership gets a reference to the given bool and assigns it to the IsOpenMembership field.
func (o *AbbreviatedTeamAttributes) SetIsOpenMembership(v bool) {
	o.IsOpenMembership = &v
}


// GetName returns the Name field value.
func (o *AbbreviatedTeamAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AbbreviatedTeamAttributes) SetName(v string) {
	o.Name = v
}


// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AbbreviatedTeamAttributes) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbbreviatedTeamAttributes) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AbbreviatedTeamAttributes) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AbbreviatedTeamAttributes) SetSummary(v string) {
	o.Summary = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o AbbreviatedTeamAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Banner != nil {
		toSerialize["banner"] = o.Banner
	}
	toSerialize["handle"] = o.Handle
	if o.Handles != nil {
		toSerialize["handles"] = o.Handles
	}
	if o.IsOpenMembership != nil {
		toSerialize["is_open_membership"] = o.IsOpenMembership
	}
	toSerialize["name"] = o.Name
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AbbreviatedTeamAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar datadog.NullableString `json:"avatar,omitempty"`
		Banner *int64 `json:"banner,omitempty"`
		Handle *string `json:"handle"`
		Handles *string `json:"handles,omitempty"`
		IsOpenMembership *bool `json:"is_open_membership,omitempty"`
		Name *string `json:"name"`
		Summary *string `json:"summary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "avatar", "banner", "handle", "handles", "is_open_membership", "name", "summary",  })
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Banner = all.Banner
	o.Handle = *all.Handle
	o.Handles = all.Handles
	o.IsOpenMembership = all.IsOpenMembership
	o.Name = *all.Name
	o.Summary = all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
