package entity

type Roles struct {
	AuditEntity
	Name        string `db:"name"`
	Description string `db:"description"`
}
