// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleChatConfigurationDataAttributesResponse Attributes of a Google Chat configuration.
type IncidentGoogleChatConfigurationDataAttributesResponse struct {
	// Timestamp when the configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// The Google Chat domain ID.
	DomainId string `json:"domain_id"`
	// Timestamp when the configuration was last modified.
	ModifiedAt time.Time `json:"modified_at"`
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

// NewIncidentGoogleChatConfigurationDataAttributesResponse instantiates a new IncidentGoogleChatConfigurationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleChatConfigurationDataAttributesResponse(createdAt time.Time, domainId string, modifiedAt time.Time, spaceNameTemplate string, spaceTargetAudienceId string, spaceTimeZone string) *IncidentGoogleChatConfigurationDataAttributesResponse {
	this := IncidentGoogleChatConfigurationDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.DomainId = domainId
	this.ModifiedAt = modifiedAt
	this.SpaceNameTemplate = spaceNameTemplate
	this.SpaceTargetAudienceId = spaceTargetAudienceId
	this.SpaceTimeZone = spaceTimeZone
	return &this
}

// NewIncidentGoogleChatConfigurationDataAttributesResponseWithDefaults instantiates a new IncidentGoogleChatConfigurationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleChatConfigurationDataAttributesResponseWithDefaults() *IncidentGoogleChatConfigurationDataAttributesResponse {
	this := IncidentGoogleChatConfigurationDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDomainId returns the DomainId field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetDomainId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetDomainId(v string) {
	o.DomainId = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetSpaceNameTemplate returns the SpaceNameTemplate field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceNameTemplate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceNameTemplate
}

// GetSpaceNameTemplateOk returns a tuple with the SpaceNameTemplate field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceNameTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceNameTemplate, true
}

// SetSpaceNameTemplate sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetSpaceNameTemplate(v string) {
	o.SpaceNameTemplate = v
}

// GetSpaceTargetAudienceId returns the SpaceTargetAudienceId field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceTargetAudienceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceTargetAudienceId
}

// GetSpaceTargetAudienceIdOk returns a tuple with the SpaceTargetAudienceId field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceTargetAudienceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceTargetAudienceId, true
}

// SetSpaceTargetAudienceId sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetSpaceTargetAudienceId(v string) {
	o.SpaceTargetAudienceId = v
}

// GetSpaceTimeZone returns the SpaceTimeZone field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceTimeZone
}

// GetSpaceTimeZoneOk returns a tuple with the SpaceTimeZone field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) GetSpaceTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceTimeZone, true
}

// SetSpaceTimeZone sets field value.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) SetSpaceTimeZone(v string) {
	o.SpaceTimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleChatConfigurationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["domain_id"] = o.DomainId
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["space_name_template"] = o.SpaceNameTemplate
	toSerialize["space_target_audience_id"] = o.SpaceTargetAudienceId
	toSerialize["space_time_zone"] = o.SpaceTimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleChatConfigurationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt             *time.Time `json:"created_at"`
		DomainId              *string    `json:"domain_id"`
		ModifiedAt            *time.Time `json:"modified_at"`
		SpaceNameTemplate     *string    `json:"space_name_template"`
		SpaceTargetAudienceId *string    `json:"space_target_audience_id"`
		SpaceTimeZone         *string    `json:"space_time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DomainId == nil {
		return fmt.Errorf("required field domain_id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "domain_id", "modified_at", "space_name_template", "space_target_audience_id", "space_time_zone"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.DomainId = *all.DomainId
	o.ModifiedAt = *all.ModifiedAt
	o.SpaceNameTemplate = *all.SpaceNameTemplate
	o.SpaceTargetAudienceId = *all.SpaceTargetAudienceId
	o.SpaceTimeZone = *all.SpaceTimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
