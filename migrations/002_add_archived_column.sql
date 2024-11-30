ALTER TABLE tickets ADD archived BOOLEAN NOT NULL DEFAULT FALSE;

---- create above / drop below ----

ALTER TABLE tickets DROP COLUMN archived;
