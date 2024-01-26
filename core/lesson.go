package core

type Lesson struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Type     string `db:"type"`
	Duration string `db:"duration"`
	Author   string `db:"author"`
	Link     string `db:"link"`
	Time     string `db:"time"`
	Location string `db:"location"`
}
