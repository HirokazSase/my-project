import { Category } from '../../domain/models/Category';
import { CategoryRepository } from '../../domain/repositories/CategoryRepository';
import { CategoryApiService } from '../../infrastructure/api/CategoryApiService';

export class CategoryViewModel {
  private categoryRepository: CategoryRepository;

  constructor() {
    this.categoryRepository = new CategoryApiService();
  }

  async getAllCategories(): Promise<Category[]> {
    try {
      return await this.categoryRepository.getAll();
    } catch (error) {
      console.error('Failed to fetch categories:', error);
      throw new Error('カテゴリの取得に失敗しました');
    }
  }

  async getCategoryById(id: string): Promise<Category | null> {
    try {
      return await this.categoryRepository.getById(id);
    } catch (error) {
      console.error('Failed to fetch category:', error);
      throw new Error('カテゴリの取得に失敗しました');
    }
  }

  async createCategory(
    name: string,
    description?: string,
    color?: string
  ): Promise<Category> {
    try {
      // バリデーション
      const validationErrors = this.validateCategory(name, description, color);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const categoryData = {
        name: name.trim(),
        description: description?.trim() || '',
        color: color || '#3b82f6',
      };

      return await this.categoryRepository.create(categoryData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to create category:', error);
      throw new Error('カテゴリの作成に失敗しました');
    }
  }

  async updateCategory(
    id: string,
    name: string,
    description?: string,
    color?: string
  ): Promise<Category> {
    try {
      // バリデーション
      const validationErrors = this.validateCategory(name, description, color);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const categoryData = {
        name: name.trim(),
        description: description?.trim() || '',
        color: color || '#3b82f6',
      };

      return await this.categoryRepository.update(id, categoryData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to update category:', error);
      throw new Error('カテゴリの更新に失敗しました');
    }
  }

  async deleteCategory(id: string): Promise<void> {
    try {
      await this.categoryRepository.delete(id);
    } catch (error) {
      console.error('Failed to delete category:', error);
      throw new Error('カテゴリの削除に失敗しました');
    }
  }

  validateCategoryName(name: string): string[] {
    const errors: string[] = [];

    if (!name || name.trim().length === 0) {
      errors.push('カテゴリ名は必須です');
    } else if (name.trim().length > 50) {
      errors.push('カテゴリ名は50文字以内で入力してください');
    }

    return errors;
  }

  validateCategoryDescription(description?: string): string[] {
    const errors: string[] = [];

    if (description && description.length > 200) {
      errors.push('説明は200文字以内で入力してください');
    }

    return errors;
  }

  validateCategoryColor(color?: string): string[] {
    const errors: string[] = [];

    if (color && !/^#[0-9A-Fa-f]{6}$/.test(color)) {
      errors.push('無効なカラーコードです');
    }

    return errors;
  }

  validateCategory(name: string, description?: string, color?: string): string[] {
    const errors: string[] = [];

    errors.push(...this.validateCategoryName(name));
    errors.push(...this.validateCategoryDescription(description));
    errors.push(...this.validateCategoryColor(color));

    return errors;
  }
}