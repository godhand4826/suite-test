package user_test

import (
	mockUser "suite-test/mocks/suite-test/user"
	"suite-test/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Test entry
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

type UserSuite struct {
	suite.Suite

	// mocks
	repo   *mockUser.MockUserRepository
	sender *mockUser.MockEmailSender

	// SUT
	svc user.UserService
}

func (s *UserSuite) SetupTest() {
	// initialize all mocks
	// mocks will automatically assert expectation calls when cleaning up
	s.repo = mockUser.NewMockUserRepository(s.T())
	s.sender = mockUser.NewMockEmailSender(s.T())

	// initialize SUT
	s.svc = user.NewUserService(s.repo, s.sender)
}

func (s *UserSuite) TearDownTest() {
	// clear state to avoid accidentally reused
	s.repo = nil
	s.sender = nil

	s.svc = nil
}

func (s *UserSuite) setupMockExpect(expect MockExpect) {
	expect.UserRepository.applyOn(s.repo)
	expect.EmailSenderSend.applyOn(s.sender)
}

type MockExpect struct {
	UserRepository  *UserRepositoryMockExpect
	EmailSenderSend *EmailSenderSendMockExpect
}

type UserRepositoryMockExpect struct {
	Get  *UserRepositoryGetMockExpect
	Save *UserRepositorySaveMockExpect
}

func (m *UserRepositoryMockExpect) applyOn(repo *mockUser.MockUserRepository) {
	if m == nil {
		return
	}

	m.Get.applyOn(repo)
	m.Save.applyOn(repo)
}

type UserRepositorySaveMockExpect struct {
	ExpectUser user.User

	ReturnUser user.User
	ReturnErr  error
}

func (m *UserRepositorySaveMockExpect) applyOn(repo *mockUser.MockUserRepository) {
	if m == nil {
		return
	}

	repo.EXPECT().Save(m.ExpectUser).Return(m.ReturnUser, m.ReturnErr).Once()
}

type UserRepositoryGetMockExpect struct {
	ExpectID int

	ReturnUser user.User
	ReturnErr  error
}

func (m *UserRepositoryGetMockExpect) applyOn(repo *mockUser.MockUserRepository) {
	if m == nil {
		return
	}

	repo.EXPECT().Get(m.ExpectID).Return(m.ReturnUser, m.ReturnErr).Once()
}

type EmailSenderSendMockExpect struct {
	ExpectEmail   string
	ExpectContent string

	Return error
}

func (m *EmailSenderSendMockExpect) applyOn(sender *mockUser.MockEmailSender) {
	if m == nil {
		return
	}

	sender.EXPECT().Send(m.ExpectEmail, m.ExpectContent).Return(m.Return).Once()
}
