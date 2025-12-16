package main

import (
	"encoding/json"
	"os"
	"sync"
)

type Persistence struct {
	filename string
	mu       sync.RWMutex
}

func NewPersistence(filename string) *Persistence {
	return &Persistence{filename: filename}
}

func (p *Persistence) Save(tasks map[string]*Task) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(p.filename, data, 0644)
}

func (p *Persistence) Load() (map[string]*Task, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	data, err := os.ReadFile(p.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]*Task), nil
		}
		return nil, err
	}

	var tasks map[string]*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p *Persistence) AutoSave(tasks map[string]*Task, interval int) {
	// In production, implement periodic auto-save
}


