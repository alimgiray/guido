package user

import "golang.org/x/crypto/bcrypt"

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Register(username, password, email, userType string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password = string(passwordHash)
	return u.userRepository.Insert(username, password, email, userType)
}

func (u *UserService) Login(username, password string) (*User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
