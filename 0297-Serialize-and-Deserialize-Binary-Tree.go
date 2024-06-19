/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() Codec {
    return Codec{}    
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "null"
    }
    var result []string
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        if node != nil {
            result = append(result, strconv.Itoa(node.Val))
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        } else {
            result = append(result, "null")
        }
    }   
    return strings.Join(result, ",")    
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "null" {
        return nil
    }
    values := strings.Split(data, ",")
    rootVal, _ := strconv.Atoi(values[0])
    root := &TreeNode{Val: rootVal}
    queue := []*TreeNode{root}
    index := 1
    for len(queue) > 0 && index < len(values) {
        node := queue[0]
        queue = queue[1:]
        if values[index] != "null" {
            leftVal, _ := strconv.Atoi(values[index])
            node.Left = &TreeNode{Val: leftVal}
            queue = append(queue, node.Left)
        }
        index++
        if index < len(values) && values[index] != "null" {
            rightVal, _ := strconv.Atoi(values[index])
            node.Right = &TreeNode{Val: rightVal}
            queue = append(queue, node.Right)
        }
        index++
    }  
    return root    
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
