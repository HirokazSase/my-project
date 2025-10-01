import { CategoryViewModel } from '../CategoryViewModel';
import { Category } from '../../../domain/models/Category';
import { CategoryRepository } from '../../../domain/repositories/CategoryRepository';

// Mock CategoryRepository
class MockCategoryRepository implements CategoryRepository {
  private categories: Category[] = [];

  async getAll(): Promise<Category[]> {
    return [...this.categories];
  }

  async getById(id: string): Promise<Category | null> {
    return this.categories.find(cat => cat.id === id) || null;
  }

  async create(data: { name: string; description?: string; color?: string }): Promise<Category> {
    const category = new Category(
      `cat-${Date.now()}`,
      data.name,
      data.description || '',
      data.color || '#3b82f6',
      new Date(),
      new Date()
    );
    this.categories.push(category);
    return category;
  }

  async update(id: string, data: { name: string; description?: string; color?: string }): Promise<Category> {
    const index = this.categories.findIndex(cat => cat.id === id);
    if (index === -1) {
      throw new Error('Category not found');
    }
    
    const updatedCategory = new Category(
      id,
      data.name,
      data.description || '',
      data.color || '#3b82f6',
      this.categories[index].createdAt,
      new Date()
    );
    
    this.categories[index] = updatedCategory;
    return updatedCategory;
  }

  async delete(id: string): Promise<void> {
    const index = this.categories.findIndex(cat => cat.id === id);
    if (index === -1) {
      throw new Error('Category not found');
    }
    this.categories.splice(index, 1);
  }
}

// Mock CategoryApiService
jest.mock('../../../infrastructure/api/CategoryApiService', () => {
  return {
    CategoryApiService: jest.fn().mockImplementation(() => new MockCategoryRepository()),
  };
});

describe('CategoryViewModel', () => {
  let viewModel: CategoryViewModel;
  let mockRepo: MockCategoryRepository;

  beforeEach(() => {
    viewModel = new CategoryViewModel();
    // @ts-ignore
    mockRepo = viewModel.categoryRepository as MockCategoryRepository;
  });

  afterEach(() => {
    jest.clearAllMocks();
  });

  describe('getAllCategories', () => {
    it('should return all categories', async () => {
      // Arrange
      await mockRepo.create({ name: 'Test Category 1' });
      await mockRepo.create({ name: 'Test Category 2' });

      // Act
      const categories = await viewModel.getAllCategories();

      // Assert
      expect(categories).toHaveLength(2);
      expect(categories[0].name).toBe('Test Category 1');
      expect(categories[1].name).toBe('Test Category 2');
    });

    it('should return empty array when no categories exist', async () => {
      // Act
      const categories = await viewModel.getAllCategories();

      // Assert
      expect(categories).toEqual([]);
    });
  });

  describe('createCategory', () => {
    it('should create category with valid data', async () => {
      // Act
      const category = await viewModel.createCategory('Food', 'Food expenses', '#ff0000');

      // Assert
      expect(category.name).toBe('Food');
      expect(category.description).toBe('Food expenses');
      expect(category.color).toBe('#ff0000');
    });

    it('should throw error when name is empty', async () => {
      // Act & Assert
      await expect(viewModel.createCategory('')).rejects.toThrow('カテゴリ名は必須です');
    });
  });

  describe('validateCategoryName', () => {
    it('should return no errors for valid name', () => {
      // Act
      const errors = viewModel.validateCategoryName('Valid Name');

      // Assert
      expect(errors).toEqual([]);
    });

    it('should return error for empty name', () => {
      // Act
      const errors = viewModel.validateCategoryName('');

      // Assert
      expect(errors).toContain('カテゴリ名は必須です');
    });
  });
});