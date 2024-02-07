-- WATCHES

-- name: GetTodo :one
SELECT * FROM todos
WHERE uuid = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY title;

-- name: CreateTodo :execresult
INSERT INTO todos (
  uuid, title, description
) VALUES (
  $1, $2, $3
);

-- name: UpdateTodo :execresult
UPDATE todos
SET title = $1,
description = $2
WHERE uuid = $3;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE uuid = $1;
