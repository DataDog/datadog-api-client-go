// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestServiceMap Request that returns nodes and edges from the service map data source.
type TopologyRequestServiceMap struct {
	// Query to the service map topology data source.
	Query *TopologyQueryServiceMap `json:"query,omitempty"`
	// Widget request type.
	RequestType *TopologyRequestType `json:"request_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopologyRequestServiceMap instantiates a new TopologyRequestServiceMap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyRequestServiceMap() *TopologyRequestServiceMap {
	this := TopologyRequestServiceMap{}
	return &this
}

// NewTopologyRequestServiceMapWithDefaults instantiates a new TopologyRequestServiceMap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyRequestServiceMapWithDefaults() *TopologyRequestServiceMap {
	this := TopologyRequestServiceMap{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TopologyRequestServiceMap) GetQuery() TopologyQueryServiceMap {
	if o == nil || o.Query == nil {
		var ret TopologyQueryServiceMap
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyRequestServiceMap) GetQueryOk() (*TopologyQueryServiceMap, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TopologyRequestServiceMap) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given TopologyQueryServiceMap and assigns it to the Query field.
func (o *TopologyRequestServiceMap) SetQuery(v TopologyQueryServiceMap) {
	o.Query = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TopologyRequestServiceMap) GetRequestType() TopologyRequestType {
	if o == nil || o.RequestType == nil {
		var ret TopologyRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyRequestServiceMap) GetRequestTypeOk() (*TopologyRequestType, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TopologyRequestServiceMap) HasRequestType() bool {
	return o != nil && o.RequestType != nil
}

// SetRequestType gets a reference to the given TopologyRequestType and assigns it to the RequestType field.
func (o *TopologyRequestServiceMap) SetRequestType(v TopologyRequestType) {
	o.RequestType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyRequestServiceMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RequestType != nil {
		toSerialize["request_type"] = o.RequestType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyRequestServiceMap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query       *TopologyQueryServiceMap `json:"query,omitempty"`
		RequestType *TopologyRequestType     `json:"request_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "request_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	if all.RequestType != nil && !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = all.RequestType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
