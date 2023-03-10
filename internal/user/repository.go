package user

import (
	"context"

	"github.com/anousoneFS/clean-architecture/config"
)

func (r UserUC) list(ctx context.Context) ([]User, error) {
	query, args, err := config.Psql().
		Select("name", "age").
		From("users").
		ToSql()
	if err != nil {
		return []User{}, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	resp := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.Name, &i.Age); err != nil {
			return []User{}, err
		}
		resp = append(resp, i)
	}
	if err != nil {
		return []User{}, err
	}
	return resp, nil
}

func (r UserUC) create(ctx context.Context, req User) error {
	query, args, err := config.Psql().
		Insert("users").
		Columns("name", "age").
		Values(req.Name, req.Age).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
