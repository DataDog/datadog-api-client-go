// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListRequest Test-only generated schema for the dashboard versioning proof of concept.
type DatasetListRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []DatasetListRequestColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query DatasetListRequest20270101Query `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType DatasetListRequestRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetListRequest instantiates a new DatasetListRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetListRequest(columns []DatasetListRequestColumnsItems, query DatasetListRequest20270101Query, requestType DatasetListRequestRequestType) *DatasetListRequest {
	this := DatasetListRequest{}
	this.Columns = columns
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewDatasetListRequestWithDefaults instantiates a new DatasetListRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetListRequestWithDefaults() *DatasetListRequest {
	this := DatasetListRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *DatasetListRequest) GetColumns() []DatasetListRequestColumnsItems {
	if o == nil {
		var ret []DatasetListRequestColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *DatasetListRequest) GetColumnsOk() (*[]DatasetListRequestColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *DatasetListRequest) SetColumns(v []DatasetListRequestColumnsItems) {
	o.Columns = v
}

// GetQuery returns the Query field value.
func (o *DatasetListRequest) GetQuery() DatasetListRequest20270101Query {
	if o == nil {
		var ret DatasetListRequest20270101Query
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DatasetListRequest) GetQueryOk() (*DatasetListRequest20270101Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DatasetListRequest) SetQuery(v DatasetListRequest20270101Query) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *DatasetListRequest) GetRequestType() DatasetListRequestRequestType {
	if o == nil {
		var ret DatasetListRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *DatasetListRequest) GetRequestTypeOk() (*DatasetListRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *DatasetListRequest) SetRequestType(v DatasetListRequestRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetListRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     *[]DatasetListRequestColumnsItems `json:"columns"`
		Query       *DatasetListRequest20270101Query  `json:"query"`
		RequestType *DatasetListRequestRequestType    `json:"request_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
