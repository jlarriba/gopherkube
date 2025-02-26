/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenStackFloatingIPSpec defines the desired state of OpenStackFloatingIP
type OpenStackFloatingIPSpec struct {
	// Cloud is the OpenStackCloud hosting this resource
	Cloud string `json:"cloud"`

	// ID is the OpenStack UUID of the resource. If left empty, the
	// controller will create a new resource and populate this field. If
	// manually populated, the controller will adopt the corresponding
	// resource.
	ID string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	// FloatingNetwork is the external OpenStackNetwork where the floating
	// IP is to be created.
	FloatingNetwork string `json:"floatingNetwork,omitempty"`

	FloatingIP string `json:"floatingIPAddress,omitempty"`
	PortID     string `json:"portID,omitempty"`
	FixedIP    string `json:"fixedIPAddress,omitempty"`
	SubnetID   string `json:"subnetID,omitempty"`
	TenantID   string `json:"tenantID,omitempty"`
	ProjectID  string `json:"projectID,omitempty"`

	// Unmanaged, when true, means that no action will be performed in
	// OpenStack against this resource. This is false by default, except
	// for pre-existing resources that are adopted by passing ID on
	// creation.
	Unmanaged *bool `json:"unmanaged,omitempty"`
}

// OpenStackFloatingIPStatus defines the observed state of OpenStackFloatingIP
type OpenStackFloatingIPStatus struct {
	// ID is the unique identifier for the floating IP instance.
	ID string `json:"id,omitempty"`

	// Description for the floating IP instance.
	Description string `json:"description,omitempty"`

	// FloatingNetworkID is the UUID of the external network where the floating
	// IP is to be created.
	FloatingNetworkID string `json:"floatingNetworkID,omitempty"`

	// FloatingIP is the address of the floating IP on the external network.
	FloatingIP string `json:"floatingIPAddress,omitempty"`

	// PortID is the UUID of the port on an internal network that is associated
	// with the floating IP.
	PortID string `json:"portIP,omitempty"`

	// FixedIP is the specific IP address of the internal port which should be
	// associated with the floating IP.
	FixedIP string `json:"fixedIPAddress,omitempty"`

	// TenantID is the project owner of the floating IP. Only admin users can
	// specify a project identifier other than its own.
	TenantID string `json:"tenantID,omitempty"`

	// UpdatedAt contains the timestamp of when the resource was last
	// changed.
	UpdatedAt string `json:"updatedAt,omitempty"`

	// CreatedAt contains the timestamp of when the resource was created.
	CreatedAt string `json:"createdAt,omitempty"`

	// ProjectID is the project owner of the floating IP.
	ProjectID string `json:"projectID,omitempty"`

	// Status is the condition of the API resource.
	Status string `json:"status,omitempty"`

	// RouterID is the ID of the router used for this floating IP.
	RouterID string `json:"routerID,omitempty"`

	// Tags optionally set via extensions/attributestags
	Tags []string `json:"tags,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OpenStackFloatingIP is the Schema for the openstackfloatingips API
type OpenStackFloatingIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenStackFloatingIPSpec   `json:"spec,omitempty"`
	Status OpenStackFloatingIPStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OpenStackFloatingIPList contains a list of OpenStackFloatingIP
type OpenStackFloatingIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenStackFloatingIP `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenStackFloatingIP{}, &OpenStackFloatingIPList{})
}
