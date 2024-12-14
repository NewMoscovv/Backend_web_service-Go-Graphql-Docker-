CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       comments_enabled BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          post_id INT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
                          text TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


