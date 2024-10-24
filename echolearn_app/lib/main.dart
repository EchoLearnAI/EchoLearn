// A simple Flutter app that integrates Azure AI Services for real-time conversation assistance.

import 'package:flutter/material.dart';
import 'package:speech_to_text/speech_to_text.dart' as stt;
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'EchoLearn',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  late stt.SpeechToText _speech;
  bool _isListening = false;
  String _text = "";
  final String _apiKey = '';
  final String _endpoint = '';

  @override
  void initState() {
    super.initState();
    _speech = stt.SpeechToText();
    _initializeSpeech();
  }

  void _initializeSpeech() async {
    bool hasSpeech = await _speech.initialize(
      onStatus: (val) => print('onStatus: $val'),
      onError: (val) => print('onError: $val'),
    );
    if (!hasSpeech) {
      print("Speech recognition initialization failed");
    }
  }

  void _listen() async {
    if (!_isListening) {
      bool available = await _speech.initialize(
        onStatus: (val) => print('onStatus: $val'),
        onError: (val) => print('onError: $val'),
      );
      if (available && await _speech.hasPermission) {
        setState(() => _isListening = true);
        _speech.listen(
          onResult: (val) => setState(() {
            _text = val.recognizedWords;
          }),
        );
      } else {
        print("Speech recognition not available or permission not granted");
      }
    } else {
      setState(() => _isListening = false);
      _speech.stop();
      _analyzeText();
    }
  }

  void _analyzeText() async {
    if (_text.isEmpty) {
      print("No text to analyze");
      return;
    }
    try {
      String prompt = "You are an AI assistant that helps users with their grammar and pronunciation issues during a conversation. "
          "When the user speaks or types, identify any grammar or pronunciation problems they may have and offer clear, helpful corrections in real-time. "
          "After assisting, continue engaging with the user naturally. Your aim is to correct, assist, and then facilitate an ongoing conversation, making the experience as comfortable and engaging as possible.";

      print("Sending request to Azure with prompt: $prompt\nUser: $_text");

      var response = await http.post(
        Uri.parse('$_endpoint/openai/deployments/gpt-4o-realtime-preview/completions?api-version=2023-05-15'),
        headers: {
          'Content-Type': 'application/json',
          'api-key': _apiKey,
        },
        body: jsonEncode({
          'prompt': "$prompt\nUser: $_text\nAI:",
          'max_tokens': 100,
        }),
      );

      if (response.statusCode == 200) {
        var responseData = jsonDecode(response.body);
        setState(() {
          _text += "\nAI: ${responseData['choices'][0]['text']}";
        });
      } else {
        print("Error: ${response.statusCode}, ${response.body}");
      }
    } catch (e) {
      print("Error: $e");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Grammar and Pronunciation Helper'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            Expanded(
              child: SingleChildScrollView(
                reverse: true,
                child: Text(
                  _text,
                  style: TextStyle(fontSize: 18.0),
                ),
              ),
            ),
            SizedBox(height: 10),
            FloatingActionButton(
              onPressed: _listen,
              child: Icon(_isListening ? Icons.mic : Icons.mic_none),
            ),
          ],
        ),
      ),
    );
  }
}
