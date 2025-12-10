package main

import (
	"strings"
	"time"
)

func searchTasks(tasks []*Task, query string) []*Task {
	if query == "" {
		return tasks
	}
	
	query = strings.ToLower(query)
	var results []*Task
	
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), query) ||
			strings.Contains(strings.ToLower(task.Description), query) {
			results = append(results, task)
		}
	}
	
	return results
}

func sortTasks(tasks []*Task, sortBy string) []*Task {
	switch sortBy {
	case "created_at":
		// Sort by creation date (newest first)
		for i := 0; i < len(tasks)-1; i++ {
			for j := i + 1; j < len(tasks); j++ {
				if tasks[i].CreatedAt.Before(tasks[j].CreatedAt) {
					tasks[i], tasks[j] = tasks[j], tasks[i]
				}
			}
		}
	case "priority":
		priorityOrder := map[string]int{"high": 3, "medium": 2, "low": 1}
		for i := 0; i < len(tasks)-1; i++ {
			for j := i + 1; j < len(tasks); j++ {
				if priorityOrder[tasks[i].Priority] < priorityOrder[tasks[j].Priority] {
					tasks[i], tasks[j] = tasks[j], tasks[i]
				}
			}
		}
	}
	return tasks
}

