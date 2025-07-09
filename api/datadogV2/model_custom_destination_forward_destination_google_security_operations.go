// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationGoogleSecurityOperations The Google Security Operations destination.
type CustomDestinationForwardDestinationGoogleSecurityOperations struct {
	// Google Security Operations destination authentication.
	Auth CustomDestinationGoogleSecurityOperationsDestinationAuth `json:"auth"`
	// The customer ID of the Google Security Operations account.
	CustomerId string `json:"customer_id"`
	// The namespace of the Google Security Operations account.
	Namespace string `json:"namespace"`
	// The `CustomDestinationForwardDestinationGoogleSecurityOperations` `regional_endpoint`.
	RegionalEndpoint string `json:"regional_endpoint"`
	// Type of the Google Security Operations destination.
	Type CustomDestinationForwardDestinationGoogleSecurityOperationsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationForwardDestinationGoogleSecurityOperations instantiates a new CustomDestinationForwardDestinationGoogleSecurityOperations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationForwardDestinationGoogleSecurityOperations(auth CustomDestinationGoogleSecurityOperationsDestinationAuth, customerId string, namespace string, regionalEndpoint string, typeVar CustomDestinationForwardDestinationGoogleSecurityOperationsType) *CustomDestinationForwardDestinationGoogleSecurityOperations {
	this := CustomDestinationForwardDestinationGoogleSecurityOperations{}
	this.Auth = auth
	this.CustomerId = customerId
	this.Namespace = namespace
	this.RegionalEndpoint = regionalEndpoint
	this.Type = typeVar
	return &this
}

// NewCustomDestinationForwardDestinationGoogleSecurityOperationsWithDefaults instantiates a new CustomDestinationForwardDestinationGoogleSecurityOperations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationForwardDestinationGoogleSecurityOperationsWithDefaults() *CustomDestinationForwardDestinationGoogleSecurityOperations {
	this := CustomDestinationForwardDestinationGoogleSecurityOperations{}
	var typeVar CustomDestinationForwardDestinationGoogleSecurityOperationsType = CUSTOMDESTINATIONFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetAuth() CustomDestinationGoogleSecurityOperationsDestinationAuth {
	if o == nil {
		var ret CustomDestinationGoogleSecurityOperationsDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetAuthOk() (*CustomDestinationGoogleSecurityOperationsDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) SetAuth(v CustomDestinationGoogleSecurityOperationsDestinationAuth) {
	o.Auth = v
}

// GetCustomerId returns the CustomerId field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetNamespace returns the Namespace field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) SetNamespace(v string) {
	o.Namespace = v
}

// GetRegionalEndpoint returns the RegionalEndpoint field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetRegionalEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RegionalEndpoint
}

// GetRegionalEndpointOk returns a tuple with the RegionalEndpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetRegionalEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionalEndpoint, true
}

// SetRegionalEndpoint sets field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) SetRegionalEndpoint(v string) {
	o.RegionalEndpoint = v
}

// GetType returns the Type field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetType() CustomDestinationForwardDestinationGoogleSecurityOperationsType {
	if o == nil {
		var ret CustomDestinationForwardDestinationGoogleSecurityOperationsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) GetTypeOk() (*CustomDestinationForwardDestinationGoogleSecurityOperationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) SetType(v CustomDestinationForwardDestinationGoogleSecurityOperationsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationForwardDestinationGoogleSecurityOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["namespace"] = o.Namespace
	toSerialize["regional_endpoint"] = o.RegionalEndpoint
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationForwardDestinationGoogleSecurityOperations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth             *CustomDestinationGoogleSecurityOperationsDestinationAuth        `json:"auth"`
		CustomerId       *string                                                          `json:"customer_id"`
		Namespace        *string                                                          `json:"namespace"`
		RegionalEndpoint *string                                                          `json:"regional_endpoint"`
		Type             *CustomDestinationForwardDestinationGoogleSecurityOperationsType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.CustomerId == nil {
		return fmt.Errorf("required field customer_id missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.RegionalEndpoint == nil {
		return fmt.Errorf("required field regional_endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "customer_id", "namespace", "regional_endpoint", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = *all.Auth
	o.CustomerId = *all.CustomerId
	o.Namespace = *all.Namespace
	o.RegionalEndpoint = *all.RegionalEndpoint
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
