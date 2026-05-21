// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentMicrosoftTeamsConfigurationDataAttributesResponse Attributes of a Microsoft Teams configuration.
type IncidentMicrosoftTeamsConfigurationDataAttributesResponse struct {
	// Timestamp when the configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether manual meeting creation is enabled.
	ManualMeetingCreation bool `json:"manual_meeting_creation"`
	// Timestamp when the configuration was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Whether post-meeting summary is enabled.
	PostMeetingSummary bool `json:"post_meeting_summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentMicrosoftTeamsConfigurationDataAttributesResponse instantiates a new IncidentMicrosoftTeamsConfigurationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentMicrosoftTeamsConfigurationDataAttributesResponse(createdAt time.Time, manualMeetingCreation bool, modifiedAt time.Time, postMeetingSummary bool) *IncidentMicrosoftTeamsConfigurationDataAttributesResponse {
	this := IncidentMicrosoftTeamsConfigurationDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.ManualMeetingCreation = manualMeetingCreation
	this.ModifiedAt = modifiedAt
	this.PostMeetingSummary = postMeetingSummary
	return &this
}

// NewIncidentMicrosoftTeamsConfigurationDataAttributesResponseWithDefaults instantiates a new IncidentMicrosoftTeamsConfigurationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentMicrosoftTeamsConfigurationDataAttributesResponseWithDefaults() *IncidentMicrosoftTeamsConfigurationDataAttributesResponse {
	this := IncidentMicrosoftTeamsConfigurationDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetManualMeetingCreation returns the ManualMeetingCreation field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetManualMeetingCreation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ManualMeetingCreation
}

// GetManualMeetingCreationOk returns a tuple with the ManualMeetingCreation field value
// and a boolean to check if the value has been set.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetManualMeetingCreationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualMeetingCreation, true
}

// SetManualMeetingCreation sets field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) SetManualMeetingCreation(v bool) {
	o.ManualMeetingCreation = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetPostMeetingSummary returns the PostMeetingSummary field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetPostMeetingSummary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PostMeetingSummary
}

// GetPostMeetingSummaryOk returns a tuple with the PostMeetingSummary field value
// and a boolean to check if the value has been set.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) GetPostMeetingSummaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostMeetingSummary, true
}

// SetPostMeetingSummary sets field value.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) SetPostMeetingSummary(v bool) {
	o.PostMeetingSummary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentMicrosoftTeamsConfigurationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["manual_meeting_creation"] = o.ManualMeetingCreation
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["post_meeting_summary"] = o.PostMeetingSummary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentMicrosoftTeamsConfigurationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt             *time.Time `json:"created_at"`
		ManualMeetingCreation *bool      `json:"manual_meeting_creation"`
		ModifiedAt            *time.Time `json:"modified_at"`
		PostMeetingSummary    *bool      `json:"post_meeting_summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ManualMeetingCreation == nil {
		return fmt.Errorf("required field manual_meeting_creation missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.PostMeetingSummary == nil {
		return fmt.Errorf("required field post_meeting_summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "manual_meeting_creation", "modified_at", "post_meeting_summary"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ManualMeetingCreation = *all.ManualMeetingCreation
	o.ModifiedAt = *all.ModifiedAt
	o.PostMeetingSummary = *all.PostMeetingSummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
