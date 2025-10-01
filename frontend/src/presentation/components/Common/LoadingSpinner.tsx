import React from 'react';

interface LoadingSpinnerProps {
  size?: 'small' | 'medium' | 'large';
  message?: string;
}

export const LoadingSpinner: React.FC<LoadingSpinnerProps> = ({ 
  size = 'medium', 
  message 
}) => {
  const getSpinnerSize = () => {
    switch (size) {
      case 'small':
        return { width: '16px', height: '16px', borderWidth: '2px' };
      case 'large':
        return { width: '32px', height: '32px', borderWidth: '4px' };
      default:
        return { width: '24px', height: '24px', borderWidth: '3px' };
    }
  };

  const spinnerSize = getSpinnerSize();

  return (
    <div style={styles.container} data-testid="loading-spinner">
      <div 
        style={{
          ...styles.spinner,
          ...spinnerSize,
        }}
      />
      {message && <p style={styles.message}>{message}</p>}
    </div>
  );
};

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column' as const,
    alignItems: 'center',
    justifyContent: 'center',
    padding: '2rem',
  } as React.CSSProperties,
  spinner: {
    border: '3px solid rgba(59, 130, 246, 0.3)',
    borderRadius: '50%',
    borderTop: '3px solid #3b82f6',
    animation: 'spin 1s linear infinite',
  } as React.CSSProperties,
  message: {
    marginTop: '1rem',
    fontSize: '0.9rem',
    color: '#6b7280',
    textAlign: 'center' as const,
  } as React.CSSProperties,
};

// Add keyframes for spin animation using CSS-in-JS
const styleSheet = document.styleSheets[0];
const keyframes = `
  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
`;

try {
  styleSheet.insertRule(keyframes, styleSheet.cssRules.length);
} catch (e) {
  // Ignore if rule already exists
}