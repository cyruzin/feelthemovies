package model

// Genres Queries

const queryGenresSelect = "SELECT * FROM genres ORDER BY id DESC LIMIT ?"

const queryGenreSelectByID = "SELECT * FROM genres WHERE id = ?"

const queryGenreInsert = "INSERT INTO genres (name, created_at, updated_at) VALUES (?, ?, ?)"

const queryGenreUpdate = "UPDATE genres SET name = ?, updated_at = ? WHERE id = ?"

const queryGenreDelete = "DELETE FROM genres WHERE id = ?"
