[堆和堆的应用](http://blog.jobbole.com/113552/)

堆的操作：
1. insert()
2. extractMin()
3. peek(findMin)
4. delete(i) // i是index
5. siftUp(i)
6. siftDown(i)

小顶堆是降序排列，大顶堆是升序排列。

堆的应用：
1. 优先队列（通常用堆来实现优先队列）
    * Dijkstra's algorithm（单源最短路问题中需要在邻接表中找到某一点的最短邻接边，这可以将复杂度降低）
    * Huffman coding（贪心算法的一个典型例子，采用优先队列构建最优的前缀编码树(prefixEncodeTree)）
    * Prim's algorithm for minimum spanning tree
    * Best-first search algorithms
2. 堆排序
3. 海量数据里（一亿以上）找到 TopK（一万以下）的数集合