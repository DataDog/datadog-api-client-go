// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmStatsQuery Test-only generated schema for the dashboard versioning proof of concept.
type ApmStatsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []QueryTableColumnRequest `json:"columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Env string `json:"env"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PrimaryTag string `json:"primary_tag"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Resource *string `json:"resource,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RowType ApmStatsQueryRowType `json:"row_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Service string `json:"service"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewApmStatsQuery instantiates a new ApmStatsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmStatsQuery(env string, name string, primaryTag string, rowType ApmStatsQueryRowType, service string) *ApmStatsQuery {
	this := ApmStatsQuery{}
	this.Env = env
	this.Name = name
	this.PrimaryTag = primaryTag
	this.RowType = rowType
	this.Service = service
	return &this
}

// NewApmStatsQueryWithDefaults instantiates a new ApmStatsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmStatsQueryWithDefaults() *ApmStatsQuery {
	this := ApmStatsQuery{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ApmStatsQuery) GetColumns() []QueryTableColumnRequest {
	if o == nil || o.Columns == nil {
		var ret []QueryTableColumnRequest
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetColumnsOk() (*[]QueryTableColumnRequest, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ApmStatsQuery) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []QueryTableColumnRequest and assigns it to the Columns field.
func (o *ApmStatsQuery) SetColumns(v []QueryTableColumnRequest) {
	o.Columns = v
}

// GetEnv returns the Env field value.
func (o *ApmStatsQuery) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ApmStatsQuery) SetEnv(v string) {
	o.Env = v
}

// GetName returns the Name field value.
func (o *ApmStatsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApmStatsQuery) SetName(v string) {
	o.Name = v
}

// GetPrimaryTag returns the PrimaryTag field value.
func (o *ApmStatsQuery) GetPrimaryTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryTag
}

// GetPrimaryTagOk returns a tuple with the PrimaryTag field value
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetPrimaryTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryTag, true
}

// SetPrimaryTag sets field value.
func (o *ApmStatsQuery) SetPrimaryTag(v string) {
	o.PrimaryTag = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ApmStatsQuery) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ApmStatsQuery) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ApmStatsQuery) SetResource(v string) {
	o.Resource = &v
}

// GetRowType returns the RowType field value.
func (o *ApmStatsQuery) GetRowType() ApmStatsQueryRowType {
	if o == nil {
		var ret ApmStatsQueryRowType
		return ret
	}
	return o.RowType
}

// GetRowTypeOk returns a tuple with the RowType field value
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetRowTypeOk() (*ApmStatsQueryRowType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowType, true
}

// SetRowType sets field value.
func (o *ApmStatsQuery) SetRowType(v ApmStatsQueryRowType) {
	o.RowType = v
}

// GetService returns the Service field value.
func (o *ApmStatsQuery) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ApmStatsQuery) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ApmStatsQuery) SetService(v string) {
	o.Service = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmStatsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["env"] = o.Env
	toSerialize["name"] = o.Name
	toSerialize["primary_tag"] = o.PrimaryTag
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	toSerialize["row_type"] = o.RowType
	toSerialize["service"] = o.Service
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmStatsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns    []QueryTableColumnRequest `json:"columns,omitempty"`
		Env        *string                   `json:"env"`
		Name       *string                   `json:"name"`
		PrimaryTag *string                   `json:"primary_tag"`
		Resource   *string                   `json:"resource,omitempty"`
		RowType    *ApmStatsQueryRowType     `json:"row_type"`
		Service    *string                   `json:"service"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.PrimaryTag == nil {
		return fmt.Errorf("required field primary_tag missing")
	}
	if all.RowType == nil {
		return fmt.Errorf("required field row_type missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.Env = *all.Env
	o.Name = *all.Name
	o.PrimaryTag = *all.PrimaryTag
	o.Resource = all.Resource
	if !all.RowType.IsValid() {
		hasInvalidField = true
	} else {
		o.RowType = *all.RowType
	}
	o.Service = *all.Service

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
