// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapQuery Journey Paths query definition.
type JourneyMapQuery struct {
	// Whether to show the paths that convert to the end step or drop off before it.
	ConversionType *JourneyMapConversionType `json:"conversion_type,omitempty"`
	// Whether to count unique journeys or all journeys.
	CountingMethod *JourneyMapCountingMethod `json:"counting_method,omitempty"`
	// Data source for user journey funnel queries.
	DataSource ProductAnalyticsFunnelDataSource `json:"data_source"`
	// Maximum number of paths shown per step.
	EntriesPerStep *int64 `json:"entries_per_step,omitempty"`
	// Filter applied to the events shown along the paths.
	EventFilter *string `json:"event_filter,omitempty"`
	// Paths expanded to show the events that follow them.
	ExpandedPrefixes []JourneyMapExpandedPrefix `json:"expanded_prefixes,omitempty"`
	// Group by configuration.
	GroupBy []JourneyMapGroupBy `json:"group_by,omitempty"`
	// User journey search configuration.
	Search UserJourneySearch `json:"search"`
	// The pair of journey steps the paths are visualized between.
	Target JourneyMapTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJourneyMapQuery instantiates a new JourneyMapQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJourneyMapQuery(dataSource ProductAnalyticsFunnelDataSource, search UserJourneySearch, target JourneyMapTarget) *JourneyMapQuery {
	this := JourneyMapQuery{}
	this.DataSource = dataSource
	this.Search = search
	this.Target = target
	return &this
}

// NewJourneyMapQueryWithDefaults instantiates a new JourneyMapQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJourneyMapQueryWithDefaults() *JourneyMapQuery {
	this := JourneyMapQuery{}
	return &this
}

// GetConversionType returns the ConversionType field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetConversionType() JourneyMapConversionType {
	if o == nil || o.ConversionType == nil {
		var ret JourneyMapConversionType
		return ret
	}
	return *o.ConversionType
}

// GetConversionTypeOk returns a tuple with the ConversionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetConversionTypeOk() (*JourneyMapConversionType, bool) {
	if o == nil || o.ConversionType == nil {
		return nil, false
	}
	return o.ConversionType, true
}

// HasConversionType returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasConversionType() bool {
	return o != nil && o.ConversionType != nil
}

// SetConversionType gets a reference to the given JourneyMapConversionType and assigns it to the ConversionType field.
func (o *JourneyMapQuery) SetConversionType(v JourneyMapConversionType) {
	o.ConversionType = &v
}

// GetCountingMethod returns the CountingMethod field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetCountingMethod() JourneyMapCountingMethod {
	if o == nil || o.CountingMethod == nil {
		var ret JourneyMapCountingMethod
		return ret
	}
	return *o.CountingMethod
}

// GetCountingMethodOk returns a tuple with the CountingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetCountingMethodOk() (*JourneyMapCountingMethod, bool) {
	if o == nil || o.CountingMethod == nil {
		return nil, false
	}
	return o.CountingMethod, true
}

// HasCountingMethod returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasCountingMethod() bool {
	return o != nil && o.CountingMethod != nil
}

// SetCountingMethod gets a reference to the given JourneyMapCountingMethod and assigns it to the CountingMethod field.
func (o *JourneyMapQuery) SetCountingMethod(v JourneyMapCountingMethod) {
	o.CountingMethod = &v
}

// GetDataSource returns the DataSource field value.
func (o *JourneyMapQuery) GetDataSource() ProductAnalyticsFunnelDataSource {
	if o == nil {
		var ret ProductAnalyticsFunnelDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetDataSourceOk() (*ProductAnalyticsFunnelDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *JourneyMapQuery) SetDataSource(v ProductAnalyticsFunnelDataSource) {
	o.DataSource = v
}

// GetEntriesPerStep returns the EntriesPerStep field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetEntriesPerStep() int64 {
	if o == nil || o.EntriesPerStep == nil {
		var ret int64
		return ret
	}
	return *o.EntriesPerStep
}

// GetEntriesPerStepOk returns a tuple with the EntriesPerStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetEntriesPerStepOk() (*int64, bool) {
	if o == nil || o.EntriesPerStep == nil {
		return nil, false
	}
	return o.EntriesPerStep, true
}

// HasEntriesPerStep returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasEntriesPerStep() bool {
	return o != nil && o.EntriesPerStep != nil
}

// SetEntriesPerStep gets a reference to the given int64 and assigns it to the EntriesPerStep field.
func (o *JourneyMapQuery) SetEntriesPerStep(v int64) {
	o.EntriesPerStep = &v
}

