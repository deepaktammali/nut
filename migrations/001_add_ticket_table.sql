CREATE TABLE tickets (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    -- could be 'open' | 'closed'
    status VARCHAR(32) NOT NULL,
    -- could be 'p1' | 'p2' | 'p3' | 'p4'
    priority VARCHAR(32) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT pk_tickets PRIMARY KEY ("id")
);

---- create above / drop below ----

DROP TABLE tickets;
