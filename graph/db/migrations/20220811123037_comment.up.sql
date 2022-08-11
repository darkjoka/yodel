CREATE TABLE
    comments(
        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
        user_id UUID NOT NULL,
        post_id UUID NOT NULL,
        content VARCHAR(250) NOT NULL,
        is_flagged BOOLEAN DEFAULT false,
        number_of_votes SMALLINT DEFAULT 0,
        location VARCHAR(250) NOT NULL,
        longitude FLOAT NOT NULL,
        latitude FLOAT NOT NULL,
        created_at TIMESTAMP DEFAULT current_timestamp
    )