import React from 'react';
import { Expense } from '../../../domain/models/Expense';

interface ExpenseCardProps {
  expense: Expense;
  showActions?: boolean;
  showApprovalActions?: boolean;
  onClick?: () => void;
  onEdit?: () => void;
  onDelete?: () => void;
  onApprove?: () => void;
  onReject?: () => void;
}

export const ExpenseCard: React.FC<ExpenseCardProps> = ({
  expense,
  showActions = true,
  showApprovalActions = false,
  onClick,
  onEdit,
  onDelete,
  onApprove,
  onReject,
}) => {
  const getStatusStyle = (status: string) => {
    switch (status) {
      case 'pending':
        return { backgroundColor: '#fef3c7', color: '#92400e' };
      case 'approved':
        return { backgroundColor: '#d1fae5', color: '#065f46' };
      case 'rejected':
        return { backgroundColor: '#fee2e2', color: '#991b1b' };
      default:
        return { backgroundColor: '#f3f4f6', color: '#374151' };
    }
  };

  const handleDeleteClick = () => {
    if (window.confirm('この経費を削除してもよろしいですか？')) {
      onDelete?.();
    }
  };

  const handleApproveClick = () => {
    if (window.confirm('この経費を承認してもよろしいですか？')) {
      onApprove?.();
    }
  };

  const handleRejectClick = () => {
    if (window.confirm('この経費を却下してもよろしいですか？')) {
      onReject?.();
    }
  };

  return (
    <div 
      style={{
        ...styles.card,
        ...(onClick ? styles.cardClickable : {}),
      }}
      onClick={onClick}
    >
      <div style={styles.cardHeader}>
        <div style={styles.titleSection}>
          <h4 style={styles.title}>{expense.title}</h4>
          <span 
            style={{
              ...styles.statusBadge,
              ...getStatusStyle(expense.status),
            }}
          >
            {expense.getStatusDisplay()}
          </span>
        </div>
        <div style={styles.amount}>
          {expense.getFormattedAmount()}
        </div>
      </div>

      {expense.description && (
        <p style={styles.description}>{expense.description}</p>
      )}

      <div style={styles.cardFooter}>
        <div style={styles.metadata}>
          <span style={styles.date}>{expense.getFormattedDate()}</span>
          <span style={styles.categoryId}>カテゴリID: {expense.categoryId}</span>
        </div>

        {showActions && (
          <div style={styles.actions}>
            {showApprovalActions && (
              <>
                <button
                  onClick={(e) => {
                    e.stopPropagation();
                    handleApproveClick();
                  }}
                  style={styles.approveButton}
                  disabled={!onApprove}
                >
                  承認
                </button>
                <button
                  onClick={(e) => {
                    e.stopPropagation();
                    handleRejectClick();
                  }}
                  style={styles.rejectButton}
                  disabled={!onReject}
                >
                  却下
                </button>
              </>
            )}
            <button
              onClick={(e) => {
                e.stopPropagation();
                onEdit?.();
              }}
              style={styles.editButton}
              disabled={!onEdit}
            >
              編集
            </button>
            <button
              onClick={(e) => {
                e.stopPropagation();
                handleDeleteClick();
              }}
              style={styles.deleteButton}
              disabled={!onDelete}
            >
              削除
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

const styles = {
  card: {
    border: '1px solid #e5e7eb',
    borderRadius: '8px',
    padding: '1rem',
    backgroundColor: 'white',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  cardClickable: {
    cursor: 'pointer',
  } as React.CSSProperties,
  cardHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: '0.75rem',
  } as React.CSSProperties,
  titleSection: {
    display: 'flex',
    flexDirection: 'column' as const,
    gap: '0.5rem',
    flex: 1,
  } as React.CSSProperties,
  title: {
    margin: 0,
    fontSize: '1rem',
    fontWeight: '600',
    color: '#111827',
  } as React.CSSProperties,
  statusBadge: {
    alignSelf: 'flex-start',
    padding: '0.25rem 0.5rem',
    borderRadius: '12px',
    fontSize: '0.75rem',
    fontWeight: '500',
    textTransform: 'uppercase' as const,
    letterSpacing: '0.05em',
  } as React.CSSProperties,
  amount: {
    fontSize: '1.25rem',
    fontWeight: 'bold',
    color: '#059669',
    textAlign: 'right' as const,
  } as React.CSSProperties,
  description: {
    margin: '0 0 0.75rem 0',
    fontSize: '0.875rem',
    color: '#6b7280',
    lineHeight: '1.4',
  } as React.CSSProperties,
  cardFooter: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    gap: '1rem',
  } as React.CSSProperties,
  metadata: {
    display: 'flex',
    flexDirection: 'column' as const,
    gap: '0.25rem',
  } as React.CSSProperties,
  date: {
    fontSize: '0.875rem',
    color: '#374151',
    fontWeight: '500',
  } as React.CSSProperties,
  categoryId: {
    fontSize: '0.75rem',
    color: '#9ca3af',
  } as React.CSSProperties,
  actions: {
    display: 'flex',
    gap: '0.5rem',
    flexWrap: 'wrap' as const,
  } as React.CSSProperties,
  approveButton: {
    padding: '0.375rem 0.75rem',
    border: '1px solid #10b981',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#10b981',
    fontSize: '0.75rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  rejectButton: {
    padding: '0.375rem 0.75rem',
    border: '1px solid #ef4444',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#ef4444',
    fontSize: '0.75rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  editButton: {
    padding: '0.375rem 0.75rem',
    border: '1px solid #6b7280',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#6b7280',
    fontSize: '0.75rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  deleteButton: {
    padding: '0.375rem 0.75rem',
    border: '1px solid #ef4444',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#ef4444',
    fontSize: '0.75rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
};