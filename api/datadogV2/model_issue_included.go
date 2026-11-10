// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueIncluded - An array of related resources, returned when the `include` query parameter is used.
type IssueIncluded struct {
	IssueCase        *IssueCase
	IssueUser        *IssueUser
	IssueTeam        *IssueTeam
	IssuePullRequest *IssuePullRequest
	IssueWatch       *IssueWatch
	IssueSample      *IssueSample

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IssueCaseAsIssueIncluded is a convenience function that returns IssueCase wrapped in IssueIncluded.
func IssueCaseAsIssueIncluded(v *IssueCase) IssueIncluded {
	return IssueIncluded{IssueCase: v}
}

// IssueUserAsIssueIncluded is a convenience function that returns IssueUser wrapped in IssueIncluded.
func IssueUserAsIssueIncluded(v *IssueUser) IssueIncluded {
	return IssueIncluded{IssueUser: v}
}

// IssueTeamAsIssueIncluded is a convenience function that returns IssueTeam wrapped in IssueIncluded.
func IssueTeamAsIssueIncluded(v *IssueTeam) IssueIncluded {
	return IssueIncluded{IssueTeam: v}
}

// IssuePullRequestAsIssueIncluded is a convenience function that returns IssuePullRequest wrapped in IssueIncluded.
func IssuePullRequestAsIssueIncluded(v *IssuePullRequest) IssueIncluded {
	return IssueIncluded{IssuePullRequest: v}
}

// IssueWatchAsIssueIncluded is a convenience function that returns IssueWatch wrapped in IssueIncluded.
func IssueWatchAsIssueIncluded(v *IssueWatch) IssueIncluded {
	return IssueIncluded{IssueWatch: v}
}

// IssueSampleAsIssueIncluded is a convenience function that returns IssueSample wrapped in IssueIncluded.
func IssueSampleAsIssueIncluded(v *IssueSample) IssueIncluded {
	return IssueIncluded{IssueSample: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IssueIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IssueCase
	err = datadog.Unmarshal(data, &obj.IssueCase)
	if err == nil {
		if obj.IssueCase != nil && obj.IssueCase.UnparsedObject == nil {
			jsonIssueCase, _ := datadog.Marshal(obj.IssueCase)
			if string(jsonIssueCase) == "{}" { // empty struct
				obj.IssueCase = nil
			} else {
				match++
			}
		} else {
			obj.IssueCase = nil
		}
	} else {
		obj.IssueCase = nil
	}

	// try to unmarshal data into IssueUser
	err = datadog.Unmarshal(data, &obj.IssueUser)
	if err == nil {
		if obj.IssueUser != nil && obj.IssueUser.UnparsedObject == nil {
			jsonIssueUser, _ := datadog.Marshal(obj.IssueUser)
			if string(jsonIssueUser) == "{}" { // empty struct
				obj.IssueUser = nil
			} else {
				match++
			}
		} else {
			obj.IssueUser = nil
		}
	} else {
		obj.IssueUser = nil
	}

	// try to unmarshal data into IssueTeam
	err = datadog.Unmarshal(data, &obj.IssueTeam)
	if err == nil {
		if obj.IssueTeam != nil && obj.IssueTeam.UnparsedObject == nil {
			jsonIssueTeam, _ := datadog.Marshal(obj.IssueTeam)
			if string(jsonIssueTeam) == "{}" { // empty struct
				obj.IssueTeam = nil
			} else {
				match++
			}
		} else {
			obj.IssueTeam = nil
		}
	} else {
		obj.IssueTeam = nil
	}

	// try to unmarshal data into IssuePullRequest
	err = datadog.Unmarshal(data, &obj.IssuePullRequest)
	if err == nil {
		if obj.IssuePullRequest != nil && obj.IssuePullRequest.UnparsedObject == nil {
			jsonIssuePullRequest, _ := datadog.Marshal(obj.IssuePullRequest)
			if string(jsonIssuePullRequest) == "{}" { // empty struct
				obj.IssuePullRequest = nil
			} else {
				match++
			}
		} else {
			obj.IssuePullRequest = nil
		}
	} else {
		obj.IssuePullRequest = nil
	}

	// try to unmarshal data into IssueWatch
	err = datadog.Unmarshal(data, &obj.IssueWatch)
	if err == nil {
		if obj.IssueWatch != nil && obj.IssueWatch.UnparsedObject == nil {
			jsonIssueWatch, _ := datadog.Marshal(obj.IssueWatch)
			if string(jsonIssueWatch) == "{}" { // empty struct
				obj.IssueWatch = nil
			} else {
				match++
			}
		} else {
			obj.IssueWatch = nil
		}
	} else {
		obj.IssueWatch = nil
	}

	// try to unmarshal data into IssueSample
	err = datadog.Unmarshal(data, &obj.IssueSample)
	if err == nil {
		if obj.IssueSample != nil && obj.IssueSample.UnparsedObject == nil {
			jsonIssueSample, _ := datadog.Marshal(obj.IssueSample)
			if string(jsonIssueSample) == "{}" { // empty struct
				obj.IssueSample = nil
			} else {
				match++
			}
		} else {
			obj.IssueSample = nil
		}
	} else {
		obj.IssueSample = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IssueCase = nil
		obj.IssueUser = nil
		obj.IssueTeam = nil
		obj.IssuePullRequest = nil
		obj.IssueWatch = nil
		obj.IssueSample = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IssueIncluded) MarshalJSON() ([]byte, error) {
	if obj.IssueCase != nil {
		return datadog.Marshal(&obj.IssueCase)
	}

	if obj.IssueUser != nil {
		return datadog.Marshal(&obj.IssueUser)
	}

	if obj.IssueTeam != nil {
		return datadog.Marshal(&obj.IssueTeam)
	}

	if obj.IssuePullRequest != nil {
		return datadog.Marshal(&obj.IssuePullRequest)
	}

	if obj.IssueWatch != nil {
		return datadog.Marshal(&obj.IssueWatch)
	}

	if obj.IssueSample != nil {
		return datadog.Marshal(&obj.IssueSample)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IssueIncluded) GetActualInstance() interface{} {
	if obj.IssueCase != nil {
		return obj.IssueCase
	}

	if obj.IssueUser != nil {
		return obj.IssueUser
	}

	if obj.IssueTeam != nil {
		return obj.IssueTeam
	}

	if obj.IssuePullRequest != nil {
		return obj.IssuePullRequest
	}

	if obj.IssueWatch != nil {
		return obj.IssueWatch
	}

	if obj.IssueSample != nil {
		return obj.IssueSample
	}

	// all schemas are nil
	return nil
}
