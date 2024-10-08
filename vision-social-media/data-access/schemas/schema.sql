CREATE TABLE posts (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL,
    author_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    repo_link VARCHAR(255),
    demo_link VARCHAR(255),
    post_image VARCHAR(255),
    like_count INT NOT NULL DEFAULT 0,
    comment_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (author_id) REFERENCES accounts(id)
);

CREATE TABLE reactions (
id UUID PRIMARY KEY,
post_id UUID REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
user_id UUID NOT NULL,
reaction_type VARCHAR(50) NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
FOREIGN KEY (user_id) REFERENCES accounts(id)
);

CREATE TABLE comments (
id UUID PRIMARY KEY,
post_id UUID REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
parent_id UUID REFERENCES comments(id) ON DELETE CASCADE,
user_id UUID NOT NULL,
content TEXT NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
FOREIGN KEY (user_id) REFERENCES accounts(id)
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    account_level VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    account_id UUID REFERENCES accounts(id) ON DELETE CASCADE NOT NULL,
    name VARCHAR(255) NOT NULL,
    title VARCHAR(255),
    description TEXT,
    image VARCHAR(255),
    link VARCHAR(255),
    star_project UUID,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (star_project) REFERENCES projects(id)
);

