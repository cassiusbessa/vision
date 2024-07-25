package data

import "context"

func CreateTable(db DBTX) {
	_, err := db.Exec(context.Background(), `CREATE TABLE posts (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL,
    author_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    repo_link VARCHAR(255),
    demo_link VARCHAR(255),
    post_image VARCHAR(255),
    like_count INT DEFAULT 0,
    comment_count INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (author_id) REFERENCES accounts(id)
    );`)
	if err != nil {
		panic(err)
	}
}
