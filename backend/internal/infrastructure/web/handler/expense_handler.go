package handler

import (
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExpenseHandler 経費ハンドラー
type ExpenseHandler struct {
	expenseUseCase *usecase.ExpenseUseCase
}

// NewExpenseHandler ExpenseHandlerのコンストラクタ
func NewExpenseHandler(expenseUseCase *usecase.ExpenseUseCase) *ExpenseHandler {
	return &ExpenseHandler{
		expenseUseCase: expenseUseCase,
	}
}

// CreateExpense 経費作成
// @Summary 経費作成
// @Description 新しい経費を作成します
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Param expense body dto.CreateExpenseRequest true "経費作成リクエスト"
// @Success 201 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/expenses [post]
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	userID := c.Param("id")

	var req dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	expense, err := h.expenseUseCase.CreateExpense(c.Request.Context(), userID, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, expense)
}

// GetExpense 経費取得
// @Summary 経費取得
// @Description 指定されたIDの経費を取得します
// @Tags expenses
// @Produce json
// @Param id path string true "経費ID"
// @Success 200 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id} [get]
func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	expenseID := c.Param("id")

	expense, err := h.expenseUseCase.GetExpense(c.Request.Context(), expenseID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expense)
}

// UpdateExpense 経費更新
// @Summary 経費更新
// @Description 指定されたIDの経費を更新します
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path string true "経費ID"
// @Param expense body dto.UpdateExpenseRequest true "経費更新リクエスト"
// @Success 200 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id} [put]
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	expenseID := c.Param("id")

	var req dto.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "リクエストの形式が正しくありません",
			Details: err.Error(),
		})
		return
	}

	expense, err := h.expenseUseCase.UpdateExpense(c.Request.Context(), expenseID, &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expense)
}

// DeleteExpense 経費削除
// @Summary 経費削除
// @Description 指定されたIDの経費を削除します
// @Tags expenses
// @Param id path string true "経費ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	expenseID := c.Param("id")

	err := h.expenseUseCase.DeleteExpense(c.Request.Context(), expenseID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetExpensesByUser ユーザーの経費一覧取得
// @Summary ユーザーの経費一覧取得
// @Description 指定されたユーザーの経費一覧を取得します
// @Tags expenses
// @Produce json
// @Param id path string true "ユーザーID"
// @Param status query string false "ステータス"
// @Success 200 {array} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id}/expenses [get]
func (h *ExpenseHandler) GetExpensesByUser(c *gin.Context) {
	userID := c.Param("id")
	status := c.Query("status")

	var expenses []*dto.ExpenseResponse
	var err error

	if status != "" {
		expenses, err = h.expenseUseCase.GetExpensesByUserAndStatus(c.Request.Context(), userID, status)
	} else {
		expenses, err = h.expenseUseCase.GetExpensesByUser(c.Request.Context(), userID)
	}

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expenses)
}

// SubmitExpense 経費申請
// @Summary 経費申請
// @Description 経費を申請状態に変更します
// @Tags expenses
// @Param id path string true "経費ID"
// @Success 200 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id}/submit [post]
func (h *ExpenseHandler) SubmitExpense(c *gin.Context) {
	expenseID := c.Param("id")

	expense, err := h.expenseUseCase.SubmitExpense(c.Request.Context(), expenseID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expense)
}

// ApproveExpense 経費承認
// @Summary 経費承認
// @Description 経費を承認状態に変更します
// @Tags expenses
// @Param id path string true "経費ID"
// @Success 200 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id}/approve [post]
func (h *ExpenseHandler) ApproveExpense(c *gin.Context) {
	expenseID := c.Param("id")

	expense, err := h.expenseUseCase.ApproveExpense(c.Request.Context(), expenseID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expense)
}

// RejectExpense 経費却下
// @Summary 経費却下
// @Description 経費を却下状態に変更します
// @Tags expenses
// @Param id path string true "経費ID"
// @Success 200 {object} dto.ExpenseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /expenses/{id}/reject [post]
func (h *ExpenseHandler) RejectExpense(c *gin.Context) {
	expenseID := c.Param("id")

	expense, err := h.expenseUseCase.RejectExpense(c.Request.Context(), expenseID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, expense)
}
