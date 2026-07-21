// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentConfigurationPatchDataAttributesRequest Attributes for patching an incident configuration. All fields are optional.
type IncidentConfigurationPatchDataAttributesRequest struct {
	// Whether to execute integrations for this incident.
	ExecuteIntegrations *bool `json:"execute_integrations,omitempty"`
	// Whether to execute notification rules for this incident.
	ExecuteNotificationRules *bool `json:"execute_notification_rules,omitempty"`
	// Whether to include this incident in analytics.
	IncludeInAnalytics *bool `json:"include_in_analytics,omitempty"`
	// Whether to include this incident in search results.
	IncludeInSearch *bool `json:"include_in_search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentConfigurationPatchDataAttributesRequest instantiates a new IncidentConfigurationPatchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentConfigurationPatchDataAttributesRequest() *IncidentConfigurationPatchDataAttributesRequest {
	this := IncidentConfigurationPatchDataAttributesRequest{}
	return &this
}

// NewIncidentConfigurationPatchDataAttributesRequestWithDefaults instantiates a new IncidentConfigurationPatchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentConfigurationPatchDataAttributesRequestWithDefaults() *IncidentConfigurationPatchDataAttributesRequest {
	this := IncidentConfigurationPatchDataAttributesRequest{}
	return &this
}

// GetExecuteIntegrations returns the ExecuteIntegrations field value if set, zero value otherwise.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetExecuteIntegrations() bool {
	if o == nil || o.ExecuteIntegrations == nil {
		var ret bool
		return ret
	}
	return *o.ExecuteIntegrations
}

// GetExecuteIntegrationsOk returns a tuple with the ExecuteIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetExecuteIntegrationsOk() (*bool, bool) {
	if o == nil || o.ExecuteIntegrations == nil {
		return nil, false
	}
	return o.ExecuteIntegrations, true
}

// HasExecuteIntegrations returns a boolean if a field has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) HasExecuteIntegrations() bool {
	return o != nil && o.ExecuteIntegrations != nil
}

// SetExecuteIntegrations gets a reference to the given bool and assigns it to the ExecuteIntegrations field.
func (o *IncidentConfigurationPatchDataAttributesRequest) SetExecuteIntegrations(v bool) {
	o.ExecuteIntegrations = &v
}

// GetExecuteNotificationRules returns the ExecuteNotificationRules field value if set, zero value otherwise.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetExecuteNotificationRules() bool {
	if o == nil || o.ExecuteNotificationRules == nil {
		var ret bool
		return ret
	}
	return *o.ExecuteNotificationRules
}

// GetExecuteNotificationRulesOk returns a tuple with the ExecuteNotificationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetExecuteNotificationRulesOk() (*bool, bool) {
	if o == nil || o.ExecuteNotificationRules == nil {
		return nil, false
	}
	return o.ExecuteNotificationRules, true
}

// HasExecuteNotificationRules returns a boolean if a field has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) HasExecuteNotificationRules() bool {
	return o != nil && o.ExecuteNotificationRules != nil
}

// SetExecuteNotificationRules gets a reference to the given bool and assigns it to the ExecuteNotificationRules field.
func (o *IncidentConfigurationPatchDataAttributesRequest) SetExecuteNotificationRules(v bool) {
	o.ExecuteNotificationRules = &v
}

// GetIncludeInAnalytics returns the IncludeInAnalytics field value if set, zero value otherwise.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetIncludeInAnalytics() bool {
	if o == nil || o.IncludeInAnalytics == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInAnalytics
}

// GetIncludeInAnalyticsOk returns a tuple with the IncludeInAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetIncludeInAnalyticsOk() (*bool, bool) {
	if o == nil || o.IncludeInAnalytics == nil {
		return nil, false
	}
	return o.IncludeInAnalytics, true
}

// HasIncludeInAnalytics returns a boolean if a field has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) HasIncludeInAnalytics() bool {
	return o != nil && o.IncludeInAnalytics != nil
}

// SetIncludeInAnalytics gets a reference to the given bool and assigns it to the IncludeInAnalytics field.
func (o *IncidentConfigurationPatchDataAttributesRequest) SetIncludeInAnalytics(v bool) {
	o.IncludeInAnalytics = &v
}

// GetIncludeInSearch returns the IncludeInSearch field value if set, zero value otherwise.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetIncludeInSearch() bool {
	if o == nil || o.IncludeInSearch == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInSearch
}

// GetIncludeInSearchOk returns a tuple with the IncludeInSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) GetIncludeInSearchOk() (*bool, bool) {
	if o == nil || o.IncludeInSearch == nil {
		return nil, false
	}
	return o.IncludeInSearch, true
}

// HasIncludeInSearch returns a boolean if a field has been set.
func (o *IncidentConfigurationPatchDataAttributesRequest) HasIncludeInSearch() bool {
	return o != nil && o.IncludeInSearch != nil
}

// SetIncludeInSearch gets a reference to the given bool and assigns it to the IncludeInSearch field.
func (o *IncidentConfigurationPatchDataAttributesRequest) SetIncludeInSearch(v bool) {
	o.IncludeInSearch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentConfigurationPatchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExecuteIntegrations != nil {
		toSerialize["execute_integrations"] = o.ExecuteIntegrations
	}
	if o.ExecuteNotificationRules != nil {
		toSerialize["execute_notification_rules"] = o.ExecuteNotificationRules
	}
	if o.IncludeInAnalytics != nil {
		toSerialize["include_in_analytics"] = o.IncludeInAnalytics
	}
	if o.IncludeInSearch != nil {
		toSerialize["include_in_search"] = o.IncludeInSearch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentConfigurationPatchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExecuteIntegrations      *bool `json:"execute_integrations,omitempty"`
		ExecuteNotificationRules *bool `json:"execute_notification_rules,omitempty"`
		IncludeInAnalytics       *bool `json:"include_in_analytics,omitempty"`
		IncludeInSearch          *bool `json:"include_in_search,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"execute_integrations", "execute_notification_rules", "include_in_analytics", "include_in_search"})
	} else {
		return err
	}
	o.ExecuteIntegrations = all.ExecuteIntegrations
	o.ExecuteNotificationRules = all.ExecuteNotificationRules
	o.IncludeInAnalytics = all.IncludeInAnalytics
	o.IncludeInSearch = all.IncludeInSearch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
