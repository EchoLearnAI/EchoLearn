import { createGlobalStyle } from 'styled-components';

const GlobalStyle = createGlobalStyle`
  *,
  *::before,
  *::after {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
  }

  html {
    font-size: 16px; // Base font size, aids responsive typography with rem units
    scroll-behavior: smooth;
  }

  body {
    font-family: ${props => props.theme.fonts.main};
    color: ${props => props.theme.colors.text};
    background-color: ${props => props.theme.colors.background};
    line-height: 1.6;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    min-height: 100vh;
    display: flex;
    flex-direction: column;    
  }

  #root {
    flex: 1;
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 100vw; /* Ensure it doesn't exceed viewport width */
  }

  h1, h2, h3, h4, h5, h6 {
    font-family: ${props => props.theme.fonts.headings};
    font-weight: bold;
    margin-bottom: ${props => props.theme.spacings.medium};
  }

  p {
    margin-bottom: ${props => props.theme.spacings.medium};
  }

  a {
    color: ${props => props.theme.colors.primary};
    text-decoration: none;
    &:hover {
      text-decoration: underline;
    }
  }

  img, picture, video, canvas, svg {
    display: block;
    max-width: 100%;
  }

  input, button, textarea, select {
    font: inherit;
  }

  button {
    cursor: pointer;
    border: none;
    background-color: transparent;
  }

  ul, ol {
    list-style: none;
  }

  // Responsive container utility
  .container {
    width: 100%;
    max-width: 1200px; // Max width of the content area
    margin-left: auto;
    margin-right: auto;
    padding-left: ${props => props.theme.spacings.medium}; // Default padding
    padding-right: ${props => props.theme.spacings.medium};

    // Adapts padding for smaller screens
    ${props => props.theme.media.sm`
      padding-left: ${props.theme.spacings.large};
      padding-right: ${props.theme.spacings.large};
    `}
  }

  // Accessibility: focus outline
  *:focus-visible {
    outline: 2px solid ${props => props.theme.colors.primary};
    outline-offset: 2px;
  }
`;

export default GlobalStyle; 