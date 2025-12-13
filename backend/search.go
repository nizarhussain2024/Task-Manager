package main

import (
	"strings"
)

func searchTasks(query string) []*Task {
	store.mu.RLock()
	defer store.mu.RUnlock()

	query = strings.ToLower(query)
	results := []*Task{}

	for _, task := range store.tasks {
		if strings.Contains(strings.ToLower(task.Title), query) ||
			strings.Contains(strings.ToLower(task.Description), query) {
			results = append(results, task)
		}
	}

	return results
}

func filterTasksByStatus(status string) []*Task {
	store.mu.RLock()
	defer store.mu.RUnlock()

	results := []*Task{}
	for _, task := range store.tasks {
		if task.Status == status {
			results = append(results, task)
		}
	}
	return results
}

func filterTasksByPriority(priority string) []*Task {
	store.mu.RLock()
	defer store.mu.RUnlock()

	results := []*Task{}
	for _, task := range store.tasks {
		if task.Priority == priority {
			results = append(results, task)
		}
	}
	return results
}


