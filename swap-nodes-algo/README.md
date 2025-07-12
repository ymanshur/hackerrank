# [swap-nodes-algo](https://www.hackerrank.com/challenges/swap-nodes-algo/problem)

## NOTES

1. Swap operation all nodes at (h % k) == 0, h = current depth
2. Swap operation actually do by the roots at h-1           -> `swap(root, k, h)`
3. Needs tree structure to swap the children of a root node -> `build(indexes) root`
4. Needs to in-order traverse the tree in an array by       -> `traverse(root, *result)`
