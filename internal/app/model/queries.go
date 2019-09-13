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

// Recommendations Queries

const (
	queryRecommendationsSelect = `
		SELECT
		r.id, 
		r.user_id, 
		r.title, 
		r.type, 
		r.body, 
		r.poster, 
		r.backdrop, 
		r.status, 
		r.created_at, 
		r.updated_at,
		GROUP_CONCAT(DISTINCT g.name SEPARATOR ', ') AS genres,
		GROUP_CONCAT(DISTINCT k.name SEPARATOR ', ') AS keywords
		FROM recommendations AS r
		JOIN keyword_recommendation AS kr ON kr.recommendation_id = r.id
		JOIN keywords AS k ON k.id = kr.keyword_id
		JOIN genre_recommendation AS gr ON gr.recommendation_id = r.id
		JOIN genres AS g ON g.id = gr.genre_id	
		WHERE r.status = ?
		GROUP BY r.id
		ORDER BY r.id DESC
		LIMIT ?,?
	`
)
