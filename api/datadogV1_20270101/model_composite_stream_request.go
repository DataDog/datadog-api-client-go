// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeStreamRequest Test-only generated schema for the dashboard versioning proof of concept.
type CompositeStreamRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []CompositeStreamRequestColumn `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query ListCompositeQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType CompositeStreamRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ListStreamSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeStreamRequest instantiates a new CompositeStreamRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeStreamRequest(columns []CompositeStreamRequestColumn, query ListCompositeQuery, requestType CompositeStreamRequestRequestType) *CompositeStreamRequest {
	this := CompositeStreamRequest{}
	this.Columns = columns
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewCompositeStreamRequestWithDefaults instantiates a new CompositeStreamRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeStreamRequestWithDefaults() *CompositeStreamRequest {
	this := CompositeStreamRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CompositeStreamRequest) GetColumns() []CompositeStreamRequestColumn {
	if o == nil {
		var ret []CompositeStreamRequestColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CompositeStreamRequest) GetColumnsOk() (*[]CompositeStreamRequestColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CompositeStreamRequest) SetColumns(v []CompositeStreamRequestColumn) {
	o.Columns = v
}

// GetQuery returns the Query field value.
func (o *CompositeStreamRequest) GetQuery() ListCompositeQuery {
	if o == nil {
		var ret ListCompositeQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CompositeStreamRequest) GetQueryOk() (*ListCompositeQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CompositeStreamRequest) SetQuery(v ListCompositeQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *CompositeStreamRequest) GetRequestType() CompositeStreamRequestRequestType {
	if o == nil {
		var ret CompositeStreamRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *CompositeStreamRequest) GetRequestTypeOk() (*CompositeStreamRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *CompositeStreamRequest) SetRequestType(v CompositeStreamRequestRequestType) {
	o.RequestType = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CompositeStreamRequest) GetSort() ListStreamSort {
	if o == nil || o.Sort == nil {
		var ret ListStreamSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeStreamRequest) GetSortOk() (*ListStreamSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CompositeStreamRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ListStreamSort and assigns it to the Sort field.
func (o *CompositeStreamRequest) SetSort(v ListStreamSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompositeStreamRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     *[]CompositeStreamRequestColumn    `json:"columns"`
		Query       *ListCompositeQuery                `json:"query"`
		RequestType *CompositeStreamRequestRequestType `json:"request_type"`
		Sort        *ListStreamSort                    `json:"sort,omitempty"`
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
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
