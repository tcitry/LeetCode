function ListNode(val, next) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
}

var addTwoNumbers = function (l1, l2) {
    let carry = 0;
    let head = new ListNode();
    let curr = head;
    while (l1 || l2) {
        let sum = (l1 ? l1.val : 0) + (l2 ? l2.val : 0) + carry;
        carry = Math.floor(sum / 10);
        curr.next = new ListNode(sum % 10);
        curr = curr.next;
        if (l1) l1 = l1.next;
        if (l2) l2 = l2.next;
    }
    if (carry) curr.next = new ListNode(carry);
    return head.next;
};

(function main() {
    let l1 = new ListNode(2, ListNode(4, ListNode(3)));
    let l2 = new ListNode(5, ListNode(6, ListNode(4)));
    let l3 = addTwoNumbers(l1, l2);
    console.log(l3.val);
})()
