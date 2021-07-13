package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"go-backend/db"
	"go-backend/model"
	"time"
)

type UserRepoImpl struct {
	sql *db.SQL
}

func NewUserRepo(sql *db.SQL) UserPepo {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {

	statement := `
		INSERT INTO users(id, email, phone, password, address, full_name, avatar, role, created_at, updated_at)
		VALUES(:id, :email, :phone, :password, :address, :full_name, :avatar, :role, :created_at, :updated_at)
	`
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAT = now

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "Unique_violation" {
				return user, errors.New("Người dùng đã tồn tại")
			}
		}
		return user, errors.New("Đăng ký tài khoản thất bại")
	}
	return user, nil
}

func (u UserRepoImpl) SelectUserByEmail(context context.Context, email string) (model.User, error) {
	var user = model.User{}
	statemnt := `SELECT * FROM users WHERE email=$1`
	err := u.sql.Db.GetContext(context, &user, statemnt, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("lỗi, không tồn tại trường hợp này")
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUsserById(context context.Context, userID string) (model.User, error) {
	var user = model.User{}
	statemnt := `SELECT * FROM users WHERE id=$1`
	err := u.sql.Db.GetContext(context, &user, statemnt, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("lỗi, không tồn tại d trường hợp này")
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUsers(context context.Context) ([]model.User, error) {
	var users = []model.User{}
	statemnt := `SELECT * FROM users`
	err := u.sql.Db.SelectContext(context, &users, statemnt)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, errors.New("lỗi, không tồn tại d trường hợp này")
		}
		log.Error(err.Error())
		return users, err
	}

	return users, nil
}


