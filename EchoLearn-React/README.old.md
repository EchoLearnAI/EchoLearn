# EchoLearn React Web Frontend

This project is the React web frontend for the EchoLearn application.

## Prerequisites

- Node.js (v16 or later recommended)
- npm or yarn

## Backend API

This frontend consumes the API from the EchoLearn Go backend. Ensure the backend server is running (typically on `http://localhost:8080`). The API base URL is configured in `src/api/apiClient.js`.

## Available Scripts

In the project directory, you can run:

### `npm start` or `yarn start`

Runs the app in development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm test` or `yarn test`

Launches the test runner in interactive watch mode.

### `npm run build` or `yarn run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

### `npm run eject` or `yarn run eject`

**Note: this is a one-way operation. Once you `eject`, you can't go back!**

If you aren't satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you're on your own.

## Project Structure

- `public/`: Contains static assets and the `index.html` template.
- `src/`: Contains the application's source code.
  - `api/`: Modules for interacting with the backend API (Axios client, service files).
  - `assets/`: Static assets like images, fonts.
  - `components/`: Reusable UI components.
    - `common/`: General-purpose components (buttons, inputs, layout elements).
    - `layout/`: Components defining the overall page structure (Header, Footer, Navbar).
    - `quiz/`: Components specific to the quiz functionality.
  - `contexts/`: React Context API for global state management (UserContext, SessionContext).
  - `hooks/`: Custom React hooks.
  - `pages/`: Top-level page components (HomePage, QuizPage, LoginPage, etc.).
  - `routes/`: Routing configuration.
  - `services/`: (Alternative to `api/` or for business logic not directly tied to API calls).
  - `styles/`: Global styles, themes, and CSS modules.
  - `utils/`: Utility functions.
  - `App.js`: Main application component, sets up routing and providers.
  - `index.js`: Entry point of the application.
  - `reportWebVitals.js`: For measuring performance.
  - `setupTests.js`: Jest setup.

## Responsiveness

The application is designed to be responsive for mobile resolutions. This is achieved using:
- CSS Flexbox and Grid
- Media Queries
- Relative units (%, vw, vh, rem, em)
- Mobile-first design approach where appropriate.

## Next Steps
- Implement API client and services.
- Set up User and Session contexts.
- Create page components for user creation, home, quiz, and summary.
- Implement routing.
- Style components for a good user experience on mobile devices. 