import { Category } from '../../domain/models/Category';
import { CategoryRepository } from '../../domain/repositories/CategoryRepository';
import { ApiClient } from './ApiClient';

export interface CategoryApiData {
  id: string;
  name: string;
  description: string;
  color: string;
  createdAt: string;
  updatedAt: string;
}

export class CategoryApiService implements CategoryRepository {
  private apiClient: ApiClient;

  constructor() {
    this.apiClient = new ApiClient();
  }

  private mapToCategory(data: CategoryApiData): Category {
    return new Category(
      data.id,
      data.name,
      data.description,
      data.color,
      new Date(data.createdAt),
      new Date(data.updatedAt)
    );
  }

  async getAll(): Promise<Category[]> {
    const response = await this.apiClient.get<CategoryApiData[]>('/categories');
    
    if (response.error) {
      throw new Error(response.error);
    }

    return response.data?.map(data => this.mapToCategory(data)) || [];
  }

  async getById(id: string): Promise<Category | null> {
    const response = await this.apiClient.get<CategoryApiData>(`/categories/${id}`);
    
    if (response.error) {
      if (response.status === 404) {
        return null;
      }
      throw new Error(response.error);
    }

    return response.data ? this.mapToCategory(response.data) : null;
  }

  async create(data: { name: string; description?: string; color?: string }): Promise<Category> {
    const requestData = {
      name: data.name,
      description: data.description || '',
      color: data.color || '#3b82f6',
    };

    const response = await this.apiClient.post<CategoryApiData>('/categories', requestData);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToCategory(response.data);
  }

  async update(id: string, data: { name: string; description?: string; color?: string }): Promise<Category> {
    const requestData = {
      name: data.name,
      description: data.description || '',
      color: data.color || '#3b82f6',
    };

    const response = await this.apiClient.put<CategoryApiData>(`/categories/${id}`, requestData);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToCategory(response.data);
  }

  async delete(id: string): Promise<void> {
    const response = await this.apiClient.delete(`/categories/${id}`);
    
    if (response.error) {
      throw new Error(response.error);
    }
  }
}