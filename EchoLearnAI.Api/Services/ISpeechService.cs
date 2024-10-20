using Microsoft.CognitiveServices.Speech.Audio;

namespace EchoLearnAI.Api.Services
{
    public interface ISpeechService
    {
        Task<string> RecognizeSpeechAsync(AudioInputStream audioStream);
    }
}