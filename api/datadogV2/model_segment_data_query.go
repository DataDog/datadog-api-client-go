// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataQuery Defines who belongs to the segment. It contains one or more blocks, up to 20 in total.
// Use `combination` to combine named blocks with a boolean formula such as `A AND NOT B`.
type SegmentDataQuery struct {
	// Account profile blocks.
	AccountStore []SegmentProfileStoreBlock `json:"account_store,omitempty"`
	// Boolean formula that combines blocks by name.
	Combination *string `json:"combination,omitempty"`
	// Event blocks.
	EventPlatform []SegmentEventPlatformBlock `json:"event_platform,omitempty"`
	// Funnel blocks.
	Journeys []SegmentJourneyBlock `json:"journeys,omitempty"`
	// Reference table blocks.
	ReferenceTable []SegmentReferenceTableBlock `json:"reference_table,omitempty"`
	// Retention blocks.
	Retentions []map[string]interface{} `json:"retentions,omitempty"`
	// Uploaded list blocks.
	Static []SegmentStaticBlock `json:"static,omitempty"`
	// Template blocks.
	Templates []SegmentTemplateBlock `json:"templates,omitempty"`
	// User profile blocks.
	UserStore []SegmentProfileStoreBlock `json:"user_store,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDataQuery instantiates a new SegmentDataQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDataQuery() *SegmentDataQuery {
	this := SegmentDataQuery{}
	return &this
}

// NewSegmentDataQueryWithDefaults instantiates a new SegmentDataQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDataQueryWithDefaults() *SegmentDataQuery {
	this := SegmentDataQuery{}
	return &this
}

// GetAccountStore returns the AccountStore field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetAccountStore() []SegmentProfileStoreBlock {
	if o == nil || o.AccountStore == nil {
		var ret []SegmentProfileStoreBlock
		return ret
	}
	return o.AccountStore
}

// GetAccountStoreOk returns a tuple with the AccountStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetAccountStoreOk() (*[]SegmentProfileStoreBlock, bool) {
	if o == nil || o.AccountStore == nil {
		return nil, false
	}
	return &o.AccountStore, true
}

// HasAccountStore returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasAccountStore() bool {
	return o != nil && o.AccountStore != nil
}

// SetAccountStore gets a reference to the given []SegmentProfileStoreBlock and assigns it to the AccountStore field.
func (o *SegmentDataQuery) SetAccountStore(v []SegmentProfileStoreBlock) {
	o.AccountStore = v
}

// GetCombination returns the Combination field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetCombination() string {
	if o == nil || o.Combination == nil {
		var ret string
		return ret
	}
	return *o.Combination
}

// GetCombinationOk returns a tuple with the Combination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetCombinationOk() (*string, bool) {
	if o == nil || o.Combination == nil {
		return nil, false
	}
	return o.Combination, true
}

// HasCombination returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasCombination() bool {
	return o != nil && o.Combination != nil
}

// SetCombination gets a reference to the given string and assigns it to the Combination field.
func (o *SegmentDataQuery) SetCombination(v string) {
	o.Combination = &v
}

// GetEventPlatform returns the EventPlatform field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetEventPlatform() []SegmentEventPlatformBlock {
	if o == nil || o.EventPlatform == nil {
		var ret []SegmentEventPlatformBlock
		return ret
	}
	return o.EventPlatform
}

// GetEventPlatformOk returns a tuple with the EventPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetEventPlatformOk() (*[]SegmentEventPlatformBlock, bool) {
	if o == nil || o.EventPlatform == nil {
		return nil, false
	}
	return &o.EventPlatform, true
}

// HasEventPlatform returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasEventPlatform() bool {
	return o != nil && o.EventPlatform != nil
}

// SetEventPlatform gets a reference to the given []SegmentEventPlatformBlock and assigns it to the EventPlatform field.
func (o *SegmentDataQuery) SetEventPlatform(v []SegmentEventPlatformBlock) {
	o.EventPlatform = v
}

// GetJourneys returns the Journeys field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetJourneys() []SegmentJourneyBlock {
	if o == nil || o.Journeys == nil {
		var ret []SegmentJourneyBlock
		return ret
	}
	return o.Journeys
}

// GetJourneysOk returns a tuple with the Journeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetJourneysOk() (*[]SegmentJourneyBlock, bool) {
	if o == nil || o.Journeys == nil {
		return nil, false
	}
	return &o.Journeys, true
}

// HasJourneys returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasJourneys() bool {
	return o != nil && o.Journeys != nil
}

// SetJourneys gets a reference to the given []SegmentJourneyBlock and assigns it to the Journeys field.
func (o *SegmentDataQuery) SetJourneys(v []SegmentJourneyBlock) {
	o.Journeys = v
}

// GetReferenceTable returns the ReferenceTable field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetReferenceTable() []SegmentReferenceTableBlock {
	if o == nil || o.ReferenceTable == nil {
		var ret []SegmentReferenceTableBlock
		return ret
	}
	return o.ReferenceTable
}

// GetReferenceTableOk returns a tuple with the ReferenceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetReferenceTableOk() (*[]SegmentReferenceTableBlock, bool) {
	if o == nil || o.ReferenceTable == nil {
		return nil, false
	}
	return &o.ReferenceTable, true
}

// HasReferenceTable returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasReferenceTable() bool {
	return o != nil && o.ReferenceTable != nil
}

// SetReferenceTable gets a reference to the given []SegmentReferenceTableBlock and assigns it to the ReferenceTable field.
func (o *SegmentDataQuery) SetReferenceTable(v []SegmentReferenceTableBlock) {
	o.ReferenceTable = v
}

// GetRetentions returns the Retentions field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetRetentions() []map[string]interface{} {
	if o == nil || o.Retentions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Retentions
}

// GetRetentionsOk returns a tuple with the Retentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetRetentionsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Retentions == nil {
		return nil, false
	}
	return &o.Retentions, true
}

// HasRetentions returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasRetentions() bool {
	return o != nil && o.Retentions != nil
}

// SetRetentions gets a reference to the given []map[string]interface{} and assigns it to the Retentions field.
func (o *SegmentDataQuery) SetRetentions(v []map[string]interface{}) {
	o.Retentions = v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetStatic() []SegmentStaticBlock {
	if o == nil || o.Static == nil {
		var ret []SegmentStaticBlock
		return ret
	}
	return o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetStaticOk() (*[]SegmentStaticBlock, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return &o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasStatic() bool {
	return o != nil && o.Static != nil
}

// SetStatic gets a reference to the given []SegmentStaticBlock and assigns it to the Static field.
func (o *SegmentDataQuery) SetStatic(v []SegmentStaticBlock) {
	o.Static = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetTemplates() []SegmentTemplateBlock {
	if o == nil || o.Templates == nil {
		var ret []SegmentTemplateBlock
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetTemplatesOk() (*[]SegmentTemplateBlock, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return &o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasTemplates() bool {
	return o != nil && o.Templates != nil
}

// SetTemplates gets a reference to the given []SegmentTemplateBlock and assigns it to the Templates field.
func (o *SegmentDataQuery) SetTemplates(v []SegmentTemplateBlock) {
	o.Templates = v
}

// GetUserStore returns the UserStore field value if set, zero value otherwise.
func (o *SegmentDataQuery) GetUserStore() []SegmentProfileStoreBlock {
	if o == nil || o.UserStore == nil {
		var ret []SegmentProfileStoreBlock
		return ret
	}
	return o.UserStore
}

// GetUserStoreOk returns a tuple with the UserStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataQuery) GetUserStoreOk() (*[]SegmentProfileStoreBlock, bool) {
	if o == nil || o.UserStore == nil {
		return nil, false
	}
	return &o.UserStore, true
}

// HasUserStore returns a boolean if a field has been set.
func (o *SegmentDataQuery) HasUserStore() bool {
	return o != nil && o.UserStore != nil
}

// SetUserStore gets a reference to the given []SegmentProfileStoreBlock and assigns it to the UserStore field.
func (o *SegmentDataQuery) SetUserStore(v []SegmentProfileStoreBlock) {
	o.UserStore = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDataQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountStore != nil {
		toSerialize["account_store"] = o.AccountStore
	}
	if o.Combination != nil {
		toSerialize["combination"] = o.Combination
	}
	if o.EventPlatform != nil {
		toSerialize["event_platform"] = o.EventPlatform
	}
	if o.Journeys != nil {
		toSerialize["journeys"] = o.Journeys
	}
	if o.ReferenceTable != nil {
		toSerialize["reference_table"] = o.ReferenceTable
	}
	if o.Retentions != nil {
		toSerialize["retentions"] = o.Retentions
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.UserStore != nil {
		toSerialize["user_store"] = o.UserStore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentDataQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountStore   []SegmentProfileStoreBlock   `json:"account_store,omitempty"`
		Combination    *string                      `json:"combination,omitempty"`
		EventPlatform  []SegmentEventPlatformBlock  `json:"event_platform,omitempty"`
		Journeys       []SegmentJourneyBlock        `json:"journeys,omitempty"`
		ReferenceTable []SegmentReferenceTableBlock `json:"reference_table,omitempty"`
		Retentions     []map[string]interface{}     `json:"retentions,omitempty"`
		Static         []SegmentStaticBlock         `json:"static,omitempty"`
		Templates      []SegmentTemplateBlock       `json:"templates,omitempty"`
		UserStore      []SegmentProfileStoreBlock   `json:"user_store,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_store", "combination", "event_platform", "journeys", "reference_table", "retentions", "static", "templates", "user_store"})
	} else {
		return err
	}
	o.AccountStore = all.AccountStore
	o.Combination = all.Combination
	o.EventPlatform = all.EventPlatform
	o.Journeys = all.Journeys
	o.ReferenceTable = all.ReferenceTable
	o.Retentions = all.Retentions
	o.Static = all.Static
	o.Templates = all.Templates
	o.UserStore = all.UserStore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
