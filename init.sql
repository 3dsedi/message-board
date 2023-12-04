CREATE DATABASE chaintraced;
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_name VARCHAR(50),
    email VARCHAR(50),
    password VARCHAR(50),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		content TEXT,
		parent_id INT REFERENCES messages(id) ON DELETE CASCADE,
		created_at TIMESTAMP
);

insert into users (user_name, email, password, created_at) VALUES ('sedi','sedi@chaintraced.com', '123456', NOW());
insert into users (user_name, email, password, created_at) VALUES ('user1','user1@chaintraced.com', '123456', NOW());



