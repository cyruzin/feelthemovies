package model

// Auth Queries
const (
	queryAuthAuthenticate = "SELECT password from users WHERE email = ?"
	queryAuthGetInfo      = "SELECT id, name, email FROM users WHERE email = ?"
)

// Genres Queries
const (
	queryGenresSelect    = "SELECT * FROM genres ORDER BY id DESC LIMIT ?"
	queryGenreSelectByID = "SELECT * FROM genres WHERE id = ?"
	queryGenreInsert     = "INSERT INTO genres (name, created_at, updated_at) VALUES (?, ?, ?)"
	queryGenreUpdate     = "UPDATE genres SET name = ?, updated_at = ? WHERE id = ?"
	queryGenreDelete     = "DELETE FROM genres WHERE id = ?"
)

// Keywords Queries
const (
	queryKeywordsSelect    = "SELECT * FROM keywords ORDER BY id DESC LIMIT ?"
	queryKeywordSelectByID = "SELECT * FROM keywords WHERE id = ?"
	queryKeywordInsert     = "INSERT INTO keywords (name, created_at, updated_at) VALUES (?, ?, ?)"
	queryKeywordUpdate     = "UPDATE keywords SET name = ?, updated_at = ? WHERE id = ?"
	queryKeywordDelete     = "DELETE FROM keywords WHERE id = ?"
)
