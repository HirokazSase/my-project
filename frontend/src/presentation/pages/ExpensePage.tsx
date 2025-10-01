import React, { useState, useEffect } from 'react';
import { ExpenseList } from '../components/Expense/ExpenseList';
import { ExpenseForm, ExpenseFormData } from '../components/Expense/ExpenseForm';
import { Expense } from '../../domain/models/Expense';
import { Category } from '../../domain/models/Category';
import { ExpenseViewModel } from '../../application/viewModels/ExpenseViewModel';
import { LoadingSpinner } from '../components/Common/LoadingSpinner';
import { ErrorMessage } from '../components/Common/ErrorMessage';

interface ExpensePageProps {
  userId: string;
  categories: Category[];
}

type ExpenseStatus = 'pending' | 'approved' | 'rejected';

export const ExpensePage: React.FC<ExpensePageProps> = ({
  userId,
  categories,
}) => {
  const [expenses, setExpenses] = useState<Expense[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [currentStatus, setCurrentStatus] = useState<ExpenseStatus>('pending');
  
  // Form states
  const [showExpenseForm, setShowExpenseForm] = useState(false);
  const [editingExpense, setEditingExpense] = useState<Expense | null>(null);
  
  // ViewModel
  const [expenseViewModel] = useState(() => new ExpenseViewModel());

  // 経費一覧を取得
  useEffect(() => {
    if (userId) {
      loadExpenses();
    }
  }, [userId, currentStatus]);

  const loadExpenses = async () => {
    try {
      setLoading(true);
      setError(null);
      const expenseList = await expenseViewModel.getExpensesByUserAndStatus(userId, currentStatus);
      setExpenses(expenseList);
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の取得に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleCreateExpense = async (data: ExpenseFormData) => {
    try {
      setLoading(true);
      setError(null);
      const newExpense = await expenseViewModel.createExpense(
        userId,
        data.categoryId,
        data.amount,
        'JPY', // デフォルト通貨
        data.title,
        data.description,
        data.date
      );
      
      // 現在のステータスがpendingの場合のみリストに追加
      if (currentStatus === 'pending') {
        setExpenses(prev => [newExpense, ...prev]);
      }
      
      setShowExpenseForm(false);
      setEditingExpense(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の作成に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateExpense = async (data: ExpenseFormData) => {
    if (!editingExpense) return;
    
    try {
      setLoading(true);
      setError(null);
      const updatedExpense = await expenseViewModel.updateExpense(
        editingExpense.id,
        data.categoryId,
        data.amount,
        'JPY',
        data.title,
        data.description,
        data.date
      );
      
      setExpenses(prev => 
        prev.map(expense => 
          expense.id === updatedExpense.id ? updatedExpense : expense
        )
      );
      
      setShowExpenseForm(false);
      setEditingExpense(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の更新に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteExpense = async (expenseId: string) => {
    try {
      setLoading(true);
      setError(null);
      await expenseViewModel.deleteExpense(expenseId);
      setExpenses(prev => prev.filter(expense => expense.id !== expenseId));
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の削除に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleApproveExpense = async (expenseId: string) => {
    try {
      setLoading(true);
      setError(null);
      await expenseViewModel.approveExpense(expenseId);
      
      // リストから削除（承認済みのステータスに移動）
      setExpenses(prev => prev.filter(expense => expense.id !== expenseId));
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の承認に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleRejectExpense = async (expenseId: string) => {
    try {
      setLoading(true);
      setError(null);
      await expenseViewModel.rejectExpense(expenseId);
      
      // リストから削除（却下済みのステータスに移動）
      setExpenses(prev => prev.filter(expense => expense.id !== expenseId));
    } catch (err) {
      setError(err instanceof Error ? err.message : '経費の却下に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const getStatusLabel = (status: ExpenseStatus): string => {
    switch (status) {
      case 'pending':
        return '承認待ち';
      case 'approved':
        return '承認済み';
      case 'rejected':
        return '却下済み';
    }
  };

  if (categories.length === 0) {
    return (
      <div style={styles.emptyState}>
        <h2>カテゴリがありません</h2>
        <p>経費を作成する前に、まずカテゴリを作成してください。</p>
      </div>
    );
  }

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h1 style={styles.title}>経費管理</h1>
        <div style={styles.statusTabs}>
          {(['pending', 'approved', 'rejected'] as ExpenseStatus[]).map((status) => (
            <button
              key={status}
              onClick={() => setCurrentStatus(status)}
              style={{
                ...styles.statusTab,
                ...(currentStatus === status ? styles.statusTabActive : {}),
              }}
            >
              {getStatusLabel(status)}
            </button>
          ))}
        </div>
        <button
          onClick={() => setShowExpenseForm(true)}
          style={styles.createButton}
          disabled={loading}
        >
          新しい経費を作成
        </button>
      </div>

      {error && <ErrorMessage message={error} onRetry={loadExpenses} />}

      {showExpenseForm ? (
        <ExpenseForm
          categories={categories}
          expense={editingExpense}
          isEditing={!!editingExpense}
          loading={loading}
          errors={error ? [error] : null}
          onSubmit={editingExpense ? handleUpdateExpense : handleCreateExpense}
          onCancel={() => {
            setShowExpenseForm(false);
            setEditingExpense(null);
            setError(null);
          }}
        />
      ) : (
        <>
          {loading && <LoadingSpinner message="経費を読み込んでいます..." />}
          <ExpenseList
            expenses={expenses}
            status={currentStatus}
            onExpenseSelect={(expense) => console.log('Selected:', expense)}
            onExpenseEdit={(expense) => {
              setEditingExpense(expense);
              setShowExpenseForm(true);
            }}
            onExpenseDelete={handleDeleteExpense}
            onExpenseApprove={currentStatus === 'pending' ? handleApproveExpense : undefined}
            onExpenseReject={currentStatus === 'pending' ? handleRejectExpense : undefined}
          />
        </>
      )}
    </div>
  );
};

const styles = {
  container: {
    maxWidth: '1200px',
    margin: '0 auto',
  } as React.CSSProperties,
  header: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: '2rem',
    gap: '1rem',
    flexWrap: 'wrap' as const,
  } as React.CSSProperties,
  title: {
    margin: 0,
    fontSize: '2rem',
    fontWeight: 'bold',
    color: '#111827',
  } as React.CSSProperties,
  statusTabs: {
    display: 'flex',
    gap: '0.5rem',
    flex: '1 1 auto',
    justifyContent: 'center',
  } as React.CSSProperties,
  statusTab: {
    padding: '0.5rem 1rem',
    border: '1px solid #d1d5db',
    borderRadius: '6px',
    backgroundColor: 'white',
    color: '#6b7280',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  statusTabActive: {
    backgroundColor: '#3b82f6',
    borderColor: '#3b82f6',
    color: 'white',
  } as React.CSSProperties,
  createButton: {
    padding: '0.75rem 1.5rem',
    border: 'none',
    borderRadius: '6px',
    backgroundColor: '#10b981',
    color: 'white',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'background-color 0.2s ease-in-out',
  } as React.CSSProperties,
  emptyState: {
    textAlign: 'center' as const,
    padding: '3rem',
    color: '#6b7280',
  } as React.CSSProperties,
};