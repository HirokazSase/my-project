import { Expense } from '../../domain/models/Expense';
import { ExpenseRepository } from '../../domain/repositories/ExpenseRepository';
import { ApiClient } from './ApiClient';

export interface ExpenseApiData {
  id: string;
  userId: string;
  categoryId: string;
  amount: number;
  currency: string;
  title: string;
  description: string;
  date: string;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export class ExpenseApiService implements ExpenseRepository {
  private apiClient: ApiClient;

  constructor() {
    this.apiClient = new ApiClient();
  }

  private mapToExpense(data: ExpenseApiData): Expense {
    return new Expense(
      data.id,
      data.userId,
      data.categoryId,
      data.amount,
      data.currency,
      data.title,
      data.description,
      new Date(data.date),
      data.status,
      new Date(data.createdAt),
      new Date(data.updatedAt)
    );
  }

  async getAll(): Promise<Expense[]> {
    const response = await this.apiClient.get<ExpenseApiData[]>('/expenses');
    
    if (response.error) {
      throw new Error(response.error);
    }

    return response.data?.map(data => this.mapToExpense(data)) || [];
  }

  async getById(id: string): Promise<Expense | null> {
    const response = await this.apiClient.get<ExpenseApiData>(`/expenses/${id}`);
    
    if (response.error) {
      if (response.status === 404) {
        return null;
      }
      throw new Error(response.error);
    }

    return response.data ? this.mapToExpense(response.data) : null;
  }

  async getByUserId(userId: string): Promise<Expense[]> {
    const response = await this.apiClient.get<ExpenseApiData[]>(`/users/${userId}/expenses`);
    
    if (response.error) {
      throw new Error(response.error);
    }

    return response.data?.map(data => this.mapToExpense(data)) || [];
  }

  async getByUserAndStatus(userId: string, status: string): Promise<Expense[]> {
    const response = await this.apiClient.get<ExpenseApiData[]>(`/users/${userId}/expenses?status=${status}`);
    
    if (response.error) {
      throw new Error(response.error);
    }

    return response.data?.map(data => this.mapToExpense(data)) || [];
  }

  async create(data: {
    userId: string;
    categoryId: string;
    amount: number;
    currency: string;
    title: string;
    description: string;
    date: Date;
  }): Promise<Expense> {
    const requestData = {
      user_id: data.userId,
      category_id: data.categoryId,
      amount: data.amount,
      currency: data.currency,
      title: data.title,
      description: data.description,
      date: data.date.toISOString().split('T')[0],
    };

    const response = await this.apiClient.post<ExpenseApiData>('/expenses', requestData);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToExpense(response.data);
  }

  async update(id: string, data: {
    categoryId: string;
    amount: number;
    currency: string;
    title: string;
    description: string;
    date: Date;
  }): Promise<Expense> {
    const requestData = {
      category_id: data.categoryId,
      amount: data.amount,
      currency: data.currency,
      title: data.title,
      description: data.description,
      date: data.date.toISOString().split('T')[0],
    };

    const response = await this.apiClient.put<ExpenseApiData>(`/expenses/${id}`, requestData);
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToExpense(response.data);
  }

  async updateStatus(id: string, status: string): Promise<Expense> {
    const response = await this.apiClient.put<ExpenseApiData>(`/expenses/${id}/status`, {
      status: status,
    });
    
    if (response.error) {
      throw new Error(response.error);
    }

    if (!response.data) {
      throw new Error('No data received from server');
    }

    return this.mapToExpense(response.data);
  }

  async delete(id: string): Promise<void> {
    const response = await this.apiClient.delete(`/expenses/${id}`);
    
    if (response.error) {
      throw new Error(response.error);
    }
  }
}