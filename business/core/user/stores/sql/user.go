package sql

import (
	"bytes"
	"context"
	"fmt"
	"net/mail"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/filter"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store provides functions to execute user queries.
type Store struct {
	db     *sqlx.DB
	logger *zap.Logger
}

// NewStore creates a new User Store.
func NewStore(db *sqlx.DB, logger *zap.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) Create(ctx context.Context, user user.User) error {
	dbUser := toDBUser(user)
	query := `
	insert into users 
		(id, name, email, password_hash, roles, enabled, created_at, updated_at)
	values 
		(:id, :name, :email, :password_hash, :roles, :enabled, :created_at, :updated_at)
	`
	if _, err := s.db.NamedExecContext(ctx, query, dbUser); err != nil {
		return fmt.Errorf("db.NamedExecContext[%s]: %w", query, err)
	}
	return nil
}

func (s *Store) Update(ctx context.Context, uu user.UpdateUser) error {
	dbUser, err := toDBUpdateUser(uu)
	if err != nil {
		return fmt.Errorf("toDBUpdateUser: %w", err)
	}

	buf := bytes.Buffer{}
	buf.WriteString(`update users set `)
	buf.WriteString(`"updated_at" = :updated_at`)
	if dbUser.Name != nil {
		buf.WriteString(`",name" = :name`)
	}
	if dbUser.Roles != "" {
		buf.WriteString(`",roles" = :roles`)
	}
	if dbUser.Password != nil {
		buf.WriteString(`",password_hash" = :password`)
	}
	if dbUser.Enabled != nil {
		buf.WriteString(`",enabled" = :enabled`)
	}
	buf.WriteString(` where "id" = :id`)

	query := buf.String()
	if _, err := s.db.NamedExecContext(ctx, query, dbUser); err != nil {
		return fmt.Errorf("db.NamedExecContext[%s]: %w", query, err)
	}

	return nil
}

func (s *Store) ByID(ctx context.Context, userID string) (user.User, error) {
	query := `
	select 
		id, name, email, password_hash, roles, enabled, created_at, updated_at 
	from 
		users 
	where 
		id = :id`

	dbUser := dbUser{}
	if err := s.db.GetContext(ctx, &dbUser, query, map[string]string{"id": userID}); err != nil {
		return user.User{}, fmt.Errorf("db.GetContext[%s]: %w", query, err)
	}

	coreUser, err := toCoreUser(dbUser)
	if err != nil {
		return user.User{}, fmt.Errorf("toCoreUser[%v]: %w", dbUser, err)
	}
	return coreUser, nil
}

func (s *Store) ByIDs(ctx context.Context, ids []string) ([]user.User, error) {
	return []user.User{}, nil
}

func (s *Store) ByEmailNPassword(ctx context.Context, email mail.Address, passwordHash string) (user.User, error) {
	query := `
	select
		id, name, email, password_hash, roles, enabled, created_at, updated_at
	from
		users
	where
		email = :email and password_hash = :password_hash`

	dbUser := dbUser{}
	if err := s.db.GetContext(ctx, &dbUser, query, map[string]string{"email": email.String(), "password_hash": passwordHash}); err != nil {
		return user.User{}, fmt.Errorf("db.GetContext[%s]: %w", query, err)
	}

	coreUser, err := toCoreUser(dbUser)
	if err != nil {
		return user.User{}, fmt.Errorf("toCoreUser[%v]: %w", dbUser, err)
	}
	return coreUser, nil
}

func (s *Store) Query(context.Context, user.Filter, filter.OrderBy, filter.Page) ([]user.User, error) {
	return []user.User{}, nil
}
