// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationRequestAttributes Attributes of an anomaly investigation request.
type TimeseriesAnomalyInvestigationRequestAttributes struct {
	// Timeseries requests to investigate. This API version accepts exactly one request.
	Requests []TimeseriesAnomalyInvestigationTimeseriesRequest `json:"requests"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationRequestAttributes instantiates a new TimeseriesAnomalyInvestigationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationRequestAttributes(requests []TimeseriesAnomalyInvestigationTimeseriesRequest) *TimeseriesAnomalyInvestigationRequestAttributes {
	this := TimeseriesAnomalyInvestigationRequestAttributes{}
	this.Requests = requests
	return &this
}

// NewTimeseriesAnomalyInvestigationRequestAttributesWithDefaults instantiates a new TimeseriesAnomalyInvestigationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationRequestAttributesWithDefaults() *TimeseriesAnomalyInvestigationRequestAttributes {
	this := TimeseriesAnomalyInvestigationRequestAttributes{}
	return &this
}

// GetRequests returns the Requests field value.
func (o *TimeseriesAnomalyInvestigationRequestAttributes) GetRequests() []TimeseriesAnomalyInvestigationTimeseriesRequest {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationTimeseriesRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationRequestAttributes) GetRequestsOk() (*[]TimeseriesAnomalyInvestigationTimeseriesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *TimeseriesAnomalyInvestigationRequestAttributes) SetRequests(v []TimeseriesAnomalyInvestigationTimeseriesRequest) {
	o.Requests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["requests"] = o.Requests

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Requests *[]TimeseriesAnomalyInvestigationTimeseriesRequest `json:"requests"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"requests"})
	} else {
		return err
	}
	o.Requests = *all.Requests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
