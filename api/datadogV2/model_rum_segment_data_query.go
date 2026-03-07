// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentDataQuery Query definition for the segment. Contains one or more query blocks and an optional combination formula.
type RumSegmentDataQuery struct {
	// Boolean expression combining multiple query blocks.
	Combination *string `json:"combination,omitempty"`
	// List of event platform query blocks.
	EventPlatforms []RumSegmentEventPlatform `json:"event_platforms,omitempty"`
	// List of journey-based query blocks.
	Journeys []RumSegmentJourney `json:"journeys,omitempty"`
	// List of reference table query blocks.
	ReferenceTables []RumSegmentReferenceTable `json:"reference_tables,omitempty"`
	// List of static user list blocks.
	Static []RumSegmentStaticEntry `json:"static,omitempty"`
	// List of template-based query blocks.
	Templates []RumSegmentTemplateInstance `json:"templates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentDataQuery instantiates a new RumSegmentDataQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentDataQuery() *RumSegmentDataQuery {
	this := RumSegmentDataQuery{}
	return &this
}

// NewRumSegmentDataQueryWithDefaults instantiates a new RumSegmentDataQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentDataQueryWithDefaults() *RumSegmentDataQuery {
	this := RumSegmentDataQuery{}
	return &this
}

// GetCombination returns the Combination field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetCombination() string {
	if o == nil || o.Combination == nil {
		var ret string
		return ret
	}
	return *o.Combination
}

// GetCombinationOk returns a tuple with the Combination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetCombinationOk() (*string, bool) {
	if o == nil || o.Combination == nil {
		return nil, false
	}
	return o.Combination, true
}

// HasCombination returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasCombination() bool {
	return o != nil && o.Combination != nil
}

// SetCombination gets a reference to the given string and assigns it to the Combination field.
func (o *RumSegmentDataQuery) SetCombination(v string) {
	o.Combination = &v
}

// GetEventPlatforms returns the EventPlatforms field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetEventPlatforms() []RumSegmentEventPlatform {
	if o == nil || o.EventPlatforms == nil {
		var ret []RumSegmentEventPlatform
		return ret
	}
	return o.EventPlatforms
}

// GetEventPlatformsOk returns a tuple with the EventPlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetEventPlatformsOk() (*[]RumSegmentEventPlatform, bool) {
	if o == nil || o.EventPlatforms == nil {
		return nil, false
	}
	return &o.EventPlatforms, true
}

// HasEventPlatforms returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasEventPlatforms() bool {
	return o != nil && o.EventPlatforms != nil
}

// SetEventPlatforms gets a reference to the given []RumSegmentEventPlatform and assigns it to the EventPlatforms field.
func (o *RumSegmentDataQuery) SetEventPlatforms(v []RumSegmentEventPlatform) {
	o.EventPlatforms = v
}

// GetJourneys returns the Journeys field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetJourneys() []RumSegmentJourney {
	if o == nil || o.Journeys == nil {
		var ret []RumSegmentJourney
		return ret
	}
	return o.Journeys
}

// GetJourneysOk returns a tuple with the Journeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetJourneysOk() (*[]RumSegmentJourney, bool) {
	if o == nil || o.Journeys == nil {
		return nil, false
	}
	return &o.Journeys, true
}

// HasJourneys returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasJourneys() bool {
	return o != nil && o.Journeys != nil
}

// SetJourneys gets a reference to the given []RumSegmentJourney and assigns it to the Journeys field.
func (o *RumSegmentDataQuery) SetJourneys(v []RumSegmentJourney) {
	o.Journeys = v
}

// GetReferenceTables returns the ReferenceTables field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetReferenceTables() []RumSegmentReferenceTable {
	if o == nil || o.ReferenceTables == nil {
		var ret []RumSegmentReferenceTable
		return ret
	}
	return o.ReferenceTables
}

// GetReferenceTablesOk returns a tuple with the ReferenceTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetReferenceTablesOk() (*[]RumSegmentReferenceTable, bool) {
	if o == nil || o.ReferenceTables == nil {
		return nil, false
	}
	return &o.ReferenceTables, true
}

// HasReferenceTables returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasReferenceTables() bool {
	return o != nil && o.ReferenceTables != nil
}

// SetReferenceTables gets a reference to the given []RumSegmentReferenceTable and assigns it to the ReferenceTables field.
func (o *RumSegmentDataQuery) SetReferenceTables(v []RumSegmentReferenceTable) {
	o.ReferenceTables = v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetStatic() []RumSegmentStaticEntry {
	if o == nil || o.Static == nil {
		var ret []RumSegmentStaticEntry
		return ret
	}
	return o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetStaticOk() (*[]RumSegmentStaticEntry, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return &o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasStatic() bool {
	return o != nil && o.Static != nil
}

// SetStatic gets a reference to the given []RumSegmentStaticEntry and assigns it to the Static field.
func (o *RumSegmentDataQuery) SetStatic(v []RumSegmentStaticEntry) {
	o.Static = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *RumSegmentDataQuery) GetTemplates() []RumSegmentTemplateInstance {
	if o == nil || o.Templates == nil {
		var ret []RumSegmentTemplateInstance
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentDataQuery) GetTemplatesOk() (*[]RumSegmentTemplateInstance, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return &o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *RumSegmentDataQuery) HasTemplates() bool {
	return o != nil && o.Templates != nil
}

// SetTemplates gets a reference to the given []RumSegmentTemplateInstance and assigns it to the Templates field.
func (o *RumSegmentDataQuery) SetTemplates(v []RumSegmentTemplateInstance) {
	o.Templates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentDataQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Combination != nil {
		toSerialize["combination"] = o.Combination
	}
	if o.EventPlatforms != nil {
		toSerialize["event_platforms"] = o.EventPlatforms
	}
	if o.Journeys != nil {
		toSerialize["journeys"] = o.Journeys
	}
	if o.ReferenceTables != nil {
		toSerialize["reference_tables"] = o.ReferenceTables
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentDataQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Combination     *string                      `json:"combination,omitempty"`
		EventPlatforms  []RumSegmentEventPlatform    `json:"event_platforms,omitempty"`
		Journeys        []RumSegmentJourney          `json:"journeys,omitempty"`
		ReferenceTables []RumSegmentReferenceTable   `json:"reference_tables,omitempty"`
		Static          []RumSegmentStaticEntry      `json:"static,omitempty"`
		Templates       []RumSegmentTemplateInstance `json:"templates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"combination", "event_platforms", "journeys", "reference_tables", "static", "templates"})
	} else {
		return err
	}
	o.Combination = all.Combination
	o.EventPlatforms = all.EventPlatforms
	o.Journeys = all.Journeys
	o.ReferenceTables = all.ReferenceTables
	o.Static = all.Static
	o.Templates = all.Templates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
