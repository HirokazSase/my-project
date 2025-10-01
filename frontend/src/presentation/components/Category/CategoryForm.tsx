import React, { useState, useEffect } from 'react';
import { Category } from '../../../domain/models/Category';
import { ErrorMessage } from '../Common/ErrorMessage';

interface CategoryFormProps {
  category?: Category | null;
  isEditing?: boolean;
  loading?: boolean;
  errors?: string[] | null;
  onSubmit: (data: CategoryFormData) => void;
  onCancel: () => void;
}

export interface CategoryFormData {
  name: string;
  description: string;
  color: string;
}

const DEFAULT_COLORS = [
  '#3b82f6', // Blue
  '#ef4444', // Red
  '#10b981', // Green
  '#f59e0b', // Yellow
  '#8b5cf6', // Purple
  '#06b6d4', // Cyan
  '#f97316', // Orange
  '#84cc16', // Lime
  '#ec4899', // Pink
  '#6b7280', // Gray
];

export const CategoryForm: React.FC<CategoryFormProps> = ({
  category,
  isEditing = false,
  loading = false,
  errors,
  onSubmit,
  onCancel,
}) => {
  const [formData, setFormData] = useState<CategoryFormData>({
    name: '',
    description: '',
    color: DEFAULT_COLORS[0],
  });

  // 編集時の初期値設定
  useEffect(() => {
    if (isEditing && category) {
      setFormData({
        name: category.name,
        description: category.description || '',
        color: category.color || DEFAULT_COLORS[0],
      });
    }
  }, [isEditing, category]);

  const handleInputChange = (
    field: keyof CategoryFormData,
    value: string
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

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h2 style={styles.title}>
          {isEditing ? 'カテゴリを編集' : '新しいカテゴリを作成'}
        </h2>
      </div>

      {errors && <ErrorMessage message={errors} />}

      <form onSubmit={handleSubmit} style={styles.form}>
        {/* カテゴリ名入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="name">
            カテゴリ名 <span style={styles.required}>*</span>
          </label>
          <input
            type="text"
            id="name"
            style={styles.input}
            value={formData.name}
            onChange={(e) => handleInputChange('name', e.target.value)}
            maxLength={50}
            required
            placeholder="例: 食費、交通費、光熱費"
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
            maxLength={200}
            rows={3}
            placeholder="このカテゴリの詳細な説明を入力してください"
          />
        </div>

        {/* カラー選択 */}
        <div style={styles.field}>
          <label style={styles.label}>
            カラー
          </label>
          <div style={styles.colorGrid}>
            {DEFAULT_COLORS.map((color) => (
              <button
                key={color}
                type="button"
                style={{
                  ...styles.colorOption,
                  backgroundColor: color,
                  border: formData.color === color ? '3px solid #111827' : '2px solid #e5e7eb',
                }}
                onClick={() => handleInputChange('color', color)}
                title={`カラー: ${color}`}
              />
            ))}
          </div>
          <div style={styles.colorPreview}>
            <span style={styles.previewLabel}>選択中のカラー:</span>
            <div
              style={{
                ...styles.previewColor,
                backgroundColor: formData.color,
              }}
            />
            <span style={styles.previewCode}>{formData.color}</span>
          </div>
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
            disabled={loading || !formData.name.trim()}
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
  textarea: {
    width: '100%',
    padding: '0.75rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    fontSize: '1rem',
    resize: 'vertical' as const,
    boxSizing: 'border-box' as const,
  } as React.CSSProperties,
  colorGrid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(5, 1fr)',
    gap: '0.5rem',
    marginBottom: '1rem',
  } as React.CSSProperties,
  colorOption: {
    width: '40px',
    height: '40px',
    borderRadius: '8px',
    cursor: 'pointer',
    outline: 'none',
  } as React.CSSProperties,
  colorPreview: {
    display: 'flex',
    alignItems: 'center',
    gap: '0.5rem',
    padding: '0.5rem',
    backgroundColor: '#f9fafb',
    borderRadius: '6px',
  } as React.CSSProperties,
  previewLabel: {
    fontSize: '0.9rem',
    color: '#6b7280',
  } as React.CSSProperties,
  previewColor: {
    width: '24px',
    height: '24px',
    borderRadius: '4px',
    border: '1px solid #e5e7eb',
  } as React.CSSProperties,
  previewCode: {
    fontSize: '0.8rem',
    fontFamily: 'monospace',
    color: '#4b5563',
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