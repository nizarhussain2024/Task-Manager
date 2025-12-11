package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"time"
)

func exportTasksHandler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json"
	}

	store.mu.RLock()
	tasks := make([]*Task, 0, len(store.tasks))
	for _, task := range store.tasks {
		tasks = append(tasks, task)
	}
	store.mu.RUnlock()

	switch format {
	case "csv":
		exportCSV(w, tasks)
	case "json":
		exportJSON(w, tasks)
	default:
		http.Error(w, "Unsupported format. Use 'json' or 'csv'", http.StatusBadRequest)
	}
}

func exportJSON(w http.ResponseWriter, tasks []*Task) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=tasks_"+time.Now().Format("20060102")+".json")
	json.NewEncoder(w).Encode(tasks)
}

func exportCSV(w http.ResponseWriter, tasks []*Task) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=tasks_"+time.Now().Format("20060102")+".csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Description", "Status", "Priority", "Created At", "Updated At"})

	for _, task := range tasks {
		writer.Write([]string{
			task.ID,
			task.Title,
			task.Description,
			task.Status,
			task.Priority,
			task.CreatedAt.Format(time.RFC3339),
			task.UpdatedAt.Format(time.RFC3339),
		})
	}
}

