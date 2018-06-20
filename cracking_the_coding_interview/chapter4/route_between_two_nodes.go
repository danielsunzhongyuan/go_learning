package main

/*
检查一个图里的两个节点是否连通
输入参数：图 G，节点 Start，节点 End
输出：bool

每个节点都要有一个 state 属性，来表示它是否被读取过。state包括3个状态：未读取 Unvisited，已读取 Visited，正在读取 Visiting

1. 从 Start 开始，放入一个队列里，状态置为 Visiting
2. 从队列里取出第一个节点，如果为空，说明搜索结束，返回 false
3. 否则，遍历它的所有相邻节点，
4. 只有当相邻节点状态为 Unvisited 时，检查是否是 End，如果是返回 true
5. 如果不是，状态置为 Visiting，放入队列里
6. 把刚取出来的节点状态置为 Visited
7. 重复 2 ，直至找到 End 或者队列为空
*/
