using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;

namespace EchoLearnAI.Api.Services
{
    public class SpeechService : ISpeechService
    {
        private readonly ILogger<SpeechService> _logger;
        private readonly IConfiguration _configuration;

        public SpeechService(ILogger<SpeechService> logger, IConfiguration configuration)
        {
            _logger = logger;
            _configuration = configuration;
        }

        public async Task<string> RecognizeSpeechAsync(AudioInputStream audioStream)
        {
            var speechConfig = SpeechConfig.FromSubscription(_configuration["SpeechService:SubscriptionKey"], _configuration["SpeechService:Region"]);
            var audioConfig = AudioConfig.FromStreamInput(audioStream);
            var recognizer = new SpeechRecognizer(speechConfig, audioConfig);

            var result = await recognizer.RecognizeOnceAsync();
            return result.Text;
        }
    }
}