const theme = {
  colors: {
    primary: '#007bff', // Example primary color (blue)
    secondary: '#6c757d', // Example secondary color (gray)
    success: '#28a745', // Green for correct answers/success messages
    danger: '#dc3545', // Red for incorrect answers/error messages
    warning: '#ffc107', // Yellow for warnings
    info: '#17a2b8', // Light blue for informational messages
    light: '#f8f9fa', // Light gray for backgrounds
    dark: '#343a40', // Dark gray for text
    white: '#ffffff',
    black: '#000000',
    text: '#212529', // Default text color
    background: '#ffffff', // Default background color
    border: '#dee2e6', // Border color
    disabled: '#adb5bd', // Color for disabled elements
  },
  fonts: {
    main: 'Arial, sans-serif',
    headings: 'Helvetica, Arial, sans-serif',
  },
  fontSizes: {
    small: '0.8rem',
    medium: '1rem',
    large: '1.2rem',
    xlarge: '1.5rem',
    xxlarge: '2rem',
  },
  spacings: {
    xsmall: '0.25rem',
    small: '0.5rem',
    medium: '1rem',
    large: '1.5rem',
    xlarge: '2rem',
    xxlarge: '3rem',
  },
  breakpoints: {
    xs: '0px',
    sm: '576px',
    md: '768px',
    lg: '992px',
    xl: '1200px',
    xxl: '1400px',
  },
  shadows: {
    small: '0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24)',
    medium: '0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23)',
    large: '0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23)',
  },
  radii: {
    small: '4px',
    medium: '8px',
    large: '16px',
    round: '50%',
  },
  // Function to generate media queries easily
  media: {
    sm: (styles) => `@media (min-width: 576px) { ${styles} }`,
    md: (styles) => `@media (min-width: 768px) { ${styles} }`,
    lg: (styles) => `@media (min-width: 992px) { ${styles} }`,
    xl: (styles) => `@media (min-width: 1200px) { ${styles} }`,
  },
};

export default theme; 