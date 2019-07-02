package blueprint

// LeafType is a leaf in a tree data structure
// swagger:model
type LeafType struct {
	// Unique identifier for the leaf
	// required: true
	// unique: true
	Id *string `json:"id"`

	// An optional description for the leaf
	// required: false
	Description string `json:"description"`

	// The weight in the resolution of the constraint
	// requiered: true
	Weight int `json:"weight"`

	// The list of attributes defined in the data management section to match. All of them must comply.
	// requiered: true
	Attributes []string `json:"attributes"`
}

// TreeStructureType is a tree structure with a root and subtrees or leaves
// swagger:model
type TreeStructureType struct {

	// The operation to apply to the subtree or leaves
	// required: true
	// pattern: AND|OR
	// example: AND
	Type *string `json:"type"`

	// The subtrees pending from this node
	// required: false
	Children []TreeStructureType `json:"children"`

	// The leaves pending from this node
	// required: false
	Leaves []LeafType `json:"leaves"`
}

// GoalTreeType defines a goal tree
// swagger:model
type GoalTreeType struct {

	// Goal tree for data utility properties
	// required: false
	DataUtility TreeStructureType `json:"dataUtility"`

	// Goal tree for security properties
	// required: false
	Security TreeStructureType `json:"security"`

	// Goal tree for privacy properties
	// required: false
	Privacy TreeStructureType `json:"privacy"`
}

// AbstractPropertiesMethodType defines a goal tree for a method
// swagger:model
type AbstractPropertiesMethodType struct {

	// The method identifier this goals apply to
	// required: true
	MethodId *string `json:"method_id"`

	// The goal tree for this method
	// required: true
	GoalTrees GoalTreeType `json:"goalTrees"`
}
