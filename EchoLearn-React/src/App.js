import React, { useContext } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, Outlet } from 'react-router-dom';
import { ThemeProvider } from 'styled-components';
import GlobalStyle from './styles/GlobalStyle';
import theme from './styles/theme';
import UserContext, { UserProvider } from './contexts/UserContext';
import { SessionProvider } from './contexts/SessionContext';

import HomePage from './pages/HomePage';
import CreateUserPage from './pages/CreateUserPage';
import LoginPage from './pages/LoginPage';
import QuizPage from './pages/QuizPage';
import SummaryPage from './pages/SummaryPage';
// import NotFoundPage from './pages/NotFoundPage'; 

// Layout for routes that require authentication
const ProtectedLayout = () => {
  const { user, isLoadingUser } = useContext(UserContext);

  if (isLoadingUser) {
    return <div>Loading user session...</div>; // Or a global loader
  }

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />; // Renders the child route element
};

// Layout for public routes like Login and Create User
// Redirects to home if user is already logged in
const PublicLayout = () => {
  const { user, isLoadingUser } = useContext(UserContext);

  if (isLoadingUser) {
    return <div>Loading...</div>; // Or a global loader
  }

  if (user) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />; // Renders the child route element (LoginPage or CreateUserPage)
};

function AppContent() {
  return (
    <Router>
      <GlobalStyle />
      <SessionProvider> {/* SessionProvider can stay here */}
        <Routes>
          {/* Public Routes (Login, Create User) */}
          <Route element={<PublicLayout />}>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/create-user" element={<CreateUserPage />} />
          </Route>

          {/* Protected Routes */}
          <Route element={<ProtectedLayout />}>
            <Route path="/" element={<HomePage />} />
            <Route path="/quiz" element={<QuizPage />} />
            <Route path="/summary/:sessionId" element={<SummaryPage />} />
            {/* Add other protected routes here */}
          </Route>

          {/* Fallback: if no routes match and user is not caught by public/protected layouts,
              it implies an issue or a need for a more specific 404 page. 
              For now, if not user, ProtectedLayout sends to /login.
              If user, and no path matches, redirect to home. Or show NotFoundPage.
           */}
          <Route path="*" element={<Navigate to="/" replace />} /> 
          {/* Or a more specific NotFoundPage: <Route path="*" element={<NotFoundPage />} /> */}
        </Routes>
      </SessionProvider>
    </Router>
  );
}

function App() {
  return (
    <ThemeProvider theme={theme}>
      <UserProvider>
        <AppContent />
      </UserProvider>
    </ThemeProvider>
  );
}

export default App;
