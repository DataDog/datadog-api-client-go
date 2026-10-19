// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountsResponse Response payload for a list of Temporal Cloud OpenMetrics integration accounts.
type TemporalCloudOpenMetricsIntegrationAccountsResponse struct {
	// List of Temporal Cloud OpenMetrics integration accounts.
	Data []TemporalCloudOpenMetricsIntegrationAccountResponseData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountsResponse instantiates a new TemporalCloudOpenMetricsIntegrationAccountsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountsResponse(data []TemporalCloudOpenMetricsIntegrationAccountResponseData) *TemporalCloudOpenMetricsIntegrationAccountsResponse {
	this := TemporalCloudOpenMetricsIntegrationAccountsResponse{}
	this.Data = data
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountsResponseWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountsResponseWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountsResponse {
	this := TemporalCloudOpenMetricsIntegrationAccountsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountsResponse) GetData() []TemporalCloudOpenMetricsIntegrationAccountResponseData {
	if o == nil {
		var ret []TemporalCloudOpenMetricsIntegrationAccountResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountsResponse) GetDataOk() (*[]TemporalCloudOpenMetricsIntegrationAccountResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountsResponse) SetData(v []TemporalCloudOpenMetricsIntegrationAccountResponseData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountsResponse) MarshalJSON() ([]byte, error) {
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
func (o *TemporalCloudOpenMetricsIntegrationAccountsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]TemporalCloudOpenMetricsIntegrationAccountResponseData `json:"data"`
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
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
