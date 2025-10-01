import { CategoryViewModel } from '../CategoryViewModel';
import { Category } from '../../../domain/models/Category';
import { CategoryApiService } from '../../../infrastructure/api/CategoryApiService';

// Mock the CategoryApiService module
jest.mock('../../../infrastructure/api/CategoryApiService');

describe('CategoryViewModel', () => {
  let viewModel: CategoryViewModel;
  let mockApiService: jest.Mocked<CategoryApiService>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();
    
    // Create a new instance of the mocked service
    mockApiService = new CategoryApiService() as jest.Mocked<CategoryApiService>;
    
    // Create view model
    viewModel = new CategoryViewModel();
  });

  describe('getAllCategories', () => {
    it('should return all categories', async () => {
      // Arrange
      const mockCategories = [
        new Category('1', 'Test Category 1', '', '#000000', new Date(), new Date()),
        new Category('2', 'Test Category 2', '', '#000000', new Date(), new Date()),
      ];
      
      mockApiService.getAll = jest.fn().mockResolvedValue(mockCategories);
      // @ts-ignore
      viewModel['categoryRepository'] = mockApiService;

      // Act
      const categories = await viewModel.getAllCategories();

      // Assert
      expect(categories).toHaveLength(2);
      expect(categories[0].name).toBe('Test Category 1');
      expect(categories[1].name).toBe('Test Category 2');
    });

    it('should return empty array when no categories exist', async () => {
      // Arrange
      mockApiService.getAll = jest.fn().mockResolvedValue([]);
      // @ts-ignore
      viewModel['categoryRepository'] = mockApiService;

      // Act
      const categories = await viewModel.getAllCategories();

      // Assert
      expect(categories).toEqual([]);
    });
  });

  describe('createCategory', () => {
    it('should create category with valid data', async () => {
      // Arrange
      const mockCategory = new Category('1', 'Food', 'Food expenses', '#ff0000', new Date(), new Date());
      mockApiService.create = jest.fn().mockResolvedValue(mockCategory);
      // @ts-ignore
      viewModel['categoryRepository'] = mockApiService;

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