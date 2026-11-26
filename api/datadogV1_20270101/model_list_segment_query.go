// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListSegmentQuery Test-only generated schema for the dashboard versioning proof of concept.
type ListSegmentQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeEventsPlatformQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Context *EventsContext `json:"context,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ListSegmentQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SegmentQuery ListSegmentQuerySegmentQuery `json:"segment_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewListSegmentQuery instantiates a new ListSegmentQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListSegmentQuery(baseQuery CompositeEventsPlatformQuery, dataSource ListSegmentQueryDataSource, name string, segmentQuery ListSegmentQuerySegmentQuery) *ListSegmentQuery {
	this := ListSegmentQuery{}
	this.BaseQuery = baseQuery
	this.DataSource = dataSource
	this.Name = name
	this.SegmentQuery = segmentQuery
	return &this
}

// NewListSegmentQueryWithDefaults instantiates a new ListSegmentQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListSegmentQueryWithDefaults() *ListSegmentQuery {
	this := ListSegmentQuery{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *ListSegmentQuery) GetBaseQuery() CompositeEventsPlatformQuery {
	if o == nil {
		var ret CompositeEventsPlatformQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *ListSegmentQuery) GetBaseQueryOk() (*CompositeEventsPlatformQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *ListSegmentQuery) SetBaseQuery(v CompositeEventsPlatformQuery) {
	o.BaseQuery = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListSegmentQuery) GetContext() EventsContext {
	if o == nil || o.Context == nil {
		var ret EventsContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSegmentQuery) GetContextOk() (*EventsContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListSegmentQuery) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given EventsContext and assigns it to the Context field.
func (o *ListSegmentQuery) SetContext(v EventsContext) {
	o.Context = &v
}

// GetDataSource returns the DataSource field value.
func (o *ListSegmentQuery) GetDataSource() ListSegmentQueryDataSource {
	if o == nil {
		var ret ListSegmentQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ListSegmentQuery) GetDataSourceOk() (*ListSegmentQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ListSegmentQuery) SetDataSource(v ListSegmentQueryDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *ListSegmentQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListSegmentQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ListSegmentQuery) SetName(v string) {
	o.Name = v
}

// GetSegmentQuery returns the SegmentQuery field value.
func (o *ListSegmentQuery) GetSegmentQuery() ListSegmentQuerySegmentQuery {
	if o == nil {
		var ret ListSegmentQuerySegmentQuery
		return ret
	}
	return o.SegmentQuery
}

// GetSegmentQueryOk returns a tuple with the SegmentQuery field value
// and a boolean to check if the value has been set.
func (o *ListSegmentQuery) GetSegmentQueryOk() (*ListSegmentQuerySegmentQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentQuery, true
}

// SetSegmentQuery sets field value.
func (o *ListSegmentQuery) SetSegmentQuery(v ListSegmentQuerySegmentQuery) {
	o.SegmentQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListSegmentQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["segment_query"] = o.SegmentQuery
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListSegmentQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *CompositeEventsPlatformQuery `json:"base_query"`
		Context      *EventsContext                `json:"context,omitempty"`
		DataSource   *ListSegmentQueryDataSource   `json:"data_source"`
		Name         *string                       `json:"name"`
		SegmentQuery *ListSegmentQuerySegmentQuery `json:"segment_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SegmentQuery == nil {
		return fmt.Errorf("required field segment_query missing")
	}

	hasInvalidField := false
	if all.BaseQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BaseQuery = *all.BaseQuery
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	if all.SegmentQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SegmentQuery = *all.SegmentQuery

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
