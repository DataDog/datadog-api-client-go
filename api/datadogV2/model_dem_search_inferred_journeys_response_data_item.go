// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemSearchInferredJourneysResponseDataItem A single inferred journey item, either a candidate or an ignored journey.
type DemSearchInferredJourneysResponseDataItem struct {
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemSearchInferredJourneysResponseDataItem instantiates a new DemSearchInferredJourneysResponseDataItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemSearchInferredJourneysResponseDataItem() *DemSearchInferredJourneysResponseDataItem {
	this := DemSearchInferredJourneysResponseDataItem{}
	return &this
}

// NewDemSearchInferredJourneysResponseDataItemWithDefaults instantiates a new DemSearchInferredJourneysResponseDataItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemSearchInferredJourneysResponseDataItemWithDefaults() *DemSearchInferredJourneysResponseDataItem {
	this := DemSearchInferredJourneysResponseDataItem{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o DemSearchInferredJourneysResponseDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemSearchInferredJourneysResponseDataItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{})
	} else {
		return err
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
