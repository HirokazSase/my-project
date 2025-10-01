import React from 'react';

interface ErrorMessageProps {
  message: string | string[] | null;
  onRetry?: () => void;
}

export const ErrorMessage: React.FC<ErrorMessageProps> = ({ 
  message, 
  onRetry 
}) => {
  if (!message) return null;

  const messages = Array.isArray(message) ? message : [message];

  return (
    <div style={styles.container}>
      <div style={styles.iconContainer}>
        <span style={styles.icon}>⚠️</span>
      </div>
      <div style={styles.content}>
        <h3 style={styles.title}>エラーが発生しました</h3>
        <ul style={styles.messageList}>
          {messages.map((msg, index) => (
            <li key={index} style={styles.messageItem}>
              {msg}
            </li>
          ))}
        </ul>
        {onRetry && (
          <button 
            onClick={onRetry} 
            style={styles.retryButton}
          >
            再試行
          </button>
        )}
      </div>
    </div>
  );
};

const styles = {
  container: {
    display: 'flex',
    alignItems: 'flex-start',
    padding: '1rem',
    backgroundColor: '#fee2e2',
    border: '1px solid #fecaca',
    borderRadius: '6px',
    color: '#991b1b',
    marginBottom: '1rem',
  } as React.CSSProperties,
  iconContainer: {
    marginRight: '0.75rem',
    flexShrink: 0,
  } as React.CSSProperties,
  icon: {
    fontSize: '1.25rem',
  } as React.CSSProperties,
  content: {
    flex: 1,
  } as React.CSSProperties,
  title: {
    margin: '0 0 0.5rem 0',
    fontSize: '1rem',
    fontWeight: '600',
    color: '#991b1b',
  } as React.CSSProperties,
  messageList: {
    listStyle: 'none',
    margin: 0,
    padding: 0,
  } as React.CSSProperties,
  messageItem: {
    marginBottom: '0.25rem',
    fontSize: '0.9rem',
    lineHeight: '1.4',
  } as React.CSSProperties,
  retryButton: {
    marginTop: '0.75rem',
    padding: '0.5rem 1rem',
    border: '1px solid #dc2626',
    borderRadius: '4px',
    backgroundColor: 'white',
    color: '#dc2626',
    fontSize: '0.875rem',
    fontWeight: '500',
    cursor: 'pointer',
  } as React.CSSProperties,
};