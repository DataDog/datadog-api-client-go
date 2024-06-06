// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogs AWS Logs config
type AWSLogs struct {
	// AWS Lambda forwarder
	LambdaForwarder *AWSLambdaForwarder `json:"lambda_forwarder,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSLogs instantiates a new AWSLogs object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogs() *AWSLogs {
	this := AWSLogs{}
	return &this
}

// NewAWSLogsWithDefaults instantiates a new AWSLogs object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogsWithDefaults() *AWSLogs {
	this := AWSLogs{}
	return &this
}

// GetLambdaForwarder returns the LambdaForwarder field value if set, zero value otherwise.
func (o *AWSLogs) GetLambdaForwarder() AWSLambdaForwarder {
	if o == nil || o.LambdaForwarder == nil {
		var ret AWSLambdaForwarder
		return ret
	}
	return *o.LambdaForwarder
}

// GetLambdaForwarderOk returns a tuple with the LambdaForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogs) GetLambdaForwarderOk() (*AWSLambdaForwarder, bool) {
	if o == nil || o.LambdaForwarder == nil {
		return nil, false
	}
	return o.LambdaForwarder, true
}

// HasLambdaForwarder returns a boolean if a field has been set.
func (o *AWSLogs) HasLambdaForwarder() bool {
	return o != nil && o.LambdaForwarder != nil
}

// SetLambdaForwarder gets a reference to the given AWSLambdaForwarder and assigns it to the LambdaForwarder field.
func (o *AWSLogs) SetLambdaForwarder(v AWSLambdaForwarder) {
	o.LambdaForwarder = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LambdaForwarder != nil {
		toSerialize["lambda_forwarder"] = o.LambdaForwarder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLogs) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LambdaForwarder *AWSLambdaForwarder `json:"lambda_forwarder,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lambda_forwarder"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LambdaForwarder != nil && all.LambdaForwarder.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LambdaForwarder = all.LambdaForwarder

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
