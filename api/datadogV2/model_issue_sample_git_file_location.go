// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleGitFileLocation Location of a stack frame in a repository file. Empty when the source code is unavailable.
type IssueSampleGitFileLocation struct {
	// Column of the location, starting at 1.
	Column *int64 `json:"column,omitempty"`
	// Path of the file in the repository.
	Filepath *string `json:"filepath,omitempty"`
	// First line of the location, starting at 1.
	FirstLine *int64 `json:"first_line,omitempty"`
	// Last line of the location, starting at 1.
	LastLine *int64 `json:"last_line,omitempty"`
	// Git reference of the file, usually a commit SHA.
	Ref *string `json:"ref,omitempty"`
	// Repository identifier.
	RepoId *string `json:"repo_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleGitFileLocation instantiates a new IssueSampleGitFileLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleGitFileLocation() *IssueSampleGitFileLocation {
	this := IssueSampleGitFileLocation{}
	return &this
}

// NewIssueSampleGitFileLocationWithDefaults instantiates a new IssueSampleGitFileLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleGitFileLocationWithDefaults() *IssueSampleGitFileLocation {
	this := IssueSampleGitFileLocation{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetColumn() int64 {
	if o == nil || o.Column == nil {
		var ret int64
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetColumnOk() (*int64, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given int64 and assigns it to the Column field.
func (o *IssueSampleGitFileLocation) SetColumn(v int64) {
	o.Column = &v
}

// GetFilepath returns the Filepath field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetFilepath() string {
	if o == nil || o.Filepath == nil {
		var ret string
		return ret
	}
	return *o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetFilepathOk() (*string, bool) {
	if o == nil || o.Filepath == nil {
		return nil, false
	}
	return o.Filepath, true
}

// HasFilepath returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasFilepath() bool {
	return o != nil && o.Filepath != nil
}

// SetFilepath gets a reference to the given string and assigns it to the Filepath field.
func (o *IssueSampleGitFileLocation) SetFilepath(v string) {
	o.Filepath = &v
}

// GetFirstLine returns the FirstLine field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetFirstLine() int64 {
	if o == nil || o.FirstLine == nil {
		var ret int64
		return ret
	}
	return *o.FirstLine
}

// GetFirstLineOk returns a tuple with the FirstLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetFirstLineOk() (*int64, bool) {
	if o == nil || o.FirstLine == nil {
		return nil, false
	}
	return o.FirstLine, true
}

// HasFirstLine returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasFirstLine() bool {
	return o != nil && o.FirstLine != nil
}

// SetFirstLine gets a reference to the given int64 and assigns it to the FirstLine field.
func (o *IssueSampleGitFileLocation) SetFirstLine(v int64) {
	o.FirstLine = &v
}

// GetLastLine returns the LastLine field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetLastLine() int64 {
	if o == nil || o.LastLine == nil {
		var ret int64
		return ret
	}
	return *o.LastLine
}

// GetLastLineOk returns a tuple with the LastLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetLastLineOk() (*int64, bool) {
	if o == nil || o.LastLine == nil {
		return nil, false
	}
	return o.LastLine, true
}

// HasLastLine returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasLastLine() bool {
	return o != nil && o.LastLine != nil
}

// SetLastLine gets a reference to the given int64 and assigns it to the LastLine field.
func (o *IssueSampleGitFileLocation) SetLastLine(v int64) {
	o.LastLine = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasRef() bool {
	return o != nil && o.Ref != nil
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *IssueSampleGitFileLocation) SetRef(v string) {
	o.Ref = &v
}

// GetRepoId returns the RepoId field value if set, zero value otherwise.
func (o *IssueSampleGitFileLocation) GetRepoId() string {
	if o == nil || o.RepoId == nil {
		var ret string
		return ret
	}
	return *o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleGitFileLocation) GetRepoIdOk() (*string, bool) {
	if o == nil || o.RepoId == nil {
		return nil, false
	}
	return o.RepoId, true
}

// HasRepoId returns a boolean if a field has been set.
func (o *IssueSampleGitFileLocation) HasRepoId() bool {
	return o != nil && o.RepoId != nil
}

// SetRepoId gets a reference to the given string and assigns it to the RepoId field.
func (o *IssueSampleGitFileLocation) SetRepoId(v string) {
	o.RepoId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleGitFileLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.Filepath != nil {
		toSerialize["filepath"] = o.Filepath
	}
	if o.FirstLine != nil {
		toSerialize["first_line"] = o.FirstLine
	}
	if o.LastLine != nil {
		toSerialize["last_line"] = o.LastLine
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	if o.RepoId != nil {
		toSerialize["repo_id"] = o.RepoId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleGitFileLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column    *int64  `json:"column,omitempty"`
		Filepath  *string `json:"filepath,omitempty"`
		FirstLine *int64  `json:"first_line,omitempty"`
		LastLine  *int64  `json:"last_line,omitempty"`
		Ref       *string `json:"ref,omitempty"`
		RepoId    *string `json:"repo_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "filepath", "first_line", "last_line", "ref", "repo_id"})
	} else {
		return err
	}
	o.Column = all.Column
	o.Filepath = all.Filepath
	o.FirstLine = all.FirstLine
	o.LastLine = all.LastLine
	o.Ref = all.Ref
	o.RepoId = all.RepoId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
