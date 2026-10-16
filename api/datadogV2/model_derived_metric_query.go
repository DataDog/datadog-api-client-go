// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DerivedMetricQuery A single query used in the definition of a derived metric.
type DerivedMetricQuery struct {
	// The data source that the query is run against.
	DataSource string `json:"data_source"`
	// The name used to reference this query from the formula.
	Name string `json:"name"`
	// The query string executed against the data source.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDerivedMetricQuery instantiates a new DerivedMetricQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDerivedMetricQuery(dataSource string, name string, query string) *DerivedMetricQuery {
	this := DerivedMetricQuery{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewDerivedMetricQueryWithDefaults instantiates a new DerivedMetricQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDerivedMetricQueryWithDefaults() *DerivedMetricQuery {
	this := DerivedMetricQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *DerivedMetricQuery) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricQuery) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *DerivedMetricQuery) SetDataSource(v string) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *DerivedMetricQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DerivedMetricQuery) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *DerivedMetricQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DerivedMetricQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DerivedMetricQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DerivedMetricQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *string `json:"data_source"`
		Name       *string `json:"name"`
		Query      *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "name", "query"})
	} else {
		return err
	}
	o.DataSource = *all.DataSource
	o.Name = *all.Name
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
