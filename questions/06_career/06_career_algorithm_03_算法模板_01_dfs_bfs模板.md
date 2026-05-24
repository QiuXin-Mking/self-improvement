# q
在网格问题中如何用 dx/dy 数组表示上、右、下、左四个方向？
# a
```cpp
const int dx[4] = { 0, 1, 0, -1};
const int dy[4] = {-1, 0, 1, 0};
```
使用 `new_x = x + dx[i]`, `new_y = y + dy[i]`，分别对应：上(0,-1)、右(1,0)、下(0,1)、左(-1,0)。

# q
BFS 实现中，如何处理相邻节点的越界与重复访问？
# a
通过边界检查和访问标记数组 `st`：
```cpp
if ( nex_x < 0 || nex_x >= n )  continue;
if ( nex_y < 0 || nex_y >= m )  continue;
if ( grid[nex_x][nex_y] == '1' && !st[nex_x][nex_y]){
    st[nex_x][nex_y] = true;
    q.push({nex_x,nex_y});
}
```

# q
BFS 遍历时使用什么数据结构存储坐标？
# a
使用 `queue<PII> q;`，其中 `typedef pair<int,int> PII;`，坐标以 `{x, y}` 形式入队。

# q
DFS 版本中如何标记已访问的陆地单元格？
# a
直接将当前单元格修改为 `'0'`：`grid[x][y] = '0';`，然后递归处理四个方向的相邻单元格，无需额外的 `st` 数组。

# q
`numIslands` 函数如何统计岛屿数量并触发遍历？
# a
1. 初始化 `st` 数组（BFS 版本用）或直接修改 `grid`（DFS 版本）。
2. 双重循环遍历每个单元格：若 `grid[i][j] != '0'` 且未被访问，则 `res++` 并调用 `bfs` 或 `dfs` 标记整个岛屿。
3. 最后返回 `res`。

