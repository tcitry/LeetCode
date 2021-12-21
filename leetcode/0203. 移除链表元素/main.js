/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    if (head === null) {
        return head
    }
    curr = head.next
    pre = head
    while (curr !== null) {
        if (curr.val === val) {
            pre.next = curr.next
        } else {
            pre = pre.next
        }
        curr = curr.next
    }
    if (head.val === val) {
        head = head.next
    }
    return head
};