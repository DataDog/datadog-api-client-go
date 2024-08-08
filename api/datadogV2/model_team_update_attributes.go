// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamUpdateAttributes Team update attributes
type TeamUpdateAttributes struct {
	// Unicode representation of the avatar for the team, limited to a single grapheme
	Avatar datadog.NullableString `json:"avatar,omitempty"`
	// Banner selection for the team
	Banner datadog.NullableInt64 `json:"banner,omitempty"`
	// Free-form markdown description/content for the team's homepage
	Description *string `json:"description,omitempty"`
	// The team's identifier
	Handle string `json:"handle"`
	// The TeamUpdateAttributes handles.
	Handles *string `json:"handles,omitempty"`
	// Collection of hidden modules for the team
	HiddenModules []string `json:"hidden_modules,omitempty"`
	// The number of links belonging to the team
	LinkCount *int32 `json:"link_count,omitempty"`
	// The name of the team
	Name string `json:"name"`
	// A brief summary of the team
	Summary *string `json:"summary,omitempty"`
	// Collection of visible modules for the team
	VisibleModules []string `json:"visible_modules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamUpdateAttributes instantiates a new TeamUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamUpdateAttributes(handle string, name string) *TeamUpdateAttributes {
	this := TeamUpdateAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewTeamUpdateAttributesWithDefaults instantiates a new TeamUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamUpdateAttributesWithDefaults() *TeamUpdateAttributes {
	this := TeamUpdateAttributes{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamUpdateAttributes) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamUpdateAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasAvatar() bool {
	return o != nil && o.Avatar.IsSet()
}

// SetAvatar gets a reference to the given datadog.NullableString and assigns it to the Avatar field.
func (o *TeamUpdateAttributes) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil.
func (o *TeamUpdateAttributes) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil.
func (o *TeamUpdateAttributes) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamUpdateAttributes) GetBanner() int64 {
	if o == nil || o.Banner.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamUpdateAttributes) GetBannerOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasBanner() bool {
	return o != nil && o.Banner.IsSet()
}

// SetBanner gets a reference to the given datadog.NullableInt64 and assigns it to the Banner field.
func (o *TeamUpdateAttributes) SetBanner(v int64) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil.
func (o *TeamUpdateAttributes) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil.
func (o *TeamUpdateAttributes) UnsetBanner() {
	o.Banner.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value.
func (o *TeamUpdateAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TeamUpdateAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetHandles returns the Handles field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetHandles() string {
	if o == nil || o.Handles == nil {
		var ret string
		return ret
	}
	return *o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetHandlesOk() (*string, bool) {
	if o == nil || o.Handles == nil {
		return nil, false
	}
	return o.Handles, true
}

// HasHandles returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasHandles() bool {
	return o != nil && o.Handles != nil
}

// SetHandles gets a reference to the given string and assigns it to the Handles field.
func (o *TeamUpdateAttributes) SetHandles(v string) {
	o.Handles = &v
}

// GetHiddenModules returns the HiddenModules field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetHiddenModules() []string {
	if o == nil || o.HiddenModules == nil {
		var ret []string
		return ret
	}
	return o.HiddenModules
}

// GetHiddenModulesOk returns a tuple with the HiddenModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetHiddenModulesOk() (*[]string, bool) {
	if o == nil || o.HiddenModules == nil {
		return nil, false
	}
	return &o.HiddenModules, true
}

// HasHiddenModules returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasHiddenModules() bool {
	return o != nil && o.HiddenModules != nil
}

// SetHiddenModules gets a reference to the given []string and assigns it to the HiddenModules field.
func (o *TeamUpdateAttributes) SetHiddenModules(v []string) {
	o.HiddenModules = v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetLinkCount() int32 {
	if o == nil || o.LinkCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetLinkCountOk() (*int32, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int32 and assigns it to the LinkCount field.
func (o *TeamUpdateAttributes) SetLinkCount(v int32) {
	o.LinkCount = &v
}

// GetName returns the Name field value.
func (o *TeamUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TeamUpdateAttributes) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *TeamUpdateAttributes) SetSummary(v string) {
	o.Summary = &v
}

// GetVisibleModules returns the VisibleModules field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetVisibleModules() []string {
	if o == nil || o.VisibleModules == nil {
		var ret []string
		return ret
	}
	return o.VisibleModules
}

// GetVisibleModulesOk returns a tuple with the VisibleModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetVisibleModulesOk() (*[]string, bool) {
	if o == nil || o.VisibleModules == nil {
		return nil, false
	}
	return &o.VisibleModules, true
}

// HasVisibleModules returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasVisibleModules() bool {
	return o != nil && o.VisibleModules != nil
}

// SetVisibleModules gets a reference to the given []string and assigns it to the VisibleModules field.
func (o *TeamUpdateAttributes) SetVisibleModules(v []string) {
	o.VisibleModules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["handle"] = o.Handle
	if o.Handles != nil {
		toSerialize["handles"] = o.Handles
	}
	if o.HiddenModules != nil {
		toSerialize["hidden_modules"] = o.HiddenModules
	}
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	toSerialize["name"] = o.Name
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.VisibleModules != nil {
		toSerialize["visible_modules"] = o.VisibleModules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar         datadog.NullableString `json:"avatar,omitempty"`
		Banner         datadog.NullableInt64  `json:"banner,omitempty"`
		Description    *string                `json:"description,omitempty"`
		Handle         *string                `json:"handle"`
		Handles        *string                `json:"handles,omitempty"`
		HiddenModules  []string               `json:"hidden_modules,omitempty"`
		LinkCount      *int32                 `json:"link_count,omitempty"`
		Name           *string                `json:"name"`
		Summary        *string                `json:"summary,omitempty"`
		VisibleModules []string               `json:"visible_modules,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"avatar", "banner", "description", "handle", "handles", "hidden_modules", "link_count", "name", "summary", "visible_modules"})
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Banner = all.Banner
	o.Description = all.Description
	o.Handle = *all.Handle
	o.Handles = all.Handles
	o.HiddenModules = all.HiddenModules
	o.LinkCount = all.LinkCount
	o.Name = *all.Name
	o.Summary = all.Summary
	o.VisibleModules = all.VisibleModules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
