package dao

type UserDAO struct {
	Email          string `validate:"email"`
	Name           string `validate:""`
	HashedPassword string `validate:""`
	Balance        int    `validate:"gte:0"`
	InuseBalance   int    `validate:"gte:0"`
}

func CreateUser(email string, name string, hashedPassword string) (*UserDAO, error) {
	userDao := &UserDAO{
		email,
		name,
		hashedPassword,
		0,
		0,
	}
	
	return userDao, nil
}
