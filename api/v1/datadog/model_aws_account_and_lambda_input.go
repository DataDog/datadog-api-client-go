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
)

// AwsAccountAndLambdaInput struct for AwsAccountAndLambdaInput
type AwsAccountAndLambdaInput struct {
	// Your AWS Account ID without dashes.
	AccountId string `json:"account_id"`
	// ARN of the Datadog Lambda created during the Datadog-Amazon Web services Log collection setup.
	LambdaArn string `json:"lambda_arn"`
}

// GetAccountId returns the AccountId field value
func (o *AwsAccountAndLambdaInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *AwsAccountAndLambdaInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetLambdaArn returns the LambdaArn field value
func (o *AwsAccountAndLambdaInput) GetLambdaArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LambdaArn
}

// SetLambdaArn sets field value
func (o *AwsAccountAndLambdaInput) SetLambdaArn(v string) {
	o.LambdaArn = v
}

type NullableAwsAccountAndLambdaInput struct {
	Value        AwsAccountAndLambdaInput
	ExplicitNull bool
}

func (v NullableAwsAccountAndLambdaInput) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsAccountAndLambdaInput) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
