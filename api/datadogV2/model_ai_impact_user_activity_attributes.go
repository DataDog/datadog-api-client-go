// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIImpactUserActivityAttributes Daily AI coding tool activity for a single user. Each entry reports whether the user was
// active on a given day and which AI tools and models they used.
type AIImpactUserActivityAttributes struct {
	// The day the activity refers to, in `YYYY-MM-DD` format.
	Day string `json:"day"`
	// Whether the user actively used the listed AI tools on that day.
	IsActive bool `json:"is_active"`
	// The AI models the user used on that day, for example `claude-sonnet-4.5` or `gpt-5`.
	// Values are lowercased and duplicates are removed.
	Models []string `json:"models,omitempty"`
	// The AI coding tools the user used on that day, for example `Claude Code`, `Cursor`, or
	// `GitHub Copilot`. Known tools are normalized to a canonical name (`claude_code`, `cursor`,
	// `copilot`), and other values are converted to snake case. Entries must not be empty.
	Tools []string `json:"tools"`
	// The email address of the user. It is case-insensitive and is matched against the
	// email addresses of commit authors.
	UserEmail string `json:"user_email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAIImpactUserActivityAttributes instantiates a new AIImpactUserActivityAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIImpactUserActivityAttributes(day string, isActive bool, tools []string, userEmail string) *AIImpactUserActivityAttributes {
	this := AIImpactUserActivityAttributes{}
	this.Day = day
	this.IsActive = isActive
	this.Tools = tools
	this.UserEmail = userEmail
	return &this
}

// NewAIImpactUserActivityAttributesWithDefaults instantiates a new AIImpactUserActivityAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIImpactUserActivityAttributesWithDefaults() *AIImpactUserActivityAttributes {
	this := AIImpactUserActivityAttributes{}
	return &this
}

// GetDay returns the Day field value.
func (o *AIImpactUserActivityAttributes) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityAttributes) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value.
func (o *AIImpactUserActivityAttributes) SetDay(v string) {
	o.Day = v
}

// GetIsActive returns the IsActive field value.
func (o *AIImpactUserActivityAttributes) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityAttributes) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value.
func (o *AIImpactUserActivityAttributes) SetIsActive(v bool) {
	o.IsActive = v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *AIImpactUserActivityAttributes) GetModels() []string {
	if o == nil || o.Models == nil {
		var ret []string
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityAttributes) GetModelsOk() (*[]string, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return &o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *AIImpactUserActivityAttributes) HasModels() bool {
	return o != nil && o.Models != nil
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *AIImpactUserActivityAttributes) SetModels(v []string) {
	o.Models = v
}

// GetTools returns the Tools field value.
func (o *AIImpactUserActivityAttributes) GetTools() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityAttributes) GetToolsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tools, true
}

// SetTools sets field value.
func (o *AIImpactUserActivityAttributes) SetTools(v []string) {
	o.Tools = v
}

// GetUserEmail returns the UserEmail field value.
func (o *AIImpactUserActivityAttributes) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityAttributes) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value.
func (o *AIImpactUserActivityAttributes) SetUserEmail(v string) {
	o.UserEmail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIImpactUserActivityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["day"] = o.Day
	toSerialize["is_active"] = o.IsActive
	if o.Models != nil {
		toSerialize["models"] = o.Models
	}
	toSerialize["tools"] = o.Tools
	toSerialize["user_email"] = o.UserEmail
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIImpactUserActivityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Day       *string   `json:"day"`
		IsActive  *bool     `json:"is_active"`
		Models    []string  `json:"models,omitempty"`
		Tools     *[]string `json:"tools"`
		UserEmail *string   `json:"user_email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Day == nil {
		return fmt.Errorf("required field day missing")
	}
	if all.IsActive == nil {
		return fmt.Errorf("required field is_active missing")
	}
	if all.Tools == nil {
		return fmt.Errorf("required field tools missing")
	}
	if all.UserEmail == nil {
		return fmt.Errorf("required field user_email missing")
	}
	o.Day = *all.Day
	o.IsActive = *all.IsActive
	o.Models = all.Models
	o.Tools = *all.Tools
	o.UserEmail = *all.UserEmail

	return nil
}
