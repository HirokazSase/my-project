import { Expense } from '../models/Expense';

export interface ExpenseRepository {
  getAll(): Promise<Expense[]>;
  getById(id: string): Promise<Expense | null>;
  getByUserId(userId: string): Promise<Expense[]>;
  getByUserAndStatus(userId: string, status: string): Promise<Expense[]>;
  create(data: {
    userId: string;
    categoryId: string;
    amount: number;
    currency: string;
    title: string;
    description: string;
    date: Date;
  }): Promise<Expense>;
  update(id: string, data: {
    categoryId: string;
    amount: number;
    currency: string;
    title: string;
    description: string;
    date: Date;
  }): Promise<Expense>;
  updateStatus(id: string, status: string): Promise<Expense>;
  delete(id: string): Promise<void>;
}