/**
 * Copyright 2018 Atos
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 *
 * This is being developed for the DITAS Project: https://www.ditas-project.eu/
 */

package blueprint

import (
	"github.com/go-openapi/spec"
)

// This is a VDC Blueprint which consists of five sections
type Blueprint struct {
	ID                 string                         `json:"_id"`
	AbstractProperties []AbstractPropertiesMethodType `json:"ABSTRACT_PROPERTIES"`
	CookbookAppendix   CookbookAppendix               `json:"COOKBOOK_APPENDIX"`  // CookbookAppendix is the definition of the Cookbook Appendix section in the blueprint
	DataManagement     []DataManagement               `json:"DATA_MANAGEMENT"`    // list of methods
	ExposedAPI         spec.Swagger                   `json:"EXPOSED_API"`        // The CAF RESTful API of the VDC, written according to the current version (3.0.1) of the; OpenAPI Specification (OAS), but also adapted to DITAS requirements
	InternalStructure  InternalStructure              `json:"INTERNAL_STRUCTURE"` // General information about the VDC Blueprint
}
type COOKBOOKAPPENDIXIdentityAccessManagement struct {
	Mapping        []Mapping                `json:"mapping"`
	ValidationKeys []map[string]interface{} `json:"validation_keys"`
}

type Mapping struct {
	MappingURL *string       `json:"mapping_url,omitempty"`
	Provider   *string       `json:"provider,omitempty"`
	RoleMap    *RoleMapUnion `json:"role_map"`
	Roles      []string      `json:"roles"`
}

type RoleMapElement struct {
	Matcher  *string  `json:"matcher,omitempty"`
	Priority *float64 `json:"priority,omitempty"`
	Roles    []string `json:"roles"`
}

type DataManagement struct {
	Attributes Attributes `json:"attributes"` // goal trees
	MethodID   string     `json:"method_id"`  // The id (operationId) of the method (as indicated in the EXPOSED_API.paths field)
}

// goal trees
type Attributes struct {
	DataUtility []DataUtility `json:"dataUtility"`
	Privacy     []DataUtility `json:"privacy"`
	Security    []DataUtility `json:"security"`
}

// definition of the metric
type DataUtility struct {
	ID         *string             `json:"id,omitempty"`         // id of the metric
	Name       *string             `json:"name,omitempty"`       // name of the metric
	Properties map[string]Property `json:"properties,omitempty"` // properties related to the metric
	Type       *string             `json:"type,omitempty"`       // type of the metric
}

// properties related to the metric
type Property struct {
	Maximum *float64     `json:"maximum,omitempty"` // lower limit of the offered property
	Minimum *float64     `json:"minimum,omitempty"` // upper limit of the offered property
	Unit    *string      `json:"unit,omitempty"`    // unit of measure of the property
	Value   *interface{} `json:"value"`             // value of the property
}

// General information about the VDC Blueprint
type InternalStructure struct {
	DALImages                map[string]DALImage                        `json:"DAL_Images,omitempty"` // Docker images that must be deployed in the DAL indexed by DAL name. It will be used to; compose the service name and the DNS entry that other images in the cluster can access to.
	DataSources              []DataSource                               `json:"Data_Sources"`
	Flow                     *Flow                                      `json:"Flow,omitempty"` // The data flow that implements the VDC
	IdentityAccessManagement *INTERNALSTRUCTUREIdentityAccessManagement `json:"Identity_Access_Management,omitempty"`
	MethodsInput             *MethodsInput                              `json:"Methods_Input,omitempty"` // This filed contains the part of the data source that each method needs to be executed
	Overview                 Overview                                   `json:"Overview"`
	TestingOutputData        []TestingOutputDatum                       `json:"Testing_Output_Data"`
	VDCImages                map[string]Image                           `json:"VDC_Images,omitempty"`
}

// DALImage about the DAL including its original location
type DALImage struct {
	Images     map[string]Image `json:"images,omitempty"` // Set of images to deploy indexed by the image identifier
	OriginalIP string           `json:"original_ip"`      // IP of the original DAL's location
}

