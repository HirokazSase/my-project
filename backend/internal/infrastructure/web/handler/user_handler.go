package handler

import (
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler ユーザーハンドラー
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler UserHandlerのコンストラクタ
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUser ユーザー作成
// @Summary ユーザー作成
// @Description 新しいユーザーを作成します
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "ユーザー作成リクエスト"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	user, err := h.userUseCase.CreateUser(c.Request.Context(), &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser ユーザー取得
// @Summary ユーザー取得
// @Description 指定されたIDのユーザーを取得します
// @Tags users
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.userUseCase.GetUser(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser ユーザー更新
// @Summary ユーザー更新
// @Description 指定されたIDのユーザーを更新します
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Param user body dto.UpdateUserRequest true "ユーザー更新リクエスト"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	user, err := h.userUseCase.UpdateUser(c.Request.Context(), userID, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser ユーザー削除
// @Summary ユーザー削除
// @Description 指定されたIDのユーザーを削除します
// @Tags users
// @Param id path string true "ユーザーID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.userUseCase.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllUsers 全ユーザー取得
// @Summary 全ユーザー取得
// @Description 全てのユーザーを取得します
// @Tags users
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAllUsers(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
