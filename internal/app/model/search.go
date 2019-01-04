package model

import (
	"database/sql"
	"log"
)

// Search type is a struct for seach queries.
type Search struct {
	Query string `json:"query" validate:"required"`
	Type  int    `json:"type"`
}

// SearchRecommendation search for recommendations.
func SearchRecommendation(s string, t int, db *sql.DB) (*ResultRecommendation, error) {

	stmt, err := db.Prepare(`
		SELECT 
		r.id, r.user_id, r.title, r.type,
		r.body, r.backdrop, r.poster, r.status,
		r.created_at, r.updated_at
		FROM recommendations AS r
		JOIN keyword_recommendation AS kr ON kr.recommendation_id = r.id
		JOIN genre_recommendation AS gr ON gr.recommendation_id = r.id
		JOIN genres AS g ON g.id = gr.genre_id
		JOIN keywords AS k ON k.id = kr.keyword_id
		WHERE r.title LIKE ? AND r.type = ?
		OR k.name LIKE ?
		OR g.name LIKE ?
		ORDER BY r.id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query("%"+s+"%", t, "%"+s+"%", "%"+s+"%", 20)

	if err != nil {
		log.Println(err)
	}

	res := ResultRecommendation{}

	for rows.Next() {
		rec := Recommendation{}

		err = rows.Scan(
			&rec.ID, &rec.UserID, &rec.Title, &rec.Type,
			&rec.Body, &rec.Backdrop, &rec.Poster, &rec.Status,
			&rec.CreatedAt, &rec.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &rec)

	}

	return &res, err

}

// SearchUser search for users.
func SearchUser(s string, db *sql.DB) (*ResultUser, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password, 
		api_token, created_at, updated_at
		FROM users
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query("%"+s+"%", 20)

	if err != nil {
		log.Println(err)
	}

	res := ResultUser{}

	for rows.Next() {
		u := User{}

		err = rows.Scan(
			&u.ID, &u.Name, &u.Email, &u.Password,
			&u.APIToken, &u.CreatedAt, &u.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &u)

	}

	return &res, err

}

// SearchGenre search for genres.
func SearchGenre(s string, db *sql.DB) (*ResultGenre, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query("%"+s+"%", 20)

	if err != nil {
		log.Println(err)
	}

	res := ResultGenre{}

	for rows.Next() {
		g := Genre{}

		err = rows.Scan(
			&g.ID, &g.Name, &g.CreatedAt, &g.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &g)

	}

	return &res, err

}

// SearchKeyword search for keywords.
func SearchKeyword(s string, db *sql.DB) (*ResultKeyword, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query("%"+s+"%", 20)

	if err != nil {
		log.Println(err)
	}

	res := ResultKeyword{}

	for rows.Next() {
		k := Keyword{}

		err = rows.Scan(
			&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &k)

	}

	return &res, err

}

// SearchSource search for sources.
func SearchSource(s string, db *sql.DB) (*ResultSource, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query("%"+s+"%", 20)

	if err != nil {
		log.Println(err)
	}

	res := ResultSource{}

	for rows.Next() {
		s := Source{}

		err = rows.Scan(
			&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &s)

	}

	return &res, err

}
