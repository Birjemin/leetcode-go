package course_schedule

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
	for _, v := range graph[start] {
		if stop[v] {
			continue
		}
		// had visited
		if visited[v] {
			return false
		}
		visited[v] = true
		if !dfs(v, graph, visited, stop) {
			return false
		}
		visited[v] = false
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

func canFinish1(numCourses int, prerequisites [][]int) bool {

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
