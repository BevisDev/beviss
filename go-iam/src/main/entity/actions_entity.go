package entity

type Actions struct {
	AuditEntity
	Name        string `db:"name"`
	Description string `db:"description"`
}
