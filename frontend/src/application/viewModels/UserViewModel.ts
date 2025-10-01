import { User } from '../../domain/models/User';
import { UserRepository } from '../../domain/repositories/UserRepository';
import { UserApiService } from '../../infrastructure/api/UserApiService';

export class UserViewModel {
  private userRepository: UserRepository;

  constructor() {
    this.userRepository = new UserApiService();
  }

  async getAllUsers(): Promise<User[]> {
    try {
      return await this.userRepository.getAll();
    } catch (error) {
      console.error('Failed to fetch users:', error);
      throw new Error('ユーザーの取得に失敗しました');
    }
  }

  async getUserById(id: string): Promise<User | null> {
    try {
      return await this.userRepository.getById(id);
    } catch (error) {
      console.error('Failed to fetch user:', error);
      throw new Error('ユーザーの取得に失敗しました');
    }
  }

  async createUser(name: string, email: string): Promise<User> {
    try {
      // バリデーション
      const validationErrors = this.validateUser(name, email);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const userData = {
        name: name.trim(),
        email: email.trim().toLowerCase(),
      };

      return await this.userRepository.create(userData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to create user:', error);
      throw new Error('ユーザーの作成に失敗しました');
    }
  }

  async updateUser(id: string, name: string, email: string): Promise<User> {
    try {
      // バリデーション
      const validationErrors = this.validateUser(name, email);
      if (validationErrors.length > 0) {
        throw new Error(validationErrors.join(', '));
      }

      const userData = {
        name: name.trim(),
        email: email.trim().toLowerCase(),
      };

      return await this.userRepository.update(id, userData);
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      console.error('Failed to update user:', error);
      throw new Error('ユーザーの更新に失敗しました');
    }
  }

  async deleteUser(id: string): Promise<void> {
    try {
      await this.userRepository.delete(id);
    } catch (error) {
      console.error('Failed to delete user:', error);
      throw new Error('ユーザーの削除に失敗しました');
    }
  }

  validateUser(name: string, email: string): string[] {
    const errors: string[] = [];

    // 名前のバリデーション
    if (!name || name.trim().length === 0) {
      errors.push('名前は必須です');
    } else if (name.trim().length > 100) {
      errors.push('名前は100文字以内で入力してください');
    }

    // メールのバリデーション
    if (!email || email.trim().length === 0) {
      errors.push('メールアドレスは必須です');
    } else if (!this.isValidEmail(email)) {
      errors.push('有効なメールアドレスを入力してください');
    }

    return errors;
  }

  private isValidEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email.trim());
  }

  validateUserName(name: string): string[] {
    const errors: string[] = [];

    if (!name || name.trim().length === 0) {
      errors.push('名前は必須です');
    } else if (name.trim().length > 100) {
      errors.push('名前は100文字以内で入力してください');
    }

    return errors;
  }

  validateUserEmail(email: string): string[] {
    const errors: string[] = [];

    if (!email || email.trim().length === 0) {
      errors.push('メールアドレスは必須です');
    } else if (!this.isValidEmail(email)) {
      errors.push('有効なメールアドレスを入力してください');
    }

    return errors;
  }
}