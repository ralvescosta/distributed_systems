package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
	"webapi/pkg/domain/usecases"
)

type mockConfigure struct {
	method       string
	customResult interface{}
	customError  error
}
type createUserUseCaseToTest struct {
	useCase      usecases.ICreateUserUseCase
	repo         interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
}

func newCreateUserUseCaseToTest(configs map[string]mockConfigure) createUserUseCaseToTest {
	repoConfig, ok := configs["userRepository"]
	var repo interfaces.IUserRepository
	if ok {
		repo = userRepositorySpy{config: &repoConfig}
	} else {
		repo = userRepositorySpy{}
	}

	hasherConfig, ok := configs["hasher"]
	var hasher interfaces.IHasher
	if ok {
		hasher = hasherSpy{config: &hasherConfig}
	} else {
		hasher = hasherSpy{}
	}

	tokenManagerConfig, ok := configs["tokenManager"]
	var tokenManager interfaces.ITokenManager
	if ok {
		tokenManager = tokenManagerSpy{config: &tokenManagerConfig}
	} else {
		tokenManager = tokenManagerSpy{}
	}

	useCase := NewCreateUserUseCase(repo, hasher, tokenManager)
	return createUserUseCaseToTest{useCase, repo, hasher, tokenManager}
}

type authenticationUseCaseToTest struct {
	useCase      usecases.IValidationTokenUseCase
	repo         interfaces.IUserRepository
	tokenManager interfaces.ITokenManager
}

func newAuthenticationUseCaseToTest(configs map[string]mockConfigure) authenticationUseCaseToTest {
	repoConfig, ok := configs["userRepository"]
	var repo interfaces.IUserRepository
	if ok {
		repo = userRepositorySpy{config: &repoConfig}
	} else {
		repo = userRepositorySpy{}
	}

	tokenManagerConfig, ok := configs["tokenManager"]
	var tokenManager interfaces.ITokenManager
	if ok {
		tokenManager = tokenManagerSpy{config: &tokenManagerConfig}
	} else {
		tokenManager = tokenManagerSpy{}
	}

	useCase := NewAuthenticationUseCase(repo, tokenManager)
	return authenticationUseCaseToTest{useCase, repo, tokenManager}
}

type userRepositorySpy struct {
	config *mockConfigure
}

func (pst userRepositorySpy) FindById(ctx context.Context, txn interface{}, id int) (*entities.User, error) {
	if pst.config != nil && pst.config.method == "FindById" {
		user, ok := pst.config.customResult.(*entities.User)
		if ok {
			return user, pst.config.customError
		}
		return nil, pst.config.customError
	}

	return &entities.User{}, nil
}
func (pst userRepositorySpy) FindByEmail(ctx context.Context, txn interface{}, email string) (*entities.User, error) {
	if pst.config != nil && pst.config.method == "FindByEmail" {
		user, ok := pst.config.customResult.(*entities.User)
		if ok {
			return user, pst.config.customError
		}

		return nil, pst.config.customError
	}

	return nil, nil
}
func (pst userRepositorySpy) Create(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (*entities.User, error) {
	if pst.config != nil && pst.config.method == "Create" {
		user, ok := pst.config.customResult.(*entities.User)
		if ok {
			return user, pst.config.customError
		}
		return nil, pst.config.customError
	}

	return &entities.User{}, nil
}

type hasherSpy struct {
	config *mockConfigure
}

func (pst hasherSpy) Hahser(text string) (string, error) {
	if pst.config != nil && pst.config.method == "Hahser" {
		return pst.config.customResult.(string), pst.config.customError
	}

	return "", nil
}
func (pst hasherSpy) Verify(originalText, hashedText string) bool {
	if pst.config != nil && pst.config.method == "Verify" {
		return pst.config.customResult.(bool)
	}

	return true
}

type tokenManagerSpy struct {
	config *mockConfigure
}

func (pst tokenManagerSpy) GenerateToken(tokenData dtos.TokenDataDto) (string, error) {
	if pst.config != nil && pst.config.method == "GenerateToken" {
		return pst.config.customResult.(string), pst.config.customError
	}

	return "", nil
}
func (pst tokenManagerSpy) VerifyToken(token string) (*dtos.SessionDto, error) {
	if pst.config != nil && pst.config.method == "VerifyToken" {
		return pst.config.customResult.(*dtos.SessionDto), pst.config.customError
	}

	return &dtos.SessionDto{}, nil
}