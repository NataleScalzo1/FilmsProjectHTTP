package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/natale/codermine/models"
)

type FilmModel struct {
	DB *sql.DB
}

func (m *FilmModel) Insert(titlename, content, expires string) (int, error) {

	stmt := `INSERT INTO film (titlename, content, created, expires)
VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, titlename, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *FilmModel) Get(id int) (*models.Film, error) {

	stmt := `SELECT id, titlename, content, created, expires FROM film
WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Film{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	filmObj := &models.Film{ID: s.ID, Title: s.Title, Content: s.Content, Created: s.Created, Expires: s.Expires}
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return filmObj, nil
}

func (m *FilmModel) GetName(titlename string) (*models.Film, error) {

	stmt := `SELECT id, titlename, content, created, expires FROM film
WHERE expires > UTC_TIMESTAMP() AND titlename = ?`

	row := m.DB.QueryRow(stmt, titlename)

	s := &models.Film{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	filmObj := &models.Film{ID: s.ID, Title: s.Title, Content: s.Content, Created: s.Created, Expires: s.Expires}

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return filmObj, err
}
func (m *FilmModel) GetContent(content string) ([]*models.Film, error) {

	rows, _ := m.DB.Query(`select * from film WHERE content = ?;`, content)

	films := []*models.Film{}

	for rows.Next() {
		s := new(models.Film)
		_ = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		films = append(films, s)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return films, nil
}

func (m *FilmModel) Delete(id int) (int64, error) {
	result, err := m.DB.Exec("delete from film where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (m *FilmModel) DeletebyName(titlename string) (int64, error) {

	result, err := m.DB.Exec("delete from film where titlename = ?", titlename)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
