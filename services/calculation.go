package services

import (
	"binary-tree/types"
	"math"
)

// Helper function to find the maximum path sum and update the result.
func maxPathSumHelper(root *types.TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	// Recursively find the maximum path sum in the left and right children.
	leftMax := maxPathSumHelper(root.Left, result)
	rightMax := maxPathSumHelper(root.Right, result)

	// bypass negative children
	currentMax := root.Value
	if leftMax > 0 {
		currentMax += leftMax
	}
	if rightMax > 0 {
		currentMax += rightMax
	}

	// Update the result with the maximum path sum found so far.
	*result = int(math.Max(float64(*result), float64(currentMax)))

	// Return the maximum path sum that can continue from this node to its parent.
	return int(math.Max(float64(root.Value), float64(root.Value+int(math.Max(float64(leftMax), float64(rightMax))))))
}

// MaxPathSum calculates the maximum path sum
func MaxPathSum(req types.CalculationRequest) int {
	root := types.Request2TreeNode(req)
	if root == nil {
		return 0
	}

	result := math.MinInt32 // Initialize with the smallest possible integer.
	maxPathSumHelper(root, &result)
	return result
}
