# EchoLearn React Native App

This is the React Native mobile application for EchoLearn, an English learning app.

## Prerequisites

- Node.js (LTS version recommended)
- npm or Yarn
- Watchman (for macOS users)
- React Native CLI (`npx react-native --version`)
- An Android Emulator/Device or iOS Simulator/Device
- Backend server from the `EchoLearn` project running.

## Setup Instructions

1.  **Clone the Repository (if applicable) or Create Project:**
    If you haven't initialized a React Native project yet:
    ```bash
    npx react-native init EchoLearnApp
    cd EchoLearnApp
    ```
    Then, you can integrate the files provided by the assistant into this structure.
    If the `EchoLearnApp` directory is already provided with a `package.json`:
    ```bash
    cd EchoLearnApp 
    ```

2.  **Install Dependencies:**
    Navigate to the `EchoLearnApp` directory if you aren't already there:
    ```bash
    # If you prefer npm
    npm install
    # OR if you prefer yarn
    yarn install
    ```

3.  **Configure Backend URL:**
    Open `src/api/apiClient.js`.
    Modify the `API_BASE_URL` to point to your running backend:
    - For Android Emulator (if backend is on `localhost:8080` on your PC): `http://10.0.2.2:8080/api/v1`
    - For iOS Simulator (if backend is on `localhost:8080` on your PC): `http://localhost:8080/api/v1`
    - For physical device: `http://<YOUR_COMPUTER_NETWORK_IP>:8080/api/v1` (Ensure your backend is accessible on your local network).

4.  **Run on iOS:**
    ```bash
    npx react-native run-ios
    # or
    yarn ios 
    ```

5.  **Run on Android:**
    Ensure you have an emulator running or a device connected.
    ```bash
    npx react-native run-android
    # or
    yarn android
    ```

## Project Structure (within EchoLearnApp)

```
EchoLearnApp/
├── android/
├── ios/
├── src/
│   ├── api/
│   │   ├── apiClient.js         // Axios instance and base URL
│   │   ├── userService.js
│   │   ├── questionService.js
│   │   └── sessionService.js
│   ├── components/              // Reusable UI components
│   │   ├── QuestionDisplay.js
│   │   └── OptionButton.js
│   ├── navigation/
│   │   └── AppNavigator.js      // React Navigation setup
│   ├── screens/                 // App screens
│   │   ├── CreateUserScreen.js
│   │   ├── HomeScreen.js
│   │   ├── QuizScreen.js
│   │   └── SummaryScreen.js
│   ├── context/                 // React Context for global state
│   │   ├── UserContext.js
│   │   └── SessionContext.js
│   └── App.js                   // Main app component using contexts and navigator
├── app.json
├── babel.config.js
├── index.js                     // Entry point
├── metro.config.js
├── package.json
└── ... (other React Native boilerplate)
``` 