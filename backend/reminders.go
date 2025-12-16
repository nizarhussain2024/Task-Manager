package main

import (
	"time"
)

type Reminder struct {
	TaskID    string
	Message   string
	RemindAt  time.Time
	Notified  bool
}

var reminders = make(map[string]*Reminder)

func checkReminders() {
	now := time.Now()
	store.mu.RLock()
	defer store.mu.RUnlock()

	for _, task := range store.tasks {
		if task.DueDate != nil && !task.DueDate.After(now.Add(1*time.Hour)) && task.Status != "completed" {
			reminderID := task.ID + "_due"
			if _, exists := reminders[reminderID]; !exists {
				reminders[reminderID] = &Reminder{
					TaskID:   task.ID,
					Message:  "Task '" + task.Title + "' is due soon",
					RemindAt: *task.DueDate,
				}
			}
		}
	}
}

func getPendingReminders() []*Reminder {
	pending := []*Reminder{}
	for _, reminder := range reminders {
		if !reminder.Notified && time.Now().After(reminder.RemindAt) {
			pending = append(pending, reminder)
		}
	}
	return pending
}

