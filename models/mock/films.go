package mock

import (
	"github.com/natale/codermine/models"
	"time"
)

var mockFilm = &models.Film{
	ID:      1,
	Title:   "Ciao",
	Content: "CiaoCiao",
	Created: time.Now(),
	Expires: "7",
}

type FilmModel struct{}

func (m *FilmModel) Insert(titlename, content, expires string) (int, error) {

	return 2, nil

}

func (m *FilmModel) Get(id int) (*models.Film, error) {

	switch id {
	case 1:
		return mockFilm, nil
	default:
		return nil, models.ErrNoRecord
	}

}

func (m *FilmModel) GetName(titlename string) (*models.Film, error) {

	return mockFilm, nil
}
func (m *FilmModel) Delete(id int) (int64, error) {
	return int64(2), nil
}

func (m *FilmModel) DeletebyName(titlename string) (int64, error) {
	return int64(2), nil
}
func (m *FilmModel) GetContent(string) ([]*models.Film, error) {

	return nil, nil
}
