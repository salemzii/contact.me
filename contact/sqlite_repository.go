package contact

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {

	query := `
		CREATE TABLE IF NOT EXISTS contact(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			phone TEXT NOT NULL UNIQUE,
			address TEXT
		);
	`
	_, err := repo.db.Exec(query)
	return err
}

func (repo *SQLiteRepository) Create(contact ContactInfo) (*ContactInfo, error) {
	res, err := repo.db.Exec("INSERT INTO contact(name, email, phone, address) values(?,?,?, ?)",
		contact.Name, contact.Email, contact.PhoneNumber, contact.HomeAddress)

	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	contact.ID = id
	return &contact, nil
}

func (repo *SQLiteRepository) All() ([]ContactInfo, error) {
	rows, err := repo.db.Query("SELECT * FROM contact")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []ContactInfo

	for rows.Next() {
		var contact ContactInfo
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email,
			&contact.PhoneNumber, &contact.HomeAddress); err != nil {
			return nil, err
		}
		all = append(all, contact)
	}
	return all, nil
}

func (repo *SQLiteRepository) GetByName(name string) (*ContactInfo, error) {
	row := repo.db.QueryRow("SELECT * FROM contact WHERE name = ?", name)

	var contact ContactInfo
	if err := row.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.PhoneNumber, &contact.HomeAddress); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}

		return nil, err
	}
	return &contact, nil
}

func (repo *SQLiteRepository) Update(id int64, newContact ContactInfo) (*ContactInfo, error) {

	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}

	res, err := repo.db.Exec("UPDATE contact SET name = ?, email = ?, phone = ?, address = ? WHERE id = ? ",
		newContact.Name, newContact.Email, newContact.PhoneNumber, newContact.HomeAddress, id)

	if err != nil {
		return nil, err
	}

	rowsaffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsaffected == 0 {
		return nil, ErrUpdateFailed
	}
	return &newContact, nil
}

func (repo *SQLiteRepository) Delete(id int64) error {
	res, err := repo.db.Exec("DELETE FROM websites WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
