package database

const createTopicTableQuery = `
	CREATE TABLE topic (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"title" TEXT NOT NULL,
		"url" TEXT NOT NULL,
		"user" INTEGER NOT NULL,
		"createdAt" TEXT NOT NULL,
		"updatedAt" TEXT,
		FOREIGN KEY(user) REFERENCES user(id)
	);
`
const createPostTableQuery = `
	CREATE TABLE post (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"text" TEXT NOT NULL,
		"user" INTEGER NOT NULL,
		"topic" INTEGER NOT NULL,
		"createdAt" TEXT NOT NULL,
		"updatedAt" TEXT,
		FOREIGN KEY(user) REFERENCES user(id),
		FOREIGN KEY(topic) REFERENCES topic(id)
	);
`
