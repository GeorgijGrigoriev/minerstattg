package db

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
