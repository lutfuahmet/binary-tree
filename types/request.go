package types

type CalculationRequest struct {
	Tree Tree `json:"tree"`
}

// Tree represents a json tree payload.
type Tree struct {
	Nodes []TreeNode `json:"nodes"`
	Root  string     `json:"root"`
}

// Request2TreeNode converts request to tree of nodes
func Request2TreeNode(request CalculationRequest) *TreeNode {
	nodeMap := map[string]TreeNode{}
	// iterate all nodes
	for _, node := range request.Tree.Nodes {
		nodeMap[node.ID] = node
	}
	return GetNodeByID(request.Tree.Root, nodeMap)
}

// GetNodeByID get node by given id
func GetNodeByID(id string, nodeMap map[string]TreeNode) *TreeNode {
	node, exist := nodeMap[id]
	if !exist {
		return nil
	}
	if node.LeftNode != "" {
		node.Left = GetNodeByID(node.LeftNode, nodeMap)
	}
	if node.RightNode != "" {
		node.Right = GetNodeByID(node.RightNode, nodeMap)
	}
	return &node
}
