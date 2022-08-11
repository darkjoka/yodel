CREATE TABLE
    commentors(
        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
        user_id UUID NOT NULL,
        post_id UUID NOT NULL,
        increment_id SERIAL
    )