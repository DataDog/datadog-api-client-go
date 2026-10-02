// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateExposureSQLModelV2Response Updated exposure SQL model and its removed entries.
type ExperimentsUpdateExposureSQLModelV2Response struct {
	// Exposure SQL model resource with its identifier and configuration.
	Data ExperimentsExposureSQLModelV2DTOData `json:"data"`
	// Removed model entries. Present only when the update removes an entry.
	Meta *ExperimentsUpdateExposureSQLModelV2ResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateExposureSQLModelV2Response instantiates a new ExperimentsUpdateExposureSQLModelV2Response object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateExposureSQLModelV2Response(data ExperimentsExposureSQLModelV2DTOData) *ExperimentsUpdateExposureSQLModelV2Response {
	this := ExperimentsUpdateExposureSQLModelV2Response{}
	this.Data = data
	return &this
}

// NewExperimentsUpdateExposureSQLModelV2ResponseWithDefaults instantiates a new ExperimentsUpdateExposureSQLModelV2Response object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateExposureSQLModelV2ResponseWithDefaults() *ExperimentsUpdateExposureSQLModelV2Response {
	this := ExperimentsUpdateExposureSQLModelV2Response{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsUpdateExposureSQLModelV2Response) GetData() ExperimentsExposureSQLModelV2DTOData {
	if o == nil {
		var ret ExperimentsExposureSQLModelV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2Response) GetDataOk() (*ExperimentsExposureSQLModelV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2Response) SetData(v ExperimentsExposureSQLModelV2DTOData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsUpdateExposureSQLModelV2Response) GetMeta() ExperimentsUpdateExposureSQLModelV2ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ExperimentsUpdateExposureSQLModelV2ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2Response) GetMetaOk() (*ExperimentsUpdateExposureSQLModelV2ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2Response) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsUpdateExposureSQLModelV2ResponseMeta and assigns it to the Meta field.
func (o *ExperimentsUpdateExposureSQLModelV2Response) SetMeta(v ExperimentsUpdateExposureSQLModelV2ResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateExposureSQLModelV2Response) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsUpdateExposureSQLModelV2Response) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ExperimentsExposureSQLModelV2DTOData            `json:"data"`
		Meta *ExperimentsUpdateExposureSQLModelV2ResponseMeta `json:"meta,omitempty"`
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
