// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf4461521920 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf4461521920 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection_20270101 `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query ScatterplotBackendTableRequest20270101OneOfQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf4461521920 instantiates a new ScatterplotBackendTableRequest20270101OneOf4461521920 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf4461521920(query ScatterplotBackendTableRequest20270101OneOfQuery, requestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101, responseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101) *ScatterplotBackendTableRequest20270101OneOf4461521920 {
	this := ScatterplotBackendTableRequest20270101OneOf4461521920{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOf4461521920WithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf4461521920 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOf4461521920WithDefaults() *ScatterplotBackendTableRequest20270101OneOf4461521920 {
	this := ScatterplotBackendTableRequest20270101OneOf4461521920{}
	return &this
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetProjection() TabularProjection_20270101 {
	if o == nil || o.Projection == nil {
		var ret TabularProjection_20270101
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetProjectionOk() (*TabularProjection_20270101, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection_20270101 and assigns it to the Projection field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) SetProjection(v TabularProjection_20270101) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetQuery() ScatterplotBackendTableRequest20270101OneOfQuery {
	if o == nil {
		var ret ScatterplotBackendTableRequest20270101OneOfQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetQueryOk() (*ScatterplotBackendTableRequest20270101OneOfQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) SetQuery(v ScatterplotBackendTableRequest20270101OneOfQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetRequestType() ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetRequestTypeOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) SetRequestType(v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetResponseFormat() ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) GetResponseFormatOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) SetResponseFormat(v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf4461521920) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendTableRequest20270101OneOf4461521920) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Projection     *TabularProjection_20270101                                                    `json:"projection,omitempty"`
		Query          *ScatterplotBackendTableRequest20270101OneOfQuery                              `json:"query"`
		RequestType    *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320_20270101    `json:"request_type"`
		ResponseFormat *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451204640_20270101 `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Projection = all.Projection
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
