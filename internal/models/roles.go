package models

type ConfluentRole string

const (
	DeveloperRead  ConfluentRole = "DeveloperRead"
	DeveloperWrite ConfluentRole = "DeveloperWrite"
	ResourceOwner  ConfluentRole = "ResourceOwner"
	Operator       ConfluentRole = "Operator"
	SecurityAdmin  ConfluentRole = "SecurityAdmin"
	SystemAdmin    ConfluentRole = "SystemAdmin"
	UserAdmin      ConfluentRole = "UserAdmin"
	AuditAdmin     ConfluentRole = "AuditAdmin"
)

var AllRoles = []ConfluentRole{
	DeveloperRead,
	DeveloperWrite,
	ResourceOwner,
	Operator,
	SecurityAdmin,
	SystemAdmin,
	UserAdmin,
	AuditAdmin,
}
