// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleChatConfigurationDataAttributesRequest Attributes for creating a Google Chat configuration.
type IncidentGoogleChatConfigurationDataAttributesRequest struct {
	// The Google Chat domain ID.
	DomainId string `json:"domain_id"`
	// The template for the Google Chat space name.
	SpaceNameTemplate string `json:"space_name_template"`
	// The target audience ID for the Google Chat space.
	SpaceTargetAudienceId string `json:"space_target_audience_id"`
	// The time zone for the Google Chat space.
	SpaceTimeZone string `json:"space_time_zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleChatConfigurationDataAttributesRequest instantiates a new IncidentGoogleChatConfigurationDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleChatConfigurationDataAttributesRequest(domainId string, spaceNameTemplate string, spaceTargetAudienceId string, spaceTimeZone string) *IncidentGoogleChatConfigurationDataAttributesRequest {
	this := IncidentGoogleChatConfigurationDataAttributesRequest{}
	this.DomainId = domainId
	this.SpaceNameTemplate = spaceNameTemplate
	this.SpaceTargetAudienceId = spaceTargetAudienceId
	this.SpaceTimeZone = spaceTimeZone
	return &this
}

// NewIncidentGoogleChatConfigurationDataAttributesRequestWithDefaults instantiates a new IncidentGoogleChatConfigurationDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleChatConfigurationDataAttributesRequestWithDefaults() *IncidentGoogleChatConfigurationDataAttributesRequest {
	this := IncidentGoogleChatConfigurationDataAttributesRequest{}
	return &this
}

// GetDomainId returns the DomainId field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetDomainId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) SetDomainId(v string) {
	o.DomainId = v
}

// GetSpaceNameTemplate returns the SpaceNameTemplate field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceNameTemplate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceNameTemplate
}

// GetSpaceNameTemplateOk returns a tuple with the SpaceNameTemplate field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceNameTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceNameTemplate, true
}

// SetSpaceNameTemplate sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) SetSpaceNameTemplate(v string) {
	o.SpaceNameTemplate = v
}

// GetSpaceTargetAudienceId returns the SpaceTargetAudienceId field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceTargetAudienceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceTargetAudienceId
}

// GetSpaceTargetAudienceIdOk returns a tuple with the SpaceTargetAudienceId field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceTargetAudienceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceTargetAudienceId, true
}

// SetSpaceTargetAudienceId sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) SetSpaceTargetAudienceId(v string) {
	o.SpaceTargetAudienceId = v
}

// GetSpaceTimeZone returns the SpaceTimeZone field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceTimeZone
}

// GetSpaceTimeZoneOk returns a tuple with the SpaceTimeZone field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) GetSpaceTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceTimeZone, true
}

// SetSpaceTimeZone sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) SetSpaceTimeZone(v string) {
	o.SpaceTimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleChatConfigurationDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["domain_id"] = o.DomainId
	toSerialize["space_name_template"] = o.SpaceNameTemplate
	toSerialize["space_target_audience_id"] = o.SpaceTargetAudienceId
	toSerialize["space_time_zone"] = o.SpaceTimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleChatConfigurationDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DomainId              *string `json:"domain_id"`
		SpaceNameTemplate     *string `json:"space_name_template"`
		SpaceTargetAudienceId *string `json:"space_target_audience_id"`
		SpaceTimeZone         *string `json:"space_time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DomainId == nil {
		return fmt.Errorf("required field domain_id missing")
	}
	if all.SpaceNameTemplate == nil {
		return fmt.Errorf("required field space_name_template missing")
	}
	if all.SpaceTargetAudienceId == nil {
		return fmt.Errorf("required field space_target_audience_id missing")
	}
	if all.SpaceTimeZone == nil {
		return fmt.Errorf("required field space_time_zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain_id", "space_name_template", "space_target_audience_id", "space_time_zone"})
	} else {
		return err
	}
	o.DomainId = *all.DomainId
	o.SpaceNameTemplate = *all.SpaceNameTemplate
	o.SpaceTargetAudienceId = *all.SpaceTargetAudienceId
	o.SpaceTimeZone = *all.SpaceTimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
