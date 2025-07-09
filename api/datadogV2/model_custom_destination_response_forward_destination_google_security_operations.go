// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationGoogleSecurityOperations The Google Security Operations destination.
type CustomDestinationResponseForwardDestinationGoogleSecurityOperations struct {
	// Google Security Operations destination authentication.
	Auth CustomDestinationResponseGoogleSecurityOperationsDestinationAuth `json:"auth"`
	// The customer ID of the Google Security Operations account.
	CustomerId string `json:"customer_id"`
	// The namespace of the Google Security Operations account.
	Namespace string `json:"namespace"`
	// The `CustomDestinationResponseForwardDestinationGoogleSecurityOperations` `regional_endpoint`.
	RegionalEndpoint string `json:"regional_endpoint"`
	// Type of the Google Security Operations destination.
	Type CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseForwardDestinationGoogleSecurityOperations instantiates a new CustomDestinationResponseForwardDestinationGoogleSecurityOperations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseForwardDestinationGoogleSecurityOperations(auth CustomDestinationResponseGoogleSecurityOperationsDestinationAuth, customerId string, namespace string, regionalEndpoint string, typeVar CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) *CustomDestinationResponseForwardDestinationGoogleSecurityOperations {
	this := CustomDestinationResponseForwardDestinationGoogleSecurityOperations{}
	this.Auth = auth
	this.CustomerId = customerId
	this.Namespace = namespace
	this.RegionalEndpoint = regionalEndpoint
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseForwardDestinationGoogleSecurityOperationsWithDefaults instantiates a new CustomDestinationResponseForwardDestinationGoogleSecurityOperations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseForwardDestinationGoogleSecurityOperationsWithDefaults() *CustomDestinationResponseForwardDestinationGoogleSecurityOperations {
	this := CustomDestinationResponseForwardDestinationGoogleSecurityOperations{}
	var typeVar CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType = CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetAuth() CustomDestinationResponseGoogleSecurityOperationsDestinationAuth {
	if o == nil {
		var ret CustomDestinationResponseGoogleSecurityOperationsDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetAuthOk() (*CustomDestinationResponseGoogleSecurityOperationsDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) SetAuth(v CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) {
	o.Auth = v
}

// GetCustomerId returns the CustomerId field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetNamespace returns the Namespace field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) SetNamespace(v string) {
	o.Namespace = v
}

// GetRegionalEndpoint returns the RegionalEndpoint field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetRegionalEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RegionalEndpoint
}

// GetRegionalEndpointOk returns a tuple with the RegionalEndpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetRegionalEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionalEndpoint, true
}

// SetRegionalEndpoint sets field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) SetRegionalEndpoint(v string) {
	o.RegionalEndpoint = v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetType() CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType {
	if o == nil {
		var ret CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) GetTypeOk() (*CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) SetType(v CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseForwardDestinationGoogleSecurityOperations) MarshalJSON() ([]byte, error) {
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
func (o *CustomDestinationResponseForwardDestinationGoogleSecurityOperations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth             *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth        `json:"auth"`
		CustomerId       *string                                                                  `json:"customer_id"`
		Namespace        *string                                                                  `json:"namespace"`
		RegionalEndpoint *string                                                                  `json:"regional_endpoint"`
		Type             *CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType `json:"type"`
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
