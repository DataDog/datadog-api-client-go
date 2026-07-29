// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentConfigurationDataAttributesResponse Attributes of an incident configuration in a response.
type IncidentConfigurationDataAttributesResponse struct {
	// Timestamp when the configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether integrations are executed for this incident.
	ExecuteIntegrations *bool `json:"execute_integrations,omitempty"`
	// Whether notification rules are executed for this incident.
	ExecuteNotificationRules *bool `json:"execute_notification_rules,omitempty"`
	// The incident identifier.
	IncidentId string `json:"incident_id"`
	// Whether this incident is included in analytics.
	IncludeInAnalytics *bool `json:"include_in_analytics,omitempty"`
	// Whether this incident is included in search results.
	IncludeInSearch *bool `json:"include_in_search,omitempty"`
	// Timestamp when the configuration was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentConfigurationDataAttributesResponse instantiates a new IncidentConfigurationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentConfigurationDataAttributesResponse(createdAt time.Time, incidentId string, modifiedAt time.Time) *IncidentConfigurationDataAttributesResponse {
	this := IncidentConfigurationDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.IncidentId = incidentId
	this.ModifiedAt = modifiedAt
	return &this
}

// NewIncidentConfigurationDataAttributesResponseWithDefaults instantiates a new IncidentConfigurationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentConfigurationDataAttributesResponseWithDefaults() *IncidentConfigurationDataAttributesResponse {
	this := IncidentConfigurationDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentConfigurationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentConfigurationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExecuteIntegrations returns the ExecuteIntegrations field value if set, zero value otherwise.
func (o *IncidentConfigurationDataAttributesResponse) GetExecuteIntegrations() bool {
	if o == nil || o.ExecuteIntegrations == nil {
		var ret bool
		return ret
	}
	return *o.ExecuteIntegrations
}

// GetExecuteIntegrationsOk returns a tuple with the ExecuteIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetExecuteIntegrationsOk() (*bool, bool) {
	if o == nil || o.ExecuteIntegrations == nil {
		return nil, false
	}
	return o.ExecuteIntegrations, true
}

// HasExecuteIntegrations returns a boolean if a field has been set.
func (o *IncidentConfigurationDataAttributesResponse) HasExecuteIntegrations() bool {
	return o != nil && o.ExecuteIntegrations != nil
}

// SetExecuteIntegrations gets a reference to the given bool and assigns it to the ExecuteIntegrations field.
func (o *IncidentConfigurationDataAttributesResponse) SetExecuteIntegrations(v bool) {
	o.ExecuteIntegrations = &v
}

// GetExecuteNotificationRules returns the ExecuteNotificationRules field value if set, zero value otherwise.
func (o *IncidentConfigurationDataAttributesResponse) GetExecuteNotificationRules() bool {
	if o == nil || o.ExecuteNotificationRules == nil {
		var ret bool
		return ret
	}
	return *o.ExecuteNotificationRules
}

// GetExecuteNotificationRulesOk returns a tuple with the ExecuteNotificationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetExecuteNotificationRulesOk() (*bool, bool) {
	if o == nil || o.ExecuteNotificationRules == nil {
		return nil, false
	}
	return o.ExecuteNotificationRules, true
}

// HasExecuteNotificationRules returns a boolean if a field has been set.
func (o *IncidentConfigurationDataAttributesResponse) HasExecuteNotificationRules() bool {
	return o != nil && o.ExecuteNotificationRules != nil
}

// SetExecuteNotificationRules gets a reference to the given bool and assigns it to the ExecuteNotificationRules field.
func (o *IncidentConfigurationDataAttributesResponse) SetExecuteNotificationRules(v bool) {
	o.ExecuteNotificationRules = &v
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentConfigurationDataAttributesResponse) GetIncidentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetIncidentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentConfigurationDataAttributesResponse) SetIncidentId(v string) {
	o.IncidentId = v
}

// GetIncludeInAnalytics returns the IncludeInAnalytics field value if set, zero value otherwise.
func (o *IncidentConfigurationDataAttributesResponse) GetIncludeInAnalytics() bool {
	if o == nil || o.IncludeInAnalytics == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInAnalytics
}

// GetIncludeInAnalyticsOk returns a tuple with the IncludeInAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetIncludeInAnalyticsOk() (*bool, bool) {
	if o == nil || o.IncludeInAnalytics == nil {
		return nil, false
	}
	return o.IncludeInAnalytics, true
}

// HasIncludeInAnalytics returns a boolean if a field has been set.
func (o *IncidentConfigurationDataAttributesResponse) HasIncludeInAnalytics() bool {
	return o != nil && o.IncludeInAnalytics != nil
}

// SetIncludeInAnalytics gets a reference to the given bool and assigns it to the IncludeInAnalytics field.
func (o *IncidentConfigurationDataAttributesResponse) SetIncludeInAnalytics(v bool) {
	o.IncludeInAnalytics = &v
}

// GetIncludeInSearch returns the IncludeInSearch field value if set, zero value otherwise.
func (o *IncidentConfigurationDataAttributesResponse) GetIncludeInSearch() bool {
	if o == nil || o.IncludeInSearch == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInSearch
}

// GetIncludeInSearchOk returns a tuple with the IncludeInSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetIncludeInSearchOk() (*bool, bool) {
	if o == nil || o.IncludeInSearch == nil {
		return nil, false
	}
	return o.IncludeInSearch, true
}

// HasIncludeInSearch returns a boolean if a field has been set.
func (o *IncidentConfigurationDataAttributesResponse) HasIncludeInSearch() bool {
	return o != nil && o.IncludeInSearch != nil
}

// SetIncludeInSearch gets a reference to the given bool and assigns it to the IncludeInSearch field.
func (o *IncidentConfigurationDataAttributesResponse) SetIncludeInSearch(v bool) {
	o.IncludeInSearch = &v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentConfigurationDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentConfigurationDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentConfigurationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExecuteIntegrations != nil {
		toSerialize["execute_integrations"] = o.ExecuteIntegrations
	}
	if o.ExecuteNotificationRules != nil {
		toSerialize["execute_notification_rules"] = o.ExecuteNotificationRules
	}
	toSerialize["incident_id"] = o.IncidentId
	if o.IncludeInAnalytics != nil {
		toSerialize["include_in_analytics"] = o.IncludeInAnalytics
	}
	if o.IncludeInSearch != nil {
		toSerialize["include_in_search"] = o.IncludeInSearch
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentConfigurationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                *time.Time `json:"created_at"`
		ExecuteIntegrations      *bool      `json:"execute_integrations,omitempty"`
		ExecuteNotificationRules *bool      `json:"execute_notification_rules,omitempty"`
		IncidentId               *string    `json:"incident_id"`
		IncludeInAnalytics       *bool      `json:"include_in_analytics,omitempty"`
		IncludeInSearch          *bool      `json:"include_in_search,omitempty"`
		ModifiedAt               *time.Time `json:"modified_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "execute_integrations", "execute_notification_rules", "incident_id", "include_in_analytics", "include_in_search", "modified_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ExecuteIntegrations = all.ExecuteIntegrations
	o.ExecuteNotificationRules = all.ExecuteNotificationRules
	o.IncidentId = *all.IncidentId
	o.IncludeInAnalytics = all.IncludeInAnalytics
	o.IncludeInSearch = all.IncludeInSearch
	o.ModifiedAt = *all.ModifiedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
