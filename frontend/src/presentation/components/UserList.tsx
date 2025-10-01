import React from 'react';
import { User } from '../../domain/models/User';
import { LoadingSpinner } from './Common/LoadingSpinner';
import { ErrorMessage } from './Common/ErrorMessage';

interface UserListProps {
  users: User[];
  currentUser: User | null;
  loading?: boolean;
  error?: string | null;
  onUserSelect: (user: User) => void;
  onUserEdit: (user: User) => void;
  onUserDelete: (userId: string) => void;
  onCreateNew: () => void;
}

export const UserList: React.FC<UserListProps> = ({
  users,
  currentUser,
  loading = false,
  error,
  onUserSelect,
  onUserEdit,
  onUserDelete,
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
        <h2 style={styles.title}>ユーザー管理</h2>
        <button 
          onClick={onCreateNew} 
          style={styles.createButton}
        >
          新しいユーザーを作成
        </button>
      </div>

      {users.length === 0 ? (
        <div style={styles.emptyState}>
          <p style={styles.emptyMessage}>ユーザーがありません</p>
          <button 
            onClick={onCreateNew} 
            style={styles.emptyButton}
          >
            最初のユーザーを作成
          </button>
        </div>
      ) : (
        <div style={styles.grid}>
          {users.map((user) => (
            <div 
              key={user.id} 
              style={{
                ...styles.card,
                ...(currentUser && currentUser.id === user.id ? styles.cardActive : {}),
              }}
            >
              <div style={styles.cardHeader}>
                <h3 style={styles.userName}>{user.name}</h3>
                {currentUser && currentUser.id === user.id && (
                  <span style={styles.currentBadge}>現在のユーザー</span>
                )}
              </div>
              
              <p style={styles.userEmail}>{user.email}</p>
              
              <div style={styles.cardFooter}>
                <span style={styles.createdDate}>
                  作成日: {user.createdAt.toLocaleDateString('ja-JP')}
                </span>
                <div style={styles.actions}>
                  {(!currentUser || currentUser.id !== user.id) && (
                    <button
                      onClick={() => onUserSelect(user)}
                      style={styles.selectButton}
                    >
                      選択
                    </button>
                  )}
                  <button
                    onClick={() => onUserEdit(user)}
                    style={styles.editButton}
                  >
                    編集
                  </button>
                  <button
                    onClick={() => {
                      if (window.confirm('このユーザーを削除してもよろしいですか？')) {
                        onUserDelete(user.id);
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
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  cardActive: {
    borderColor: '#3b82f6',
    backgroundColor: '#eff6ff',
  } as React.CSSProperties,
  cardHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: '1rem',
  } as React.CSSProperties,
  userName: {
    margin: 0,
    fontSize: '1.1rem',
    fontWeight: '600',
    color: '#111827',
  } as React.CSSProperties,
  currentBadge: {
    padding: '0.25rem 0.5rem',
    backgroundColor: '#3b82f6',
    color: 'white',
    fontSize: '0.75rem',
    borderRadius: '4px',
    fontWeight: '500',
  } as React.CSSProperties,
  userEmail: {
    fontSize: '0.9rem',
    color: '#6b7280',
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
  selectButton: {
    padding: '0.5rem 1rem',
    border: '1px solid #3b82f6',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#3b82f6',
    fontSize: '0.8rem',
    cursor: 'pointer',
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