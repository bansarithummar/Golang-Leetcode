func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    if len(preorder) == 0 {
       return nil
   }

   root := &TreeNode{Val: preorder[0]}

   if len(preorder) == 1 {
       return root
   }
   
   leftVal := preorder[1]
   leftIdx := 0
   for i := 0; i < len(postorder); i++ {
       if postorder[i] == leftVal {
           leftIdx = i + 1 
           break
       }
   }

   root.Left = constructFromPrePost(
       preorder[1:leftIdx+1],
       postorder[:leftIdx],
   )
   
   root.Right = constructFromPrePost(
       preorder[leftIdx+1:],
       postorder[leftIdx:len(postorder)-1],
   )
   
   return root
}
