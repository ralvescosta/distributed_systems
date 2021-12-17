package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/infra/logger"
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
	logger       interfaces.ILogger
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

	logger := logger.NewLoggerSpy()

	useCase := NewCreateUserUseCase(repo, hasher, tokenManager)
	return createUserUseCaseToTest{useCase, repo, hasher, tokenManager, logger}
}

type validationTokenUseCaseToTest struct {
	useCase      usecases.IValidationTokenUseCase
	repo         interfaces.IUserRepository
	tokenManager interfaces.ITokenManager
}

func newValidationTokenUseCaseToTest(configs map[string]mockConfigure) validationTokenUseCaseToTest {
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

	useCase := NewValidatinTokenUseCase(repo, tokenManager)
	return validationTokenUseCaseToTest{useCase, repo, tokenManager}
}

type sessionUsecaseToTest struct {
	useCase      usecases.ISessionUseCase
	repo         interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
}

func newSessionUsecaseToTest(configs map[string]mockConfigure) sessionUsecaseToTest {
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

	logger := logger.NewLoggerSpy()

	useCase := NewSessionUseCase(repo, hasher, tokenManager, logger)
	return sessionUsecaseToTest{useCase, repo, hasher, tokenManager}
}

type userRepositorySpy struct {
	config *mockConfigure
}

func (pst userRepositorySpy) FindById(ctx context.Context, id int) (*entities.User, error) {
	if pst.config != nil && pst.config.method == "FindById" {
		user, ok := pst.config.customResult.(*entities.User)
		if ok {
			return user, pst.config.customError
		}
		return nil, pst.config.customError
	}

	return &entities.User{}, nil
}
func (pst userRepositorySpy) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	if pst.config != nil && pst.config.method == "FindByEmail" {
		user, ok := pst.config.customResult.(*entities.User)
		if ok {
			return user, pst.config.customError
		}

		return nil, pst.config.customError
	}

	return nil, nil
}
func (pst userRepositorySpy) Create(ctx context.Context, dto dtos.CreateUserDto) (*entities.User, error) {
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

type getProductByIdUsecaseToTest struct {
	useCase usecases.IGetProductByIdUseCase
}

func newGetProductByIdUsecaseToTest(configs map[string]mockConfigure) getProductByIdUsecaseToTest {
	inventoryClientConfig, ok := configs["inventoryClient"]
	var inventoryClient interfaces.IIventoryClient
	if ok {
		inventoryClient = inventoryClientSpy{config: &inventoryClientConfig}
	} else {
		inventoryClient = inventoryClientSpy{}
	}

	useCase := NewGetProductByIdUseCase(inventoryClient)
	return getProductByIdUsecaseToTest{useCase}
}

type inventoryClientSpy struct {
	config *mockConfigure
}

func (pst inventoryClientSpy) GetProductById(ctx context.Context, id string) (dtos.ProductDto, error) {
	if pst.config != nil && pst.config.method == "GetProductById" {
		return pst.config.customResult.(dtos.ProductDto), pst.config.customError
	}

	return dtos.ProductDto{}, nil
}

func (pst inventoryClientSpy) RegisterProduct(ctx context.Context, product dtos.ProductDto) (dtos.ProductDto, error) {
	if pst.config != nil && pst.config.method == "RegisterProduct" {
		return pst.config.customResult.(dtos.ProductDto), pst.config.customError
	}

	return dtos.ProductDto{}, nil
}
