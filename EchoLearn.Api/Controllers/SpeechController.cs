using EchoLearn.Api.Services;
using Microsoft.AspNetCore.Mvc;

namespace EchoLearn.Api.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class SpeechController : ControllerBase
    {
        private readonly ILogger<SpeechController> _logger;
        private readonly ISpeechService _speechService;

        public SpeechController(ILogger<SpeechController> logger, ISpeechService speechService)
        {
            _logger = logger;
            _speechService = speechService;
        }

        [HttpPost]
        public async Task<IActionResult> RecognizeSpeechAsync([FromForm] IFormFile audioFile)
        {
            if (audioFile == null || audioFile.Length == 0)
            {
                _logger.LogWarning("Audio file is required.");
                return BadRequest("Audio file is required.");
            }

            _logger.LogInformation("Received audio file for speech recognition.");

            using var memoryStream = audioFile.OpenReadStream();
            var recognizedText = await _speechService.RecognizeSpeechAsync(memoryStream);
            return Ok(new { Text = recognizedText });
        }
    }
}