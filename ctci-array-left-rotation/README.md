# [ctci-array-left-rotation](https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem)

## NOTES

### Naive approach ~> O(n * r) time, O(1) space

a = {1 2 3 4 5}, n, r := len(a), d % n

a = {2 3 4 5 5} -> keep a[0] in tmp; change a[i] by a[i+1]; change the a[n-1] by tmp

a = {3 4 5 1 2}

a = {4 5 1 2 3} -> done with r iteration

### Better approach ~> O(n) time, O(n) space

1. let a2 with len(a) capacity
2. copy the a to a2,   from [r] to [n-1]
3. append the a to a2, from [0] to [r-1]

a = {1 2 3 4 5} -> a2 = {4 5}

a = {1 2 3}     -> a2 = {4 5} + {1 2 3}

### Expected approach ~> O(n) time, O(1) space

a = {1 2 3 4 5} -> reverse from a[0] to a[r-1]

a = {3 2 1 5 4} -> reverse from a[r] to a[n-1]

a = {4 5 1 2 3} -> reverse from a[0] to a[n-1]
