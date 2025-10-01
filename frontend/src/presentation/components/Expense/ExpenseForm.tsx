import React, { useState, useEffect } from 'react';
import { Category } from '../../../domain/models/Category';
import { Expense } from '../../../domain/models/Expense';
import { ErrorMessage } from '../Common/ErrorMessage';

interface ExpenseFormProps {
  categories: Category[];
  expense?: Expense | null;
  isEditing?: boolean;
  loading?: boolean;
  errors?: string[] | null;
  onSubmit: (data: ExpenseFormData) => void;
  onCancel: () => void;
}

export interface ExpenseFormData {
  categoryId: string;
  amount: number;
  title: string;
  description: string;
  date: Date;
}

export const ExpenseForm: React.FC<ExpenseFormProps> = ({
  categories,
  expense,
  isEditing = false,
  loading = false,
  errors,
  onSubmit,
  onCancel,
}) => {
  const [formData, setFormData] = useState<ExpenseFormData>({
    categoryId: '',
    amount: 0,
    title: '',
    description: '',
    date: new Date(),
  });

  // 編集時の初期値設定
  useEffect(() => {
    if (isEditing && expense) {
      setFormData({
        categoryId: expense.categoryId,
        amount: expense.amount,
        title: expense.title,
        description: expense.description,
        date: expense.date,
      });
    }
  }, [isEditing, expense]);

  const handleInputChange = (
    field: keyof ExpenseFormData,
    value: string | number | Date
  ) => {
    setFormData(prev => ({
      ...prev,
      [field]: value,
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData);
  };

  const formatDateForInput = (date: Date): string => {
    return date.toISOString().split('T')[0];
  };

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h2 style={styles.title}>
          {isEditing ? '経費を編集' : '新しい経費を作成'}
        </h2>
      </div>

      {errors && <ErrorMessage message={errors} />}

      <form onSubmit={handleSubmit} style={styles.form}>
        {/* カテゴリ選択 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="category">
            カテゴリ <span style={styles.required}>*</span>
          </label>
          <select
            id="category"
            style={styles.select}
            value={formData.categoryId}
            onChange={(e) => handleInputChange('categoryId', e.target.value)}
            required
          >
            <option value="">カテゴリを選択してください</option>
            {categories.map(category => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
          </select>
        </div>

        {/* 金額入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="amount">
            金額 <span style={styles.required}>*</span>
          </label>
          <input
            type="number"
            id="amount"
            style={styles.input}
            value={formData.amount || ''}
            onChange={(e) => handleInputChange('amount', parseFloat(e.target.value) || 0)}
            min="0"
            step="1"
            required
          />
        </div>

        {/* タイトル入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="title">
            タイトル <span style={styles.required}>*</span>
          </label>
          <input
            type="text"
            id="title"
            style={styles.input}
            value={formData.title}
            onChange={(e) => handleInputChange('title', e.target.value)}
            maxLength={100}
            required
          />
        </div>

        {/* 説明入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="description">
            説明
          </label>
          <textarea
            id="description"
            style={styles.textarea}
            value={formData.description}
            onChange={(e) => handleInputChange('description', e.target.value)}
            maxLength={500}
            rows={3}
          />
        </div>

        {/* 日付入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="date">
            日付 <span style={styles.required}>*</span>
          </label>
          <input
            type="date"
            id="date"
            style={styles.input}
            value={formatDateForInput(formData.date)}
            onChange={(e) => handleInputChange('date', new Date(e.target.value))}
            required
          />
        </div>

        {/* ボタン */}
        <div style={styles.buttons}>
          <button
            type="button"
            onClick={onCancel}
            style={styles.cancelButton}
            disabled={loading}
          >
            キャンセル
          </button>
          <button
            type="submit"
            style={styles.submitButton}
            disabled={loading}
          >
            {loading ? '保存中...' : (isEditing ? '更新' : '作成')}
          </button>
        </div>
      </form>
    </div>
  );
};

const styles = {
  container: {
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
    overflow: 'hidden',
  } as React.CSSProperties,
  header: {
    padding: '1.5rem',
    borderBottom: '1px solid #e5e7eb',
    backgroundColor: '#f9fafb',
  } as React.CSSProperties,
  title: {
    margin: 0,
    fontSize: '1.25rem',
    fontWeight: 'bold',
    color: '#111827',
  } as React.CSSProperties,
  form: {
    padding: '1.5rem',
  } as React.CSSProperties,
  field: {
    marginBottom: '1.5rem',
  } as React.CSSProperties,
  label: {
    display: 'block',
    marginBottom: '0.5rem',
    fontSize: '0.9rem',
    fontWeight: '600',
    color: '#374151',
  } as React.CSSProperties,
  required: {
    color: '#ef4444',
  } as React.CSSProperties,
  input: {
    width: '100%',
    padding: '0.75rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    fontSize: '1rem',
    boxSizing: 'border-box' as const,
  } as React.CSSProperties,
  select: {
    width: '100%',
    padding: '0.75rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    fontSize: '1rem',
    backgroundColor: 'white',
    boxSizing: 'border-box' as const,
  } as React.CSSProperties,
  textarea: {
    width: '100%',
    padding: '0.75rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    fontSize: '1rem',
    resize: 'vertical' as const,
    boxSizing: 'border-box' as const,
  } as React.CSSProperties,
  buttons: {
    display: 'flex',
    gap: '1rem',
    justifyContent: 'flex-end',
    paddingTop: '1rem',
  } as React.CSSProperties,
  cancelButton: {
    padding: '0.75rem 1.5rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    backgroundColor: 'white',
    color: '#374151',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
  } as React.CSSProperties,
  submitButton: {
    padding: '0.75rem 1.5rem',
    border: 'none',
    borderRadius: '6px',
    backgroundColor: '#3b82f6',
    color: 'white',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
  } as React.CSSProperties,
};