import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { CategoryForm, CategoryFormData } from '../CategoryForm';
import { Category } from '../../../../domain/models/Category';

describe('CategoryForm', () => {
  const mockProps = {
    onSubmit: jest.fn(),
    onCancel: jest.fn(),
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('should render create form', () => {
    render(<CategoryForm {...mockProps} />);

    expect(screen.getByText('新しいカテゴリを作成')).toBeInTheDocument();
    expect(screen.getByLabelText(/カテゴリ名/)).toBeInTheDocument();
    expect(screen.getByLabelText(/説明/)).toBeInTheDocument();
    expect(screen.getByText('カラー')).toBeInTheDocument();
    expect(screen.getByText('作成')).toBeInTheDocument();
    expect(screen.getByText('キャンセル')).toBeInTheDocument();
  });

  it('should render edit form with existing category', () => {
    const category = new Category(
      'cat-1',
      'Food',
      'Food expenses',
      '#ff0000',
      new Date(),
      new Date()
    );

    render(<CategoryForm {...mockProps} category={category} isEditing={true} />);

    expect(screen.getByText('カテゴリを編集')).toBeInTheDocument();
    expect(screen.getByDisplayValue('Food')).toBeInTheDocument();
    expect(screen.getByDisplayValue('Food expenses')).toBeInTheDocument();
    expect(screen.getByText('更新')).toBeInTheDocument();
  });

  it('should handle form submission with valid data', async () => {
    render(<CategoryForm {...mockProps} />);

    // Fill out form
    fireEvent.change(screen.getByLabelText(/カテゴリ名/), {
      target: { value: 'Test Category' },
    });
    fireEvent.change(screen.getByLabelText(/説明/), {
      target: { value: 'Test description' },
    });

    // Submit form
    fireEvent.click(screen.getByText('作成'));

    await waitFor(() => {
      expect(mockProps.onSubmit).toHaveBeenCalledWith({
        name: 'Test Category',
        description: 'Test description',
        color: '#3b82f6', // Default color
      } as CategoryFormData);
    });
  });

  it('should handle cancel button click', () => {
    render(<CategoryForm {...mockProps} />);

    fireEvent.click(screen.getByText('キャンセル'));
    expect(mockProps.onCancel).toHaveBeenCalledTimes(1);
  });

  it('should show loading state', () => {
    render(<CategoryForm {...mockProps} loading={true} />);

    expect(screen.getByText('保存中...')).toBeInTheDocument();
    expect(screen.getByText('保存中...')).toBeDisabled();
    expect(screen.getByText('キャンセル')).toBeDisabled();
  });

  it('should show error messages', () => {
    const errors = ['カテゴリ名は必須です', '説明が長すぎます'];
    render(<CategoryForm {...mockProps} errors={errors} />);

    expect(screen.getByText('カテゴリ名は必須です')).toBeInTheDocument();
    expect(screen.getByText('説明が長すぎます')).toBeInTheDocument();
  });

  it('should handle color selection', () => {
    render(<CategoryForm {...mockProps} />);

    const colorButtons = screen.getAllByRole('button');
    const redColorButton = colorButtons.find(
      button => button.style.backgroundColor === 'rgb(239, 68, 68)' // #ef4444 in RGB
    );

    if (redColorButton) {
      fireEvent.click(redColorButton);
    }

    // Submit form to verify color was selected
    fireEvent.change(screen.getByLabelText(/カテゴリ名/), {
      target: { value: 'Test Category' },
    });
    fireEvent.click(screen.getByText('作成'));

    expect(mockProps.onSubmit).toHaveBeenCalledWith(
      expect.objectContaining({
        color: '#ef4444',
      })
    );
  });

  it('should prevent submission with empty name', () => {
    render(<CategoryForm {...mockProps} />);

    // Try to submit without filling name
    fireEvent.click(screen.getByText('作成'));

    // The submit button should be disabled or form should not submit
    expect(mockProps.onSubmit).not.toHaveBeenCalled();
  });
});