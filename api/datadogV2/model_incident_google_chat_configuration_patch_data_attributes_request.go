// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleChatConfigurationPatchDataAttributesRequest Attributes for patching a Google Chat configuration. All fields are optional.
type IncidentGoogleChatConfigurationPatchDataAttributesRequest struct {
	// The Google Chat domain ID.
	DomainId *string `json:"domain_id,omitempty"`
	// The template for the Google Chat space name.
	SpaceNameTemplate *string `json:"space_name_template,omitempty"`
	// The target audience ID for the Google Chat space.
	SpaceTargetAudienceId *string `json:"space_target_audience_id,omitempty"`
	// The time zone for the Google Chat space.
	SpaceTimeZone *string `json:"space_time_zone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleChatConfigurationPatchDataAttributesRequest instantiates a new IncidentGoogleChatConfigurationPatchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleChatConfigurationPatchDataAttributesRequest() *IncidentGoogleChatConfigurationPatchDataAttributesRequest {
	this := IncidentGoogleChatConfigurationPatchDataAttributesRequest{}
	return &this
}

// NewIncidentGoogleChatConfigurationPatchDataAttributesRequestWithDefaults instantiates a new IncidentGoogleChatConfigurationPatchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleChatConfigurationPatchDataAttributesRequestWithDefaults() *IncidentGoogleChatConfigurationPatchDataAttributesRequest {
	this := IncidentGoogleChatConfigurationPatchDataAttributesRequest{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) HasDomainId() bool {
	return o != nil && o.DomainId != nil
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) SetDomainId(v string) {
	o.DomainId = &v
}

// GetSpaceNameTemplate returns the SpaceNameTemplate field value if set, zero value otherwise.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceNameTemplate() string {
	if o == nil || o.SpaceNameTemplate == nil {
		var ret string
		return ret
	}
	return *o.SpaceNameTemplate
}

// GetSpaceNameTemplateOk returns a tuple with the SpaceNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceNameTemplateOk() (*string, bool) {
	if o == nil || o.SpaceNameTemplate == nil {
		return nil, false
	}
	return o.SpaceNameTemplate, true
}

// HasSpaceNameTemplate returns a boolean if a field has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) HasSpaceNameTemplate() bool {
	return o != nil && o.SpaceNameTemplate != nil
}

// SetSpaceNameTemplate gets a reference to the given string and assigns it to the SpaceNameTemplate field.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) SetSpaceNameTemplate(v string) {
	o.SpaceNameTemplate = &v
}

// GetSpaceTargetAudienceId returns the SpaceTargetAudienceId field value if set, zero value otherwise.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceTargetAudienceId() string {
	if o == nil || o.SpaceTargetAudienceId == nil {
		var ret string
		return ret
	}
	return *o.SpaceTargetAudienceId
}

// GetSpaceTargetAudienceIdOk returns a tuple with the SpaceTargetAudienceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceTargetAudienceIdOk() (*string, bool) {
	if o == nil || o.SpaceTargetAudienceId == nil {
		return nil, false
	}
	return o.SpaceTargetAudienceId, true
}

// HasSpaceTargetAudienceId returns a boolean if a field has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) HasSpaceTargetAudienceId() bool {
	return o != nil && o.SpaceTargetAudienceId != nil
}

// SetSpaceTargetAudienceId gets a reference to the given string and assigns it to the SpaceTargetAudienceId field.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) SetSpaceTargetAudienceId(v string) {
	o.SpaceTargetAudienceId = &v
}

// GetSpaceTimeZone returns the SpaceTimeZone field value if set, zero value otherwise.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceTimeZone() string {
	if o == nil || o.SpaceTimeZone == nil {
		var ret string
		return ret
	}
	return *o.SpaceTimeZone
}

// GetSpaceTimeZoneOk returns a tuple with the SpaceTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) GetSpaceTimeZoneOk() (*string, bool) {
	if o == nil || o.SpaceTimeZone == nil {
		return nil, false
	}
	return o.SpaceTimeZone, true
}

// HasSpaceTimeZone returns a boolean if a field has been set.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) HasSpaceTimeZone() bool {
	return o != nil && o.SpaceTimeZone != nil
}

// SetSpaceTimeZone gets a reference to the given string and assigns it to the SpaceTimeZone field.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) SetSpaceTimeZone(v string) {
	o.SpaceTimeZone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleChatConfigurationPatchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.SpaceNameTemplate != nil {
		toSerialize["space_name_template"] = o.SpaceNameTemplate
	}
	if o.SpaceTargetAudienceId != nil {
		toSerialize["space_target_audience_id"] = o.SpaceTargetAudienceId
	}
	if o.SpaceTimeZone != nil {
		toSerialize["space_time_zone"] = o.SpaceTimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleChatConfigurationPatchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DomainId              *string `json:"domain_id,omitempty"`
		SpaceNameTemplate     *string `json:"space_name_template,omitempty"`
		SpaceTargetAudienceId *string `json:"space_target_audience_id,omitempty"`
		SpaceTimeZone         *string `json:"space_time_zone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain_id", "space_name_template", "space_target_audience_id", "space_time_zone"})
	} else {
		return err
	}
	o.DomainId = all.DomainId
	o.SpaceNameTemplate = all.SpaceNameTemplate
	o.SpaceTargetAudienceId = all.SpaceTargetAudienceId
	o.SpaceTimeZone = all.SpaceTimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
