// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerView - Test-only generated schema for the dashboard versioning proof of concept.
type ContainerView struct {
	OrchestrationView_20270101       *OrchestrationView
	DashboardInlineSchema37_20270101 *DashboardInlineSchema37

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OrchestrationView_20270101AsContainerView is a convenience function that returns OrchestrationView wrapped in ContainerView.
func OrchestrationView_20270101AsContainerView(v *OrchestrationView) ContainerView {
	return ContainerView{OrchestrationView_20270101: v}
}

// DashboardInlineSchema37_20270101AsContainerView is a convenience function that returns DashboardInlineSchema37 wrapped in ContainerView.
func DashboardInlineSchema37_20270101AsContainerView(v *DashboardInlineSchema37) ContainerView {
	return ContainerView{DashboardInlineSchema37_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ContainerView) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OrchestrationView_20270101
	err = datadog.Unmarshal(data, &obj.OrchestrationView_20270101)
	if err == nil {
		if obj.OrchestrationView_20270101 != nil {
			jsonOrchestrationView_20270101, _ := datadog.Marshal(obj.OrchestrationView_20270101)
			if string(jsonOrchestrationView_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.OrchestrationView_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.OrchestrationView_20270101 = nil
		}
	} else {
		obj.OrchestrationView_20270101 = nil
	}

	// try to unmarshal data into DashboardInlineSchema37_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema37_20270101)
	if err == nil {
		if obj.DashboardInlineSchema37_20270101 != nil {
			jsonDashboardInlineSchema37_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema37_20270101)
			if string(jsonDashboardInlineSchema37_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema37_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema37_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema37_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OrchestrationView_20270101 = nil
		obj.DashboardInlineSchema37_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ContainerView) MarshalJSON() ([]byte, error) {
	if obj.OrchestrationView_20270101 != nil {
		return datadog.Marshal(&obj.OrchestrationView_20270101)
	}

	if obj.DashboardInlineSchema37_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema37_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ContainerView) GetActualInstance() interface{} {
	if obj.OrchestrationView_20270101 != nil {
		return obj.OrchestrationView_20270101
	}

	if obj.DashboardInlineSchema37_20270101 != nil {
		return obj.DashboardInlineSchema37_20270101
	}

	// all schemas are nil
	return nil
}
