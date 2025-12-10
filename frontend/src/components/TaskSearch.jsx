export default function TaskSearch({ searchQuery, onSearchChange }) {
  return (
    <div className="task-search">
      <input
        type="text"
        placeholder="Search tasks..."
        value={searchQuery}
        onChange={(e) => onSearchChange(e.target.value)}
        className="search-input"
      />
    </div>
  )
}

