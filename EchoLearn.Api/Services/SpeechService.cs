using System.Linq.Expressions;
using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;

namespace EchoLearn.Api.Services
{
    public class SpeechService : ISpeechService
    {
        private readonly ILogger<SpeechService> _logger;
        private readonly IConfiguration _configuration;
        private readonly ISpeechRecognizerFactory _speechRecognizerFactory;

        public SpeechService(ILogger<SpeechService> logger, IConfiguration configuration, ISpeechRecognizerFactory speechRecognizerFactory)
        {
            _logger = logger;
            _configuration = configuration;
            _speechRecognizerFactory = speechRecognizerFactory;
        }

        public async Task<string> RecognizeSpeechAsync(Stream audioStream)
        {
            try {
                _logger.LogInformation("Starting speech recognition.");

                var subscriptionKey = _configuration["SpeechService:SubscriptionKey"];
                var region = _configuration["SpeechService:Region"];
                var speechConfig = SpeechConfig.FromSubscription(subscriptionKey, region);

                // Create an AudioConfig from the stream
                var audioFormat = AudioStreamFormat.GetWaveFormatPCM(16000, 16, 1);
                var pushStream = AudioInputStream.CreatePushStream(audioFormat);

                // Read audio data from the and push it to the recognition service
                byte[] buffer = new byte[4096];
                int bytesRead;
                while ((bytesRead = await audioStream.ReadAsync(buffer, 0, buffer.Length)) > 0)
                {
                    byte[] actualBytes = new byte[bytesRead];
                    Array.Copy(buffer, actualBytes, bytesRead);
                    pushStream.Write(actualBytes);
                }
                pushStream.Close();

                var audioConfig = AudioConfig.FromStreamInput(pushStream);

                using var recognizer = _speechRecognizerFactory.CreateSpeechRecognizer(speechConfig, audioConfig);
                var result = await recognizer.RecognizeOnceAsync();

                if (result.Reason == ResultReason.RecognizedSpeech)
                {
                    _logger.LogInformation($"Recognized: {result.Text}");
                    return result.Text;
                } else {
                    _logger.LogWarning($"Recognition failed. Reason: {result.Reason}");
                    return $"Recognition failed. Reason: {result.Reason}";
                }
            } catch (Exception ex) {
                _logger.LogError(ex, "An error occurred during speech recognition.");
                return "An error occurred during speech recognition.";
            }
        }
    }
}