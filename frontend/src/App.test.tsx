import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import App from './App';

// Mock API calls
jest.mock('./infrastructure/api/UserApiService');
jest.mock('./infrastructure/api/CategoryApiService');
jest.mock('./infrastructure/api/ExpenseApiService');

describe('App Component', () => {
  it('renders without crashing', async () => {
    render(<App />);
    await waitFor(() => {
      expect(screen.getByText(/経費管理システム/i)).toBeInTheDocument();
    });
  });

  it('renders header with title', async () => {
    render(<App />);
    await waitFor(() => {
      expect(screen.getByText(/経費管理システム/i)).toBeInTheDocument();
    });
  });
});