import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

// Mock API calls
jest.mock('./infrastructure/api/UserApiService');
jest.mock('./infrastructure/api/CategoryApiService');
jest.mock('./infrastructure/api/ExpenseApiService');

describe('App Component', () => {
  it('renders without crashing', () => {
    render(<App />);
  });

  it('renders header with title', () => {
    render(<App />);
    expect(screen.getByText(/経費管理システム/i)).toBeInTheDocument();
  });
});