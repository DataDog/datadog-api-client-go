// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// UserTeamIncluded - Included resources related to the team membership
type UserTeamIncluded struct {
	User *User
	Team *Team
	AbbreviatedTeam *AbbreviatedTeam
	UserTeamUser *UserTeamUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsUserTeamIncluded is a convenience function that returns User wrapped in UserTeamIncluded.
func UserAsUserTeamIncluded(v *User) UserTeamIncluded {
	return UserTeamIncluded{User: v}
}

// TeamAsUserTeamIncluded is a convenience function that returns Team wrapped in UserTeamIncluded.
func TeamAsUserTeamIncluded(v *Team) UserTeamIncluded {
	return UserTeamIncluded{Team: v}
}

// AbbreviatedTeamAsUserTeamIncluded is a convenience function that returns AbbreviatedTeam wrapped in UserTeamIncluded.
func AbbreviatedTeamAsUserTeamIncluded(v *AbbreviatedTeam) UserTeamIncluded {
	return UserTeamIncluded{AbbreviatedTeam: v}
}

// UserTeamUserAsUserTeamIncluded is a convenience function that returns UserTeamUser wrapped in UserTeamIncluded.
func UserTeamUserAsUserTeamIncluded(v *UserTeamUser) UserTeamIncluded {
	return UserTeamIncluded{UserTeamUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UserTeamIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = datadog.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := datadog.Marshal(obj.User)
			if string(jsonUser) == "{}" { // empty struct
				obj.User = nil
			} else {
				match++
			}
		} else {
			obj.User = nil
		}
	} else {
		obj.User = nil
	}

	// try to unmarshal data into Team
	err = datadog.Unmarshal(data, &obj.Team)
	if err == nil {
		if obj.Team != nil && obj.Team.UnparsedObject == nil {
			jsonTeam, _ := datadog.Marshal(obj.Team)
			if string(jsonTeam) == "{}" { // empty struct
				obj.Team = nil
			} else {
				match++
			}
		} else {
			obj.Team = nil
		}
	} else {
		obj.Team = nil
	}

	// try to unmarshal data into AbbreviatedTeam
	err = datadog.Unmarshal(data, &obj.AbbreviatedTeam)
	if err == nil {
		if obj.AbbreviatedTeam != nil && obj.AbbreviatedTeam.UnparsedObject == nil {
			jsonAbbreviatedTeam, _ := datadog.Marshal(obj.AbbreviatedTeam)
			if string(jsonAbbreviatedTeam) == "{}" { // empty struct
				obj.AbbreviatedTeam = nil
			} else {
				match++
			}
		} else {
			obj.AbbreviatedTeam = nil
		}
	} else {
		obj.AbbreviatedTeam = nil
	}

	// try to unmarshal data into UserTeamUser
	err = datadog.Unmarshal(data, &obj.UserTeamUser)
	if err == nil {
		if obj.UserTeamUser != nil && obj.UserTeamUser.UnparsedObject == nil {
			jsonUserTeamUser, _ := datadog.Marshal(obj.UserTeamUser)
			if string(jsonUserTeamUser) == "{}" { // empty struct
				obj.UserTeamUser = nil
			} else {
				match++
			}
		} else {
			obj.UserTeamUser = nil
		}
	} else {
		obj.UserTeamUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.Team = nil
		obj.AbbreviatedTeam = nil
		obj.UserTeamUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UserTeamIncluded) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}


	if obj.Team != nil {
		return datadog.Marshal(&obj.Team)
	}


	if obj.AbbreviatedTeam != nil {
		return datadog.Marshal(&obj.AbbreviatedTeam)
	}


	if obj.UserTeamUser != nil {
		return datadog.Marshal(&obj.UserTeamUser)
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UserTeamIncluded) GetActualInstance() (interface{}) {
	if obj.User != nil {
		return obj.User
	}


	if obj.Team != nil {
		return obj.Team
	}


	if obj.AbbreviatedTeam != nil {
		return obj.AbbreviatedTeam
	}


	if obj.UserTeamUser != nil {
		return obj.UserTeamUser
	}


	// all schemas are nil
	return nil
}
