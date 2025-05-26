import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { ThemeProvider } from 'styled-components';
import GlobalStyle from './styles/GlobalStyle';
import theme from './styles/theme';
import { UserProvider } from './contexts/UserContext';
import { SessionProvider } from './contexts/SessionContext';

// Page Components
import HomePage from './pages/HomePage';
import CreateUserPage from './pages/CreateUserPage';
import QuizPage from './pages/QuizPage';
import SummaryPage from './pages/SummaryPage';

// Placeholder pages (we will create these later in src/pages/)
// const HomePage = () => <div><h1>Home Page</h1><nav><Link to="/create-user">Create User</Link> | <Link to="/quiz">Start Quiz</Link></nav></div>; // Replaced
// const QuizPage = () => <div><h1>Quiz Page</h1><p>Quiz interface will be here.</p><Link to="/">Go Home</Link> | <Link to="/summary/123">Go to Summary (Test)</Link></div>; // Replaced
// const SummaryPage = () => <div><h1>Summary Page</h1><p>Session summary will be here.</p><Link to="/">Go Home</Link></div>; // Replaced
const NotFoundPage = () => (
  <div style={{ textAlign: 'center', marginTop: '50px' }}>
    <h1>404 - Page Not Found</h1>
    <p>Sorry, the page you are looking for does not exist.</p>
    <Link to="/">Go Home</Link>
  </div>
);

function App() {
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyle />
      <UserProvider>
        <SessionProvider>
          <Router>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/create-user" element={<CreateUserPage />} />
              <Route path="/quiz" element={<QuizPage />} />
              <Route path="/summary/:sessionId" element={<SummaryPage />} />
              <Route path="*" element={<NotFoundPage />} />
            </Routes>
          </Router>
        </SessionProvider>
      </UserProvider>
    </ThemeProvider>
  );
}

export default App;
