namespace EchoLearnAI.Api.Services
{
    public interface ISpeechService
    {
        Task<string> RecognizeSpeechAsync(Stream audioStream);
    }
}