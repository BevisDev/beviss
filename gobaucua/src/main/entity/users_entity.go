package entity

type Users struct {
	AuditEntity
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
}
