// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InfrastructureListRequest Test-only generated schema for the dashboard versioning proof of concept.
type InfrastructureListRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []InfrastructureListRequestColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query InfrastructureListQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType InfrastructureListRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ListStreamSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewInfrastructureListRequest instantiates a new InfrastructureListRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfrastructureListRequest(columns []InfrastructureListRequestColumnsItems, query InfrastructureListQuery, requestType InfrastructureListRequestRequestType) *InfrastructureListRequest {
	this := InfrastructureListRequest{}
	this.Columns = columns
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewInfrastructureListRequestWithDefaults instantiates a new InfrastructureListRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfrastructureListRequestWithDefaults() *InfrastructureListRequest {
	this := InfrastructureListRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *InfrastructureListRequest) GetColumns() []InfrastructureListRequestColumnsItems {
	if o == nil {
		var ret []InfrastructureListRequestColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListRequest) GetColumnsOk() (*[]InfrastructureListRequestColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *InfrastructureListRequest) SetColumns(v []InfrastructureListRequestColumnsItems) {
	o.Columns = v
}

// GetQuery returns the Query field value.
func (o *InfrastructureListRequest) GetQuery() InfrastructureListQuery {
	if o == nil {
		var ret InfrastructureListQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListRequest) GetQueryOk() (*InfrastructureListQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *InfrastructureListRequest) SetQuery(v InfrastructureListQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *InfrastructureListRequest) GetRequestType() InfrastructureListRequestRequestType {
	if o == nil {
		var ret InfrastructureListRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListRequest) GetRequestTypeOk() (*InfrastructureListRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *InfrastructureListRequest) SetRequestType(v InfrastructureListRequestRequestType) {
	o.RequestType = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *InfrastructureListRequest) GetSort() ListStreamSort {
	if o == nil || o.Sort == nil {
		var ret ListStreamSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfrastructureListRequest) GetSortOk() (*ListStreamSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *InfrastructureListRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ListStreamSort and assigns it to the Sort field.
func (o *InfrastructureListRequest) SetSort(v ListStreamSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfrastructureListRequest) MarshalJSON() ([]byte, error) {
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
func (o *InfrastructureListRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     *[]InfrastructureListRequestColumnsItems `json:"columns"`
		Query       *InfrastructureListQuery                 `json:"query"`
		RequestType *InfrastructureListRequestRequestType    `json:"request_type"`
		Sort        *ListStreamSort                          `json:"sort,omitempty"`
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
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
