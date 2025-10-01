export class User {
  constructor(
    public readonly id: string,
    public readonly name: string,
    public readonly email: string,
    public readonly createdAt: Date,
    public readonly updatedAt: Date
  ) {
    this.validate();
  }

  private validate(): void {
    if (!this.id || this.id.trim().length === 0) {
      throw new Error('User ID is required');
    }

    if (!this.name || this.name.trim().length === 0) {
      throw new Error('User name is required');
    }

    if (this.name.length > 100) {
      throw new Error('User name must be 100 characters or less');
    }

    if (!this.email || this.email.trim().length === 0) {
      throw new Error('Email is required');
    }

    if (!this.isValidEmail(this.email)) {
      throw new Error('Invalid email format');
    }
  }

  private isValidEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }

  public getDisplayName(): string {
    return this.name;
  }

  public getDisplayInfo(): string {
    return `${this.name} (${this.email})`;
  }

  public equals(other: User): boolean {
    return this.id === other.id;
  }
}