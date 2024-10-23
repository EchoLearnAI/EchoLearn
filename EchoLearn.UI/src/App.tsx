import React, { useState } from 'react';
import { OpenAIClient, AzureKeyCredential, ChatCompletionsOptions } from "@azure/openai";
import { SpeechRecognizer, AudioConfig, SpeechConfig, ResultReason } from 'microsoft-cognitiveservices-speech-sdk';
import './App.css';

const App: React.FC = () => {
  const [transcription, setTranscription] = useState<string>('');
  const [isListening, setIsListening] = useState<boolean>(false);

  // Replace these values with your Azure Speech Service credentials
  const azureSpeechKey: string = 'YOUR_AZURE_SPEECH_KEY';
  const azureRegion: string = 'YOUR_AZURE_REGION';
  const azureOpenAIEndpoint: string = 'YOUR_OPENAI_ENDPOINT';
  const azureOpenAIKey: string = 'YOUR_OPENAI_KEY';

  const client = new OpenAIClient(azureOpenAIEndpoint, new AzureKeyCredential(azureOpenAIKey));

  const handleStartListening = (): void => {
    if (isListening) return;

    setIsListening(true);

    const speechConfig = SpeechConfig.fromSubscription(azureSpeechKey, azureRegion);
    const audioConfig = AudioConfig.fromDefaultMicrophoneInput();

    const recognizer = new SpeechRecognizer(speechConfig, audioConfig);

    recognizer.recognizing = (s, e) => {
      console.log(`Recognizing: ${e.result.text}`);
      setTranscription(e.result.text);
    };

    recognizer.recognized = (s, e) => {
      if (e.result.reason === ResultReason.RecognizedSpeech) {
        console.log(`Recognized: ${e.result.text}`);
        setTranscription(e.result.text);
      } else if (e.result.reason === ResultReason.NoMatch) {
        console.log('No speech could be recognized.');
      }
    };

    recognizer.sessionStopped = (s, e) => {
      recognizer.stopContinuousRecognitionAsync();
      setIsListening(false);
    };

    recognizer.canceled = (s, e) => {
      console.error(`Canceled: ${e.errorDetails}`);
      recognizer.stopContinuousRecognitionAsync();
      setIsListening(false);
    };

    recognizer.startContinuousRecognitionAsync();
  };

  const handleStopListening = (): void => {
    setIsListening(false);
  };

  const handleSendToAI = async (): Promise<void> => {
    try {
      const options: ChatCompletionsOptions = {
        messages: [
          { role: "user", content: transcription }
        ],
        maxTokens: 100,
        deploymentName: "YOUR_DEPLOYMENT_NAME"
      };

      const response = await client.getChatCompletions(options);
      setTranscription(response.choices[0].message?.content || '');
    } catch (error) {
      console.error('Error sending message to AI:', error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Azure OpenAI Real-time Audio Chat</h1>
        <button onClick={handleStartListening} disabled={isListening}>
          {isListening ? 'Listening...' : 'Start Listening'}
        </button>
        <button onClick={handleStopListening} disabled={!isListening}>
          Stop Listening
        </button>
        <button onClick={handleSendToAI} disabled={!transcription}>
          Send to AI
        </button>
        <p>Transcription: {transcription}</p>
      </header>
    </div>
  );
};

export default App;
