CREATE TYPE todo_status AS ENUM (
    'pending',
    'in_progress',
    'done',
    'archived'
);

CREATE TABLE todos (
    id        BIGSERIAL PRIMARY KEY,
    title     TEXT NOT NULL,
    status    todo_status NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_todos_status ON todos(status);
