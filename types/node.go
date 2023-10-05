package types

// TreeNode represents a binary tree node.
type TreeNode struct {
	Value     int       `json:"value"`
	ID        string    `json:"id"`
	Left      *TreeNode `json:"-"`
	LeftNode  string    `json:"left"`
	Right     *TreeNode `json:"-"`
	RightNode string    `json:"right"`
}
