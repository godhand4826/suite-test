package user

import "errors"

type User struct {
	Id    int
	Name  string
	Email string
}

type UserService interface {
	Register(name, email string) (User, error)
	Find(id int) (User, error)
}

type UserRepository interface {
	Save(User) (User, error)
	Get(id int) (User, error)
}

type EmailSender interface {
	Send(email, content string) error
}

type userService struct {
	repo   UserRepository
	sender EmailSender
}

func NewUserService(repo UserRepository, sender EmailSender) UserService {
	return &userService{
		repo:   repo,
		sender: sender,
	}
}

func (s *userService) Register(name, email string) (User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	if len(name) == 0 || len(email) == 0 {
		return User{}, errors.New("invalid argument")
	}

	user, err := s.repo.Save(user)
	if err != nil {
		return User{}, err
	}

	if err := s.sender.Send(user.Email, "welcome to join us"); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) Find(id int) (User, error) {
	return s.repo.Get(id)
}
