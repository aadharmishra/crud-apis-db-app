package users

import (
	"context"
	"crud-apis-db-app/modules/users/models"
	"crud-apis-db-app/shared"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type UsersDal interface {
	IsUserAvailable(ctx context.Context, id int) (bool, error)
	InsertUser(ctx context.Context, user *models.User) error
	SelectUsers(ctx context.Context) ([]models.User, error)
	SelectUsersById(ctx context.Context, id int) (models.User, error)
	UpdateUserById(ctx context.Context, user *models.User) (models.User, error)
	DeleteUserById(ctx context.Context, id int) (bool, error)
}

type User struct {
	Deps *shared.Deps
}

func NewUsersDal(deps *shared.Deps) UsersDal {
	return &User{
		Deps: deps,
	}
}

func (u *User) IsUserAvailable(ctx context.Context, id int) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("empty id")
	}

	query := `SELECT * FROM users WHERE id = $1`

	rows, err := u.Deps.Database.PostgresDb.Read(ctx, query, id)
	if err != nil {
		fmt.Print("error", map[string]interface{}{"error": err})
		return false, err
	}

	if rows == nil || !rows.Next() {
		return false, fmt.Errorf("no records")
	}

	return true, nil
}

// InsertUser implements UsersDal.
func (u *User) InsertUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, username, email, age, isadmin, lastlogin, preferences) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	err := u.Deps.Database.PostgresDb.Create(ctx, query, user.ID, user.Username, user.Email, user.Age, user.IsAdmin, time.Now(), user.Preferences)
	if err != nil {
		return err
	}
	return nil
}

// SelectUsersById implements UsersDal.
func (u *User) SelectUsersById(ctx context.Context, id int) (models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	rows, err := u.Deps.Database.PostgresDb.Read(ctx, query, id)
	if err != nil {
		return models.User{}, err
	}

	if rows != nil {
		result, err := pgx.CollectOneRow[models.User](rows, pgx.RowToStructByName)
		if err != nil {
			return models.User{}, err
		}
		return result, nil
	}

	return models.User{}, fmt.Errorf("no records available")
}

// SelectUsers implements UsersDal.
func (u *User) SelectUsers(ctx context.Context) ([]models.User, error) {
	query := `SELECT * FROM users`

	rows, err := u.Deps.Database.PostgresDb.Read(ctx, query)
	if err != nil {
		return []models.User{}, err
	}

	if rows != nil {
		result, err := pgx.CollectRows[models.User](rows, pgx.RowToStructByName)
		if err != nil {
			return []models.User{}, err
		}
		return result, nil
	}

	return []models.User{}, fmt.Errorf("no records available")
}

// UpdateUserById implements UsersDal.
func (u *User) UpdateUserById(ctx context.Context, user *models.User) (models.User, error) {
	query := `UPDATE users SET username = $2, email = $3, age = $4, isadmin = $5, lastlogin = $6, preferences = $7 WHERE id = $1`
	err := u.Deps.Database.PostgresDb.Update(ctx, query, user.ID, user.Username, user.Email, user.Age, user.IsAdmin, time.Now(), user.Preferences)
	if err != nil {
		return models.User{}, err
	}

	updated, err := u.SelectUsersById(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return updated, nil
}

// DeleteUserById implements UsersDal.
func (u *User) DeleteUserById(ctx context.Context, id int) (bool, error) {
	query := `DELETE FROM users WHERE id = $1`
	err := u.Deps.Database.PostgresDb.Delete(ctx, query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
