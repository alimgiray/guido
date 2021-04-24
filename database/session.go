package database

const createSessionTableQuery = `
	CREATE TABLE session (
		"session_id" TEXT NOT NULL,
		"user" INTEGER NOT NULL,
		"createdAt" TEXT NOT NULL,
		FOREIGN KEY(user) REFERENCES user(id)
	);
`
