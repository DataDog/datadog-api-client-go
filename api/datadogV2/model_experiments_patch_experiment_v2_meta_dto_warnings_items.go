// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2MetaDTOWarningsItems A warning returned after an experiment update.
type ExperimentsPatchExperimentV2MetaDTOWarningsItems struct {
	// Code that identifies the warning.
	Code string `json:"code"`
	// Explanation of the warning and its effect on the update.
	Detail string `json:"detail"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2MetaDTOWarningsItems instantiates a new ExperimentsPatchExperimentV2MetaDTOWarningsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2MetaDTOWarningsItems(code string, detail string) *ExperimentsPatchExperimentV2MetaDTOWarningsItems {
	this := ExperimentsPatchExperimentV2MetaDTOWarningsItems{}
	this.Code = code
	this.Detail = detail
	return &this
}

// NewExperimentsPatchExperimentV2MetaDTOWarningsItemsWithDefaults instantiates a new ExperimentsPatchExperimentV2MetaDTOWarningsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2MetaDTOWarningsItemsWithDefaults() *ExperimentsPatchExperimentV2MetaDTOWarningsItems {
	this := ExperimentsPatchExperimentV2MetaDTOWarningsItems{}
	return &this
}

// GetCode returns the Code field value.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) SetCode(v string) {
	o.Code = v
}

// GetDetail returns the Detail field value.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) SetDetail(v string) {
	o.Detail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2MetaDTOWarningsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["detail"] = o.Detail

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentV2MetaDTOWarningsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code   *string `json:"code"`
		Detail *string `json:"detail"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Detail == nil {
		return fmt.Errorf("required field detail missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "detail"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.Detail = *all.Detail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
