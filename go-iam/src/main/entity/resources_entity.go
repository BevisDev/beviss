package entity

type Resources struct {
	AuditEntity
	Name        string `db:"name"`
	Description string `db:"description"`
}
