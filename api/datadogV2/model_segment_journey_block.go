// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentJourneyBlock Selects the users who converted or dropped off in a funnel.
type SegmentJourneyBlock struct {
	// Whether to return the entities that converted at the target step, or those that dropped off.
	ConversionType *ProductAnalyticsJourneyConversionType `json:"conversion_type,omitempty"`
	// Whether to leave out anonymous users.
	ExcludeAnonymousTraffic *bool `json:"exclude_anonymous_traffic,omitempty"`
	// Start of the time range, in epoch milliseconds.
	From *int64 `json:"from,omitempty"`
	// Group-by definitions of the funnel query.
	GroupBy []map[string]interface{} `json:"group_by,omitempty"`
	// Name of the block, used in `combination`.
	Name *string `json:"name,omitempty"`
	// Funnel search, with the same shape as the `search` of a Product Analytics journey query.
	Search map[string]interface{} `json:"search,omitempty"`
	// Funnel step or step range to evaluate.
	Target map[string]interface{} `json:"target,omitempty"`
	// End of the time range, in epoch milliseconds.
	To *int64 `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentJourneyBlock instantiates a new SegmentJourneyBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentJourneyBlock() *SegmentJourneyBlock {
	this := SegmentJourneyBlock{}
	return &this
}

// NewSegmentJourneyBlockWithDefaults instantiates a new SegmentJourneyBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentJourneyBlockWithDefaults() *SegmentJourneyBlock {
	this := SegmentJourneyBlock{}
	return &this
}

// GetConversionType returns the ConversionType field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetConversionType() ProductAnalyticsJourneyConversionType {
	if o == nil || o.ConversionType == nil {
		var ret ProductAnalyticsJourneyConversionType
		return ret
	}
	return *o.ConversionType
}

// GetConversionTypeOk returns a tuple with the ConversionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetConversionTypeOk() (*ProductAnalyticsJourneyConversionType, bool) {
	if o == nil || o.ConversionType == nil {
		return nil, false
	}
	return o.ConversionType, true
}

// HasConversionType returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasConversionType() bool {
	return o != nil && o.ConversionType != nil
}

// SetConversionType gets a reference to the given ProductAnalyticsJourneyConversionType and assigns it to the ConversionType field.
func (o *SegmentJourneyBlock) SetConversionType(v ProductAnalyticsJourneyConversionType) {
	o.ConversionType = &v
}

// GetExcludeAnonymousTraffic returns the ExcludeAnonymousTraffic field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetExcludeAnonymousTraffic() bool {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAnonymousTraffic
}

// GetExcludeAnonymousTrafficOk returns a tuple with the ExcludeAnonymousTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetExcludeAnonymousTrafficOk() (*bool, bool) {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		return nil, false
	}
	return o.ExcludeAnonymousTraffic, true
}

// HasExcludeAnonymousTraffic returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasExcludeAnonymousTraffic() bool {
	return o != nil && o.ExcludeAnonymousTraffic != nil
}

// SetExcludeAnonymousTraffic gets a reference to the given bool and assigns it to the ExcludeAnonymousTraffic field.
func (o *SegmentJourneyBlock) SetExcludeAnonymousTraffic(v bool) {
	o.ExcludeAnonymousTraffic = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *SegmentJourneyBlock) SetFrom(v int64) {
	o.From = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetGroupBy() []map[string]interface{} {
	if o == nil || o.GroupBy == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetGroupByOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []map[string]interface{} and assigns it to the GroupBy field.
func (o *SegmentJourneyBlock) SetGroupBy(v []map[string]interface{}) {
	o.GroupBy = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentJourneyBlock) SetName(v string) {
	o.Name = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetSearch() map[string]interface{} {
	if o == nil || o.Search == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetSearchOk() (*map[string]interface{}, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return &o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given map[string]interface{} and assigns it to the Search field.
func (o *SegmentJourneyBlock) SetSearch(v map[string]interface{}) {
	o.Search = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetTarget() map[string]interface{} {
	if o == nil || o.Target == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetTargetOk() (*map[string]interface{}, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given map[string]interface{} and assigns it to the Target field.
func (o *SegmentJourneyBlock) SetTarget(v map[string]interface{}) {
	o.Target = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SegmentJourneyBlock) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentJourneyBlock) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SegmentJourneyBlock) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *SegmentJourneyBlock) SetTo(v int64) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentJourneyBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConversionType != nil {
		toSerialize["conversion_type"] = o.ConversionType
	}
	if o.ExcludeAnonymousTraffic != nil {
		toSerialize["exclude_anonymous_traffic"] = o.ExcludeAnonymousTraffic
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentJourneyBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversionType          *ProductAnalyticsJourneyConversionType `json:"conversion_type,omitempty"`
		ExcludeAnonymousTraffic *bool                                  `json:"exclude_anonymous_traffic,omitempty"`
		From                    *int64                                 `json:"from,omitempty"`
		GroupBy                 []map[string]interface{}               `json:"group_by,omitempty"`
		Name                    *string                                `json:"name,omitempty"`
		Search                  map[string]interface{}                 `json:"search,omitempty"`
		Target                  map[string]interface{}                 `json:"target,omitempty"`
		To                      *int64                                 `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conversion_type", "exclude_anonymous_traffic", "from", "group_by", "name", "search", "target", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConversionType != nil && !all.ConversionType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConversionType = all.ConversionType
	}
	o.ExcludeAnonymousTraffic = all.ExcludeAnonymousTraffic
	o.From = all.From
	o.GroupBy = all.GroupBy
	o.Name = all.Name
	o.Search = all.Search
	o.Target = all.Target
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
