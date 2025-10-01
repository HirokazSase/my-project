import { User } from '../models/User';

export interface UserRepository {
  getAll(): Promise<User[]>;
  getById(id: string): Promise<User | null>;
  create(data: { name: string; email: string }): Promise<User>;
  update(id: string, data: { name: string; email: string }): Promise<User>;
  delete(id: string): Promise<void>;
}