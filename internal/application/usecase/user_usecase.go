package usecase

import (
	"context"
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/repository"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
)

// UserUseCase ユーザーユースケース
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase UserUseCaseのコンストラクタ
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// CreateUser ユーザーを作成
func (uc *UserUseCase) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	// メールアドレスの重複チェック
	existingUser, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.NewApplicationError("EMAIL_ALREADY_EXISTS", "このメールアドレスは既に使用されています")
	}

	// 新しいユーザーを作成
	user, err := entity.NewUser(req.Name, req.Email)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// ユーザーを保存
	if err := uc.userRepo.Save(ctx, user); err != nil {
		return nil, errors.NewApplicationError("USER_CREATION_FAILED", "ユーザーの作成に失敗しました")
	}

	return &dto.UserResponse{
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

// GetUser ユーザーを取得
func (uc *UserUseCase) GetUser(ctx context.Context, userID string) (*dto.UserResponse, error) {
	id, err := valueobject.NewUserID(userID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	return &dto.UserResponse{
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

// UpdateUser ユーザーを更新
func (uc *UserUseCase) UpdateUser(ctx context.Context, userID string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	id, err := valueobject.NewUserID(userID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	// メールアドレスの重複チェック（自分以外で同じメールアドレスが存在するか）
	if req.Email != user.Email() {
		existingUser, err := uc.userRepo.FindByEmail(ctx, req.Email)
		if err == nil && existingUser != nil && !existingUser.ID().Equals(user.ID()) {
			return nil, errors.NewApplicationError("EMAIL_ALREADY_EXISTS", "このメールアドレスは既に使用されています")
		}
	}

	// ユーザー情報を更新
	if err := user.UpdateProfile(req.Name, req.Email); err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// ユーザーを保存
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, errors.NewApplicationError("USER_UPDATE_FAILED", "ユーザーの更新に失敗しました")
	}

	return &dto.UserResponse{
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

// DeleteUser ユーザーを削除
func (uc *UserUseCase) DeleteUser(ctx context.Context, userID string) error {
	id, err := valueobject.NewUserID(userID)
	if err != nil {
		return errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	exists, err := uc.userRepo.Exists(ctx, id)
	if err != nil {
		return errors.NewApplicationError("USER_DELETE_FAILED", "ユーザーの削除チェックに失敗しました")
	}

	if !exists {
		return errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return errors.NewApplicationError("USER_DELETE_FAILED", "ユーザーの削除に失敗しました")
	}

	return nil
}

// GetAllUsers 全てのユーザーを取得
func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.NewApplicationError("USER_FETCH_FAILED", "ユーザー一覧の取得に失敗しました")
	}

	responses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = &dto.UserResponse{
			ID:        user.ID().String(),
			Name:      user.Name(),
			Email:     user.Email(),
			CreatedAt: user.CreatedAt(),
			UpdatedAt: user.UpdatedAt(),
		}
	}

	return responses, nil
}