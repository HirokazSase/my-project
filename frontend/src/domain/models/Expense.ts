export type ExpenseStatus = 'pending' | 'approved' | 'rejected';

export class Expense {
  constructor(
    public readonly id: string,
    public readonly userId: string,
    public readonly categoryId: string,
    public readonly amount: number,
    public readonly currency: string,
    public readonly title: string,
    public readonly description: string,
    public readonly date: Date,
    public readonly status: string,
    public readonly createdAt: Date,
    public readonly updatedAt: Date
  ) {
    this.validate();
  }

  private validate(): void {
    if (!this.id || this.id.trim().length === 0) {
      throw new Error('Expense ID is required');
    }

    if (!this.userId || this.userId.trim().length === 0) {
      throw new Error('User ID is required');
    }

    if (!this.categoryId || this.categoryId.trim().length === 0) {
      throw new Error('Category ID is required');
    }

    if (!this.amount || this.amount <= 0) {
      throw new Error('Amount must be greater than 0');
    }

    if (this.amount > 10000000) {
      throw new Error('Amount must be 10,000,000 or less');
    }

    if (!this.currency || this.currency.trim().length === 0) {
      throw new Error('Currency is required');
    }

    if (!this.title || this.title.trim().length === 0) {
      throw new Error('Title is required');
    }

    if (this.title.length > 100) {
      throw new Error('Title must be 100 characters or less');
    }

    if (this.description && this.description.length > 500) {
      throw new Error('Description must be 500 characters or less');
    }

    if (!this.date) {
      throw new Error('Date is required');
    }

    if (!this.isValidStatus(this.status)) {
      throw new Error('Invalid status');
    }
  }

  private isValidStatus(status: string): boolean {
    return ['pending', 'approved', 'rejected'].includes(status);
  }

  public isPending(): boolean {
    return this.status === 'pending';
  }

  public isApproved(): boolean {
    return this.status === 'approved';
  }

  public isRejected(): boolean {
    return this.status === 'rejected';
  }

  public getStatusDisplay(): string {
    switch (this.status) {
      case 'pending':
        return '承認待ち';
      case 'approved':
        return '承認済み';
      case 'rejected':
        return '却下済み';
      default:
        return '不明';
    }
  }

  public getFormattedAmount(): string {
    return new Intl.NumberFormat('ja-JP', {
      style: 'currency',
      currency: 'JPY',
    }).format(this.amount);
  }

  public getFormattedDate(): string {
    return new Intl.DateTimeFormat('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    }).format(this.date);
  }

  public getDisplayInfo(): string {
    return `${this.title} - ${this.getFormattedAmount()} (${this.getFormattedDate()})`;
  }

  public equals(other: Expense): boolean {
    return this.id === other.id;
  }
}