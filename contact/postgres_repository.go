package contact

import (
	"fmt"
	"errors"
	"database/sql"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {

	return &PostgresRepository{
		db: db,
	}
}


func (repo *PostgresRepository)Migrate()error {
	query := `

		CREATE TABLE IF NOT EXISTS contact(

			id SERIAL PRIMARY KEY,
			name varchar(200) NOT NULL,
			email varchar(104) NOT NULL UNIQUE,
			phone varchar(24) UNIQUE,
			address varchar(125)
		);
	`

	_, err := repo.db.Exec(query)

	return err
}


func (repo *PostgresRepository) Create(contact ContactInfo) (*ContactInfo, error) {
	var contactResp ContactInfo

	stmt, err := repo.db.Prepare("INSERT INTO contact(name, email, phone, address) values($1, $2, $3, $4)")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := stmt.Exec(contact.Name, contact.Email, 
		contact.PhoneNumber, contact.HomeAddress)

	if err != nil {
		log.Println(err)
		return nil, err
	}		
	fmt.Println(res)

	fmt.Println(contactResp)
	return &contactResp, nil
}

func (repo *PostgresRepository) All() ([]ContactInfo, error) {
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

func (repo *PostgresRepository) GetByName(name string) (*ContactInfo, error) {
	row := repo.db.QueryRow("SELECT * FROM contact WHERE name=$1", name)

	var contact ContactInfo
	if err := row.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.PhoneNumber, &contact.HomeAddress); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}

		return nil, err
	}
	return &contact, nil
}

func (repo *PostgresRepository) Update(id int64, newContact ContactInfo) (*ContactInfo, error) {

	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}

	res, err := repo.db.Exec("UPDATE contact SET name=$1, email=$2, phone=$3, address=$4 WHERE id=$5",
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

func (repo *PostgresRepository) Delete(id int64) error {
	res, err := repo.db.Exec("DELETE FROM contact WHERE id=$1", id)
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

	return nil
}
