// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateMetricSQLModelV2Response Updated metric SQL model and its removed entries.
type ExperimentsUpdateMetricSQLModelV2Response struct {
	// JSON:API resource containing the metric SQL model identity and fields.
	Data ExperimentsMetricSQLModelV2DTOData `json:"data"`
	// Model entries removed by the update. Empty arrays mean no entries were removed.
	Meta *ExperimentsUpdateMetricSQLModelV2ResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateMetricSQLModelV2Response instantiates a new ExperimentsUpdateMetricSQLModelV2Response object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateMetricSQLModelV2Response(data ExperimentsMetricSQLModelV2DTOData) *ExperimentsUpdateMetricSQLModelV2Response {
	this := ExperimentsUpdateMetricSQLModelV2Response{}
	this.Data = data
	return &this
}

// NewExperimentsUpdateMetricSQLModelV2ResponseWithDefaults instantiates a new ExperimentsUpdateMetricSQLModelV2Response object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateMetricSQLModelV2ResponseWithDefaults() *ExperimentsUpdateMetricSQLModelV2Response {
	this := ExperimentsUpdateMetricSQLModelV2Response{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsUpdateMetricSQLModelV2Response) GetData() ExperimentsMetricSQLModelV2DTOData {
	if o == nil {
		var ret ExperimentsMetricSQLModelV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2Response) GetDataOk() (*ExperimentsMetricSQLModelV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsUpdateMetricSQLModelV2Response) SetData(v ExperimentsMetricSQLModelV2DTOData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2Response) GetMeta() ExperimentsUpdateMetricSQLModelV2ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ExperimentsUpdateMetricSQLModelV2ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2Response) GetMetaOk() (*ExperimentsUpdateMetricSQLModelV2ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2Response) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsUpdateMetricSQLModelV2ResponseMeta and assigns it to the Meta field.
func (o *ExperimentsUpdateMetricSQLModelV2Response) SetMeta(v ExperimentsUpdateMetricSQLModelV2ResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateMetricSQLModelV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateMetricSQLModelV2Response) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ExperimentsMetricSQLModelV2DTOData            `json:"data"`
		Meta *ExperimentsUpdateMetricSQLModelV2ResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
