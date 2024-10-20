using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;

namespace EchoLearn.Api.Services
{
    public interface ISpeechRecognizerFactory
    {
        SpeechRecognizer CreateSpeechRecognizer(SpeechConfig speechConfig, AudioConfig audioConfig);
    }
}