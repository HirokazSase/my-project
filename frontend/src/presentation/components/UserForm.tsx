import React, { useState, useEffect } from 'react';
import { User } from '../../domain/models/User';
import { ErrorMessage } from './Common/ErrorMessage';

interface UserFormProps {
  user?: User | null;
  isEditing?: boolean;
  loading?: boolean;
  errors?: string[] | null;
  onSubmit: (data: UserFormData) => void;
  onCancel: () => void;
}

export interface UserFormData {
  name: string;
  email: string;
}

export const UserForm: React.FC<UserFormProps> = ({
  user,
  isEditing = false,
  loading = false,
  errors,
  onSubmit,
  onCancel,
}) => {
  const [formData, setFormData] = useState<UserFormData>({
    name: '',
    email: '',
  });

  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  // 編集時の初期値設定
  useEffect(() => {
    if (isEditing && user) {
      setFormData({
        name: user.name,
        email: user.email,
      });
    }
  }, [isEditing, user]);

  const handleInputChange = (
    field: keyof UserFormData,
    value: string
  ) => {
    setFormData(prev => ({
      ...prev,
      [field]: value,
    }));
    
    // Clear field error when user starts typing
    if (formErrors[field]) {
      setFormErrors(prev => ({
        ...prev,
        [field]: '',
      }));
    }
  };

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    // Name validation
    if (!formData.name.trim()) {
      newErrors.name = 'ユーザー名は必須です';
    } else if (formData.name.trim().length < 2) {
      newErrors.name = 'ユーザー名は2文字以上で入力してください';
    } else if (formData.name.trim().length > 50) {
      newErrors.name = 'ユーザー名は50文字以内で入力してください';
    }

    // Email validation
    if (!formData.email.trim()) {
      newErrors.email = 'メールアドレスは必須です';
    } else {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailRegex.test(formData.email.trim())) {
        newErrors.email = '有効なメールアドレスを入力してください';
      } else if (formData.email.trim().length > 100) {
        newErrors.email = 'メールアドレスは100文字以内で入力してください';
      }
    }

    setFormErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm()) {
      return;
    }
    
    onSubmit({
      name: formData.name.trim(),
      email: formData.email.trim(),
    });
  };

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h2 style={styles.title}>
          {isEditing ? 'ユーザーを編集' : '新しいユーザーを作成'}
        </h2>
      </div>

      {errors && <ErrorMessage message={errors} />}

      <form onSubmit={handleSubmit} style={styles.form}>
        {/* ユーザー名入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="name">
            ユーザー名 <span style={styles.required}>*</span>
          </label>
          <input
            type="text"
            id="name"
            style={{
              ...styles.input,
              ...(formErrors.name ? styles.inputError : {}),
            }}
            value={formData.name}
            onChange={(e) => handleInputChange('name', e.target.value)}
            maxLength={50}
            required
            placeholder="例: 田中太郎"
          />
          {formErrors.name && (
            <span style={styles.fieldError}>{formErrors.name}</span>
          )}
        </div>

        {/* メールアドレス入力 */}
        <div style={styles.field}>
          <label style={styles.label} htmlFor="email">
            メールアドレス <span style={styles.required}>*</span>
          </label>
          <input
            type="email"
            id="email"
            style={{
              ...styles.input,
              ...(formErrors.email ? styles.inputError : {}),
            }}
            value={formData.email}
            onChange={(e) => handleInputChange('email', e.target.value)}
            maxLength={100}
            required
            placeholder="例: tanaka@example.com"
          />
          {formErrors.email && (
            <span style={styles.fieldError}>{formErrors.email}</span>
          )}
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
            disabled={loading || !formData.name.trim() || !formData.email.trim()}
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
    transition: 'border-color 0.2s ease-in-out',
  } as React.CSSProperties,
  inputError: {
    borderColor: '#ef4444',
  } as React.CSSProperties,
  fieldError: {
    display: 'block',
    marginTop: '0.5rem',
    fontSize: '0.875rem',
    color: '#ef4444',
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