/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
        carry := 0
        dummy_head := new(ListNode)
        tail := dummy_head

        for l1 != nil ||  l2 != nil || carry > 0 {
            sum := carry
            if (l1 != nil) {
                sum += l1.Val
                l1 = l1.Next
            }
            if (l2 != nil) {
                sum += l2.Val
                l2 = l2.Next
            }
            carry = sum / 10
            tail.Next = &ListNode{Val:sum % 10}
            tail = tail.Next
        }
        return dummy_head.Next    
}