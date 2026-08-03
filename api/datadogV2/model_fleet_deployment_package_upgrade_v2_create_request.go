// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentPackageUpgradeV2CreateRequest Request payload for creating a new v2 package upgrade deployment.
type FleetDeploymentPackageUpgradeV2CreateRequest struct {
	// Data for creating a new v2 package upgrade deployment.
	Data FleetDeploymentPackageUpgradeV2Create `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentPackageUpgradeV2CreateRequest instantiates a new FleetDeploymentPackageUpgradeV2CreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentPackageUpgradeV2CreateRequest(data FleetDeploymentPackageUpgradeV2Create) *FleetDeploymentPackageUpgradeV2CreateRequest {
	this := FleetDeploymentPackageUpgradeV2CreateRequest{}
	this.Data = data
	return &this
}

// NewFleetDeploymentPackageUpgradeV2CreateRequestWithDefaults instantiates a new FleetDeploymentPackageUpgradeV2CreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentPackageUpgradeV2CreateRequestWithDefaults() *FleetDeploymentPackageUpgradeV2CreateRequest {
	this := FleetDeploymentPackageUpgradeV2CreateRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *FleetDeploymentPackageUpgradeV2CreateRequest) GetData() FleetDeploymentPackageUpgradeV2Create {
	if o == nil {
		var ret FleetDeploymentPackageUpgradeV2Create
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentPackageUpgradeV2CreateRequest) GetDataOk() (*FleetDeploymentPackageUpgradeV2Create, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *FleetDeploymentPackageUpgradeV2CreateRequest) SetData(v FleetDeploymentPackageUpgradeV2Create) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentPackageUpgradeV2CreateRequest) MarshalJSON() ([]byte, error) {
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
func (o *FleetDeploymentPackageUpgradeV2CreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *FleetDeploymentPackageUpgradeV2Create `json:"data"`
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
