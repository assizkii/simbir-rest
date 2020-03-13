package storages

import (
	"errors"
	"fmt"
	"github.com/assizkii/simbir-rest/internal/domain/interfaces"
	"github.com/assizkii/simbir-rest/internal/domain/usecases"
	"github.com/assizkii/simbir-rest/internal/entities"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type PgStorage struct {
	connection *sqlx.DB
}

type UsersDb struct {
	Id       int
	Login    string
	Password string
}

func (pg PgStorage) Validate(user entities.User) error {
	switch "" {
	case strings.TrimSpace(user.Login):
		return errors.New("login field cannot be empty string")
	case strings.TrimSpace(user.Password):
		return errors.New("password field cannot be empty string")
	}
	return nil
}

func (pg PgStorage) Get(login string) (*entities.User, error) {
	var user entities.User
	var userRow UsersDb
	query := `select *	from users where login=$1`
	err := pg.connection.Get(&userRow, query, login)
	if err != nil {
		return nil, err
	}

	user = entities.User{
		Login:    userRow.Login,
		Password: userRow.Password,
	}

	return &user, nil
}

func (pg PgStorage) Add(user entities.User) error {
	if err := pg.Validate(user); err != nil {
		return fmt.Errorf("validate error: %s", err)
	}

	currentUser, _ := pg.Get(user.Login)
	if currentUser != nil {
		return fmt.Errorf("user already exist")
	}

	query := `insert into users(login, password)
				 values($1, $2) RETURNING id`

	var id int
	err := pg.connection.QueryRow(query, user.Login, usecases.HashPassword(user.Password)).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func NewPgStorage(dsn string) interfaces.AppStorage {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	log.Printf(dsn)
	err = Migrate(db)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	return &PgStorage{db}
}

// sql migrations from migrations folder
func Migrate(db *sqlx.DB) error {
	err := filepath.Walk("./migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			log.Println(path)
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fileValue := string(fileData)
			_, err = db.Exec(fileValue)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
