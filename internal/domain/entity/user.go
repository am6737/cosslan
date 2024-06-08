package entity

type User struct {
	UserID    string
	Name      string
	Password  string
	Email     string
	Status    UserStatus
	Role      UserRole
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

// UserStatus 是用户的状态枚举
type UserStatus int

const (
	Active    UserStatus = iota + 1 //正常
	Suspended                       //冻结
	Deleted                         //删除
)

// IsActive 返回用户是否处于活跃状态
func (s UserStatus) IsActive() bool {
	return s == Active
}

// IsAdmin 返回是否是管理员
func (r UserRole) IsAdmin() bool {
	return r == Admin
}

// UserRole 是用户的角色枚举
type UserRole int

const (
	Admin  UserRole = iota + 1 //超管
	Member                     //用户
)
