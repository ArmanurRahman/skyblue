package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/ArmanurRahman/skyblue/internal/models"
)

func (m *postgresRepo) InsetAddress(addr models.Address) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `
		insert into address (country, city, word, street, other_info, create_at, update_at)
		values
		($1, $2, $3, $4, $5, $6, $7) returning id
	`

	var id int
	row := m.DB.QueryRowContext(ctx, sql,
		addr.Country,
		addr.City,
		addr.Word,
		addr.Street,
		addr.OtherInfo,
		addr.CreateAt,
		addr.UpdateAt,
	)

	err := row.Scan(
		&id,
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *postgresRepo) InsetUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `
		insert into users (first_name, last_name, phone, email, password, address_id, create_at, update_at)
		values
		($1, $2, $3, $4, $5, $6, $7, $8) 
	`

	_, err := m.DB.ExecContext(ctx, sql,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Email,
		user.Password,
		user.AddressId,
		user.CreateAt,
		user.UpdateAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresRepo) InsetSaler(saler models.Saler) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sql := `
		insert into salers (name, details, phone, email, password, address_id, create_at, update_at)
		values
		($1, $2, $3, $4, $5, $6, $7, $8) 
	`

	_, err := m.DB.ExecContext(ctx, sql,
		saler.Name,
		saler.Details,
		saler.Phone,
		saler.Email,
		saler.Password,
		saler.AddressId,
		saler.CreateAt,
		saler.UpdateAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresRepo) Login(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	sql := `select first_name, last_name, email, phone, password, 
			country, city, word, street, other_info
			from users U
			inner join address AD on U.address_id = AD.id 
			where email = $1
	`

	row := m.DB.QueryRowContext(ctx, sql, email)

	err := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Address.Country,
		&user.Address.City,
		&user.Address.Word,
		&user.Address.Street,
		&user.Address.OtherInfo,
	)

	if user.FirstName == "" {
		return user, errors.New("user not found")
	}
	if err != nil {
		return user, err
	}

	return user, nil

}
