package main

import (
	"sort"
	"time"
)

type TaskSorter struct {
	tasks []*Task
	by    func(t1, t2 *Task) bool
}

func (ts *TaskSorter) Len() int {
	return len(ts.tasks)
}

func (ts *TaskSorter) Swap(i, j int) {
	ts.tasks[i], ts.tasks[j] = ts.tasks[j], ts.tasks[i]
}

func (ts *TaskSorter) Less(i, j int) bool {
	return ts.by(ts.tasks[i], ts.tasks[j])
}

func sortTasksByCreated(tasks []*Task, ascending bool) []*Task {
	sorted := make([]*Task, len(tasks))
	copy(sorted, tasks)

	sorter := &TaskSorter{
		tasks: sorted,
		by: func(t1, t2 *Task) bool {
			if ascending {
				return t1.CreatedAt.Before(t2.CreatedAt)
			}
			return t1.CreatedAt.After(t2.CreatedAt)
		},
	}
	sort.Sort(sorter)
	return sorted
}

func sortTasksByPriority(tasks []*Task) []*Task {
	priorityOrder := map[string]int{
		"high":   3,
		"medium": 2,
		"low":    1,
	}

	sorted := make([]*Task, len(tasks))
	copy(sorted, tasks)

	sort.Slice(sorted, func(i, j int) bool {
		return priorityOrder[sorted[i].Priority] > priorityOrder[sorted[j].Priority]
	})
	return sorted
}

func sortTasksByStatus(tasks []*Task) []*Task {
	statusOrder := map[string]int{
		"pending":     1,
		"in-progress": 2,
		"completed":   3,
	}

	sorted := make([]*Task, len(tasks))
	copy(sorted, tasks)

	sort.Slice(sorted, func(i, j int) bool {
		return statusOrder[sorted[i].Status] < statusOrder[sorted[j].Status]
	})
	return sorted
}


