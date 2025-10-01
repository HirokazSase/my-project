import { Expense } from '../../domain/models/Expense';
import { ExpenseRepository } from '../../domain/repositories/ExpenseRepository';
import { ExpenseApiService } from '../../infrastructure/api/ExpenseApiService';

export class ExpenseViewModel {
  private expenseRepository: ExpenseRepository;

  constructor() {
    this.expenseRepository = new ExpenseApiService();
  }

  async getExpensesByUserAndStatus(userId: string, status: string): Promise<Expense[]> {
    try {
      return await this.expenseRepository.getByUserAndStatus(userId, status);
    } catch (error) {
      console.error('Failed to fetch expenses:', error);
      throw new Error('経費の取得に失敗しました');
    }
  }

  async getAllExpensesByUser(userId: string): Promise<Expense[]> {
    try {
      return await this.expenseRepository.getByUserId(userId);
    } catch (error) {
      console.error('Failed to fetch expenses:', error);
      throw new Error('経費の取得に失敗しました');
    }
  }

  async getExpenseById(id: string): Promise<Expense | null> {
    try {
      return await this.expenseRepository.getById(id);
    } catch (error) {
      console.error('Failed to fetch expense:', error);
      throw new Error('経費の取得に失敗しました');
    }
  }

  async createExpense(
    userId: string,
    categoryId: string,
    amount: number,
    currency: string,
    title: string,
    description: string,
    date: Date
  ): Promise<Expense> {
    try {
      // バリデーション
      const validationErrors = this.validateExpense(userId, categoryId, amount, title, date);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const expenseData = {
        userId,
        categoryId,
        amount,
        currency,
        title: title.trim(),
        description: description.trim(),
        date,
      };

      return await this.expenseRepository.create(expenseData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to create expense:', error);
      throw new Error('経費の作成に失敗しました');
    }
  }

  async updateExpense(
    id: string,
    categoryId: string,
    amount: number,
    currency: string,
    title: string,
    description: string,
    date: Date
  ): Promise<Expense> {
    try {
      // バリデーション
      const validationErrors = this.validateExpenseUpdate(categoryId, amount, title, date);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const expenseData = {
        categoryId,
        amount,
        currency,
        title: title.trim(),
        description: description.trim(),
        date,
      };

      return await this.expenseRepository.update(id, expenseData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to update expense:', error);
      throw new Error('経費の更新に失敗しました');
    }
  }

  async deleteExpense(id: string): Promise<void> {
    try {
      await this.expenseRepository.delete(id);
    } catch (error) {
      console.error('Failed to delete expense:', error);
      throw new Error('経費の削除に失敗しました');
    }
  }

  async approveExpense(id: string): Promise<Expense> {
    try {
      return await this.expenseRepository.updateStatus(id, 'approved');
    } catch (error) {
      console.error('Failed to approve expense:', error);
      throw new Error('経費の承認に失敗しました');
    }
  }

  async rejectExpense(id: string): Promise<Expense> {
    try {
      return await this.expenseRepository.updateStatus(id, 'rejected');
    } catch (error) {
      console.error('Failed to reject expense:', error);
      throw new Error('経費の却下に失敗しました');
    }
  }

  // バリデーションメソッド
  validateExpense(userId: string, categoryId: string, amount: number, title: string, date: Date): string[] {
    const errors: string[] = [];

    if (!userId || userId.trim().length === 0) {
      errors.push('ユーザーIDは必須です');
    }

    if (!categoryId || categoryId.trim().length === 0) {
      errors.push('カテゴリは必須です');
    }

    if (!amount || amount <= 0) {
      errors.push('金額は0より大きい値を入力してください');
    }

    if (amount && amount > 10000000) {
      errors.push('金額は10,000,000円以下で入力してください');
    }

    if (!title || title.trim().length === 0) {
      errors.push('タイトルは必須です');
    }

    if (title && title.trim().length > 100) {
      errors.push('タイトルは100文字以内で入力してください');
    }

    if (!date) {
      errors.push('日付は必須です');
    }

    if (date && date > new Date()) {
      errors.push('未来の日付は入力できません');
    }

    return errors;
  }

  validateExpenseUpdate(categoryId: string, amount: number, title: string, date: Date): string[] {
    const errors: string[] = [];

    if (!categoryId || categoryId.trim().length === 0) {
      errors.push('カテゴリは必須です');
    }

    if (!amount || amount <= 0) {
      errors.push('金額は0より大きい値を入力してください');
    }

    if (amount && amount > 10000000) {
      errors.push('金額は10,000,000円以下で入力してください');
    }

    if (!title || title.trim().length === 0) {
      errors.push('タイトルは必須です');
    }

    if (title && title.trim().length > 100) {
      errors.push('タイトルは100文字以内で入力してください');
    }

    if (!date) {
      errors.push('日付は必須です');
    }

    if (date && date > new Date()) {
      errors.push('未来の日付は入力できません');
    }

    return errors;
  }

  validateAmount(amount: number): string[] {
    const errors: string[] = [];

    if (!amount || amount <= 0) {
      errors.push('金額は0より大きい値を入力してください');
    }

    if (amount && amount > 10000000) {
      errors.push('金額は10,000,000円以下で入力してください');
    }

    return errors;
  }

  validateTitle(title: string): string[] {
    const errors: string[] = [];

    if (!title || title.trim().length === 0) {
      errors.push('タイトルは必須です');
    }

    if (title && title.trim().length > 100) {
      errors.push('タイトルは100文字以内で入力してください');
    }

    return errors;
  }

  validateDescription(description: string): string[] {
    const errors: string[] = [];

    if (description && description.length > 500) {
      errors.push('説明は500文字以内で入力してください');
    }

    return errors;
  }

  formatAmount(amount: number): string {
    return new Intl.NumberFormat('ja-JP', {
      style: 'currency',
      currency: 'JPY',
    }).format(amount);
  }

  formatDate(date: Date): string {
    return new Intl.DateTimeFormat('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    }).format(date);
  }
}