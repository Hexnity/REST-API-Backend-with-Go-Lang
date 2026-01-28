-- 1. Create the update function (run this once per database)
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 2. Create the Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL, -- Unique constraint handled below
    email VARCHAR(100) NOT NULL,    -- Unique constraint handled below
    email_verified BOOLEAN DEFAULT FALSE,
    remember_token VARCHAR(255),
    status VARCHAR(20) DEFAULT 'active',
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 3. Create Partial Unique Indexes (Allows re-use of email if old account is deleted)
CREATE UNIQUE INDEX idx_users_email_active ON users (email) WHERE (deleted_at IS NULL);
CREATE UNIQUE INDEX idx_users_username_active ON users (username) WHERE (deleted_at IS NULL);

-- 4. Attach the Update Trigger
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();