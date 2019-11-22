/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
	"os"
)

// IdpFormData struct for IdpFormData
type IdpFormData struct {
	// The path to the XML metadata file you wish to upload.
	IdpFile *os.File `json:"idp_file"`
}

// GetIdpFile returns the IdpFile field value
func (o *IdpFormData) GetIdpFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.IdpFile
}

// SetIdpFile sets field value
func (o *IdpFormData) SetIdpFile(v *os.File) {
	o.IdpFile = v
}

type NullableIdpFormData struct {
	Value        IdpFormData
	ExplicitNull bool
}

func (v NullableIdpFormData) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIdpFormData) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
