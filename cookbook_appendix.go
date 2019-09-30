package blueprint

// ExtraPropertiesType represents extra properties to define for resources, infrastructures or deployments. This properties are provisioner or deployment specific and they should document them when they expect any.
// swagger:model
type ExtraPropertiesType map[string]string

// Drive holds information about a data drive attached to a node
// swagger:model
type Drive struct {
	// Unique name for the drive
	// required:true
	Name string `json:"name"`
	// Type of the drive. It can be "SSD" or "HDD"
	// pattern: SSD|HDD
	// example: SSD
	Type string `json:"type"`
	// Size of the disk in Mb
	// required:true
	Size int64 `json:"size"`
}

// ResourceType has information about a node that needs to be created by a deployer.
// swagger:model
type ResourceType struct {
	// Suffix for the hostname. The real hostname will be formed of the infrastructure name + resource name
	// required:true
	// unique:true
	Name string `json:"name"`
	// Type of the VM to create i.e. n1-small
	// example: n1-small
	Type string `json:"type"`
	// CPU speed in Mhz. Ignored if type is provided
	CPU int `json:"cpu"`
	// Number of cores. Ignored if type is provided
	Cores int `json:"cores"`
	// RAM quantity in Mb. Ignored if type is provided
	RAM int64 `json:"ram"`
	// Boot disk size in Mb
	// required:true
	Disk int64 `json:"disk"`
	// Role that this VM plays. In case of a Kubernetes deployment at least one "master" is needed.
	Role string `json:"role"`
	// Boot image ID to use
	// required:true
	ImageId string `json:"image_id"`
	// IP to assign this VM. In case it's not specified, the first available one will be used.
	IP string `json:"ip,omitempty"`
	// List of data drives to attach to this VM
	Drives []Drive `json:"drives"`
	// Extra properties to pass to the provider or the provisioner
	ExtraProperties ExtraPropertiesType `json:"extra_properties"`
}

// CloudProviderInfo contains information about a cloud provider
// swagger:model
type CloudProviderInfo struct {
	// Endpoint to use for this infrastructure
	// required:true
	APIEndpoint string `json:"api_endpoint"`
	// Type of the infrastructure. i.e AWS, Cloudsigma, GCP or Edge
	APIType string `json:"api_type"`
	// Secret identifier to use to log in to the infrastructure manager.
	SecretID string `json:"secret_id"`
	// Credentials to access the cloud provider. Either this or secret_id is mandatory. Each cloud provider should define the format of this element.
	Credentials map[string]interface{} `json:"credentials"`
}

// InfrastructureType is a set of resources that need to be created or configured to form a cluster
// swagger:model
type InfrastructureType struct {
	// Unique name for the infrastructure
	// required:true
	// unique:true
	Name string `json:"name"`
	// Optional description for the infrastructure
	Description string `json:"description"`
	// Type of the infrastructure: Cloud or Edge: Cloud infrastructures mean that the resources will be VMs that need to be instantiated. Edge means that the infrastructure is already in place and its information will be added to the database but no further work will be done by a deployer.
	Type string `json:"type"`
	// Provider information. Required in case of Cloud type
	Provider CloudProviderInfo `json:"provider"`
	// List of resources to deploy
	// required:true
	Resources []ResourceType `json:"resources"`
	// Extra properties to pass to the provider or the provisioner
	ExtraProperties ExtraPropertiesType `json:"extra_properties"`
}

// Deployment is a set of infrastructures that need to be instantiated or configurated to form clusters
// swagger:model
type Deployment struct {
	// Name for this deployment
	// required:true
	// unique:true
	Name string `json:"name"`
	// Optional description
	Description string `json:"description"`
	// List of infrastructures to deploy for this hybrid deployment
	// required:true
	Infrastructures []InfrastructureType `json:"infrastructures"`
}

// DriveInfo is the information of a drive that has been instantiated
// swagger:model
type DriveInfo struct {
	// Name of the data drive
	// unique:true
	// required:true
	Name string `json:"name"`
	// Size of the disk in bytes
	// required:true
	Size int64 `json:"size"`
}

// NodeInfo is the information of a virtual machine that has been instantiated or a physical one that was pre-existing
// swagger:model
type NodeInfo struct {
	// Hostname of the node.
	// requiered:true
	// unique:true
	Hostname string `json:"hostname"`
	// Role of the node. Master or slave in case of Kubernetes.
	// example:master
	Role string `json:"role"`
	// CPU speed in Mhz.
	CPU int `json:"cpu"`
	// Number of cores.
	Cores int `json:"cores"`
	// RAM quantity in bytes.
	RAM int64 `json:"ram"`
	// IP assigned to this node.
	// required:true
	// unique:true
	IP string `json:"ip"`
	// Size of the boot disk in bytes
	// required:true
	// unique:true
	DriveSize int64 `json:"drive_size" bson:"drive_size"`
	// Data drives information
	DataDrives []DriveInfo `json:"data_drives" bson:"data_drives"`
	// Extra properties to pass to the provider or the provisioner
	ExtraProperties ExtraPropertiesType `json:"extra_properties"`
}

// VDCInfo contains information about related to a VDC running in a kubernetes cluster
// swagger:model
type VDCInfo struct {
	Ports map[string]int
}

// InfrastructureDeploymentInfo contains information about a cluster of nodes that has been instantiated or were already existing.
// swagger:model
type InfrastructureDeploymentInfo struct {
	// Unique infrastructure ID on the deployment
	// required:true
	// unique:true
	ID string `json:"id"`
	// Name of the infrastructure
	Name string `json:"name"`
	// Type of the infrastructure: cloud or edge
	// pattern:cloud|edge
	// required:true
	Type string `json:"type"`
	// Provider information
	// required:true
	Provider CloudProviderInfo `json:"provider"`
	// Set of nodes in the infrastructure indexed by role
	// required:true
	Nodes map[string][]NodeInfo
	// Status of the infrastructure
	Status string `json:"status"`
	// Configuration of VDCs running in the cluster, indexed by VDC identifier.
	VDCs map[string]VDCInfo `json:"vdcs"`
	// Set weather the VDM is running in this cluster or not
	VDM bool
	// Extra properties to pass to the provider or the provisioner
	ExtraProperties ExtraPropertiesType `json:"extra_properties"`
}

// DeploymentInfo contains information of a deployment than may compromise several clusters
// swagger:model
type DeploymentInfo struct {
	// Unique ID for the deployment
	// required:true
	// unique:true
	ID string `json:"id" bson:"_id"`
	// Name of the deployment
	Name string `json:"name"`
	// Lisf of infrastructures, each one representing a different cluster.
	Infrastructures map[string]InfrastructureDeploymentInfo `json:"infrastructures"`
	// Extra properties bound to the deployment
	ExtraProperties ExtraPropertiesType `json:"extra_properties"`
	// Global status of the deployment
	Status string `json:"status"`
}

// CookbookAppendix is the definition of the Cookbook Appendix section in the blueprint
// swagger:model
type CookbookAppendix struct {

	// Information about the resources which are available to deploy VDCs
	// required:true
	Resources Deployment `json:"Resources"`

	// Information about the clusters that were deployed with this blueprint
	// required: true
	Deployment DeploymentInfo `json:"Deployment"`

	IdentityAccessManagement *COOKBOOKAPPENDIXIdentityAccessManagement `json:"Identity_Access_Management,omitempty"`
}
