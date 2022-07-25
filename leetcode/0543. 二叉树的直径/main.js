var diameterOfBinaryTree = function (root) {
    let max = 0;
    function dfs(node) {
        if (!node) return 0;
        let left = dfs(node.left);
        let right = dfs(node.right);
        max = Math.max(max, left + right);
        return Math.max(left, right) + 1;
    }
    dfs(root);
    return max;
}

console.log(diameterOfBinaryTree([1,2,3,4,5,6,7]));