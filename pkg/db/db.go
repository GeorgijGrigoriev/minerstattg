package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v3"
)

type Op struct {
	c *sqlx.DB
}

type Com interface {
	CheckUser(id int) (bool, error)
	AddUser(user tele.User) (bool, error)
	UpdateUser(user tele.User) (bool, error)
}

func Init() *Op {
	return &Op{}
}

func (o *Op) Get() {
	db, err := sqlx.Connect("pgx", "postgress://postgres:postgres@db:5432/tgbot")
	if err != nil {
		log.Fatal(err.Error())

		o.c = nil
	}

	o.c = db
}

func (o *Op) CheckUser(id int) (bool, error) {
	res := Visitors{}

	err := o.c.Get(&res, "SELECT * FROM visitors WHERE user_id=$1", id)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if res.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (o *Op) AddUser(user *tele.User) (bool, error) {
	tx := o.c.MustBegin()
	tx.MustExec("INSERT INTO visitors (user_id, firstname, lastname, username, language_code, is_premium, is_bot, added_to_menu, can_join_groups, can_read_messages) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		int(user.ID), user.FirstName, user.LastName, user.Username, user.LanguageCode, user.IsPremium, user.IsBot, user.AddedToMenu, user.CanJoinGroups, user.CanReadMessages)

	err := tx.Commit()
	if err != nil {
		log.Fatal(err)

		return false, err
	}

	return true, nil
}

func (o *Op) UpdateUser(user *tele.User) (bool, error) {
	tx := o.c.MustBegin()
	tx.MustExec("UPDATE visitors SET (user_id, firstname, lastname, username, language_code, is_premium, is_bot, added_to_menu, can_join_groups, can_read_messages) = ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		int(user.ID), user.FirstName, user.LastName, user.Username, user.LanguageCode, user.IsPremium, user.IsBot, user.AddedToMenu, user.CanJoinGroups, user.CanReadMessages)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)

		return false, err
	}

	return true, nil
}
