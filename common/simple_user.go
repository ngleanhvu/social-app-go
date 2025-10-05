package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Role      string `json:"role" gorm:"column:role"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (s *SimpleUser) Mask(isAdmin bool) {
	s.GenUID(DbTypeUser)
}
