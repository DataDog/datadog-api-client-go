/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsBasicAuth - Object to handle basic authentication when performing the test.
type SyntheticsBasicAuth struct {
	SyntheticsBasicAuthWeb   *SyntheticsBasicAuthWeb
	SyntheticsBasicAuthSigv4 *SyntheticsBasicAuthSigv4
	SyntheticsBasicAuthNTLM  *SyntheticsBasicAuthNTLM

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsBasicAuthWebAsSyntheticsBasicAuth is a convenience function that returns SyntheticsBasicAuthWeb wrapped in SyntheticsBasicAuth
func SyntheticsBasicAuthWebAsSyntheticsBasicAuth(v *SyntheticsBasicAuthWeb) SyntheticsBasicAuth {
	return SyntheticsBasicAuth{SyntheticsBasicAuthWeb: v}
}

// SyntheticsBasicAuthSigv4AsSyntheticsBasicAuth is a convenience function that returns SyntheticsBasicAuthSigv4 wrapped in SyntheticsBasicAuth
func SyntheticsBasicAuthSigv4AsSyntheticsBasicAuth(v *SyntheticsBasicAuthSigv4) SyntheticsBasicAuth {
	return SyntheticsBasicAuth{SyntheticsBasicAuthSigv4: v}
}

// SyntheticsBasicAuthNTLMAsSyntheticsBasicAuth is a convenience function that returns SyntheticsBasicAuthNTLM wrapped in SyntheticsBasicAuth
func SyntheticsBasicAuthNTLMAsSyntheticsBasicAuth(v *SyntheticsBasicAuthNTLM) SyntheticsBasicAuth {
	return SyntheticsBasicAuth{SyntheticsBasicAuthNTLM: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyntheticsBasicAuth) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsBasicAuthWeb
	err = json.Unmarshal(data, &dst.SyntheticsBasicAuthWeb)
	if err == nil {
		if dst.SyntheticsBasicAuthWeb != nil && dst.SyntheticsBasicAuthWeb.UnparsedObject == nil {
			jsonSyntheticsBasicAuthWeb, _ := json.Marshal(dst.SyntheticsBasicAuthWeb)
			if string(jsonSyntheticsBasicAuthWeb) == "{}" { // empty struct
				dst.SyntheticsBasicAuthWeb = nil
			} else {
				match++
			}
		} else {
			dst.SyntheticsBasicAuthWeb = nil
		}
	} else {
		dst.SyntheticsBasicAuthWeb = nil
	}

	// try to unmarshal data into SyntheticsBasicAuthSigv4
	err = json.Unmarshal(data, &dst.SyntheticsBasicAuthSigv4)
	if err == nil {
		if dst.SyntheticsBasicAuthSigv4 != nil && dst.SyntheticsBasicAuthSigv4.UnparsedObject == nil {
			jsonSyntheticsBasicAuthSigv4, _ := json.Marshal(dst.SyntheticsBasicAuthSigv4)
			if string(jsonSyntheticsBasicAuthSigv4) == "{}" { // empty struct
				dst.SyntheticsBasicAuthSigv4 = nil
			} else {
				match++
			}
		} else {
			dst.SyntheticsBasicAuthSigv4 = nil
		}
	} else {
		dst.SyntheticsBasicAuthSigv4 = nil
	}

	// try to unmarshal data into SyntheticsBasicAuthNTLM
	err = json.Unmarshal(data, &dst.SyntheticsBasicAuthNTLM)
	if err == nil {
		if dst.SyntheticsBasicAuthNTLM != nil && dst.SyntheticsBasicAuthNTLM.UnparsedObject == nil {
			jsonSyntheticsBasicAuthNTLM, _ := json.Marshal(dst.SyntheticsBasicAuthNTLM)
			if string(jsonSyntheticsBasicAuthNTLM) == "{}" { // empty struct
				dst.SyntheticsBasicAuthNTLM = nil
			} else {
				match++
			}
		} else {
			dst.SyntheticsBasicAuthNTLM = nil
		}
	} else {
		dst.SyntheticsBasicAuthNTLM = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.SyntheticsBasicAuthWeb = nil
		dst.SyntheticsBasicAuthSigv4 = nil
		dst.SyntheticsBasicAuthNTLM = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyntheticsBasicAuth) MarshalJSON() ([]byte, error) {
	if src.SyntheticsBasicAuthWeb != nil {
		return json.Marshal(&src.SyntheticsBasicAuthWeb)
	}

	if src.SyntheticsBasicAuthSigv4 != nil {
		return json.Marshal(&src.SyntheticsBasicAuthSigv4)
	}

	if src.SyntheticsBasicAuthNTLM != nil {
		return json.Marshal(&src.SyntheticsBasicAuthNTLM)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyntheticsBasicAuth) GetActualInstance() interface{} {
	if obj.SyntheticsBasicAuthWeb != nil {
		return obj.SyntheticsBasicAuthWeb
	}

	if obj.SyntheticsBasicAuthSigv4 != nil {
		return obj.SyntheticsBasicAuthSigv4
	}

	if obj.SyntheticsBasicAuthNTLM != nil {
		return obj.SyntheticsBasicAuthNTLM
	}

	// all schemas are nil
	return nil
}

type NullableSyntheticsBasicAuth struct {
	value *SyntheticsBasicAuth
	isSet bool
}

func (v NullableSyntheticsBasicAuth) Get() *SyntheticsBasicAuth {
	return v.value
}

func (v *NullableSyntheticsBasicAuth) Set(val *SyntheticsBasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBasicAuth(val *SyntheticsBasicAuth) *NullableSyntheticsBasicAuth {
	return &NullableSyntheticsBasicAuth{value: val, isSet: true}
}

func (v NullableSyntheticsBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
