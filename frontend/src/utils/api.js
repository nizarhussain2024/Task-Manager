const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

export async function fetchTasks() {
  const response = await fetch(`${API_URL}/tasks`)
  if (!response.ok) {
    throw new Error('Failed to fetch tasks')
  }
  return response.json()
}

export async function createTask(task) {
  const response = await fetch(`${API_URL}/tasks`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(task)
  })
  if (!response.ok) {
    throw new Error('Failed to create task')
  }
  return response.json()
}

export async function updateTask(id, updates) {
  const response = await fetch(`${API_URL}/tasks/${id}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(updates)
  })
  if (!response.ok) {
    throw new Error('Failed to update task')
  }
  return response.json()
}

export async function deleteTask(id) {
  const response = await fetch(`${API_URL}/tasks/${id}`, {
    method: 'DELETE'
  })
  if (!response.ok) {
    throw new Error('Failed to delete task')
  }
}



