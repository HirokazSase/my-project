import React from 'react';
import { Category } from '../../../domain/models/Category';
import { LoadingSpinner } from '../Common/LoadingSpinner';
import { ErrorMessage } from '../Common/ErrorMessage';

interface CategoryListProps {
  categories: Category[];
  loading?: boolean;
  error?: string | null;
  onCategoryEdit: (category: Category) => void;
  onCategoryDelete: (categoryId: string) => void;
  onCreateNew: () => void;
}

export const CategoryList: React.FC<CategoryListProps> = ({
  categories,
  loading = false,
  error,
  onCategoryEdit,
  onCategoryDelete,
  onCreateNew,
}) => {
  if (loading) {
    return (
      <div style={styles.container}>
        <LoadingSpinner />
      </div>
    );
  }

  if (error) {
    return (
      <div style={styles.container}>
        <ErrorMessage message={error} />
      </div>
    );
  }

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h2 style={styles.title}>カテゴリ管理</h2>
        <button 
          onClick={onCreateNew} 
          style={styles.createButton}
        >
          新しいカテゴリを作成
        </button>
      </div>

      {categories.length === 0 ? (
        <div style={styles.emptyState}>
          <p style={styles.emptyMessage}>カテゴリがありません</p>
          <button 
            onClick={onCreateNew} 
            style={styles.emptyButton}
          >
            最初のカテゴリを作成
          </button>
        </div>
      ) : (
        <div style={styles.grid}>
          {categories.map((category) => (
            <div key={category.id} style={styles.card}>
              <div style={styles.cardHeader}>
                <h3 style={styles.categoryName}>{category.name}</h3>
                {category.color && (
                  <div 
                    style={{
                      ...styles.colorIndicator,
                      backgroundColor: category.color,
                    }}
                  />
                )}
              </div>
              
              {category.description && (
                <p style={styles.categoryDescription}>
                  {category.description}
                </p>
              )}
              
              <div style={styles.cardFooter}>
                <span style={styles.createdDate}>
                  作成日: {category.createdAt.toLocaleDateString('ja-JP')}
                </span>
                <div style={styles.actions}>
                  <button
                    onClick={() => onCategoryEdit(category)}
                    style={styles.editButton}
                  >
                    編集
                  </button>
                  <button
                    onClick={() => {
                      if (window.confirm('このカテゴリを削除してもよろしいですか？')) {
                        onCategoryDelete(category.id);
                      }
                    }}
                    style={styles.deleteButton}
                  >
                    削除
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
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
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
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
  createButton: {
    padding: '0.5rem 1rem',
    border: 'none',
    borderRadius: '6px',
    backgroundColor: '#3b82f6',
    color: 'white',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
  } as React.CSSProperties,
  emptyState: {
    textAlign: 'center' as const,
    padding: '3rem',
  } as React.CSSProperties,
  emptyMessage: {
    fontSize: '1.1rem',
    color: '#6b7280',
    marginBottom: '1rem',
  } as React.CSSProperties,
  emptyButton: {
    padding: '0.75rem 1.5rem',
    border: 'none',
    borderRadius: '6px',
    backgroundColor: '#3b82f6',
    color: 'white',
    fontSize: '1rem',
    fontWeight: '500',
    cursor: 'pointer',
  } as React.CSSProperties,
  grid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))',
    gap: '1rem',
    padding: '1.5rem',
  } as React.CSSProperties,
  card: {
    border: '1px solid #e5e7eb',
    borderRadius: '8px',
    padding: '1.5rem',
    backgroundColor: 'white',
  } as React.CSSProperties,
  cardHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: '1rem',
  } as React.CSSProperties,
  categoryName: {
    margin: 0,
    fontSize: '1.1rem',
    fontWeight: '600',
    color: '#111827',
  } as React.CSSProperties,
  colorIndicator: {
    width: '24px',
    height: '24px',
    borderRadius: '50%',
    border: '2px solid #e5e7eb',
  } as React.CSSProperties,
  categoryDescription: {
    fontSize: '0.9rem',
    color: '#6b7280',
    lineHeight: '1.5',
    marginBottom: '1rem',
  } as React.CSSProperties,
  cardFooter: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
  } as React.CSSProperties,
  createdDate: {
    fontSize: '0.8rem',
    color: '#9ca3af',
  } as React.CSSProperties,
  actions: {
    display: 'flex',
    gap: '0.5rem',
  } as React.CSSProperties,
  editButton: {
    padding: '0.5rem 1rem',
    border: '1px solid #d1d5db',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#374151',
    fontSize: '0.8rem',
    cursor: 'pointer',
  } as React.CSSProperties,
  deleteButton: {
    padding: '0.5rem 1rem',
    border: '1px solid #ef4444',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#ef4444',
    fontSize: '0.8rem',
    cursor: 'pointer',
  } as React.CSSProperties,
};