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

// Sources Queries
const (
	querySourcesSelect     = "SELECT * FROM sources ORDER BY id DESC LIMIT ?"
	querySourcesSelectByID = "SELECT * FROM sources WHERE id = ?"
	querySourcesInsert     = "INSERT INTO sources (name, created_at, updated_at) VALUES (?, ?, ?)"
	querySourcesUpdate     = "UPDATE sources SET name = ?, updated_at = ? WHERE id = ?"
	querySourcesDelete     = "DELETE FROM sources WHERE id = ?"
)

// Users Queries
const (
	queryUsersSelect = `
		SELECT 
		id, 
		name, 
		email, 
		password,
		api_token, 
		created_at, 
		updated_at
		FROM users
		ORDER BY id DESC
		LIMIT ?
	`

	queryUserSelectByID = `
		SELECT 
		id, 
		name, 
		email, 
		password,
		api_token, 
		created_at, 
		updated_at
		FROM users
		WHERE id = ?
	`

	queryUserInsert = `
		INSERT INTO users (
		name, 
		email, 
		password,
		api_token, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	queryUserUpdate = `
		UPDATE users
		SET 
		name = ?, 
		email = ?, 
		password = ?,
		api_token = ?, 
		updated_at = ?
		WHERE id = ?
	`

	queryUserDelete = "DELETE FROM users WHERE id = ?"
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
		WHERE r.status = 1
		GROUP BY r.id
		ORDER BY r.id DESC
		LIMIT ?,?
	`

	queryRecommendationsAdminSelect = `
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
		GROUP BY r.id
		ORDER BY r.id DESC
		LIMIT ?
	`

	queryRecommendationSelectByID = `
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
		WHERE r.id = ?
		GROUP BY r.id
	`

	queryRecommendationInsert = `
	    INSERT INTO recommendations (
		user_id, 
		title, 
		type, 
		body, 
		poster, 
		backdrop, 
		status, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	queryRecommendationUpdate = `
		UPDATE recommendations
		SET 
		title = ?, 
		type = ?, 
		body = ?, 
		poster = ?,
		backdrop = ?, 
		status = ?, 
		updated_at = ?
		WHERE id = ?
	`

	queryRecommendationDelete = "DELETE FROM recommendations WHERE id = ?"

	queryRecommendationGenres = `
		SELECT 
		g.id, 
		g.name 
		FROM genres AS g
		JOIN genre_recommendation AS gr ON gr.genre_id = g.id
		JOIN recommendations AS r ON r.id = gr.recommendation_id
		WHERE r.id = ?
	`

	queryRecommendationKeywords = `
		SELECT 
		k.id, 
		k.name 
		FROM keywords AS k
		JOIN keyword_recommendation AS kr ON kr.keyword_id = k.id
		JOIN recommendations AS r ON r.id = kr.recommendation_id
		WHERE r.id = ?
	`

	queryRecommendationsTotalRows = "SELECT COUNT(*) FROM recommendations"
)

// Recommendation Items Queries
const (
	queryRecommendationItems = `
		SELECT 
		ri.id, 
		ri.recommendation_id, 
		ri.name, 
		ri.tmdb_id, 
		ri.year, 
		ri.overview, 
		ri.poster, 
		ri.backdrop, 
		ri.trailer, 
		ri.commentary, 
		ri.media_type,
		ri.created_at, 
		ri.updated_at,
		GROUP_CONCAT(DISTINCT s.name SEPARATOR ', ') AS sources
		FROM recommendation_items AS ri
		JOIN recommendation_item_source AS ris ON ris.recommendation_item_id = ri.id
		JOIN sources AS s ON s.id = ris.source_id
		WHERE ri.recommendation_id = ?
		GROUP BY ri.id
	`

	queryRecommendationItem = `
		SELECT 
		ri.id, 
		ri.recommendation_id, 
		ri.name, 
		ri.tmdb_id, 
		ri.year, 
		ri.overview, 
		ri.poster, 
		ri.backdrop, 
		ri.trailer, 
		ri.commentary, 
		ri.media_type,
		ri.created_at, 
		ri.updated_at,
		GROUP_CONCAT(DISTINCT s.name SEPARATOR ', ') AS sources
		FROM recommendation_items AS ri
		JOIN recommendation_item_source AS ris ON ris.recommendation_item_id = ri.id
		JOIN sources AS s ON s.id = ris.source_id
		WHERE ri.id = ?
		GROUP BY ri.id
	`

	queryRecommendationItemInsert = `
		INSERT INTO recommendation_items (
		recommendation_id, 
		name, 
		tmdb_id, 
		year, 
		overview, 
		poster, 
		backdrop, 
		trailer, 
		commentary, 
		media_type, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	queryRecommendationItemUpdate = `
		UPDATE recommendation_items
		SET 
		name = ?, 
		tmdb_id = ?, 
		year = ?, 
		overview = ?,
		poster=?, 
		backdrop = ?, 
		trailer = ?, 
		commentary = ?,
		media_type = ?, 
		updated_at = ?
		WHERE id = ?
	`

	queryRecommendationItemDelete = "DELETE FROM recommendation_items WHERE id = ?"

	queryRecommendationItemSources = `
		SELECT 
		s.id, 
		s.name 
		FROM sources AS s
		JOIN recommendation_item_source AS ris ON ris.source_id = s.id 	
		JOIN recommendation_items AS ri ON ri.id = ris.recommendation_item_id	
		WHERE ri.id = ?
	`

	queryRecommendationItemsTotalRows = `
		SELECT 
		COUNT(*) 
		FROM recommendation_items 
		WHERE recommendation_id = ?
	`
)

// Search Queries
const (
	querySearchRecommendations = `
		SELECT DISTINCT
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
		JOIN genre_recommendation AS gr ON gr.recommendation_id = r.id
		JOIN genres AS g ON g.id = gr.genre_id
		JOIN keywords AS k ON k.id = kr.keyword_id
		WHERE r.title LIKE ?
		OR k.name LIKE ?
		OR g.name LIKE ?
		GROUP BY r.id
		ORDER BY r.id DESC
		LIMIT ?,?
	`

	querySearchRecommendationsTotalRows = `
		SELECT 
		COUNT(DISTINCT r.id)
		FROM recommendations AS r
		JOIN keyword_recommendation AS kr ON kr.recommendation_id = r.id
		JOIN genre_recommendation AS gr ON gr.recommendation_id = r.id
		JOIN genres AS g ON g.id = gr.genre_id
		JOIN keywords AS k ON k.id = kr.keyword_id
		WHERE r.title LIKE ?
		OR k.name LIKE ?
		OR g.name LIKE ?
	`

	querySearchUsers = `
		SELECT 
		id, 
		name, 
		email, 
		password, 
		api_token, 
		created_at, 
		updated_at
		FROM users
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`

	querySearchGenres = `
		SELECT 
		id, 
		name, 
		created_at, 
		updated_at
		FROM genres
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`

	querySearchKeywords = `
		SELECT 
		id, 
		name, 
		created_at, 
		updated_at
		FROM keywords
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`

	querySearchSources = `
		SELECT 
		id, 
		name, 
		created_at, 
		updated_at
		FROM sources
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`
)
