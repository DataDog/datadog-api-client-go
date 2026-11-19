// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshV2DTODataAttributes The bandit being refreshed.
type ExperimentsBanditRefreshV2DTODataAttributes struct {
	// The ID of the bandit being refreshed.
	BanditId uuid.UUID `json:"bandit_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditRefreshV2DTODataAttributes instantiates a new ExperimentsBanditRefreshV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditRefreshV2DTODataAttributes(banditId uuid.UUID) *ExperimentsBanditRefreshV2DTODataAttributes {
	this := ExperimentsBanditRefreshV2DTODataAttributes{}
	this.BanditId = banditId
	return &this
}

// NewExperimentsBanditRefreshV2DTODataAttributesWithDefaults instantiates a new ExperimentsBanditRefreshV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditRefreshV2DTODataAttributesWithDefaults() *ExperimentsBanditRefreshV2DTODataAttributes {
	this := ExperimentsBanditRefreshV2DTODataAttributes{}
	return &this
}

// GetBanditId returns the BanditId field value.
func (o *ExperimentsBanditRefreshV2DTODataAttributes) GetBanditId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BanditId
}

// GetBanditIdOk returns a tuple with the BanditId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2DTODataAttributes) GetBanditIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BanditId, true
}

// SetBanditId sets field value.
func (o *ExperimentsBanditRefreshV2DTODataAttributes) SetBanditId(v uuid.UUID) {
	o.BanditId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditRefreshV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bandit_id"] = o.BanditId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditRefreshV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BanditId *uuid.UUID `json:"bandit_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BanditId == nil {
		return fmt.Errorf("required field bandit_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bandit_id"})
	} else {
		return err
	}
	o.BanditId = *all.BanditId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
