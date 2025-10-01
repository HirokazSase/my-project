import React, { useState, useEffect } from 'react';
import { Layout } from './presentation/components/Layout/Layout';
import { ExpensePage } from './presentation/pages/ExpensePage';
import { CategoryList } from './presentation/components/Category/CategoryList';
import { CategoryForm, CategoryFormData } from './presentation/components/Category/CategoryForm';
import { UserList } from './presentation/components/UserList';
import { UserForm, UserFormData } from './presentation/components/UserForm';
import { User } from './domain/models/User';
import { Category } from './domain/models/Category';
import { UserViewModel } from './application/viewModels/UserViewModel';
import { CategoryViewModel } from './application/viewModels/CategoryViewModel';
import './App.css';

type CurrentPage = 'expenses' | 'categories' | 'users';

function App() {
  const [currentPage, setCurrentPage] = useState<CurrentPage>('expenses');
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [users, setUsers] = useState<User[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  // Form states
  const [showUserForm, setShowUserForm] = useState(false);
  const [showCategoryForm, setShowCategoryForm] = useState(false);
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const [editingCategory, setEditingCategory] = useState<Category | null>(null);
  
  // ViewModels
  const [userViewModel] = useState(() => new UserViewModel());
  const [categoryViewModel] = useState(() => new CategoryViewModel());

  // ユーザー一覧を取得
  useEffect(() => {
    loadUsers();
  }, []);

  // 現在のユーザーのカテゴリを取得
  useEffect(() => {
    if (currentUser) {
      loadCategories();
    }
  }, [currentUser]);

  const loadUsers = async () => {
    try {
      setLoading(true);
      const userList = await userViewModel.getAllUsers();
      setUsers(userList);
      
      // 最初のユーザーを選択
      if (userList.length > 0 && !currentUser) {
        setCurrentUser(userList[0]);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const loadCategories = async () => {
    if (!currentUser) return;
    
    try {
      setLoading(true);
      const categoryList = await categoryViewModel.getAllCategories();
      setCategories(categoryList);
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  // User operations
  const handleCreateUser = async (data: UserFormData) => {
    try {
      setLoading(true);
      const newUser = await userViewModel.createUser(data.name, data.email);
      setUsers(prev => [...prev, newUser]);
      setShowUserForm(false);
      setCurrentUser(newUser);
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateUser = async (data: UserFormData) => {
    if (!editingUser) return;
    
    try {
      setLoading(true);
      const updatedUser = await userViewModel.updateUser(editingUser.id, data.name, data.email);
      setUsers(prev => prev.map(user => user.id === updatedUser.id ? updatedUser : user));
      setEditingUser(null);
      setShowUserForm(false);
      
      if (currentUser && currentUser.id === updatedUser.id) {
        setCurrentUser(updatedUser);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteUser = async (userId: string) => {
    try {
      setLoading(true);
      await userViewModel.deleteUser(userId);
      setUsers(prev => prev.filter(user => user.id !== userId));
      
      if (currentUser && currentUser.id === userId) {
        const remainingUsers = users.filter(user => user.id !== userId);
        setCurrentUser(remainingUsers.length > 0 ? remainingUsers[0] : null);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  // Category operations
  const handleCreateCategory = async (data: CategoryFormData) => {
    if (!currentUser) return;
    
    try {
      setLoading(true);
      const newCategory = await categoryViewModel.createCategory(
        data.name,
        data.description,
        data.color
      );
      setCategories(prev => [...prev, newCategory]);
      setShowCategoryForm(false);
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateCategory = async (data: CategoryFormData) => {
    if (!editingCategory) return;
    
    try {
      setLoading(true);
      const updatedCategory = await categoryViewModel.updateCategory(
        editingCategory.id,
        data.name,
        data.description,
        data.color
      );
      setCategories(prev => prev.map(cat => cat.id === updatedCategory.id ? updatedCategory : cat));
      setEditingCategory(null);
      setShowCategoryForm(false);
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteCategory = async (categoryId: string) => {
    try {
      setLoading(true);
      await categoryViewModel.deleteCategory(categoryId);
      setCategories(prev => prev.filter(cat => cat.id !== categoryId));
    } catch (err) {
      setError(err instanceof Error ? err.message : '予期しないエラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  const renderCurrentPage = () => {
    switch (currentPage) {
      case 'expenses':
        return currentUser ? (
          <ExpensePage 
            userId={currentUser.id} 
            categories={categories}
          />
        ) : (
          <div style={styles.emptyState}>
            <p>ユーザーを選択してください</p>
          </div>
        );
      
      case 'categories':
        if (showCategoryForm) {
          return (
            <CategoryForm
              category={editingCategory}
              isEditing={!!editingCategory}
              loading={loading}
              errors={error ? [error] : null}
              onSubmit={editingCategory ? handleUpdateCategory : handleCreateCategory}
              onCancel={() => {
                setShowCategoryForm(false);
                setEditingCategory(null);
                setError(null);
              }}
            />
          );
        }
        
        return (
          <CategoryList
            categories={categories}
            loading={loading}
            error={error}
            onCategoryEdit={(category) => {
              setEditingCategory(category);
              setShowCategoryForm(true);
            }}
            onCategoryDelete={handleDeleteCategory}
            onCreateNew={() => setShowCategoryForm(true)}
          />
        );
      
      case 'users':
        if (showUserForm) {
          return (
            <UserForm
              user={editingUser}
              isEditing={!!editingUser}
              loading={loading}
              errors={error ? [error] : null}
              onSubmit={editingUser ? handleUpdateUser : handleCreateUser}
              onCancel={() => {
                setShowUserForm(false);
                setEditingUser(null);
                setError(null);
              }}
            />
          );
        }
        
        return (
          <UserList
            users={users}
            currentUser={currentUser}
            loading={loading}
            error={error}
            onUserSelect={setCurrentUser}
            onUserEdit={(user) => {
              setEditingUser(user);
              setShowUserForm(true);
            }}
            onUserDelete={handleDeleteUser}
            onCreateNew={() => setShowUserForm(true)}
          />
        );
      
      default:
        return <div>Page not found</div>;
    }
  };

  return (
    <Layout
      currentUser={currentUser}
      onUserChange={setCurrentUser}
      users={users}
      currentPage={currentPage}
      onPageChange={(page) => setCurrentPage(page as CurrentPage)}
    >
      {renderCurrentPage()}
    </Layout>
  );
}

const styles = {
  emptyState: {
    textAlign: 'center' as const,
    padding: '3rem',
    color: '#6b7280',
  },
};

export default App;