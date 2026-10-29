// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SendSlackMessageActionCustomizableContent Options that customize the content of the Slack message sent for a Page.
type SendSlackMessageActionCustomizableContent struct {
	// Whether to include the Page description in the Slack message.
	IncludeDescription *bool `json:"include_description,omitempty"`
	// Whether to include the source that triggered the Page in the Slack message.
	IncludeSource *bool `json:"include_source,omitempty"`
	// Whether to mirror replies posted in the Slack message thread to the Page timeline.
	MirrorReplies *bool `json:"mirror_replies,omitempty"`
	// Whether to post Page timeline updates as replies in the Slack message thread.
	ThreadTimeline *bool `json:"thread_timeline,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSendSlackMessageActionCustomizableContent instantiates a new SendSlackMessageActionCustomizableContent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSendSlackMessageActionCustomizableContent() *SendSlackMessageActionCustomizableContent {
	this := SendSlackMessageActionCustomizableContent{}
	return &this
}

// NewSendSlackMessageActionCustomizableContentWithDefaults instantiates a new SendSlackMessageActionCustomizableContent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSendSlackMessageActionCustomizableContentWithDefaults() *SendSlackMessageActionCustomizableContent {
	this := SendSlackMessageActionCustomizableContent{}
	return &this
}

// GetIncludeDescription returns the IncludeDescription field value if set, zero value otherwise.
func (o *SendSlackMessageActionCustomizableContent) GetIncludeDescription() bool {
	if o == nil || o.IncludeDescription == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDescription
}

// GetIncludeDescriptionOk returns a tuple with the IncludeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSlackMessageActionCustomizableContent) GetIncludeDescriptionOk() (*bool, bool) {
	if o == nil || o.IncludeDescription == nil {
		return nil, false
	}
	return o.IncludeDescription, true
}

// HasIncludeDescription returns a boolean if a field has been set.
func (o *SendSlackMessageActionCustomizableContent) HasIncludeDescription() bool {
	return o != nil && o.IncludeDescription != nil
}

// SetIncludeDescription gets a reference to the given bool and assigns it to the IncludeDescription field.
func (o *SendSlackMessageActionCustomizableContent) SetIncludeDescription(v bool) {
	o.IncludeDescription = &v
}

// GetIncludeSource returns the IncludeSource field value if set, zero value otherwise.
func (o *SendSlackMessageActionCustomizableContent) GetIncludeSource() bool {
	if o == nil || o.IncludeSource == nil {
		var ret bool
		return ret
	}
	return *o.IncludeSource
}

// GetIncludeSourceOk returns a tuple with the IncludeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSlackMessageActionCustomizableContent) GetIncludeSourceOk() (*bool, bool) {
	if o == nil || o.IncludeSource == nil {
		return nil, false
	}
	return o.IncludeSource, true
}

// HasIncludeSource returns a boolean if a field has been set.
func (o *SendSlackMessageActionCustomizableContent) HasIncludeSource() bool {
	return o != nil && o.IncludeSource != nil
}

// SetIncludeSource gets a reference to the given bool and assigns it to the IncludeSource field.
func (o *SendSlackMessageActionCustomizableContent) SetIncludeSource(v bool) {
	o.IncludeSource = &v
}

// GetMirrorReplies returns the MirrorReplies field value if set, zero value otherwise.
func (o *SendSlackMessageActionCustomizableContent) GetMirrorReplies() bool {
	if o == nil || o.MirrorReplies == nil {
		var ret bool
		return ret
	}
	return *o.MirrorReplies
}

// GetMirrorRepliesOk returns a tuple with the MirrorReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSlackMessageActionCustomizableContent) GetMirrorRepliesOk() (*bool, bool) {
	if o == nil || o.MirrorReplies == nil {
		return nil, false
	}
	return o.MirrorReplies, true
}

// HasMirrorReplies returns a boolean if a field has been set.
func (o *SendSlackMessageActionCustomizableContent) HasMirrorReplies() bool {
	return o != nil && o.MirrorReplies != nil
}

// SetMirrorReplies gets a reference to the given bool and assigns it to the MirrorReplies field.
func (o *SendSlackMessageActionCustomizableContent) SetMirrorReplies(v bool) {
	o.MirrorReplies = &v
}

// GetThreadTimeline returns the ThreadTimeline field value if set, zero value otherwise.
func (o *SendSlackMessageActionCustomizableContent) GetThreadTimeline() bool {
	if o == nil || o.ThreadTimeline == nil {
		var ret bool
		return ret
	}
	return *o.ThreadTimeline
}

// GetThreadTimelineOk returns a tuple with the ThreadTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSlackMessageActionCustomizableContent) GetThreadTimelineOk() (*bool, bool) {
	if o == nil || o.ThreadTimeline == nil {
		return nil, false
	}
	return o.ThreadTimeline, true
}

// HasThreadTimeline returns a boolean if a field has been set.
func (o *SendSlackMessageActionCustomizableContent) HasThreadTimeline() bool {
	return o != nil && o.ThreadTimeline != nil
}

// SetThreadTimeline gets a reference to the given bool and assigns it to the ThreadTimeline field.
func (o *SendSlackMessageActionCustomizableContent) SetThreadTimeline(v bool) {
	o.ThreadTimeline = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SendSlackMessageActionCustomizableContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeDescription != nil {
		toSerialize["include_description"] = o.IncludeDescription
	}
	if o.IncludeSource != nil {
		toSerialize["include_source"] = o.IncludeSource
	}
	if o.MirrorReplies != nil {
		toSerialize["mirror_replies"] = o.MirrorReplies
	}
	if o.ThreadTimeline != nil {
		toSerialize["thread_timeline"] = o.ThreadTimeline
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SendSlackMessageActionCustomizableContent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeDescription *bool `json:"include_description,omitempty"`
		IncludeSource      *bool `json:"include_source,omitempty"`
		MirrorReplies      *bool `json:"mirror_replies,omitempty"`
		ThreadTimeline     *bool `json:"thread_timeline,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_description", "include_source", "mirror_replies", "thread_timeline"})
	} else {
		return err
	}
	o.IncludeDescription = all.IncludeDescription
	o.IncludeSource = all.IncludeSource
	o.MirrorReplies = all.MirrorReplies
	o.ThreadTimeline = all.ThreadTimeline

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
