package main

import (
	"encoding/json"
	"net/http"
)

type TaskStats struct {
	TotalTasks      int            `json:"total_tasks"`
	ByStatus        map[string]int `json:"by_status"`
	ByPriority      map[string]int `json:"by_priority"`
	CompletedTasks  int            `json:"completed_tasks"`
	PendingTasks    int            `json:"pending_tasks"`
	InProgressTasks int            `json:"in_progress_tasks"`
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	stats := TaskStats{
		ByStatus:   make(map[string]int),
		ByPriority: make(map[string]int),
	}

	for _, task := range store.tasks {
		stats.TotalTasks++
		stats.ByStatus[task.Status]++
		stats.ByPriority[task.Priority]++

		switch task.Status {
		case "completed":
			stats.CompletedTasks++
		case "pending":
			stats.PendingTasks++
		case "in-progress":
			stats.InProgressTasks++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}



