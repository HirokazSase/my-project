import { User } from '../../domain/models/User';
import { UserRepository } from '../../domain/repositories/UserRepository';
import { ApiClient } from './ApiClient';

export interface UserApiData {
  id: string;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

export class UserApiService implements UserRepository {
  private apiClient: ApiClient;

  constructor() {
    this.apiClient = new ApiClient();
  }

  private mapToUser(data: UserApiData): User {
    return new User(
      data.id,
      data.name,
      data.email,
      new Date(data.createdAt),
      new Date(data.updatedAt)
    );
  }

  async getAll(): Promise<User[]> {
    const response = await this.apiClient.get<UserApiData[]>('/users');
    
    if (response.error) {
      throw new Error(response.error);
    }

    return response.data?.map(data => this.mapToUser(data)) || [];
  }

  async getById(id: string): Promise<User | null> {
    const response = await this.apiClient.get<UserApiData>(`/users/${id}`);
    
    if (response.error) {
      if (response.status === 404) {
        return null;
      }
      throw new Error(response.error);
    }

    return response.data ? this.mapToUser(response.data) : null;
  }

  async create(data: { name: string; email: string }): Promise<User> {
    const response = await this.apiClient.post<UserApiData>('/users', data);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToUser(response.data);
  }

  async update(id: string, data: { name: string; email: string }): Promise<User> {
    const response = await this.apiClient.put<UserApiData>(`/users/${id}`, data);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToUser(response.data);
  }

  async delete(id: string): Promise<void> {
    const response = await this.apiClient.delete(`/users/${id}`);
    
    if (response.error) {
      throw new Error(response.error);
    }
  }
}