// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelRequestDataAttributes
type FunnelRequestDataAttributes struct {
	//
	DataSource *string `json:"data_source,omitempty"`
	//
	EnforcedExecutionType *string `json:"enforced_execution_type,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
	//
	Search *FunnelRequestDataAttributesSearch `json:"search,omitempty"`
	//
	Time *FunnelRequestDataAttributesTime `json:"time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelRequestDataAttributes instantiates a new FunnelRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelRequestDataAttributes() *FunnelRequestDataAttributes {
	this := FunnelRequestDataAttributes{}
	return &this
}

// NewFunnelRequestDataAttributesWithDefaults instantiates a new FunnelRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelRequestDataAttributesWithDefaults() *FunnelRequestDataAttributes {
	this := FunnelRequestDataAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *FunnelRequestDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelRequestDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *FunnelRequestDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *FunnelRequestDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetEnforcedExecutionType returns the EnforcedExecutionType field value if set, zero value otherwise.
func (o *FunnelRequestDataAttributes) GetEnforcedExecutionType() string {
	if o == nil || o.EnforcedExecutionType == nil {
		var ret string
		return ret
	}
	return *o.EnforcedExecutionType
}

// GetEnforcedExecutionTypeOk returns a tuple with the EnforcedExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelRequestDataAttributes) GetEnforcedExecutionTypeOk() (*string, bool) {
	if o == nil || o.EnforcedExecutionType == nil {
		return nil, false
	}
	return o.EnforcedExecutionType, true
}

// HasEnforcedExecutionType returns a boolean if a field has been set.
func (o *FunnelRequestDataAttributes) HasEnforcedExecutionType() bool {
	return o != nil && o.EnforcedExecutionType != nil
}

// SetEnforcedExecutionType gets a reference to the given string and assigns it to the EnforcedExecutionType field.
func (o *FunnelRequestDataAttributes) SetEnforcedExecutionType(v string) {
	o.EnforcedExecutionType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *FunnelRequestDataAttributes) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelRequestDataAttributes) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *FunnelRequestDataAttributes) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *FunnelRequestDataAttributes) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FunnelRequestDataAttributes) GetSearch() FunnelRequestDataAttributesSearch {
	if o == nil || o.Search == nil {
		var ret FunnelRequestDataAttributesSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelRequestDataAttributes) GetSearchOk() (*FunnelRequestDataAttributesSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FunnelRequestDataAttributes) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FunnelRequestDataAttributesSearch and assigns it to the Search field.
func (o *FunnelRequestDataAttributes) SetSearch(v FunnelRequestDataAttributesSearch) {
	o.Search = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *FunnelRequestDataAttributes) GetTime() FunnelRequestDataAttributesTime {
	if o == nil || o.Time == nil {
		var ret FunnelRequestDataAttributesTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelRequestDataAttributes) GetTimeOk() (*FunnelRequestDataAttributesTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *FunnelRequestDataAttributes) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given FunnelRequestDataAttributesTime and assigns it to the Time field.
func (o *FunnelRequestDataAttributes) SetTime(v FunnelRequestDataAttributesTime) {
	o.Time = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataSource != nil {
		toSerialize["data_source"] = o.DataSource
	}
	if o.EnforcedExecutionType != nil {
		toSerialize["enforced_execution_type"] = o.EnforcedExecutionType
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource            *string                            `json:"data_source,omitempty"`
		EnforcedExecutionType *string                            `json:"enforced_execution_type,omitempty"`
		RequestId             *string                            `json:"request_id,omitempty"`
		Search                *FunnelRequestDataAttributesSearch `json:"search,omitempty"`
		Time                  *FunnelRequestDataAttributesTime   `json:"time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "enforced_execution_type", "request_id", "search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = all.DataSource
	o.EnforcedExecutionType = all.EnforcedExecutionType
	o.RequestId = all.RequestId
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
