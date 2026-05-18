// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotDataProjectionQuery The query for a scatterplot data projection request.
type ScatterplotDataProjectionQuery struct {
	// Data source for event platform-based queries.
	DataSource FormulaAndFunctionEventsDataSource `json:"data_source"`
	// Indexes to search.
	Indexes []string `json:"indexes,omitempty"`
	// The search query string.
	QueryString string `json:"query_string"`
	// Storage tier to query.
	Storage *ScatterplotDataProjectionQueryStorage `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScatterplotDataProjectionQuery instantiates a new ScatterplotDataProjectionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotDataProjectionQuery(dataSource FormulaAndFunctionEventsDataSource, queryString string) *ScatterplotDataProjectionQuery {
	this := ScatterplotDataProjectionQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewScatterplotDataProjectionQueryWithDefaults instantiates a new ScatterplotDataProjectionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotDataProjectionQueryWithDefaults() *ScatterplotDataProjectionQuery {
	this := ScatterplotDataProjectionQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ScatterplotDataProjectionQuery) GetDataSource() FormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret FormulaAndFunctionEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionQuery) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ScatterplotDataProjectionQuery) SetDataSource(v FormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ScatterplotDataProjectionQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ScatterplotDataProjectionQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *ScatterplotDataProjectionQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *ScatterplotDataProjectionQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *ScatterplotDataProjectionQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ScatterplotDataProjectionQuery) GetStorage() ScatterplotDataProjectionQueryStorage {
	if o == nil || o.Storage == nil {
		var ret ScatterplotDataProjectionQueryStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionQuery) GetStorageOk() (*ScatterplotDataProjectionQueryStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ScatterplotDataProjectionQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given ScatterplotDataProjectionQueryStorage and assigns it to the Storage field.
func (o *ScatterplotDataProjectionQuery) SetStorage(v ScatterplotDataProjectionQueryStorage) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotDataProjectionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query_string"] = o.QueryString
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotDataProjectionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *FormulaAndFunctionEventsDataSource    `json:"data_source"`
		Indexes     []string                               `json:"indexes,omitempty"`
		QueryString *string                                `json:"query_string"`
		Storage     *ScatterplotDataProjectionQueryStorage `json:"storage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "indexes", "query_string", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Indexes = all.Indexes
	o.QueryString = *all.QueryString
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
