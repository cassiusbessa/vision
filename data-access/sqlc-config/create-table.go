package data

import "context"

func CreateTable(db DBTX) {
	_, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY,
    project_id UUID,
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
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES accounts(id)
    );`)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
    parent_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES accounts(id)
    );`)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS reactions (
        id UUID PRIMARY KEY,
        post_id UUID REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
        comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
        user_id UUID NOT NULL,
        reaction_type VARCHAR(50) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES accounts(id)
        );`)
	if err != nil {
		panic(err.Error())
	}

}
