package graphs


func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for i:=0; i<len(prerequisites); i++ {
        item:= prerequisites[i]
        graph[item[0]] = append(graph[item[0]], item[1])
    }

    statuses := make([]int, numCourses)
    
    for i := 0; i<numCourses; i++ {
        if hasCycle(i, &statuses, &graph) {
            return false
        }
    }
    
    return true
}

func hasCycle(course int, statuses *[]int, graph *[][]int) bool {
    if (*statuses)[course] == 1 {
        return true
    }
    if (*statuses)[course] == 2 {
        return false
    }
    (*statuses)[course] = 1
    for _, v := range (*graph)[course] {
        if hasCycle(v, statuses, graph) {
            return true
        }
    }
    (*statuses)[course] = 2
    return false
}