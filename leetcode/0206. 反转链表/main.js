var reverseList = function (head) {
  if (head === null || head.next === null) {
    return head;
  }
  let curr = head;
  let prev = null;
  while (curr != null) {
    next = curr.next;
    curr.next = prev;
    prev = curr;
    curr = next;
  }
  return prev;
};
