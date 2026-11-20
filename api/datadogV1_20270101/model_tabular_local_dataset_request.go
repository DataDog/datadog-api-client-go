// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TabularLocalDatasetRequest Test-only generated schema for the dashboard versioning proof of concept.
type TabularLocalDatasetRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query TabularLocalDatasetRequest20270101Query `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType TabularLocalDatasetRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TabularLocalDatasetRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTabularLocalDatasetRequest instantiates a new TabularLocalDatasetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTabularLocalDatasetRequest(query TabularLocalDatasetRequest20270101Query, requestType TabularLocalDatasetRequestRequestType, responseFormat TabularLocalDatasetRequestResponseFormat) *TabularLocalDatasetRequest {
	this := TabularLocalDatasetRequest{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewTabularLocalDatasetRequestWithDefaults instantiates a new TabularLocalDatasetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTabularLocalDatasetRequestWithDefaults() *TabularLocalDatasetRequest {
	this := TabularLocalDatasetRequest{}
	return &this
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *TabularLocalDatasetRequest) GetProjection() TabularProjection {
	if o == nil || o.Projection == nil {
		var ret TabularProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TabularLocalDatasetRequest) GetProjectionOk() (*TabularProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *TabularLocalDatasetRequest) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection and assigns it to the Projection field.
func (o *TabularLocalDatasetRequest) SetProjection(v TabularProjection) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *TabularLocalDatasetRequest) GetQuery() TabularLocalDatasetRequest20270101Query {
	if o == nil {
		var ret TabularLocalDatasetRequest20270101Query
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TabularLocalDatasetRequest) GetQueryOk() (*TabularLocalDatasetRequest20270101Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *TabularLocalDatasetRequest) SetQuery(v TabularLocalDatasetRequest20270101Query) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *TabularLocalDatasetRequest) GetRequestType() TabularLocalDatasetRequestRequestType {
	if o == nil {
		var ret TabularLocalDatasetRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *TabularLocalDatasetRequest) GetRequestTypeOk() (*TabularLocalDatasetRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *TabularLocalDatasetRequest) SetRequestType(v TabularLocalDatasetRequestRequestType) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TabularLocalDatasetRequest) GetResponseFormat() TabularLocalDatasetRequestResponseFormat {
	if o == nil {
		var ret TabularLocalDatasetRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TabularLocalDatasetRequest) GetResponseFormatOk() (*TabularLocalDatasetRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TabularLocalDatasetRequest) SetResponseFormat(v TabularLocalDatasetRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TabularLocalDatasetRequest) MarshalJSON() ([]byte, error) {
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
func (o *TabularLocalDatasetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Projection     *TabularProjection                        `json:"projection,omitempty"`
		Query          *TabularLocalDatasetRequest20270101Query  `json:"query"`
		RequestType    *TabularLocalDatasetRequestRequestType    `json:"request_type"`
		ResponseFormat *TabularLocalDatasetRequestResponseFormat `json:"response_format"`
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
