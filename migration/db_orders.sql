BEGIN;

CREATE TABLE IF NOT EXISTS orders(
    uuid varchar,
    data jsonb
);

END;

