// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf4461756480 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf4461756480 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection ScatterplotProjection_20270101 `json:"projection"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query EventsPlatformDataProjectionQuery_20270101 `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101 `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf4461756480 instantiates a new ScatterplotBackendTableRequest20270101OneOf4461756480 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf4461756480(projection ScatterplotProjection_20270101, query EventsPlatformDataProjectionQuery_20270101, requestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101) *ScatterplotBackendTableRequest20270101OneOf4461756480 {
	this := ScatterplotBackendTableRequest20270101OneOf4461756480{}
	this.Projection = projection
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOf4461756480WithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf4461756480 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOf4461756480WithDefaults() *ScatterplotBackendTableRequest20270101OneOf4461756480 {
	this := ScatterplotBackendTableRequest20270101OneOf4461756480{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) SetLimit(v float64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetProjection() ScatterplotProjection_20270101 {
	if o == nil {
		var ret ScatterplotProjection_20270101
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetProjectionOk() (*ScatterplotProjection_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) SetProjection(v ScatterplotProjection_20270101) {
	o.Projection = v
}

// GetQuery returns the Query field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetQuery() EventsPlatformDataProjectionQuery_20270101 {
	if o == nil {
		var ret EventsPlatformDataProjectionQuery_20270101
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetQueryOk() (*EventsPlatformDataProjectionQuery_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) SetQuery(v EventsPlatformDataProjectionQuery_20270101) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetRequestType() ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) GetRequestTypeOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) SetRequestType(v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf4461756480) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["projection"] = o.Projection
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendTableRequest20270101OneOf4461756480) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit       *float64                                                                    `json:"limit,omitempty"`
		Projection  *ScatterplotProjection_20270101                                             `json:"projection"`
		Query       *EventsPlatformDataProjectionQuery_20270101                                 `json:"query"`
		RequestType *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104_20270101 `json:"request_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	o.Limit = all.Limit
	o.Projection = *all.Projection
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
