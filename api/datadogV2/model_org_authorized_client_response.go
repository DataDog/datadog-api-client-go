// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientResponse Response containing a single org authorized client.
type OrgAuthorizedClientResponse struct {
	// Data object representing an org authorized client.
	Data OrgAuthorizedClientData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAuthorizedClientResponse instantiates a new OrgAuthorizedClientResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAuthorizedClientResponse(data OrgAuthorizedClientData) *OrgAuthorizedClientResponse {
	this := OrgAuthorizedClientResponse{}
	this.Data = data
	return &this
}

// NewOrgAuthorizedClientResponseWithDefaults instantiates a new OrgAuthorizedClientResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAuthorizedClientResponseWithDefaults() *OrgAuthorizedClientResponse {
	this := OrgAuthorizedClientResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *OrgAuthorizedClientResponse) GetData() OrgAuthorizedClientData {
	if o == nil {
		var ret OrgAuthorizedClientData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientResponse) GetDataOk() (*OrgAuthorizedClientData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *OrgAuthorizedClientResponse) SetData(v OrgAuthorizedClientData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAuthorizedClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAuthorizedClientResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *OrgAuthorizedClientData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
