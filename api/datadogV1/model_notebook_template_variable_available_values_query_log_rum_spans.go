// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariableAvailableValuesQueryLogRumSpans Available values query for logs, RUM, or spans data sources.
type NotebookTemplateVariableAvailableValuesQueryLogRumSpans struct {
	// The data source for the query. Must be one of `logs`, `rum`, or `spans`.
	DataSource string `json:"data_source"`
	// Group-by fields for the query.
	GroupBy []NotebookTemplateVariableAvailableValuesQueryGroupBy `json:"group_by"`
	// Search parameters for an available values query.
	Search NotebookTemplateVariableAvailableValuesQuerySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNotebookTemplateVariableAvailableValuesQueryLogRumSpans instantiates a new NotebookTemplateVariableAvailableValuesQueryLogRumSpans object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTemplateVariableAvailableValuesQueryLogRumSpans(dataSource string, groupBy []NotebookTemplateVariableAvailableValuesQueryGroupBy, search NotebookTemplateVariableAvailableValuesQuerySearch) *NotebookTemplateVariableAvailableValuesQueryLogRumSpans {
	this := NotebookTemplateVariableAvailableValuesQueryLogRumSpans{}
	this.DataSource = dataSource
	this.GroupBy = groupBy
	this.Search = search
	return &this
}

// NewNotebookTemplateVariableAvailableValuesQueryLogRumSpansWithDefaults instantiates a new NotebookTemplateVariableAvailableValuesQueryLogRumSpans object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTemplateVariableAvailableValuesQueryLogRumSpansWithDefaults() *NotebookTemplateVariableAvailableValuesQueryLogRumSpans {
	this := NotebookTemplateVariableAvailableValuesQueryLogRumSpans{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) SetDataSource(v string) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetGroupBy() []NotebookTemplateVariableAvailableValuesQueryGroupBy {
	if o == nil {
		var ret []NotebookTemplateVariableAvailableValuesQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetGroupByOk() (*[]NotebookTemplateVariableAvailableValuesQueryGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) SetGroupBy(v []NotebookTemplateVariableAvailableValuesQueryGroupBy) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetSearch() NotebookTemplateVariableAvailableValuesQuerySearch {
	if o == nil {
		var ret NotebookTemplateVariableAvailableValuesQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) GetSearchOk() (*NotebookTemplateVariableAvailableValuesQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) SetSearch(v NotebookTemplateVariableAvailableValuesQuerySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTemplateVariableAvailableValuesQueryLogRumSpans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["group_by"] = o.GroupBy
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *string                                                `json:"data_source"`
		GroupBy    *[]NotebookTemplateVariableAvailableValuesQueryGroupBy `json:"group_by"`
		Search     *NotebookTemplateVariableAvailableValuesQuerySearch    `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	o.DataSource = *all.DataSource
	o.GroupBy = *all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
