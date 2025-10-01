import React from 'react';
import { Header } from './Header';
import { User } from '../../../domain/models/User';

interface LayoutProps {
  currentUser: User | null;
  onUserChange: (user: User) => void;
  users: User[];
  currentPage: string;
  onPageChange: (page: string) => void;
  children: React.ReactNode;
}

export const Layout: React.FC<LayoutProps> = ({
  currentUser,
  onUserChange,
  users,
  currentPage,
  onPageChange,
  children,
}) => {
  return (
    <div style={styles.layout}>
      <Header
        currentUser={currentUser}
        onUserChange={onUserChange}
        users={users}
        currentPage={currentPage}
        onPageChange={onPageChange}
      />
      <main style={styles.main}>
        <div style={styles.container}>
          {children}
        </div>
      </main>
    </div>
  );
};

const styles = {
  layout: {
    minHeight: '100vh',
    backgroundColor: '#f5f5f5',
  } as React.CSSProperties,
  main: {
    paddingTop: '80px', // Account for fixed header
  } as React.CSSProperties,
  container: {
    maxWidth: '1200px',
    margin: '0 auto',
    padding: '2rem 1rem',
  } as React.CSSProperties,
};