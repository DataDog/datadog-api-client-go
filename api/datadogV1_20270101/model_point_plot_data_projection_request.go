// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotDataProjectionRequest Test-only generated schema for the dashboard versioning proof of concept.
type PointPlotDataProjectionRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection PointPlotProjection `json:"projection"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query EventsPlatformDataProjectionQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType PointPlotDataProjectionRequestRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPointPlotDataProjectionRequest instantiates a new PointPlotDataProjectionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotDataProjectionRequest(projection PointPlotProjection, query EventsPlatformDataProjectionQuery, requestType PointPlotDataProjectionRequestRequestType) *PointPlotDataProjectionRequest {
	this := PointPlotDataProjectionRequest{}
	this.Projection = projection
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewPointPlotDataProjectionRequestWithDefaults instantiates a new PointPlotDataProjectionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotDataProjectionRequestWithDefaults() *PointPlotDataProjectionRequest {
	this := PointPlotDataProjectionRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PointPlotDataProjectionRequest) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotDataProjectionRequest) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PointPlotDataProjectionRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *PointPlotDataProjectionRequest) SetLimit(v float64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value.
func (o *PointPlotDataProjectionRequest) GetProjection() PointPlotProjection {
	if o == nil {
		var ret PointPlotProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *PointPlotDataProjectionRequest) GetProjectionOk() (*PointPlotProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *PointPlotDataProjectionRequest) SetProjection(v PointPlotProjection) {
	o.Projection = v
}

// GetQuery returns the Query field value.
func (o *PointPlotDataProjectionRequest) GetQuery() EventsPlatformDataProjectionQuery {
	if o == nil {
		var ret EventsPlatformDataProjectionQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *PointPlotDataProjectionRequest) GetQueryOk() (*EventsPlatformDataProjectionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *PointPlotDataProjectionRequest) SetQuery(v EventsPlatformDataProjectionQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *PointPlotDataProjectionRequest) GetRequestType() PointPlotDataProjectionRequestRequestType {
	if o == nil {
		var ret PointPlotDataProjectionRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *PointPlotDataProjectionRequest) GetRequestTypeOk() (*PointPlotDataProjectionRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *PointPlotDataProjectionRequest) SetRequestType(v PointPlotDataProjectionRequestRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotDataProjectionRequest) MarshalJSON() ([]byte, error) {
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
func (o *PointPlotDataProjectionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit       *float64                                   `json:"limit,omitempty"`
		Projection  *PointPlotProjection                       `json:"projection"`
		Query       *EventsPlatformDataProjectionQuery         `json:"query"`
		RequestType *PointPlotDataProjectionRequestRequestType `json:"request_type"`
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
	if all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = *all.Projection
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
