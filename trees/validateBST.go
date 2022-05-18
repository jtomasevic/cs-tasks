package trees

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool{
	return validate(root, nil, nil);
}


func validate(node *TreeNode, min *int, max *int) bool{
    if node == nil {
		return true
	}
    if max!=nil && (*max) <= node.Val  {
		return false
	}
    if min!=nil && (*min) >= node.Val { 
		return false;
	}
    
    if !validate(node.Left, min, &node.Val ) {
		return false
	}
    if !validate(node.Right, &node.Val, max) { 
		return false
	}
    
    return true;
}