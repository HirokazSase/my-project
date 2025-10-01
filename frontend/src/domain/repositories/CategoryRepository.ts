import { Category } from '../models/Category';

export interface CategoryRepository {
  getAll(): Promise<Category[]>;
  getById(id: string): Promise<Category | null>;
  create(data: { name: string; description?: string; color?: string }): Promise<Category>;
  update(id: string, data: { name: string; description?: string; color?: string }): Promise<Category>;
  delete(id: string): Promise<void>;
}