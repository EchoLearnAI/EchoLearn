using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;
using Microsoft.CognitiveServices.Speech.PronunciationAssessment;
using Microsoft.CognitiveServices.Speech;
using Microsoft.CognitiveServices.Speech.Audio;
using Microsoft.CognitiveServices.Speech.PronunciationAssessment;
using OpenAI;
using OpenAI.ObjectModels.RequestModels;
using OpenAI.ObjectModels;
using OpenAI.Chat;

class Program
{
    private static async Task Main(string[] args)
    {
        var subscriptionKey = "";
        var serviceRegion = "eastus2";
        var openAiApiKey = "";

        var pronunciationAssessmentConfig = new PronunciationAssessmentConfig("", GradingSystem.HundredMark, Granularity.Phoneme);

        var speechConfig = SpeechConfig.FromSubscription(subscriptionKey, serviceRegion);
        var audioConfig = AudioConfig.FromDefaultMicrophoneInput();

        var openAiApi = new OpenAIClient(new OpenAIAuthentication(openAiApiKey));
        var chatMessages = new List<ChatMessage>();

        Console.WriteLine("Say something in English...");

        while (true)
        {
            using var recognizer = new SpeechRecognizer(speechConfig, audioConfig);
            var result = await recognizer.RecognizeOnceAsync();

            if (result.Reason == ResultReason.RecognizedSpeech)
            {
                Console.WriteLine($"You: {result.Text}");
                pronunciationAssessmentConfig.ReferenceText = result.Text;
                pronunciationAssessmentConfig.ApplyTo(recognizer);

                chatMessages.Add(ChatMessage.FromUser(result.Text));
                var chatRequest = new ChatCompletionCreateRequest
                {
                    Messages = chatMessages,
                    Model = Models.Gpt_4
                };
                var response = await openAiApi.ChatCompletion.CreateCompletion(chatRequest);

                var aiResponse = response.Choices[0].Message.Content;
                Console.WriteLine($"AI: {aiResponse}");

                chatMessages.Add(ChatMessage.FromAssistant(aiResponse));

                Console.WriteLine("Let's continue the conversation...");
            }
            else if (result.Reason == ResultReason.NoMatch)
            {
                Console.WriteLine("No speech could be recognized. Try again.");
            }
            else if (result.Reason == ResultReason.Canceled)
            {
                var cancellation = CancellationDetails.FromResult(result);
                Console.WriteLine($"CANCELED: Reason={cancellation.Reason}");
                if (cancellation.Reason == CancellationReason.Error)
                {
                    Console.WriteLine($"CANCELED: ErrorDetails={cancellation.ErrorDetails}");
                }
                break;
            }
        }
    }
}
