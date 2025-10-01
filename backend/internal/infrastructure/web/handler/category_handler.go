package handler

import (
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryHandler カテゴリハンドラー
type CategoryHandler struct {
	categoryUseCase *usecase.CategoryUseCase
}

// NewCategoryHandler CategoryHandlerのコンストラクタ
func NewCategoryHandler(categoryUseCase *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: categoryUseCase,
	}
}

// CreateCategory カテゴリ作成
// @Summary カテゴリ作成
// @Description 新しいカテゴリを作成します
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "カテゴリ作成リクエスト"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	category, err := h.categoryUseCase.CreateCategory(c.Request.Context(), &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory カテゴリ取得
// @Summary カテゴリ取得
// @Description 指定されたIDのカテゴリを取得します
// @Tags categories
// @Produce json
// @Param id path string true "カテゴリID"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	categoryID := c.Param("id")

	category, err := h.categoryUseCase.GetCategory(c.Request.Context(), categoryID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory カテゴリ更新
// @Summary カテゴリ更新
// @Description 指定されたIDのカテゴリを更新します
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "カテゴリID"
// @Param category body dto.UpdateCategoryRequest true "カテゴリ更新リクエスト"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	category, err := h.categoryUseCase.UpdateCategory(c.Request.Context(), categoryID, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory カテゴリ削除
// @Summary カテゴリ削除
// @Description 指定されたIDのカテゴリを削除します
// @Tags categories
// @Param id path string true "カテゴリID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	err := h.categoryUseCase.DeleteCategory(c.Request.Context(), categoryID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllCategories 全カテゴリ取得
// @Summary 全カテゴリ取得
// @Description 全てのカテゴリを取得します
// @Tags categories
// @Produce json
// @Success 200 {array} dto.CategoryResponse
// @Failure 500 {object} ErrorResponse
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryUseCase.GetAllCategories(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}
