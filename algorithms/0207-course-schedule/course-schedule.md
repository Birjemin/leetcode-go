## 问题
There are a total of `numCourses` courses you have to take, labeled from 0 to `numCourses-1`.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: `[0,1]`

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

 

Example 1:
```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
             To take course 1 you should have finished course 0. So it is possible.
```

Example 2:
```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
```
Constraints:

- The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
- You may assume that there are no duplicate edges in the input prerequisites.
- `1 <= numCourses <= 10^5`

## 分析
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

![图的表示](https://www.cnblogs.com/liushang0419/archive/2011/05/06/2039386.html)

这道题目内容好多~

## 最高分
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites, numCourses)

	visited := make(map[int]bool)
	curVisited := make(map[int]bool)
	for start, _ := range graph {
		if res := dfs(start, graph, visited, curVisited); !res {
			return false
		}
	}
	return true
}

func dfs(start int, graph map[int][]int, visited map[int]bool, curVisited map[int]bool) bool {
	if val, ok := visited[start]; ok && val {
		return true
	}
	curVisited[start] = true

	neighbors := graph[start]
	for _, neighbor := range neighbors {
		if val, ok := curVisited[neighbor]; ok && val {
			return false
		}
		if res := dfs(neighbor, graph, visited, curVisited); !res {
			return false
		}
	}
	curVisited[start], visited[start] = false, true
	return true
}

// buildGraph returns a graph as {courseID: [requirements courseIDs]} map
// and returns start courses which aren't required by any courses.
func buildGraph(prerequisites [][]int, numCourses int) map[int][]int {
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		if _, ok := graph[p[1]]; ok {
			graph[p[1]] = append(graph[p[1]], p[0])
		} else {
			graph[p[1]] = []int{p[0]}
		}
	}
	return graph
}
```

## 实现
和142题有点像，找到环就说明无法完成，找不到就能够完成 dfs算法
(速度略慢)
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {
	for i, v := range prerequisites {
		t := make([][]int, len(prerequisites))
		copy(t, prerequisites)
		if !dfs(v[0], v[1], append(t[0:i], t[i+1:]...)) {
			return false
		}
	}
	return true
}

func dfs(start, end int, prerequisites [][]int) bool {
	// find the circle
	if start == end {
		return false
	}
	for i, v := range prerequisites {
		if end != v[0] {
			continue
		}
		if !dfs(start, v[1], append(prerequisites[0:i], prerequisites[i+1:]...)) {
			return false
		}
	}
	return true
}
```

## 改进
dfs优化，使用空间换取时间(已经遍历过的节点，就没必要在遍历了)
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {

	visited := make([]bool, numCourses)
	stop := make([]bool, numCourses)

	graph := grouping(numCourses, prerequisites)

	for i := range graph {
		// visited[i] = true
		if !dfs(i, graph, visited, stop) {
			return false
		}
		// visited[i] = false
	}

	return true
}

func dfs(start int, graph map[int][]int, visited, stop []bool) bool {
	// stop~
	if _, ok := graph[start]; !ok {
		stop[start] = true
		return true
	}

	for _, v := range graph[start] {
		if stop[v] {
			continue
		}
		// had visited
		if ok := visited[v]; ok {
			return false
		}
		visited[v] = true
		if !dfs(v, graph, visited, stop) {
			return false
		}
		visited[v] = false
	}
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
dfs还可以改进，判断是否有环即可(和上面类似，无非是反过来)
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {

	visited := make([]bool, numCourses)
	// end,have no cycle
	stop := make([]bool, numCourses)

	graph := grouping(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		if dfs(i, graph, visited, stop) {
			return false
		}
	}

	return true
}

func dfs(start int, graph map[int][]int, visited, stop []bool) bool {
	for _, v := range graph[start] {
		if stop[v] {
			continue
		}
		// had visited
		if visited[v] {
			return true
		}
		visited[v] = true
		// had cycle
		if dfs(v, graph, visited, stop) {
			return true
		}
		visited[v] = false
	}
	stop[start] = true
	return false
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
bfs算法计算，统计每个几点的入度，将为0的推入到queue中遍历，如果最终入度为0，则无环结束，反之，则相交
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := make([][]int, numCourses)
	count := make(map[int]int, numCourses)

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		count[pre[0]]++
		if _, ok := count[pre[1]]; !ok {
			count[pre[1]] = 0
		}
	}

	return bfs(graph, count)
}

func bfs(graph [][]int, count map[int]int) bool {
	// if len(count) == 0 {
	// 	return true
	// }
	var flag1, flag2 bool
	for i, v := range count {
		flag1 = true
		if v != 0 {
			continue
		}
		flag2 = true
		delete(count, i)
		for _, k := range graph[i] {
			if _, ok := count[k]; ok {
				count[k]--
			}
		}
		break
	}
	if !flag1{
		return true
	}
	if !flag2 {
		return false
	}
	return bfs(graph, count)
}
```

## 改进
bfs改进，分情况判断
```golang
func canFinish(numCourses int, prerequisites [][]int) bool {

	graph, queue := make([][]int, numCourses), make([]int, numCourses)

	// init
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		queue[pre[0]]++
	}

	return bfs(graph, queue, len(queue))
}

func bfs(graph [][]int, queue []int, count int) bool {
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

		queue[i]--
		count--
		return bfs(graph, queue, count)
	}

	return false
}

```

## 反思
这道题目废了很多时间，感觉需要不停的推敲~~

## 参考