package db

var schema = `
CREATE TABLE visitors (
	id SERIAL,
	user_id int,
	firstname string,
	lastname string,
	username string,
	language_code string,
	is_premium bool,
	is_bot bool,
	added_to_menu bool,
	can_join_groups bool,
	can_read_messages bool
);
`

type Visitors struct {
	ID              int    `db:"id"`
	UserID          int    `db:"user_id"`
	FirstName       string `db:"firstname"`
	LastName        string `db:"lastname"`
	Username        string `db:"username"`
	LanguageCode    string `db:"language_code"`
	IsPremium       bool   `db:"is_premium"`
	IsBot           bool   `db:"is_bot"`
	AddedToMenu     bool   `db:"added_to_menu"`
	CanJoinGroups   bool   `db:"can_join_groups"`
	CanReadMessages bool   `db:"can_read_messages"`
}
