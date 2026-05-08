// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogOrgMetadata Metadata about the organization that sent the email.
type TransportWebhookLogOrgMetadata struct {
	// Country code or name used for billing purposes.
	BillingCountry *string `json:"billing_country,omitempty"`
	// The Datadog billing plan for the organization (for example, "pro", "enterprise").
	BillingPlan *string `json:"billing_plan,omitempty"`
	// Support or account tier assigned to the organization (for example, "tier-1").
	CustomerTier *string `json:"customer_tier,omitempty"`
	// Primary email domain associated with the organization (for example, "example.com").
	Domain *string `json:"domain,omitempty"`
	// Industry classification of the organization (for example, "technology", "finance").
	Industry *string `json:"industry,omitempty"`
	// Whether the organization is enrolled in the Datadog bug bounty program.
	IsBugbounty *string `json:"is_bugbounty,omitempty"`
	// Whether the organization operates as a Managed Service Provider managing child orgs.
	IsMsp *string `json:"is_msp,omitempty"`
	// Display name of the organization as configured in Datadog account settings.
	Name *string `json:"name,omitempty"`
	// Globally unique identifier for the Datadog organization (UUID v1 format).
	OrgUuid *string `json:"org_uuid,omitempty"`
	// Identifier of the immediate parent organization, if this is a child org.
	ParentOrgId *string `json:"parent_org_id,omitempty"`
	// Whether the organization has a premium support plan with Datadog.
	PremiumSupport *string `json:"premium_support,omitempty"`
	// Identifier of the top-level parent organization in a multi-org account hierarchy.
	RootOrgId *string `json:"root_org_id,omitempty"`
	// Display name of the top-level parent organization in a multi-org account hierarchy.
	RootOrgName *string `json:"root_org_name,omitempty"`
	// Country code or name used for shipping or regional assignment.
	ShippingCountry *string `json:"shipping_country,omitempty"`
	// Website URL provided during organization registration.
	Website *string `json:"website,omitempty"`
	// ISO 8601 timestamp of when the Datadog organization was created.
	WhenCreated *string `json:"when_created,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogOrgMetadata instantiates a new TransportWebhookLogOrgMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogOrgMetadata() *TransportWebhookLogOrgMetadata {
	this := TransportWebhookLogOrgMetadata{}
	return &this
}

// NewTransportWebhookLogOrgMetadataWithDefaults instantiates a new TransportWebhookLogOrgMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogOrgMetadataWithDefaults() *TransportWebhookLogOrgMetadata {
	this := TransportWebhookLogOrgMetadata{}
	return &this
}

// GetBillingCountry returns the BillingCountry field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetBillingCountry() string {
	if o == nil || o.BillingCountry == nil {
		var ret string
		return ret
	}
	return *o.BillingCountry
}

// GetBillingCountryOk returns a tuple with the BillingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetBillingCountryOk() (*string, bool) {
	if o == nil || o.BillingCountry == nil {
		return nil, false
	}
	return o.BillingCountry, true
}

// HasBillingCountry returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasBillingCountry() bool {
	return o != nil && o.BillingCountry != nil
}

// SetBillingCountry gets a reference to the given string and assigns it to the BillingCountry field.
func (o *TransportWebhookLogOrgMetadata) SetBillingCountry(v string) {
	o.BillingCountry = &v
}

// GetBillingPlan returns the BillingPlan field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetBillingPlan() string {
	if o == nil || o.BillingPlan == nil {
		var ret string
		return ret
	}
	return *o.BillingPlan
}

// GetBillingPlanOk returns a tuple with the BillingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetBillingPlanOk() (*string, bool) {
	if o == nil || o.BillingPlan == nil {
		return nil, false
	}
	return o.BillingPlan, true
}

// HasBillingPlan returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasBillingPlan() bool {
	return o != nil && o.BillingPlan != nil
}

// SetBillingPlan gets a reference to the given string and assigns it to the BillingPlan field.
func (o *TransportWebhookLogOrgMetadata) SetBillingPlan(v string) {
	o.BillingPlan = &v
}

// GetCustomerTier returns the CustomerTier field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetCustomerTier() string {
	if o == nil || o.CustomerTier == nil {
		var ret string
		return ret
	}
	return *o.CustomerTier
}

// GetCustomerTierOk returns a tuple with the CustomerTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetCustomerTierOk() (*string, bool) {
	if o == nil || o.CustomerTier == nil {
		return nil, false
	}
	return o.CustomerTier, true
}

// HasCustomerTier returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasCustomerTier() bool {
	return o != nil && o.CustomerTier != nil
}

// SetCustomerTier gets a reference to the given string and assigns it to the CustomerTier field.
func (o *TransportWebhookLogOrgMetadata) SetCustomerTier(v string) {
	o.CustomerTier = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *TransportWebhookLogOrgMetadata) SetDomain(v string) {
	o.Domain = &v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetIndustry() string {
	if o == nil || o.Industry == nil {
		var ret string
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetIndustryOk() (*string, bool) {
	if o == nil || o.Industry == nil {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasIndustry() bool {
	return o != nil && o.Industry != nil
}

// SetIndustry gets a reference to the given string and assigns it to the Industry field.
func (o *TransportWebhookLogOrgMetadata) SetIndustry(v string) {
	o.Industry = &v
}

// GetIsBugbounty returns the IsBugbounty field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetIsBugbounty() string {
	if o == nil || o.IsBugbounty == nil {
		var ret string
		return ret
	}
	return *o.IsBugbounty
}

// GetIsBugbountyOk returns a tuple with the IsBugbounty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetIsBugbountyOk() (*string, bool) {
	if o == nil || o.IsBugbounty == nil {
		return nil, false
	}
	return o.IsBugbounty, true
}

// HasIsBugbounty returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasIsBugbounty() bool {
	return o != nil && o.IsBugbounty != nil
}

// SetIsBugbounty gets a reference to the given string and assigns it to the IsBugbounty field.
func (o *TransportWebhookLogOrgMetadata) SetIsBugbounty(v string) {
	o.IsBugbounty = &v
}

// GetIsMsp returns the IsMsp field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetIsMsp() string {
	if o == nil || o.IsMsp == nil {
		var ret string
		return ret
	}
	return *o.IsMsp
}

// GetIsMspOk returns a tuple with the IsMsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetIsMspOk() (*string, bool) {
	if o == nil || o.IsMsp == nil {
		return nil, false
	}
	return o.IsMsp, true
}

// HasIsMsp returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasIsMsp() bool {
	return o != nil && o.IsMsp != nil
}

// SetIsMsp gets a reference to the given string and assigns it to the IsMsp field.
func (o *TransportWebhookLogOrgMetadata) SetIsMsp(v string) {
	o.IsMsp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransportWebhookLogOrgMetadata) SetName(v string) {
	o.Name = &v
}

// GetOrgUuid returns the OrgUuid field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetOrgUuid() string {
	if o == nil || o.OrgUuid == nil {
		var ret string
		return ret
	}
	return *o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetOrgUuidOk() (*string, bool) {
	if o == nil || o.OrgUuid == nil {
		return nil, false
	}
	return o.OrgUuid, true
}

// HasOrgUuid returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasOrgUuid() bool {
	return o != nil && o.OrgUuid != nil
}

// SetOrgUuid gets a reference to the given string and assigns it to the OrgUuid field.
func (o *TransportWebhookLogOrgMetadata) SetOrgUuid(v string) {
	o.OrgUuid = &v
}

// GetParentOrgId returns the ParentOrgId field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetParentOrgId() string {
	if o == nil || o.ParentOrgId == nil {
		var ret string
		return ret
	}
	return *o.ParentOrgId
}

// GetParentOrgIdOk returns a tuple with the ParentOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetParentOrgIdOk() (*string, bool) {
	if o == nil || o.ParentOrgId == nil {
		return nil, false
	}
	return o.ParentOrgId, true
}

// HasParentOrgId returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasParentOrgId() bool {
	return o != nil && o.ParentOrgId != nil
}

// SetParentOrgId gets a reference to the given string and assigns it to the ParentOrgId field.
func (o *TransportWebhookLogOrgMetadata) SetParentOrgId(v string) {
	o.ParentOrgId = &v
}

// GetPremiumSupport returns the PremiumSupport field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetPremiumSupport() string {
	if o == nil || o.PremiumSupport == nil {
		var ret string
		return ret
	}
	return *o.PremiumSupport
}

// GetPremiumSupportOk returns a tuple with the PremiumSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetPremiumSupportOk() (*string, bool) {
	if o == nil || o.PremiumSupport == nil {
		return nil, false
	}
	return o.PremiumSupport, true
}

// HasPremiumSupport returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasPremiumSupport() bool {
	return o != nil && o.PremiumSupport != nil
}

// SetPremiumSupport gets a reference to the given string and assigns it to the PremiumSupport field.
func (o *TransportWebhookLogOrgMetadata) SetPremiumSupport(v string) {
	o.PremiumSupport = &v
}

// GetRootOrgId returns the RootOrgId field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetRootOrgId() string {
	if o == nil || o.RootOrgId == nil {
		var ret string
		return ret
	}
	return *o.RootOrgId
}

// GetRootOrgIdOk returns a tuple with the RootOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetRootOrgIdOk() (*string, bool) {
	if o == nil || o.RootOrgId == nil {
		return nil, false
	}
	return o.RootOrgId, true
}

// HasRootOrgId returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasRootOrgId() bool {
	return o != nil && o.RootOrgId != nil
}

// SetRootOrgId gets a reference to the given string and assigns it to the RootOrgId field.
func (o *TransportWebhookLogOrgMetadata) SetRootOrgId(v string) {
	o.RootOrgId = &v
}

// GetRootOrgName returns the RootOrgName field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetRootOrgName() string {
	if o == nil || o.RootOrgName == nil {
		var ret string
		return ret
	}
	return *o.RootOrgName
}

// GetRootOrgNameOk returns a tuple with the RootOrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetRootOrgNameOk() (*string, bool) {
	if o == nil || o.RootOrgName == nil {
		return nil, false
	}
	return o.RootOrgName, true
}

// HasRootOrgName returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasRootOrgName() bool {
	return o != nil && o.RootOrgName != nil
}

// SetRootOrgName gets a reference to the given string and assigns it to the RootOrgName field.
func (o *TransportWebhookLogOrgMetadata) SetRootOrgName(v string) {
	o.RootOrgName = &v
}

// GetShippingCountry returns the ShippingCountry field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetShippingCountry() string {
	if o == nil || o.ShippingCountry == nil {
		var ret string
		return ret
	}
	return *o.ShippingCountry
}

// GetShippingCountryOk returns a tuple with the ShippingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetShippingCountryOk() (*string, bool) {
	if o == nil || o.ShippingCountry == nil {
		return nil, false
	}
	return o.ShippingCountry, true
}

// HasShippingCountry returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasShippingCountry() bool {
	return o != nil && o.ShippingCountry != nil
}

// SetShippingCountry gets a reference to the given string and assigns it to the ShippingCountry field.
func (o *TransportWebhookLogOrgMetadata) SetShippingCountry(v string) {
	o.ShippingCountry = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasWebsite() bool {
	return o != nil && o.Website != nil
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *TransportWebhookLogOrgMetadata) SetWebsite(v string) {
	o.Website = &v
}

// GetWhenCreated returns the WhenCreated field value if set, zero value otherwise.
func (o *TransportWebhookLogOrgMetadata) GetWhenCreated() string {
	if o == nil || o.WhenCreated == nil {
		var ret string
		return ret
	}
	return *o.WhenCreated
}

// GetWhenCreatedOk returns a tuple with the WhenCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogOrgMetadata) GetWhenCreatedOk() (*string, bool) {
	if o == nil || o.WhenCreated == nil {
		return nil, false
	}
	return o.WhenCreated, true
}

// HasWhenCreated returns a boolean if a field has been set.
func (o *TransportWebhookLogOrgMetadata) HasWhenCreated() bool {
	return o != nil && o.WhenCreated != nil
}

// SetWhenCreated gets a reference to the given string and assigns it to the WhenCreated field.
func (o *TransportWebhookLogOrgMetadata) SetWhenCreated(v string) {
	o.WhenCreated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogOrgMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BillingCountry != nil {
		toSerialize["billing_country"] = o.BillingCountry
	}
	if o.BillingPlan != nil {
		toSerialize["billing_plan"] = o.BillingPlan
	}
	if o.CustomerTier != nil {
		toSerialize["customer_tier"] = o.CustomerTier
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Industry != nil {
		toSerialize["industry"] = o.Industry
	}
	if o.IsBugbounty != nil {
		toSerialize["is_bugbounty"] = o.IsBugbounty
	}
	if o.IsMsp != nil {
		toSerialize["is_msp"] = o.IsMsp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgUuid != nil {
		toSerialize["org_uuid"] = o.OrgUuid
	}
	if o.ParentOrgId != nil {
		toSerialize["parent_org_id"] = o.ParentOrgId
	}
	if o.PremiumSupport != nil {
		toSerialize["premium_support"] = o.PremiumSupport
	}
	if o.RootOrgId != nil {
		toSerialize["root_org_id"] = o.RootOrgId
	}
	if o.RootOrgName != nil {
		toSerialize["root_org_name"] = o.RootOrgName
	}
	if o.ShippingCountry != nil {
		toSerialize["shipping_country"] = o.ShippingCountry
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.WhenCreated != nil {
		toSerialize["when_created"] = o.WhenCreated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogOrgMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BillingCountry  *string `json:"billing_country,omitempty"`
		BillingPlan     *string `json:"billing_plan,omitempty"`
		CustomerTier    *string `json:"customer_tier,omitempty"`
		Domain          *string `json:"domain,omitempty"`
		Industry        *string `json:"industry,omitempty"`
		IsBugbounty     *string `json:"is_bugbounty,omitempty"`
		IsMsp           *string `json:"is_msp,omitempty"`
		Name            *string `json:"name,omitempty"`
		OrgUuid         *string `json:"org_uuid,omitempty"`
		ParentOrgId     *string `json:"parent_org_id,omitempty"`
		PremiumSupport  *string `json:"premium_support,omitempty"`
		RootOrgId       *string `json:"root_org_id,omitempty"`
		RootOrgName     *string `json:"root_org_name,omitempty"`
		ShippingCountry *string `json:"shipping_country,omitempty"`
		Website         *string `json:"website,omitempty"`
		WhenCreated     *string `json:"when_created,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"billing_country", "billing_plan", "customer_tier", "domain", "industry", "is_bugbounty", "is_msp", "name", "org_uuid", "parent_org_id", "premium_support", "root_org_id", "root_org_name", "shipping_country", "website", "when_created"})
	} else {
		return err
	}
	o.BillingCountry = all.BillingCountry
	o.BillingPlan = all.BillingPlan
	o.CustomerTier = all.CustomerTier
	o.Domain = all.Domain
	o.Industry = all.Industry
	o.IsBugbounty = all.IsBugbounty
	o.IsMsp = all.IsMsp
	o.Name = all.Name
	o.OrgUuid = all.OrgUuid
	o.ParentOrgId = all.ParentOrgId
	o.PremiumSupport = all.PremiumSupport
	o.RootOrgId = all.RootOrgId
	o.RootOrgName = all.RootOrgName
	o.ShippingCountry = all.ShippingCountry
	o.Website = all.Website
	o.WhenCreated = all.WhenCreated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
