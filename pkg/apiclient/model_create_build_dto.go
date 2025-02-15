/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateBuildDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBuildDTO{}

// CreateBuildDTO struct for CreateBuildDTO
type CreateBuildDTO struct {
	Branch                string            `json:"branch"`
	EnvVars               map[string]string `json:"envVars"`
	PrebuildId            *string           `json:"prebuildId,omitempty"`
	WorkspaceTemplateName string            `json:"workspaceTemplateName"`
}

type _CreateBuildDTO CreateBuildDTO

// NewCreateBuildDTO instantiates a new CreateBuildDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBuildDTO(branch string, envVars map[string]string, workspaceTemplateName string) *CreateBuildDTO {
	this := CreateBuildDTO{}
	this.Branch = branch
	this.EnvVars = envVars
	this.WorkspaceTemplateName = workspaceTemplateName
	return &this
}

// NewCreateBuildDTOWithDefaults instantiates a new CreateBuildDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBuildDTOWithDefaults() *CreateBuildDTO {
	this := CreateBuildDTO{}
	return &this
}

// GetBranch returns the Branch field value
func (o *CreateBuildDTO) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *CreateBuildDTO) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *CreateBuildDTO) SetBranch(v string) {
	o.Branch = v
}

// GetEnvVars returns the EnvVars field value
func (o *CreateBuildDTO) GetEnvVars() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value
// and a boolean to check if the value has been set.
func (o *CreateBuildDTO) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvVars, true
}

// SetEnvVars sets field value
func (o *CreateBuildDTO) SetEnvVars(v map[string]string) {
	o.EnvVars = v
}

// GetPrebuildId returns the PrebuildId field value if set, zero value otherwise.
func (o *CreateBuildDTO) GetPrebuildId() string {
	if o == nil || IsNil(o.PrebuildId) {
		var ret string
		return ret
	}
	return *o.PrebuildId
}

// GetPrebuildIdOk returns a tuple with the PrebuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuildDTO) GetPrebuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrebuildId) {
		return nil, false
	}
	return o.PrebuildId, true
}

// HasPrebuildId returns a boolean if a field has been set.
func (o *CreateBuildDTO) HasPrebuildId() bool {
	if o != nil && !IsNil(o.PrebuildId) {
		return true
	}

	return false
}

// SetPrebuildId gets a reference to the given string and assigns it to the PrebuildId field.
func (o *CreateBuildDTO) SetPrebuildId(v string) {
	o.PrebuildId = &v
}

// GetWorkspaceTemplateName returns the WorkspaceTemplateName field value
func (o *CreateBuildDTO) GetWorkspaceTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceTemplateName
}

// GetWorkspaceTemplateNameOk returns a tuple with the WorkspaceTemplateName field value
// and a boolean to check if the value has been set.
func (o *CreateBuildDTO) GetWorkspaceTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceTemplateName, true
}

// SetWorkspaceTemplateName sets field value
func (o *CreateBuildDTO) SetWorkspaceTemplateName(v string) {
	o.WorkspaceTemplateName = v
}

func (o CreateBuildDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBuildDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["branch"] = o.Branch
	toSerialize["envVars"] = o.EnvVars
	if !IsNil(o.PrebuildId) {
		toSerialize["prebuildId"] = o.PrebuildId
	}
	toSerialize["workspaceTemplateName"] = o.WorkspaceTemplateName
	return toSerialize, nil
}

func (o *CreateBuildDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"branch",
		"envVars",
		"workspaceTemplateName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateBuildDTO := _CreateBuildDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBuildDTO)

	if err != nil {
		return err
	}

	*o = CreateBuildDTO(varCreateBuildDTO)

	return err
}

type NullableCreateBuildDTO struct {
	value *CreateBuildDTO
	isSet bool
}

func (v NullableCreateBuildDTO) Get() *CreateBuildDTO {
	return v.value
}

func (v *NullableCreateBuildDTO) Set(val *CreateBuildDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBuildDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBuildDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBuildDTO(val *CreateBuildDTO) *NullableCreateBuildDTO {
	return &NullableCreateBuildDTO{value: val, isSet: true}
}

func (v NullableCreateBuildDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBuildDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
