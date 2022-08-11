CREATE TABLE
    users (
        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(256) NOT NULL,
        karma BIGINT DEFAULT 0
    )