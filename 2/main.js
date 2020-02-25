/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
function addTwoNumbers(l1, l2) {
  const before = new ListNode();
  // pointerを渡している
  let tail = before;
  let c = 0;

  while (l1 || l2 || c) {
    const v1 = l1 ? l1.val : 0;
    const v2 = l2 ? l2.val : 0;
    const v = v1 + v2 + c;

    tail.next = new ListNode(v % 10);
    tail = tail.next;
    c = v >= 10 ? 1 : 0;
    l1 = l1 && l1.next;
    l2 = l2 && l2.next;
  }

  return before.next;
}

function ListNode(val) {
  this.val = val;
  this.next = null;
}

const arg = new ListNode(1);
arg.next = new ListNode(2);
arg.next.next = new ListNode(3);

console.log(addTwoNumbers(arg, arg));

// 参照渡しについて

let u = { a: 1, b: 2 };
let m = u;
m["a"] = 2;
m.b = 4;

console.log(u, m);

var batman = 7;
var superman = { a: batman }; //assign-by-value
superman.a++;
console.log(batman); //7
console.log(superman); //8
