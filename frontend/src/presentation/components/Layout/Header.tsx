import React from 'react';
import { User } from '../../../domain/models/User';

interface HeaderProps {
  currentUser: User | null;
  onUserChange: (user: User) => void;
  users: User[];
  currentPage: string;
  onPageChange: (page: string) => void;
}

export const Header: React.FC<HeaderProps> = ({
  currentUser,
  onUserChange,
  users,
  currentPage,
  onPageChange,
}) => {
  const navItems = [
    { key: 'expenses', label: '経費管理' },
    { key: 'categories', label: 'カテゴリ' },
    { key: 'users', label: 'ユーザー' },
  ];

  return (
    <header style={styles.header}>
      <div style={styles.container}>
        <div style={styles.brand}>
          <h1 style={styles.title}>経費管理システム</h1>
        </div>
        
        <nav style={styles.nav}>
          {navItems.map((item) => (
            <button
              key={item.key}
              onClick={() => onPageChange(item.key)}
              style={{
                ...styles.navButton,
                ...(currentPage === item.key ? styles.navButtonActive : {}),
              }}
            >
              {item.label}
            </button>
          ))}
        </nav>

        <div style={styles.userSection}>
          {currentUser && users.length > 1 && (
            <select
              value={currentUser.id}
              onChange={(e) => {
                const user = users.find(u => u.id === e.target.value);
                if (user) onUserChange(user);
              }}
              style={styles.userSelect}
            >
              {users.map(user => (
                <option key={user.id} value={user.id}>
                  {user.name}
                </option>
              ))}
            </select>
          )}
          {currentUser && (
            <div style={styles.currentUser}>
              <span style={styles.userName}>{currentUser.name}</span>
              <span style={styles.userEmail}>{currentUser.email}</span>
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

const styles = {
  header: {
    position: 'fixed' as const,
    top: 0,
    left: 0,
    right: 0,
    zIndex: 1000,
    backgroundColor: 'white',
    borderBottom: '1px solid #e5e7eb',
    boxShadow: '0 1px 3px rgba(0, 0, 0, 0.1)',
  } as React.CSSProperties,
  container: {
    maxWidth: '1200px',
    margin: '0 auto',
    padding: '0 1rem',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    height: '64px',
  } as React.CSSProperties,
  brand: {
    flex: '0 0 auto',
  } as React.CSSProperties,
  title: {
    margin: 0,
    fontSize: '1.5rem',
    fontWeight: 'bold',
    color: '#111827',
  } as React.CSSProperties,
  nav: {
    display: 'flex',
    gap: '0.5rem',
    flex: '1 1 auto',
    justifyContent: 'center',
  } as React.CSSProperties,
  navButton: {
    padding: '0.5rem 1rem',
    border: 'none',
    borderRadius: '6px',
    backgroundColor: 'transparent',
    color: '#6b7280',
    fontSize: '0.9rem',
    fontWeight: '500',
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
  } as React.CSSProperties,
  navButtonActive: {
    backgroundColor: '#3b82f6',
    color: 'white',
  } as React.CSSProperties,
  userSection: {
    display: 'flex',
    alignItems: 'center',
    gap: '1rem',
    flex: '0 0 auto',
  } as React.CSSProperties,
  userSelect: {
    padding: '0.5rem',
    border: '1px solid #d1d5db',
    borderRadius: '4px',
    fontSize: '0.9rem',
    backgroundColor: 'white',
  } as React.CSSProperties,
  currentUser: {
    display: 'flex',
    flexDirection: 'column' as const,
    alignItems: 'flex-end',
  } as React.CSSProperties,
  userName: {
    fontSize: '0.9rem',
    fontWeight: '600',
    color: '#111827',
  } as React.CSSProperties,
  userEmail: {
    fontSize: '0.8rem',
    color: '#6b7280',
  } as React.CSSProperties,
};