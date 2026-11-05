// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRestrictedDetailsMeta Visibility information for identifying instance details.
type AiGatewayRestrictedDetailsMeta struct {
	// Whether identifying rows were withheld from the returned page. Pagination uses the underlying rows.
	RestrictedDetails bool `json:"restricted_details"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRestrictedDetailsMeta instantiates a new AiGatewayRestrictedDetailsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRestrictedDetailsMeta(restrictedDetails bool) *AiGatewayRestrictedDetailsMeta {
	this := AiGatewayRestrictedDetailsMeta{}
	this.RestrictedDetails = restrictedDetails
	return &this
}

// NewAiGatewayRestrictedDetailsMetaWithDefaults instantiates a new AiGatewayRestrictedDetailsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRestrictedDetailsMetaWithDefaults() *AiGatewayRestrictedDetailsMeta {
	this := AiGatewayRestrictedDetailsMeta{}
	return &this
}

// GetRestrictedDetails returns the RestrictedDetails field value.
func (o *AiGatewayRestrictedDetailsMeta) GetRestrictedDetails() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RestrictedDetails
}

// GetRestrictedDetailsOk returns a tuple with the RestrictedDetails field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRestrictedDetailsMeta) GetRestrictedDetailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictedDetails, true
}

// SetRestrictedDetails sets field value.
func (o *AiGatewayRestrictedDetailsMeta) SetRestrictedDetails(v bool) {
	o.RestrictedDetails = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRestrictedDetailsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["restricted_details"] = o.RestrictedDetails

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRestrictedDetailsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RestrictedDetails *bool `json:"restricted_details"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RestrictedDetails == nil {
		return fmt.Errorf("required field restricted_details missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"restricted_details"})
	} else {
		return err
	}
	o.RestrictedDetails = *all.RestrictedDetails

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
