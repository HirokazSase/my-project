import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { CategoryList } from '../CategoryList';
import { Category } from '../../../../domain/models/Category';

const mockCategories: Category[] = [
  new Category(
    'cat-1',
    'Food',
    'Food expenses',
    '#ff0000',
    new Date('2024-01-01'),
    new Date('2024-01-01')
  ),
  new Category(
    'cat-2',
    'Transport',
    'Transportation expenses',
    '#00ff00',
    new Date('2024-01-02'),
    new Date('2024-01-02')
  ),
];

describe('CategoryList', () => {
  const mockProps = {
    categories: mockCategories,
    onCategoryEdit: jest.fn(),
    onCategoryDelete: jest.fn(),
    onCreateNew: jest.fn(),
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('should render category list with categories', () => {
    render(<CategoryList {...mockProps} />);

    expect(screen.getByText('カテゴリ管理')).toBeInTheDocument();
    expect(screen.getByText('Food')).toBeInTheDocument();
    expect(screen.getByText('Transport')).toBeInTheDocument();
    expect(screen.getByText('Food expenses')).toBeInTheDocument();
    expect(screen.getByText('Transportation expenses')).toBeInTheDocument();
  });

  it('should show loading state', () => {
    render(<CategoryList {...mockProps} categories={[]} loading={true} />);

    expect(screen.getByTestId('loading-spinner')).toBeInTheDocument();
  });

  it('should show error state', () => {
    render(<CategoryList {...mockProps} categories={[]} error="Test error" />);

    expect(screen.getByText('Test error')).toBeInTheDocument();
  });

  it('should show empty state when no categories', () => {
    render(<CategoryList {...mockProps} categories={[]} />);

    expect(screen.getByText('カテゴリがありません')).toBeInTheDocument();
    expect(screen.getByText('最初のカテゴリを作成')).toBeInTheDocument();
  });

  it('should call onCreateNew when create button is clicked', () => {
    render(<CategoryList {...mockProps} />);

    fireEvent.click(screen.getByText('新しいカテゴリを作成'));
    expect(mockProps.onCreateNew).toHaveBeenCalledTimes(1);
  });

  it('should call onCategoryEdit when edit button is clicked', () => {
    render(<CategoryList {...mockProps} />);

    const editButtons = screen.getAllByText('編集');
    fireEvent.click(editButtons[0]);

    expect(mockProps.onCategoryEdit).toHaveBeenCalledWith(mockCategories[0]);
  });

  it('should call onCategoryDelete when delete is confirmed', () => {
    // Mock window.confirm
    const originalConfirm = window.confirm;
    window.confirm = jest.fn(() => true);

    render(<CategoryList {...mockProps} />);

    const deleteButtons = screen.getAllByText('削除');
    fireEvent.click(deleteButtons[0]);

    expect(mockProps.onCategoryDelete).toHaveBeenCalledWith('cat-1');

    // Restore original confirm
    window.confirm = originalConfirm;
  });

  it('should not call onCategoryDelete when delete is cancelled', () => {
    // Mock window.confirm
    const originalConfirm = window.confirm;
    window.confirm = jest.fn(() => false);

    render(<CategoryList {...mockProps} />);

    const deleteButtons = screen.getAllByText('削除');
    fireEvent.click(deleteButtons[0]);

    expect(mockProps.onCategoryDelete).not.toHaveBeenCalled();

    // Restore original confirm
    window.confirm = originalConfirm;
  });
});