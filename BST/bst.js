let nodes;

const tags = [1, 2, 2, 3, 3, 4, 5];

class TreeNode {
  constructor(val, left, right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

nodes = new TreeNode(tags[0], null, null);

for (let i = 1, l = tags.length; i < l; i += 1) {
  const node = new TreeNode(tags[l], null, null);
  nodes.left = node;
}

const makeNode = node => {};
