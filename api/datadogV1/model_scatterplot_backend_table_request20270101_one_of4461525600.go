// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf4461525600 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf4461525600 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetToQuery string `json:"dataset_to_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Datasets []AdvancedQueryDataset_20270101 `json:"datasets"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *ScatterplotProjection_20270101 `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf4461525600 instantiates a new ScatterplotBackendTableRequest20270101OneOf4461525600 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf4461525600(datasetToQuery string, datasets []AdvancedQueryDataset_20270101, requestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101, responseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101) *ScatterplotBackendTableRequest20270101OneOf4461525600 {
	this := ScatterplotBackendTableRequest20270101OneOf4461525600{}
	this.DatasetToQuery = datasetToQuery
	this.Datasets = datasets
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOf4461525600WithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf4461525600 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOf4461525600WithDefaults() *ScatterplotBackendTableRequest20270101OneOf4461525600 {
	this := ScatterplotBackendTableRequest20270101OneOf4461525600{}
	return &this
}

// GetDatasetToQuery returns the DatasetToQuery field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetDatasetToQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetToQuery
}

// GetDatasetToQueryOk returns a tuple with the DatasetToQuery field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetDatasetToQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetToQuery, true
}

// SetDatasetToQuery sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetDatasetToQuery(v string) {
	o.DatasetToQuery = v
}

// GetDatasets returns the Datasets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetDatasets() []AdvancedQueryDataset_20270101 {
	if o == nil {
		var ret []AdvancedQueryDataset_20270101
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetDatasetsOk() (*[]AdvancedQueryDataset_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// SetDatasets sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetDatasets(v []AdvancedQueryDataset_20270101) {
	o.Datasets = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetLimit(v float64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetProjection() ScatterplotProjection_20270101 {
	if o == nil || o.Projection == nil {
		var ret ScatterplotProjection_20270101
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetProjectionOk() (*ScatterplotProjection_20270101, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given ScatterplotProjection_20270101 and assigns it to the Projection field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetProjection(v ScatterplotProjection_20270101) {
	o.Projection = &v
}

// GetRequestType returns the RequestType field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetRequestType() ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetRequestTypeOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetRequestType(v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetResponseFormat() ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) GetResponseFormatOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) SetResponseFormat(v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf4461525600) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_to_query"] = o.DatasetToQuery
	toSerialize["datasets"] = o.Datasets
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendTableRequest20270101OneOf4461525600) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetToQuery *string                                                                        `json:"dataset_to_query"`
		Datasets       *[]AdvancedQueryDataset_20270101                                               `json:"datasets"`
		Limit          *float64                                                                       `json:"limit,omitempty"`
		Projection     *ScatterplotProjection_20270101                                                `json:"projection,omitempty"`
		RequestType    *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184_20270101    `json:"request_type"`
		ResponseFormat *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984_20270101 `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetToQuery == nil {
		return fmt.Errorf("required field dataset_to_query missing")
	}
	if all.Datasets == nil {
		return fmt.Errorf("required field datasets missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.DatasetToQuery = *all.DatasetToQuery
	o.Datasets = *all.Datasets
	o.Limit = all.Limit
	o.Projection = all.Projection
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
