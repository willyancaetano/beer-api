package beer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Beer, error) {

	var result []*Beer

	rows, err := s.DB.Query("select id, name, type, style from beer")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b Beer
		err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.Style)
		if err != nil {
			return nil, err
		}

		result = append(result, &b)
	}

	return result, nil
}

func (s *Service) Get(ID int64) (*Beer, error) {
	var b Beer

	stmt, err := s.DB.Prepare("select id, name, type, style from beer where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&b.ID, &b.ID, &b.Type, &b.Style)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(ID int64) error {
	return nil
}
