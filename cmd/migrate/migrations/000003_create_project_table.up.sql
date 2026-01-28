CREATE TABLE IF NOT EXISTS projects (
    project_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    tech_stack VARCHAR(255),
    team_members INT[],
    created_by INT NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    FOREIGN KEY (created_by) REFERENCES user(id)
);