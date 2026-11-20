// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsPlatformDataProjectionQuery Test-only generated schema for the dashboard versioning proof of concept.
type EventsPlatformDataProjectionQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource LegacyFormulaBackedDataProjectionSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventsPlatformDataProjectionQuery instantiates a new EventsPlatformDataProjectionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsPlatformDataProjectionQuery(dataSource LegacyFormulaBackedDataProjectionSource, queryString string) *EventsPlatformDataProjectionQuery {
	this := EventsPlatformDataProjectionQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewEventsPlatformDataProjectionQueryWithDefaults instantiates a new EventsPlatformDataProjectionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsPlatformDataProjectionQueryWithDefaults() *EventsPlatformDataProjectionQuery {
	this := EventsPlatformDataProjectionQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *EventsPlatformDataProjectionQuery) GetDataSource() LegacyFormulaBackedDataProjectionSource {
	if o == nil {
		var ret LegacyFormulaBackedDataProjectionSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *EventsPlatformDataProjectionQuery) GetDataSourceOk() (*LegacyFormulaBackedDataProjectionSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *EventsPlatformDataProjectionQuery) SetDataSource(v LegacyFormulaBackedDataProjectionSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *EventsPlatformDataProjectionQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformDataProjectionQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *EventsPlatformDataProjectionQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *EventsPlatformDataProjectionQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *EventsPlatformDataProjectionQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *EventsPlatformDataProjectionQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *EventsPlatformDataProjectionQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *EventsPlatformDataProjectionQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPlatformDataProjectionQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *EventsPlatformDataProjectionQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *EventsPlatformDataProjectionQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsPlatformDataProjectionQuery) MarshalJSON() ([]byte, error) {
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
func (o *EventsPlatformDataProjectionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *LegacyFormulaBackedDataProjectionSource `json:"data_source"`
		Indexes     []string                                 `json:"indexes,omitempty"`
		QueryString *string                                  `json:"query_string"`
		Storage     *StorageType                             `json:"storage,omitempty"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
