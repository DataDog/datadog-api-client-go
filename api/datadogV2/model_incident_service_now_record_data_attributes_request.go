// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentServiceNowRecordDataAttributesRequest Attributes for creating a ServiceNow record for an incident.
type IncidentServiceNowRecordDataAttributesRequest struct {
	// The ServiceNow assignment group.
	AssignmentGroup string `json:"assignment_group"`
	// The ServiceNow configuration item mapping.
	ConfigurationItemMapping string `json:"configuration_item_mapping"`
	// The ServiceNow instance name.
	InstanceName string `json:"instance_name"`
	// An existing ServiceNow record ID (Sys ID) to link instead of creating a new record.
	RecordId *string `json:"record_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentServiceNowRecordDataAttributesRequest instantiates a new IncidentServiceNowRecordDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentServiceNowRecordDataAttributesRequest(assignmentGroup string, configurationItemMapping string, instanceName string) *IncidentServiceNowRecordDataAttributesRequest {
	this := IncidentServiceNowRecordDataAttributesRequest{}
	this.AssignmentGroup = assignmentGroup
	this.ConfigurationItemMapping = configurationItemMapping
	this.InstanceName = instanceName
	return &this
}

// NewIncidentServiceNowRecordDataAttributesRequestWithDefaults instantiates a new IncidentServiceNowRecordDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentServiceNowRecordDataAttributesRequestWithDefaults() *IncidentServiceNowRecordDataAttributesRequest {
	this := IncidentServiceNowRecordDataAttributesRequest{}
	return &this
}

// GetAssignmentGroup returns the AssignmentGroup field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetAssignmentGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AssignmentGroup
}

// GetAssignmentGroupOk returns a tuple with the AssignmentGroup field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetAssignmentGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentGroup, true
}

// SetAssignmentGroup sets field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) SetAssignmentGroup(v string) {
	o.AssignmentGroup = v
}

// GetConfigurationItemMapping returns the ConfigurationItemMapping field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetConfigurationItemMapping() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigurationItemMapping
}

// GetConfigurationItemMappingOk returns a tuple with the ConfigurationItemMapping field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetConfigurationItemMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationItemMapping, true
}

// SetConfigurationItemMapping sets field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) SetConfigurationItemMapping(v string) {
	o.ConfigurationItemMapping = v
}

// GetInstanceName returns the InstanceName field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *IncidentServiceNowRecordDataAttributesRequest) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetRecordId() string {
	if o == nil || o.RecordId == nil {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServiceNowRecordDataAttributesRequest) GetRecordIdOk() (*string, bool) {
	if o == nil || o.RecordId == nil {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *IncidentServiceNowRecordDataAttributesRequest) HasRecordId() bool {
	return o != nil && o.RecordId != nil
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *IncidentServiceNowRecordDataAttributesRequest) SetRecordId(v string) {
	o.RecordId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentServiceNowRecordDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assignment_group"] = o.AssignmentGroup
	toSerialize["configuration_item_mapping"] = o.ConfigurationItemMapping
	toSerialize["instance_name"] = o.InstanceName
	if o.RecordId != nil {
		toSerialize["record_id"] = o.RecordId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentServiceNowRecordDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentGroup          *string `json:"assignment_group"`
		ConfigurationItemMapping *string `json:"configuration_item_mapping"`
		InstanceName             *string `json:"instance_name"`
		RecordId                 *string `json:"record_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssignmentGroup == nil {
		return fmt.Errorf("required field assignment_group missing")
	}
	if all.ConfigurationItemMapping == nil {
		return fmt.Errorf("required field configuration_item_mapping missing")
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instance_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_group", "configuration_item_mapping", "instance_name", "record_id"})
	} else {
		return err
	}
	o.AssignmentGroup = *all.AssignmentGroup
	o.ConfigurationItemMapping = *all.ConfigurationItemMapping
	o.InstanceName = *all.InstanceName
	o.RecordId = all.RecordId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
