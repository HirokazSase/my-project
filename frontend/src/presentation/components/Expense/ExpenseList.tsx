import React from 'react';
import { Expense } from '../../../domain/models/Expense';
import { ExpenseCard } from './ExpenseCard';
import { LoadingSpinner } from '../Common/LoadingSpinner';
import { ErrorMessage } from '../Common/ErrorMessage';

interface ExpenseListProps {
  expenses: Expense[];
  status?: string;
  loading?: boolean;
  error?: string | null;
  onExpenseSelect: (expense: Expense) => void;
  onExpenseEdit: (expense: Expense) => void;
  onExpenseDelete: (expenseId: string) => void;
  onExpenseApprove?: (expenseId: string) => void;
  onExpenseReject?: (expenseId: string) => void;
}

export const ExpenseList: React.FC<ExpenseListProps> = ({
  expenses,
  status = 'pending',
  loading = false,
  error,
  onExpenseSelect,
  onExpenseEdit,
  onExpenseDelete,
  onExpenseApprove,
  onExpenseReject,
}) => {
  if (loading) {
    return (
      <div style={styles.container}>
        <LoadingSpinner message="経費を読み込んでいます..." />
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

  if (expenses.length === 0) {
    return (
      <div style={styles.container}>
        <div style={styles.emptyState}>
          <h3 style={styles.emptyTitle}>経費がありません</h3>
          <p style={styles.emptyMessage}>
            {status === 'pending' && '承認待ちの経費がありません'}
            {status === 'approved' && '承認済みの経費がありません'}
            {status === 'rejected' && '却下された経費がありません'}
          </p>
        </div>
      </div>
    );
  }

  // 経費を日付順（新しい順）でソート
  const sortedExpenses = [...expenses].sort((a, b) => 
    new Date(b.date).getTime() - new Date(a.date).getTime()
  );

  // 月別にグループ化
  const expensesByMonth = sortedExpenses.reduce((groups, expense) => {
    const monthKey = expense.date.toLocaleDateString('ja-JP', { 
      year: 'numeric', 
      month: 'long' 
    });
    
    if (!groups[monthKey]) {
      groups[monthKey] = [];
    }
    groups[monthKey].push(expense);
    return groups;
  }, {} as Record<string, Expense[]>);

  // 月別の合計金額を計算
  const getMonthTotal = (expenses: Expense[]): number => {
    return expenses.reduce((total, expense) => total + expense.amount, 0);
  };

  return (
    <div style={styles.container}>
      <div style={styles.summary}>
        <div style={styles.summaryItem}>
          <span style={styles.summaryLabel}>総件数:</span>
          <span style={styles.summaryValue}>{expenses.length}件</span>
        </div>
        <div style={styles.summaryItem}>
          <span style={styles.summaryLabel}>総金額:</span>
          <span style={styles.summaryValue}>
            {new Intl.NumberFormat('ja-JP', {
              style: 'currency',
              currency: 'JPY',
            }).format(expenses.reduce((total, expense) => total + expense.amount, 0))}
          </span>
        </div>
      </div>

      <div style={styles.expenseList}>
        {Object.entries(expensesByMonth).map(([monthKey, monthExpenses]) => (
          <div key={monthKey} style={styles.monthGroup}>
            <div style={styles.monthHeader}>
              <h3 style={styles.monthTitle}>{monthKey}</h3>
              <span style={styles.monthTotal}>
                {new Intl.NumberFormat('ja-JP', {
                  style: 'currency',
                  currency: 'JPY',
                }).format(getMonthTotal(monthExpenses))} 
                ({monthExpenses.length}件)
              </span>
            </div>
            
            <div style={styles.monthExpenses}>
              {monthExpenses.map((expense) => (
                <ExpenseCard
                  key={expense.id}
                  expense={expense}
                  showActions={true}
                  showApprovalActions={status === 'pending'}
                  onClick={() => onExpenseSelect(expense)}
                  onEdit={() => onExpenseEdit(expense)}
                  onDelete={() => onExpenseDelete(expense.id)}
                  onApprove={onExpenseApprove ? () => onExpenseApprove(expense.id) : undefined}
                  onReject={onExpenseReject ? () => onExpenseReject(expense.id) : undefined}
                />
              ))}
            </div>
          </div>
        ))}
      </div>
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
  summary: {
    display: 'flex',
    justifyContent: 'space-around',
    padding: '1.5rem',
    borderBottom: '1px solid #e5e7eb',
    backgroundColor: '#f9fafb',
  } as React.CSSProperties,
  summaryItem: {
    textAlign: 'center' as const,
  } as React.CSSProperties,
  summaryLabel: {
    display: 'block',
    fontSize: '0.875rem',
    color: '#6b7280',
    marginBottom: '0.25rem',
  } as React.CSSProperties,
  summaryValue: {
    fontSize: '1.25rem',
    fontWeight: 'bold',
    color: '#111827',
  } as React.CSSProperties,
  expenseList: {
    padding: '1.5rem',
  } as React.CSSProperties,
  monthGroup: {
    marginBottom: '2rem',
  } as React.CSSProperties,
  monthHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: '1rem',
    paddingBottom: '0.5rem',
    borderBottom: '2px solid #e5e7eb',
  } as React.CSSProperties,
  monthTitle: {
    margin: 0,
    fontSize: '1.125rem',
    fontWeight: '600',
    color: '#374151',
  } as React.CSSProperties,
  monthTotal: {
    fontSize: '0.875rem',
    color: '#6b7280',
    fontWeight: '500',
  } as React.CSSProperties,
  monthExpenses: {
    display: 'grid',
    gap: '0.75rem',
  } as React.CSSProperties,
  emptyState: {
    textAlign: 'center' as const,
    padding: '3rem',
  } as React.CSSProperties,
  emptyTitle: {
    fontSize: '1.25rem',
    fontWeight: '600',
    color: '#374151',
    marginBottom: '0.5rem',
  } as React.CSSProperties,
  emptyMessage: {
    fontSize: '1rem',
    color: '#6b7280',
  } as React.CSSProperties,
};