// GetEventFilter returns the EventFilter field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetEventFilter() string {
	if o == nil || o.EventFilter == nil {
		var ret string
		return ret
	}
	return *o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetEventFilterOk() (*string, bool) {
	if o == nil || o.EventFilter == nil {
		return nil, false
	}
	return o.EventFilter, true
}

// HasEventFilter returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasEventFilter() bool {
	return o != nil && o.EventFilter != nil
}

// SetEventFilter gets a reference to the given string and assigns it to the EventFilter field.
func (o *JourneyMapQuery) SetEventFilter(v string) {
	o.EventFilter = &v
}

// GetExpandedPrefixes returns the ExpandedPrefixes field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetExpandedPrefixes() []JourneyMapExpandedPrefix {
	if o == nil || o.ExpandedPrefixes == nil {
		var ret []JourneyMapExpandedPrefix
		return ret
	}
	return o.ExpandedPrefixes
}

// GetExpandedPrefixesOk returns a tuple with the ExpandedPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetExpandedPrefixesOk() (*[]JourneyMapExpandedPrefix, bool) {
	if o == nil || o.ExpandedPrefixes == nil {
		return nil, false
	}
	return &o.ExpandedPrefixes, true
}

// HasExpandedPrefixes returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasExpandedPrefixes() bool {
	return o != nil && o.ExpandedPrefixes != nil
}

// SetExpandedPrefixes gets a reference to the given []JourneyMapExpandedPrefix and assigns it to the ExpandedPrefixes field.
func (o *JourneyMapQuery) SetExpandedPrefixes(v []JourneyMapExpandedPrefix) {
	o.ExpandedPrefixes = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *JourneyMapQuery) GetGroupBy() []JourneyMapGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []JourneyMapGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetGroupByOk() (*[]JourneyMapGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *JourneyMapQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []JourneyMapGroupBy and assigns it to the GroupBy field.
func (o *JourneyMapQuery) SetGroupBy(v []JourneyMapGroupBy) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *JourneyMapQuery) GetSearch() UserJourneySearch {
	if o == nil {
		var ret UserJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetSearchOk() (*UserJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *JourneyMapQuery) SetSearch(v UserJourneySearch) {
	o.Search = v
}

// GetTarget returns the Target field value.
func (o *JourneyMapQuery) GetTarget() JourneyMapTarget {
	if o == nil {
		var ret JourneyMapTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *JourneyMapQuery) GetTargetOk() (*JourneyMapTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *JourneyMapQuery) SetTarget(v JourneyMapTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JourneyMapQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConversionType != nil {
		toSerialize["conversion_type"] = o.ConversionType
	}
	if o.CountingMethod != nil {
		toSerialize["counting_method"] = o.CountingMethod
	}
	toSerialize["data_source"] = o.DataSource
	if o.EntriesPerStep != nil {
		toSerialize["entries_per_step"] = o.EntriesPerStep
	}
	if o.EventFilter != nil {
		toSerialize["event_filter"] = o.EventFilter
	}
	if o.ExpandedPrefixes != nil {
		toSerialize["expanded_prefixes"] = o.ExpandedPrefixes
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["search"] = o.Search
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JourneyMapQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversionType   *JourneyMapConversionType         `json:"conversion_type,omitempty"`
		CountingMethod   *JourneyMapCountingMethod         `json:"counting_method,omitempty"`
		DataSource       *ProductAnalyticsFunnelDataSource `json:"data_source"`
		EntriesPerStep   *int64                            `json:"entries_per_step,omitempty"`
		EventFilter      *string                           `json:"event_filter,omitempty"`
		ExpandedPrefixes []JourneyMapExpandedPrefix        `json:"expanded_prefixes,omitempty"`
		GroupBy          []JourneyMapGroupBy               `json:"group_by,omitempty"`
		Search           *UserJourneySearch                `json:"search"`
		Target           *JourneyMapTarget                 `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conversion_type", "counting_method", "data_source", "entries_per_step", "event_filter", "expanded_prefixes", "group_by", "search", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConversionType != nil && !all.ConversionType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConversionType = all.ConversionType
	}
	if all.CountingMethod != nil && !all.CountingMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.CountingMethod = all.CountingMethod
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.EntriesPerStep = all.EntriesPerStep
	o.EventFilter = all.EventFilter
	o.ExpandedPrefixes = all.ExpandedPrefixes
	o.GroupBy = all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
