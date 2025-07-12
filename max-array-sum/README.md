# [max-array-sum](https://www.hackerrank.com/challenges/max-array-sum/problem)

## NOTES

```
arr = {3 7 4 6 5}
ans = {0 0}
idx =      i      -> ans[2] = 4 + 3
idx =        i    -> ans[3] = 6 + 7
idx =      j   i  -> ans[4] =  max(max(5 + 4, ans[2]),
idx =    j     i  ->               max(5 + 7, ans[1]),
idx =  j       i  ->               max(5 + 3, ans[0]))

arr = {3 7 4 6  5 }
ans = {0 0 7 13 12}
```
