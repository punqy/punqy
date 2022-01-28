package storage

type Role string

func (r Role) String() string {
	return string(r)
}

const (
	RoleUser       Role = "ROLE_USER"
	RoleAdmin           = "ROLE_ADMIN"
	RoleSuperAdmin      = "ROLE_SUPER_ADMIN"
)

type User struct {
	Entity
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	Roles     StringList `json:"roles" db:"roles"`
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) GetID() string {
	return u.ID.String()
}
