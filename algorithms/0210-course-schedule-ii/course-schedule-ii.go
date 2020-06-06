package course_schedule_ii

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

func findOrder1(numCourses int, prerequisites [][]int) []int {

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
