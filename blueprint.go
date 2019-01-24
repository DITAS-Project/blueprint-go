/*
Copyright 2017 Atos

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

package blueprint

import (
	"github.com/go-openapi/spec"
)

type Drive struct {
	Name string `json:"name"` //Name of the image to use. Most of the times, it will be available as /dev/disk/by-id/${name} value in the VM
	Type string `json:"type"` //Type of the drive. It can be "SSD" or "HDD"
	Size int64  `json:"size"` //Size of the disk in Mb
}

type ResourceType struct {
	Name    string  `json:"name"`         //Suffix for the hostname. The real hostname will be formed of the infrastructure name + resource name
	Type    string  `json:"type"`         //Type of the VM to create i.e. n1-small
	CPU     int     `json:"cpu"`          //CPU speed in Mhz. Ignored if type is provided
	Cores   int     `json:"cores"`        //Number of cores. Ignored if type is provided
	RAM     int64   `json:"ram"`          //RAM quantity in Mb. Ignored if type is provided
	Disk    int64   `json:"disk"`         //Boot disk size in Mb
	Role    string  `json:"role"`         //Role that this VM plays. In case of a Kubernetes deployment at least one "master" is needed.
	ImageId string  `json:"image_id"`     //Boot image ID to use
	IP      string  `json:"ip,omitempty"` //IP to assign this VM. In case it's not specified, the first available one will be used.
	Drives  []Drive `json:"drives"`       //List of data drives to attach to this VM
}

type CloudProviderInfo struct {
	APIEndpoint string `json:"api_endpoint"` //Endpoint to use for this infrastructure
	APIType     string `json:"api_type"`     //Type of the infrastructure. i.e AWS, Cloudsigma, GCP or Edge
	KeypairID   string `json:"keypair_id"`   //Keypair to use to log in to the infrastructure manager
}

type InfrastructureType struct {
	Name        string            `json:"name"`        //Unique name for the infrastructure
	Description string            `json:"description"` //Optional description for the infrastructure
	Type        string            `json:"type"`        //Type of the infrastructure: Cloud or Edge
	Provider    CloudProviderInfo `json:"provider"`    //Provider information
	Resources   []ResourceType    `json:"resources"`   //List of resources to deploy
}

type CookbookAppendix struct {
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	Infrastructure []InfrastructureType `json:"infrastructure"`
}

type LeafType struct {
	Id          *string  `json:"id"`
	Description string   `json:"description"`
	Weight      int      `json:"weight"`
	Attributes  []string `json:"attributes"`
}

type TreeStructureType struct {
	Type     *string             `json:"type"`
	Children []TreeStructureType `json:"children"`
	Leaves   []LeafType          `json:"leaves"`
}

type GoalTreeType struct {
	DataUtility TreeStructureType `json:"dataUtility`
	Security    TreeStructureType `json:"security`
	Privacy     TreeStructureType `json:"privacy`
}

type AbstractPropertiesMethodType struct {
	MethodId  *string      `json:"method_id"`
	GoalTrees GoalTreeType `json:"goalTrees"`
}

type MetricPropertyType struct {
	Unit    string       `json:"unit"`
	Minimum *float64     `json:"minimum"`
	Maximum *float64     `json:"maximum"`
	Value   *interface{} `json:"value"`
}

//IsMinimumConstraint test if the MetricPropertyType has a minimum constraint
func (m *MetricPropertyType) IsMinimumConstraint() bool {
	return m.Minimum != nil
}

//IsMaximumConstraint test if the MetricPropertyType has a maximum constraint
func (m *MetricPropertyType) IsMaximumConstraint() bool {
	return m.Maximum != nil
}

//IsEqualityConstraint test if the MetricPropertyType has only a value and no min or max constraints
func (m *MetricPropertyType) IsEqualityConstraint() bool {
	return m.Value != nil && m.Maximum == nil && m.Minimum == nil
}

type ConstraintType struct {
	ID          *string                       `json:"id"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	Properties  map[string]MetricPropertyType `json:"properties"`
}

type DataManagementAttributesType struct {
	DataUtility []ConstraintType `json:"dataUtility"`
	Security    []ConstraintType `json:"security"`
	Privacy     []ConstraintType `json:"privacy"`
}

type DataManagementMethodType struct {
	MethodId   *string                      `json:"method_id"`
	Attributes DataManagementAttributesType `json:"attributes"`
}
type MethodTagType struct {
	ID   string   `json:"method_id"`
	Tags []string `json:"tags"`
}
type OverviewType struct {
	Name *string         `json:"Name"`
	Tags []MethodTagType `json:"tags"`
}

type DataSourceType struct {
	ID         *string                `json:"id"`
	Type       *string                `json:"type"`
	Parameters map[string]interface{} `json:"parameters"`
}

//InternalStructureType is the serialization of a DITAS concrete blueprint
// swagger:model
type InternalStructureType struct {

	// The overview section
	// required: true
	Overview OverviewType `json:"Overview"`

	// The datasources description
	// required: true
	DataSources []DataSourceType `json:"Data_Sources"`
}

//BlueprintType is the serialization of a DITAS concrete blueprint
// swagger:model
type BlueprintType struct {
	// The internal structure section
	// required: true
	InternalStructure InternalStructureType `json:"INTERNAL_STRUCTURE"`

	// The data management section
	// required: true
	DataManagement []DataManagementMethodType `json:"DATA_MANAGEMENT"`

	// The abstract properties section
	// required: true
	AbstractProperties []AbstractPropertiesMethodType `json:"ABSTRACT_PROPERTIES"`

	// The blueprint API description section
	// required: true
	API spec.Swagger `json:"EXPOSED_API"`

	// The cookbook appendix section containing the available resources
	// required: true
	CookbookAppendix CookbookAppendix `json:"COOKBOOK_APPENDIX"`
}
