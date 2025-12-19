package main

import (
	"strings"
	"time"
)

func filterTasksByStatus(tasks []*Task, status string) []*Task {
	if status == "" {
		return tasks
	}
	filtered := []*Task{}
	for _, task := range tasks {
		if strings.EqualFold(task.Status, status) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func filterTasksByPriority(tasks []*Task, priority string) []*Task {
	if priority == "" {
		return tasks
	}
	filtered := []*Task{}
	for _, task := range tasks {
		if strings.EqualFold(task.Priority, priority) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func filterTasksByDueDate(tasks []*Task, filterType string) []*Task {
	now := time.Now()
	filtered := []*Task{}

	for _, task := range tasks {
		if task.DueDate == nil {
			continue
		}

		switch filterType {
		case "overdue":
			if task.DueDate.Before(now) && task.Status != "completed" {
				filtered = append(filtered, task)
			}
		case "today":
			dueDate := task.DueDate.Truncate(24 * time.Hour)
			today := now.Truncate(24 * time.Hour)
			if dueDate.Equal(today) {
				filtered = append(filtered, task)
			}
		case "upcoming":
			if task.DueDate.After(now) {
				filtered = append(filtered, task)
			}
		}
	}
	return filtered
}


