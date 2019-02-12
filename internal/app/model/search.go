package model

// Search type is a struct for search queries.
type Search struct {
	Query string `json:"query" validate:"required"`
	Type  int    `json:"type"`
}

// SearchRecommendation search for recommendations.
// o = offset | l = limit | s = search term | t = type
func (db *Conn) SearchRecommendation(
	o, l float64, s string,
) (*RecommendationResult, error) {
	stmt, err := db.Prepare(`
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
		r.updated_at
		FROM recommendations AS r
		JOIN keyword_recommendation AS kr ON kr.recommendation_id = r.id
		JOIN genre_recommendation AS gr ON gr.recommendation_id = r.id
		JOIN genres AS g ON g.id = gr.genre_id
		JOIN keywords AS k ON k.id = kr.keyword_id
		WHERE r.title LIKE ?
		OR k.name LIKE ?
		OR g.name LIKE ?
		ORDER BY r.id DESC
		LIMIT ?,?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+s+"%", "%"+s+"%", "%"+s+"%", o, l)
	if err != nil {
		return nil, err
	}
	res := RecommendationResult{}
	for rows.Next() {
		rec := Recommendation{}
		err = rows.Scan(
			&rec.ID,
			&rec.UserID,
			&rec.Title,
			&rec.Type,
			&rec.Body,
			&rec.Poster,
			&rec.Backdrop,
			&rec.Status,
			&rec.CreatedAt,
			&rec.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &rec)
	}
	return &res, nil
}

// SearchUser search for users.
func (db *Conn) SearchUser(s string) (*UserResult, error) {
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
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+s+"%", 20)
	if err != nil {
		return nil, err
	}
	res := UserResult{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(
			&u.ID, &u.Name, &u.Email, &u.Password,
			&u.APIToken, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &u)
	}
	return &res, nil
}

// SearchGenre search for genres.
func (db *Conn) SearchGenre(s string) (*GenreResult, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+s+"%", 20)
	if err != nil {
		return nil, err
	}
	res := GenreResult{}
	for rows.Next() {
		g := Genre{}
		err = rows.Scan(
			&g.ID, &g.Name, &g.CreatedAt, &g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &g)
	}
	return &res, nil
}

// SearchKeyword search for keywords.
func (db *Conn) SearchKeyword(s string) (*KeywordResult, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+s+"%", 20)
	if err != nil {
		return nil, err
	}
	res := KeywordResult{}
	for rows.Next() {
		k := Keyword{}
		err = rows.Scan(
			&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &k)
	}
	return &res, nil
}

// SearchSource search for sources.
func (db *Conn) SearchSource(s string) (*SourceResult, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		WHERE name LIKE ?
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+s+"%", 20)
	if err != nil {
		return nil, err
	}
	res := SourceResult{}
	for rows.Next() {
		s := Source{}
		err = rows.Scan(
			&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &s)
	}
	return &res, nil
}

// GetSearchRecommendationTotalRows retrieves the total
// rows of recommendations table.
func (db *Conn) GetSearchRecommendationTotalRows(
	s string,
) (float64, error) {
	stmt, err := db.Prepare(`
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
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var total float64
	err = stmt.QueryRow("%"+s+"%", "%"+s+"%", "%"+s+"%").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}