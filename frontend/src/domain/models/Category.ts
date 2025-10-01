export class Category {
  constructor(
    public readonly id: string,
    public readonly name: string,
    public readonly description: string,
    public readonly color: string,
    public readonly createdAt: Date,
    public readonly updatedAt: Date
  ) {
    this.validate();
  }

  private validate(): void {
    if (!this.id || this.id.trim().length === 0) {
      throw new Error('Category ID is required');
    }

    if (!this.name || this.name.trim().length === 0) {
      throw new Error('Category name is required');
    }

    if (this.name.length > 50) {
      throw new Error('Category name must be 50 characters or less');
    }

    if (this.description && this.description.length > 200) {
      throw new Error('Category description must be 200 characters or less');
    }

    if (this.color && !this.isValidColor(this.color)) {
      throw new Error('Invalid color format');
    }
  }

  private isValidColor(color: string): boolean {
    const colorRegex = /^#[0-9A-Fa-f]{6}$/;
    return colorRegex.test(color);
  }

  public getDisplayName(): string {
    return this.name;
  }

  public getDisplayInfo(): string {
    return this.description ? `${this.name} - ${this.description}` : this.name;
  }

  public equals(other: Category): boolean {
    return this.id === other.id;
  }
}