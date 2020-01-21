package repositories

type DbTask struct {
	Id          string `gorm:"primary_key" sql:"type:CHAR(36)"`
	Title       string `sql:"type:CHAR(36)"`
	Description string `sql:"type:TEXT"`
	Status      string `sql:"type:CHAR(36)"`
}
