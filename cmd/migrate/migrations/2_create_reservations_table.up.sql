-- Create the reservations

CREATE TABLE reservations (
    id serial PRIMARY KEY,
    item_id          INTEGER REFERENCES items(id),
    reserved_by      VARCHAR(255), -- This could reference a user's table id if you have user accounts
    start_datetime   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    end_datetime     TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT reservation_timeframe_check CHECK (end_datetime > start_datetime)
);

