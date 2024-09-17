// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// ServiceDefinitionV2Contact - Service owner's contacts information.
type ServiceDefinitionV2Contact struct {
	ServiceDefinitionV2Email *ServiceDefinitionV2Email
	ServiceDefinitionV2Slack *ServiceDefinitionV2Slack
	ServiceDefinitionV2MSTeams *ServiceDefinitionV2MSTeams

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceDefinitionV2EmailAsServiceDefinitionV2Contact is a convenience function that returns ServiceDefinitionV2Email wrapped in ServiceDefinitionV2Contact.
func ServiceDefinitionV2EmailAsServiceDefinitionV2Contact(v *ServiceDefinitionV2Email) ServiceDefinitionV2Contact {
	return ServiceDefinitionV2Contact{ServiceDefinitionV2Email: v}
}

// ServiceDefinitionV2SlackAsServiceDefinitionV2Contact is a convenience function that returns ServiceDefinitionV2Slack wrapped in ServiceDefinitionV2Contact.
func ServiceDefinitionV2SlackAsServiceDefinitionV2Contact(v *ServiceDefinitionV2Slack) ServiceDefinitionV2Contact {
	return ServiceDefinitionV2Contact{ServiceDefinitionV2Slack: v}
}

// ServiceDefinitionV2MSTeamsAsServiceDefinitionV2Contact is a convenience function that returns ServiceDefinitionV2MSTeams wrapped in ServiceDefinitionV2Contact.
func ServiceDefinitionV2MSTeamsAsServiceDefinitionV2Contact(v *ServiceDefinitionV2MSTeams) ServiceDefinitionV2Contact {
	return ServiceDefinitionV2Contact{ServiceDefinitionV2MSTeams: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceDefinitionV2Contact) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceDefinitionV2Email
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Email)
	if err == nil {
		if obj.ServiceDefinitionV2Email != nil && obj.ServiceDefinitionV2Email.UnparsedObject == nil {
			jsonServiceDefinitionV2Email, _ := datadog.Marshal(obj.ServiceDefinitionV2Email)
			if string(jsonServiceDefinitionV2Email) == "{}" { // empty struct
				obj.ServiceDefinitionV2Email = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Email = nil
		}
	} else {
		obj.ServiceDefinitionV2Email = nil
	}

	// try to unmarshal data into ServiceDefinitionV2Slack
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Slack)
	if err == nil {
		if obj.ServiceDefinitionV2Slack != nil && obj.ServiceDefinitionV2Slack.UnparsedObject == nil {
			jsonServiceDefinitionV2Slack, _ := datadog.Marshal(obj.ServiceDefinitionV2Slack)
			if string(jsonServiceDefinitionV2Slack) == "{}" { // empty struct
				obj.ServiceDefinitionV2Slack = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Slack = nil
		}
	} else {
		obj.ServiceDefinitionV2Slack = nil
	}

	// try to unmarshal data into ServiceDefinitionV2MSTeams
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2MSTeams)
	if err == nil {
		if obj.ServiceDefinitionV2MSTeams != nil && obj.ServiceDefinitionV2MSTeams.UnparsedObject == nil {
			jsonServiceDefinitionV2MSTeams, _ := datadog.Marshal(obj.ServiceDefinitionV2MSTeams)
			if string(jsonServiceDefinitionV2MSTeams) == "{}" { // empty struct
				obj.ServiceDefinitionV2MSTeams = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2MSTeams = nil
		}
	} else {
		obj.ServiceDefinitionV2MSTeams = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceDefinitionV2Email = nil
		obj.ServiceDefinitionV2Slack = nil
		obj.ServiceDefinitionV2MSTeams = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceDefinitionV2Contact) MarshalJSON() ([]byte, error) {
	if obj.ServiceDefinitionV2Email != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Email)
	}


	if obj.ServiceDefinitionV2Slack != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Slack)
	}


	if obj.ServiceDefinitionV2MSTeams != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2MSTeams)
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceDefinitionV2Contact) GetActualInstance() (interface{}) {
	if obj.ServiceDefinitionV2Email != nil {
		return obj.ServiceDefinitionV2Email
	}


	if obj.ServiceDefinitionV2Slack != nil {
		return obj.ServiceDefinitionV2Slack
	}


	if obj.ServiceDefinitionV2MSTeams != nil {
		return obj.ServiceDefinitionV2MSTeams
	}


	// all schemas are nil
	return nil
}
