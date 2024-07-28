package user_test

import (
	"errors"
	"suite-test/user"
)

func (s *UserSuite) TestSuccessfulRegister() {
	s.setupMockExpect(MockExpect{
		UserRepository: &UserRepositoryMockExpect{
			Save: &UserRepositorySaveMockExpect{
				ExpectUser: user.User{Name: "eric", Email: "email@mail.com"},
				ReturnUser: user.User{Id: 1, Name: "eric", Email: "email@mail.com"},
			},
		},
		EmailSenderSend: &EmailSenderSendMockExpect{
			ExpectEmail:   "email@mail.com",
			ExpectContent: "welcome to join us",
		},
	})

	u, err := s.svc.Register("eric", "email@mail.com")

	s.Assert().NoError(err)
	s.Assert().Equal(u, user.User{Id: 1, Name: "eric", Email: "email@mail.com"})
}

func (s *UserSuite) TestFailureRegisterRepositoryError() {
	s.setupMockExpect(MockExpect{
		UserRepository: &UserRepositoryMockExpect{
			Save: &UserRepositorySaveMockExpect{
				ExpectUser: user.User{Name: "eric", Email: "email@mail.com"},
				ReturnUser: user.User{},
				ReturnErr:  errors.New("repo is down"),
			},
		},
	})

	u, err := s.svc.Register("eric", "email@mail.com")

	s.Assert().Error(err)
	s.Assert().Zero(u)
}

func (s *UserSuite) TestFailureRegisterInvalidName() {
	s.setupMockExpect(MockExpect{})

	u, err := s.svc.Register("", "email@mail.com")

	s.Assert().Error(err)
	s.Assert().Zero(u)
}

func (s *UserSuite) TestFailureRegisterSendFailed() {
	var sendError = errors.New("sender error")
	s.setupMockExpect(MockExpect{
		UserRepository: &UserRepositoryMockExpect{
			Save: &UserRepositorySaveMockExpect{
				ExpectUser: user.User{Name: "eric", Email: "email@mail.com"},
				ReturnUser: user.User{Id: 1, Name: "eric", Email: "email@mail.com"},
			},
		},
		EmailSenderSend: &EmailSenderSendMockExpect{
			ExpectEmail:   "email@mail.com",
			ExpectContent: "welcome to join us",
			Return:        sendError,
		},
	})

	u, err := s.svc.Register("eric", "email@mail.com")

	s.Assert().Error(err)
	s.Assert().Zero(u)
}

func (s *UserSuite) TestSuccessfulFindUser() {
	s.setupMockExpect(MockExpect{
		UserRepository: &UserRepositoryMockExpect{
			Get: &UserRepositoryGetMockExpect{
				ExpectID:   1,
				ReturnUser: user.User{Id: 1, Name: "eric", Email: "email@mail.com"},
			},
		},
	})

	u, err := s.svc.Find(1)

	s.Assert().NoError(err)
	s.Assert().Equal(u, user.User{Id: 1, Name: "eric", Email: "email@mail.com"})
}
