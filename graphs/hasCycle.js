/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
 var canFinish = function(numCourses, prerequisites) {
    var graph = new Array(numCourses);
    for (var i = 0; i < numCourses; i++) graph[i] = [];
    var len = prerequisites.length;
    for (var i = 0; i < len; i++) {
        var item = prerequisites[i];
        graph[item[0]].push(item[1]);
    }
    
    // empty: other
    // 1: visiting
    // 2: visited
    let statuses = new Array(numCourses);
    const hasCycle = (course) => {
        if(statuses[course] == 1) return true;
        if(statuses[course] == 2) return false;
        statuses[course] = 1;
        for (let c of graph[course]) {
            if(hasCycle(c)) return true;
        }
        statuses[course] = 2;
    }
    
    for(let i = 0; i<numCourses; i++) {
        if(hasCycle(i)) {
            return false;
        }
    }
    return true;
};