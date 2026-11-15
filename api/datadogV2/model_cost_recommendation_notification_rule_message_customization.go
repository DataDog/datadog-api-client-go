// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageCustomization Which sections the digest shows. Sent as a whole: sections left out of a request use their defaults, which
// show everything.
type CostRecommendationNotificationRuleMessageCustomization struct {
	// Links the digest shows. `view_in_ccm` is the link to review the recommendations in Cloud Cost Management.
	Actions *CostRecommendationNotificationRuleMessageActions `json:"actions,omitempty"`
	// What the digest header shows.
	Header *CostRecommendationNotificationRuleMessageCustomizationHeader `json:"header,omitempty"`
	// Tag keys whose values are shown on each recommendation in the digest.
	IncludeTags []string `json:"include_tags,omitempty"`
	// How the message is customized.
	Mode *CostRecommendationNotificationRuleMessageCustomizationMode `json:"mode,omitempty"`
	// What each recommendation in the digest shows.
	PerRec *CostRecommendationNotificationRuleMessageCustomizationPerRec `json:"per_rec,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleMessageCustomization instantiates a new CostRecommendationNotificationRuleMessageCustomization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleMessageCustomization() *CostRecommendationNotificationRuleMessageCustomization {
	this := CostRecommendationNotificationRuleMessageCustomization{}
	return &this
}

// NewCostRecommendationNotificationRuleMessageCustomizationWithDefaults instantiates a new CostRecommendationNotificationRuleMessageCustomization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleMessageCustomizationWithDefaults() *CostRecommendationNotificationRuleMessageCustomization {
	this := CostRecommendationNotificationRuleMessageCustomization{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetActions() CostRecommendationNotificationRuleMessageActions {
	if o == nil || o.Actions == nil {
		var ret CostRecommendationNotificationRuleMessageActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetActionsOk() (*CostRecommendationNotificationRuleMessageActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given CostRecommendationNotificationRuleMessageActions and assigns it to the Actions field.
func (o *CostRecommendationNotificationRuleMessageCustomization) SetActions(v CostRecommendationNotificationRuleMessageActions) {
	o.Actions = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetHeader() CostRecommendationNotificationRuleMessageCustomizationHeader {
	if o == nil || o.Header == nil {
		var ret CostRecommendationNotificationRuleMessageCustomizationHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetHeaderOk() (*CostRecommendationNotificationRuleMessageCustomizationHeader, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) HasHeader() bool {
	return o != nil && o.Header != nil
}

// SetHeader gets a reference to the given CostRecommendationNotificationRuleMessageCustomizationHeader and assigns it to the Header field.
func (o *CostRecommendationNotificationRuleMessageCustomization) SetHeader(v CostRecommendationNotificationRuleMessageCustomizationHeader) {
	o.Header = &v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetIncludeTags() []string {
	if o == nil || o.IncludeTags == nil {
		var ret []string
		return ret
	}
	return o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetIncludeTagsOk() (*[]string, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return &o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) HasIncludeTags() bool {
	return o != nil && o.IncludeTags != nil
}

// SetIncludeTags gets a reference to the given []string and assigns it to the IncludeTags field.
func (o *CostRecommendationNotificationRuleMessageCustomization) SetIncludeTags(v []string) {
	o.IncludeTags = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetMode() CostRecommendationNotificationRuleMessageCustomizationMode {
	if o == nil || o.Mode == nil {
		var ret CostRecommendationNotificationRuleMessageCustomizationMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetModeOk() (*CostRecommendationNotificationRuleMessageCustomizationMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given CostRecommendationNotificationRuleMessageCustomizationMode and assigns it to the Mode field.
func (o *CostRecommendationNotificationRuleMessageCustomization) SetMode(v CostRecommendationNotificationRuleMessageCustomizationMode) {
	o.Mode = &v
}

// GetPerRec returns the PerRec field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetPerRec() CostRecommendationNotificationRuleMessageCustomizationPerRec {
	if o == nil || o.PerRec == nil {
		var ret CostRecommendationNotificationRuleMessageCustomizationPerRec
		return ret
	}
	return *o.PerRec
}

// GetPerRecOk returns a tuple with the PerRec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) GetPerRecOk() (*CostRecommendationNotificationRuleMessageCustomizationPerRec, bool) {
	if o == nil || o.PerRec == nil {
		return nil, false
	}
	return o.PerRec, true
}

// HasPerRec returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomization) HasPerRec() bool {
	return o != nil && o.PerRec != nil
}

// SetPerRec gets a reference to the given CostRecommendationNotificationRuleMessageCustomizationPerRec and assigns it to the PerRec field.
func (o *CostRecommendationNotificationRuleMessageCustomization) SetPerRec(v CostRecommendationNotificationRuleMessageCustomizationPerRec) {
	o.PerRec = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleMessageCustomization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.IncludeTags != nil {
		toSerialize["include_tags"] = o.IncludeTags
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.PerRec != nil {
		toSerialize["per_rec"] = o.PerRec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleMessageCustomization) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions     *CostRecommendationNotificationRuleMessageActions             `json:"actions,omitempty"`
		Header      *CostRecommendationNotificationRuleMessageCustomizationHeader `json:"header,omitempty"`
		IncludeTags []string                                                      `json:"include_tags,omitempty"`
		Mode        *CostRecommendationNotificationRuleMessageCustomizationMode   `json:"mode,omitempty"`
		PerRec      *CostRecommendationNotificationRuleMessageCustomizationPerRec `json:"per_rec,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "header", "include_tags", "mode", "per_rec"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Actions != nil && all.Actions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Actions = all.Actions
	if all.Header != nil && all.Header.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Header = all.Header
	o.IncludeTags = all.IncludeTags
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	if all.PerRec != nil && all.PerRec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PerRec = all.PerRec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
