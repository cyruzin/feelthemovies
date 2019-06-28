package model

// Auth type is a struct for authentication.
type Auth struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	APIToken string `json:"api_token,omitempty"`
}

// AuthJWT type is a struct for JWT authentication.
type AuthJWT struct {
	Token string `json:"token"`
}

// CheckAPIToken checks if the given token exists.
func (c *Conn) CheckAPIToken(token string) (bool, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		api_token
		FROM users
		WHERE api_token = ?
`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var t string
	err = stmt.QueryRow(token).Scan(&t)
	if err != nil {
		return false, err
	}
	if t != "" && t == token {
		return true, nil
	}
	return false, nil
}

// Authenticate authenticates the current user and returns it's info.
func (c *Conn) Authenticate(email string) (string, error) {
	stmt, err := c.db.Prepare(`
		SELECT password
		FROM users
		WHERE email = ?
`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var password string
	err = stmt.QueryRow(email).Scan(
		&password,
	)
	if err != nil {
		return "", err
	}
	return password, nil
}

// GetAuthInfo retrieves info for the authenticated user.
func (c *Conn) GetAuthInfo(email string) (*Auth, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, name, email, api_token
		FROM users
		WHERE email = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	a := Auth{}
	err = stmt.QueryRow(email).Scan(
		&a.ID, &a.Name, &a.Email, &a.APIToken,
	)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
