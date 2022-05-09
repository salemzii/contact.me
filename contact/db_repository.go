package contact

type Repository interface {
	Migrate() error
	Create(contact ContactInfo) (*ContactInfo, error)
	All() ([]ContactInfo, error)
	GetByName(name string)
	Update(id int64, updated ContactInfo) (*ContactInfo, error)
	Delete(id int64) error
}

//https://gosamples.dev/sqlite-intro/
//https://github.com/marketplace/actions/deploy-to-heroku
