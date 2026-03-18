# Task Tracker CLI (Go)

A simple command-line task manager built with Go.

This project lets you create, view, update, mark, and delete tasks directly from your terminal. Tasks are saved to a local JSON file so your data persists between runs.

## Features

- Add a task with a title and optional status
- List all tasks in a table format (ID, title, status, created time, last updated time)
- Update a task title by ID
- Mark task status by ID
- Delete a task by ID
- Persist tasks in `tasks.json`
- Track `created_at` and `last_updated` timestamps for each task

## Project Structure

- `main.go`: CLI entrypoint and command router
- `cmd/`: command handlers (`add`, `list`, `update`, `mark`, `delete`)
- `models/task.go`: task model, file persistence, table display
- `tasks.json`: local task database file

## Requirements

- Go 1.26+

## Installation / Running

### Option 1: Run directly with Go

From the project root:

```bash
go run . <command> [flags]
```

### Option 2: Build a binary

```bash
go build -o tasktracker .
./tasktracker <command> [flags]
```

## Commands

### Add a task

Adds a new task.

- Required: `--title`
- Optional: `--status` (defaults to `Unassigned`)

```bash
go run . add --title "Buy groceries"
go run . add --title "Finish report" --status "To-Do"
```

### List tasks

Shows all tasks in a table with these columns:

- ID
- Title
- Status
- Created At
- Last Updated

```bash
go run . list
```

### Update a task title

Updates an existing task title by ID.

- Required: `--id`, `--title`

```bash
go run . update --id 1 --title "Buy groceries and snacks"
```

### Mark task status

Updates the status of an existing task by ID.

- Required: `--id`, `--status`
- Valid statuses: `To-Do`, `In-Progress`, `Done`
- The command normalizes casing before validation (for example, `done` becomes `Done`)

```bash
go run . mark --id 1 --status "Done"
go run . mark --id 2 --status "In-Progress"
```

### Delete a task

Deletes a task by ID.

- Required: `--id`

```bash
go run . delete --id 1
```

## Data Storage

- Tasks are stored in `tasks.json` in the project root.
- If `tasks.json` does not exist yet, the CLI starts with an empty task list.

Example task record:

```json
{
	"id": 1,
	"title": "Buy groceries",
	"status": "To-Do",
	"created_at": "2026-03-18T09:00:00Z",
	"last_updated": "2026-03-18T09:00:00Z"
}
```

Notes:

- `created_at` is set when a task is first created.
- `last_updated` is refreshed when you update a title or mark a new status.

## Error Behavior

The CLI exits with an error when:

- A required flag is missing
- An unknown command is used
- A task ID does not exist
- `mark` uses an invalid status value

## Quick Example Workflow

```bash
go run . add --title "Read Go docs"
go run . list
go run . mark --id 1 --status "Done"
go run . update --id 1 --title "Read Go docs chapter 1"
go run . delete --id 1
```