// Image is the information about an image that will be deployed by the deployment engine
type Image struct {
	ExternalPort *int64            `json:"external_port,omitempty"` // Port in which this image must be exposed. It must be unique across all images in all the; ImageSets defined in this blueprint. Due to limitations in k8s, the port range must be; bewteen 30000 and 32767
	Image        string            `json:"image"`                   // Image is the image name in the standard format [group]/<image_name>:[release]
	InternalPort *int64            `json:"internal_port,omitempty"` // Port in which the docker image is listening internally. Two images inside the same; ImageSet can't have the same internal port.
	Environment  map[string]string `json:"environment,omitempty"`   // Environment is a set of environment variables to pass to this image. It can have some special variables if the value is in the form ${var}
}

type DataSource struct {
	Class       *Class                 `json:"class,omitempty"`
	Description *string                `json:"description,omitempty"`
	ID          string                 `json:"id"` // A unique identifier
	Location    *Location              `json:"location,omitempty"`
	Parameters  map[string]interface{} `json:"parameters,omitempty"` // Connection parameters
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Type        *Type                  `json:"type,omitempty"`
}

// The data flow that implements the VDC
type Flow struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Platform   *Platform              `json:"platform,omitempty"`
	SourceCode interface{}            `json:"source_code"`
}

type INTERNALSTRUCTUREIdentityAccessManagement struct {
	IamEndpoint string            `json:"iam_endpoint"`
	JwksURI     string            `json:"jwks_uri"`
	Provider    []ProviderElement `json:"provider"`
	Roles       []string          `json:"roles"`
}

type ProviderElement struct {
	LoginPortal *string `json:"loginPortal,omitempty"`
	Name        string  `json:"name"`
	Type        *string `json:"type,omitempty"`
	URI         string  `json:"uri"`
}

// This filed contains the part of the data source that each method needs to be executed
type MethodsInput struct {
	Methods []Method `json:"Methods"` // The list of methods
}

type Method struct {
	DataSources []DataSourceElement `json:"dataSources"`         // The list of data sources required by the method
	MethodID    *string             `json:"method_id,omitempty"` // The id (operationId) of the method (as indicated in the EXPOSED_API.paths field)
}

type DataSourceElement struct {
	Database       []Database `json:"database"`                  // the list of databases required by a method in a data source
	DataSourceID   *string    `json:"dataSource_id,omitempty"`   // The id of the data sources (as indicated in the Data_Sources field)
	DataSourceType *string    `json:"dataSource_type,omitempty"` // The type of the data sources (relationa/not_relational/object)
}

type Database struct {
	DatabaseID *string `json:"database_id,omitempty"` // The id of the database
	Tables     []Table `json:"tables"`                // the list of tables/collections required by a method in a data source
}

type Table struct {
	Columns []Column `json:"columns"`
	TableID *string  `json:"table_id,omitempty"` // The id of the tables/collection
}

type Column struct {
	ColumnID           *string `json:"column_id,omitempty"`          // The id of the column/field
	ComputeDataUtility *bool   `json:"computeDataUtility,omitempty"` // True if it is required for data utility computation
}

type Overview struct {
	Description string `json:"description"` // This field should contain a short description of the VDC Blueprint
	Name        string `json:"name"`        // This field should contain the name of the VDC Blueprint
	Tags        []Tag  `json:"tags"`        // Each element of this array should contain some keywords that describe the functionality; of each one exposed VDC method
}

type Tag struct {
	MethodID string   `json:"method_id"` // The id (operationId) of the method (as indicated in the EXPOSED_API.paths field)
	Tags     []string `json:"tags"`
}

type TestingOutputDatum struct {
	MethodID string `json:"method_id"` // The id (operationId) of the method (as indicated in the EXPOSED_API.paths field)
	ZipData  string `json:"zip_data"`  // The URI to the zip testing output data for each one exposed VDC method
}

type Class string

const (
	API                Class = "api"
	DataStream         Class = "data stream"
	ObjectStorage      Class = "object storage"
	RelationalDatabase Class = "relational database"
	TimeSeriesDatabase Class = "time-series database"
)

type Location string

const (
	Cloud Location = "cloud"
	Edge  Location = "edge"
)

type Type string

const (
	InfluxDB Type = "InfluxDB"
	Minio    Type = "Minio"
	MySQL    Type = "MySQL"
	Other    Type = "other"
	REST     Type = "rest"
)

type Platform string

const (
	NodeRED Platform = "Node-RED"
	Spark   Platform = "Spark"
)

type RoleMapUnion struct {
	RoleMapElementArray []RoleMapElement
	String              *string
}

// value of the property
type Value struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	String        *string
}
