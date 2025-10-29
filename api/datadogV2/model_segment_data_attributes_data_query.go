// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataAttributesDataQuery
type SegmentDataAttributesDataQuery struct {
	//
	Combination *string `json:"combination,omitempty"`
	//
	EventPlatform []SegmentDataAttributesDataQueryEventPlatformItems `json:"event_platform,omitempty"`
	//
	ReferenceTable []SegmentDataAttributesDataQueryReferenceTableItems `json:"reference_table,omitempty"`
	//
	Static []SegmentDataAttributesDataQueryStaticItems `json:"static,omitempty"`
	//
	UserStore []SegmentDataAttributesDataQueryUserStoreItems `json:"user_store,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDataAttributesDataQuery instantiates a new SegmentDataAttributesDataQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDataAttributesDataQuery() *SegmentDataAttributesDataQuery {
	this := SegmentDataAttributesDataQuery{}
	return &this
}

// NewSegmentDataAttributesDataQueryWithDefaults instantiates a new SegmentDataAttributesDataQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDataAttributesDataQueryWithDefaults() *SegmentDataAttributesDataQuery {
	this := SegmentDataAttributesDataQuery{}
	return &this
}

// GetCombination returns the Combination field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQuery) GetCombination() string {
	if o == nil || o.Combination == nil {
		var ret string
		return ret
	}
	return *o.Combination
}

// GetCombinationOk returns a tuple with the Combination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQuery) GetCombinationOk() (*string, bool) {
	if o == nil || o.Combination == nil {
		return nil, false
	}
	return o.Combination, true
}

// HasCombination returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQuery) HasCombination() bool {
	return o != nil && o.Combination != nil
}

// SetCombination gets a reference to the given string and assigns it to the Combination field.
func (o *SegmentDataAttributesDataQuery) SetCombination(v string) {
	o.Combination = &v
}

// GetEventPlatform returns the EventPlatform field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQuery) GetEventPlatform() []SegmentDataAttributesDataQueryEventPlatformItems {
	if o == nil || o.EventPlatform == nil {
		var ret []SegmentDataAttributesDataQueryEventPlatformItems
		return ret
	}
	return o.EventPlatform
}

// GetEventPlatformOk returns a tuple with the EventPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQuery) GetEventPlatformOk() (*[]SegmentDataAttributesDataQueryEventPlatformItems, bool) {
	if o == nil || o.EventPlatform == nil {
		return nil, false
	}
	return &o.EventPlatform, true
}

// HasEventPlatform returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQuery) HasEventPlatform() bool {
	return o != nil && o.EventPlatform != nil
}

// SetEventPlatform gets a reference to the given []SegmentDataAttributesDataQueryEventPlatformItems and assigns it to the EventPlatform field.
func (o *SegmentDataAttributesDataQuery) SetEventPlatform(v []SegmentDataAttributesDataQueryEventPlatformItems) {
	o.EventPlatform = v
}

// GetReferenceTable returns the ReferenceTable field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQuery) GetReferenceTable() []SegmentDataAttributesDataQueryReferenceTableItems {
	if o == nil || o.ReferenceTable == nil {
		var ret []SegmentDataAttributesDataQueryReferenceTableItems
		return ret
	}
	return o.ReferenceTable
}

// GetReferenceTableOk returns a tuple with the ReferenceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQuery) GetReferenceTableOk() (*[]SegmentDataAttributesDataQueryReferenceTableItems, bool) {
	if o == nil || o.ReferenceTable == nil {
		return nil, false
	}
	return &o.ReferenceTable, true
}

// HasReferenceTable returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQuery) HasReferenceTable() bool {
	return o != nil && o.ReferenceTable != nil
}

// SetReferenceTable gets a reference to the given []SegmentDataAttributesDataQueryReferenceTableItems and assigns it to the ReferenceTable field.
func (o *SegmentDataAttributesDataQuery) SetReferenceTable(v []SegmentDataAttributesDataQueryReferenceTableItems) {
	o.ReferenceTable = v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQuery) GetStatic() []SegmentDataAttributesDataQueryStaticItems {
	if o == nil || o.Static == nil {
		var ret []SegmentDataAttributesDataQueryStaticItems
		return ret
	}
	return o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQuery) GetStaticOk() (*[]SegmentDataAttributesDataQueryStaticItems, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return &o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQuery) HasStatic() bool {
	return o != nil && o.Static != nil
}

// SetStatic gets a reference to the given []SegmentDataAttributesDataQueryStaticItems and assigns it to the Static field.
func (o *SegmentDataAttributesDataQuery) SetStatic(v []SegmentDataAttributesDataQueryStaticItems) {
	o.Static = v
}

// GetUserStore returns the UserStore field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQuery) GetUserStore() []SegmentDataAttributesDataQueryUserStoreItems {
	if o == nil || o.UserStore == nil {
		var ret []SegmentDataAttributesDataQueryUserStoreItems
		return ret
	}
	return o.UserStore
}

// GetUserStoreOk returns a tuple with the UserStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQuery) GetUserStoreOk() (*[]SegmentDataAttributesDataQueryUserStoreItems, bool) {
	if o == nil || o.UserStore == nil {
		return nil, false
	}
	return &o.UserStore, true
}

// HasUserStore returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQuery) HasUserStore() bool {
	return o != nil && o.UserStore != nil
}

// SetUserStore gets a reference to the given []SegmentDataAttributesDataQueryUserStoreItems and assigns it to the UserStore field.
func (o *SegmentDataAttributesDataQuery) SetUserStore(v []SegmentDataAttributesDataQueryUserStoreItems) {
	o.UserStore = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDataAttributesDataQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Combination != nil {
		toSerialize["combination"] = o.Combination
	}
	if o.EventPlatform != nil {
		toSerialize["event_platform"] = o.EventPlatform
	}
	if o.ReferenceTable != nil {
		toSerialize["reference_table"] = o.ReferenceTable
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
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
func (o *SegmentDataAttributesDataQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Combination    *string                                             `json:"combination,omitempty"`
		EventPlatform  []SegmentDataAttributesDataQueryEventPlatformItems  `json:"event_platform,omitempty"`
		ReferenceTable []SegmentDataAttributesDataQueryReferenceTableItems `json:"reference_table,omitempty"`
		Static         []SegmentDataAttributesDataQueryStaticItems         `json:"static,omitempty"`
		UserStore      []SegmentDataAttributesDataQueryUserStoreItems      `json:"user_store,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"combination", "event_platform", "reference_table", "static", "user_store"})
	} else {
		return err
	}
	o.Combination = all.Combination
	o.EventPlatform = all.EventPlatform
	o.ReferenceTable = all.ReferenceTable
	o.Static = all.Static
	o.UserStore = all.UserStore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
