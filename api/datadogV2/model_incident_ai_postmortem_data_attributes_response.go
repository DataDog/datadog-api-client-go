// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAIPostmortemDataAttributesResponse Attributes of an AI-generated incident postmortem.
type IncidentAIPostmortemDataAttributesResponse struct {
	// Action items to prevent recurrence.
	ActionItems *string `json:"action_items,omitempty"`
	// The impact of the incident on customers.
	CustomerImpact *string `json:"customer_impact,omitempty"`
	// An executive summary of the incident.
	ExecutiveSummary *string `json:"executive_summary,omitempty"`
	// Key insights from the incident.
	KeyInsights *string `json:"key_insights,omitempty"`
	// Key timeline events during the incident.
	KeyTimeline *string `json:"key_timeline,omitempty"`
	// Lessons learned from the incident.
	LessonsLearned *string `json:"lessons_learned,omitempty"`
	// An overview of the affected systems.
	SystemOverview *string `json:"system_overview,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentAIPostmortemDataAttributesResponse instantiates a new IncidentAIPostmortemDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAIPostmortemDataAttributesResponse() *IncidentAIPostmortemDataAttributesResponse {
	this := IncidentAIPostmortemDataAttributesResponse{}
	return &this
}

// NewIncidentAIPostmortemDataAttributesResponseWithDefaults instantiates a new IncidentAIPostmortemDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAIPostmortemDataAttributesResponseWithDefaults() *IncidentAIPostmortemDataAttributesResponse {
	this := IncidentAIPostmortemDataAttributesResponse{}
	return &this
}

// GetActionItems returns the ActionItems field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetActionItems() string {
	if o == nil || o.ActionItems == nil {
		var ret string
		return ret
	}
	return *o.ActionItems
}

// GetActionItemsOk returns a tuple with the ActionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetActionItemsOk() (*string, bool) {
	if o == nil || o.ActionItems == nil {
		return nil, false
	}
	return o.ActionItems, true
}

// HasActionItems returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasActionItems() bool {
	return o != nil && o.ActionItems != nil
}

// SetActionItems gets a reference to the given string and assigns it to the ActionItems field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetActionItems(v string) {
	o.ActionItems = &v
}

// GetCustomerImpact returns the CustomerImpact field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetCustomerImpact() string {
	if o == nil || o.CustomerImpact == nil {
		var ret string
		return ret
	}
	return *o.CustomerImpact
}

// GetCustomerImpactOk returns a tuple with the CustomerImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetCustomerImpactOk() (*string, bool) {
	if o == nil || o.CustomerImpact == nil {
		return nil, false
	}
	return o.CustomerImpact, true
}

// HasCustomerImpact returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasCustomerImpact() bool {
	return o != nil && o.CustomerImpact != nil
}

// SetCustomerImpact gets a reference to the given string and assigns it to the CustomerImpact field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetCustomerImpact(v string) {
	o.CustomerImpact = &v
}

// GetExecutiveSummary returns the ExecutiveSummary field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetExecutiveSummary() string {
	if o == nil || o.ExecutiveSummary == nil {
		var ret string
		return ret
	}
	return *o.ExecutiveSummary
}

// GetExecutiveSummaryOk returns a tuple with the ExecutiveSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetExecutiveSummaryOk() (*string, bool) {
	if o == nil || o.ExecutiveSummary == nil {
		return nil, false
	}
	return o.ExecutiveSummary, true
}

// HasExecutiveSummary returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasExecutiveSummary() bool {
	return o != nil && o.ExecutiveSummary != nil
}

// SetExecutiveSummary gets a reference to the given string and assigns it to the ExecutiveSummary field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetExecutiveSummary(v string) {
	o.ExecutiveSummary = &v
}

// GetKeyInsights returns the KeyInsights field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetKeyInsights() string {
	if o == nil || o.KeyInsights == nil {
		var ret string
		return ret
	}
	return *o.KeyInsights
}

// GetKeyInsightsOk returns a tuple with the KeyInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetKeyInsightsOk() (*string, bool) {
	if o == nil || o.KeyInsights == nil {
		return nil, false
	}
	return o.KeyInsights, true
}

// HasKeyInsights returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasKeyInsights() bool {
	return o != nil && o.KeyInsights != nil
}

// SetKeyInsights gets a reference to the given string and assigns it to the KeyInsights field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetKeyInsights(v string) {
	o.KeyInsights = &v
}

// GetKeyTimeline returns the KeyTimeline field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetKeyTimeline() string {
	if o == nil || o.KeyTimeline == nil {
		var ret string
		return ret
	}
	return *o.KeyTimeline
}

// GetKeyTimelineOk returns a tuple with the KeyTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetKeyTimelineOk() (*string, bool) {
	if o == nil || o.KeyTimeline == nil {
		return nil, false
	}
	return o.KeyTimeline, true
}

// HasKeyTimeline returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasKeyTimeline() bool {
	return o != nil && o.KeyTimeline != nil
}

// SetKeyTimeline gets a reference to the given string and assigns it to the KeyTimeline field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetKeyTimeline(v string) {
	o.KeyTimeline = &v
}

// GetLessonsLearned returns the LessonsLearned field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetLessonsLearned() string {
	if o == nil || o.LessonsLearned == nil {
		var ret string
		return ret
	}
	return *o.LessonsLearned
}

// GetLessonsLearnedOk returns a tuple with the LessonsLearned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetLessonsLearnedOk() (*string, bool) {
	if o == nil || o.LessonsLearned == nil {
		return nil, false
	}
	return o.LessonsLearned, true
}

// HasLessonsLearned returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasLessonsLearned() bool {
	return o != nil && o.LessonsLearned != nil
}

// SetLessonsLearned gets a reference to the given string and assigns it to the LessonsLearned field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetLessonsLearned(v string) {
	o.LessonsLearned = &v
}

// GetSystemOverview returns the SystemOverview field value if set, zero value otherwise.
func (o *IncidentAIPostmortemDataAttributesResponse) GetSystemOverview() string {
	if o == nil || o.SystemOverview == nil {
		var ret string
		return ret
	}
	return *o.SystemOverview
}

// GetSystemOverviewOk returns a tuple with the SystemOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) GetSystemOverviewOk() (*string, bool) {
	if o == nil || o.SystemOverview == nil {
		return nil, false
	}
	return o.SystemOverview, true
}

// HasSystemOverview returns a boolean if a field has been set.
func (o *IncidentAIPostmortemDataAttributesResponse) HasSystemOverview() bool {
	return o != nil && o.SystemOverview != nil
}

// SetSystemOverview gets a reference to the given string and assigns it to the SystemOverview field.
func (o *IncidentAIPostmortemDataAttributesResponse) SetSystemOverview(v string) {
	o.SystemOverview = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAIPostmortemDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ActionItems != nil {
		toSerialize["action_items"] = o.ActionItems
	}
	if o.CustomerImpact != nil {
		toSerialize["customer_impact"] = o.CustomerImpact
	}
	if o.ExecutiveSummary != nil {
		toSerialize["executive_summary"] = o.ExecutiveSummary
	}
	if o.KeyInsights != nil {
		toSerialize["key_insights"] = o.KeyInsights
	}
	if o.KeyTimeline != nil {
		toSerialize["key_timeline"] = o.KeyTimeline
	}
	if o.LessonsLearned != nil {
		toSerialize["lessons_learned"] = o.LessonsLearned
	}
	if o.SystemOverview != nil {
		toSerialize["system_overview"] = o.SystemOverview
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAIPostmortemDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionItems      *string `json:"action_items,omitempty"`
		CustomerImpact   *string `json:"customer_impact,omitempty"`
		ExecutiveSummary *string `json:"executive_summary,omitempty"`
		KeyInsights      *string `json:"key_insights,omitempty"`
		KeyTimeline      *string `json:"key_timeline,omitempty"`
		LessonsLearned   *string `json:"lessons_learned,omitempty"`
		SystemOverview   *string `json:"system_overview,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_items", "customer_impact", "executive_summary", "key_insights", "key_timeline", "lessons_learned", "system_overview"})
	} else {
		return err
	}
	o.ActionItems = all.ActionItems
	o.CustomerImpact = all.CustomerImpact
	o.ExecutiveSummary = all.ExecutiveSummary
	o.KeyInsights = all.KeyInsights
	o.KeyTimeline = all.KeyTimeline
	o.LessonsLearned = all.LessonsLearned
	o.SystemOverview = all.SystemOverview

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
