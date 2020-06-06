## 问题
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.

Example 1:
```
Input: 2, [[1,0]] 
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished   
             course 0. So the correct course order is [0,1] .
```

Example 2:
```
Input: 4, [[1,0],[2,0],[3,1],[3,2]]
Output: [0,1,2,3] or [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both     
             courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0. 
             So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3] .
```

Note:

- The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
- You may assume that there are no duplicate edges in the input prerequisites.

## 分析
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

和207题一样，相同方式处理就好了

## 最高分
```golang

```

## 实现
```golang
func findOrder(numCourses int, prerequisites [][]int) []int {
	visited, stop := make([]bool, numCourses), make([]bool, numCourses)

	graph := grouping(numCourses, prerequisites)

	var ret []int
	for i := range graph {
		if !dfs(i, graph, visited, stop, &ret) {
			return []int{}
		}
	}

	for i, v := range stop {
		if !v {
			ret = append(ret, i)
		}
	}

	return ret
}

func dfs(start int, graph map[int][]int, visited, stop []bool, ret *[]int) bool {
	// stop~
	for _, v := range graph[start] {
		if stop[v] {
			continue
		}
		// had visited
		if visited[v] {
			return false
		}
		visited[v] = true
		if !dfs(v, graph, visited, stop, ret) {
			return false
		}
		visited[v] = false
	}
	if !stop[start] {
		*ret = append(*ret, start)
	}
	stop[start] = true
	return true
}

func grouping(numCourses int, prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}
	return graph
}
```

## 改进
```golang
func findOrder(numCourses int, prerequisites [][]int) []int {

	graph, queue := make([][]int, numCourses), make([]int, numCourses)
	var ret []int
	// init
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		queue[pre[0]]++
	}
	bfs(graph, queue, len(queue), &ret)

	return ret
}

func bfs(graph [][]int, queue []int, count int, ret *[]int) bool {
	// end
	if count == 0 {
		return true
	}

	for i := range queue {
		// just jump to the next circle
		if queue[i] != 0 {
			continue
		}
		for _, k := range graph[i] {
			queue[k]--
		}
		*ret = append(*ret, i)

		queue[i]--
		count--

		return bfs(graph, queue, count, ret)
	}
	*ret = []int{}
	return false
}
```

## 反思

## 参考