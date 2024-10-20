using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;

namespace EchoLearn.Api.Services
{
    public class SpeechRecognizerFactory : ISpeechRecognizerFactory
    {
        public SpeechRecognizer CreateSpeechRecognizer(SpeechConfig speechConfig, AudioConfig audioConfig)
        {
            return new SpeechRecognizer(speechConfig, audioConfig);
        }
    }
}