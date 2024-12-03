CREATE TABLE TICKET_OPTIONS (
    id SERIAL PRIMARY KEY,
    
    event_id INT8 NOT NULL,
    option_id INT8 NOT NULL,
    total_capacity int8 NOT NULL,
    current_capacity int8 NOT NULL,
    price NUMERIC(100,2) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- ALTER TABLE orders ADD CONSTRAINT users_email_unique UNIQUE (email);
