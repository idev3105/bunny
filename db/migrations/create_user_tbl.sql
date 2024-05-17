CREATE TABLE d_users (
    id BIGSERIAL PRIMARY KEY,
    user_id varchar(255) NOT NULL,
    username varchar(255),
    -- base fields
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(255),
    updated_by varchar(255),
    UNIQUE(user_id)
);