package main

import (
	"strings"
)

type Tag struct {
	Name  string
	Color string
}

func addTagsToTask(task *Task, tagNames []string) {
	if task.Tags == nil {
		task.Tags = []string{}
	}
	
	for _, tagName := range tagNames {
		tagName = strings.TrimSpace(tagName)
		if tagName != "" {
			exists := false
			for _, existingTag := range task.Tags {
				if existingTag == tagName {
					exists = true
					break
				}
			}
			if !exists {
				task.Tags = append(task.Tags, tagName)
			}
		}
	}
}

func removeTagFromTask(task *Task, tagName string) {
	if task.Tags == nil {
		return
	}
	
	tags := []string{}
	for _, tag := range task.Tags {
		if tag != tagName {
			tags = append(tags, tag)
		}
	}
	task.Tags = tags
}

func filterTasksByTag(tasks []*Task, tagName string) []*Task {
	if tagName == "" {
		return tasks
	}
	
	filtered := []*Task{}
	for _, task := range tasks {
		if task.Tags != nil {
			for _, tag := range task.Tags {
				if strings.EqualFold(tag, tagName) {
					filtered = append(filtered, task)
					break
				}
			}
		}
	}
	return filtered
}


