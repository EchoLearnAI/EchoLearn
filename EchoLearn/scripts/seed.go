package main

import (
	"echolearn/internal/db"
	"echolearn/internal/models"
	"fmt"
	"log"
)

func main() {
	// Initialize database connection
	db.InitDB()

	// Populate database with seed data
	createSampleGrammarRules()
	createSampleQuestions()

	fmt.Println("Seed data created successfully!")
}

func createSampleGrammarRules() {
	grammarRules := []models.GrammarRule{
		{
			Title:       "Present Simple",
			Description: "The present simple is used to express habits, facts, and routines.",
			Examples:    "I play tennis every weekend. She works at a bank. Water boils at 100°C.",
		},
		{
			Title:       "Articles",
			Description: "Articles are used before nouns to indicate whether the noun is specific or general. 'The' is used for specific nouns, while 'a/an' is used for non-specific nouns.",
			Examples:    "The car outside is mine (specific). I need to buy a car (non-specific).",
		},
		{
			Title:       "Prepositions of Place",
			Description: "Prepositions of place show the relationship between a person or thing and its location.",
			Examples:    "The book is on the table. She is waiting at the bus stop. There's a cat under the bed.",
		},
		{
			Title:       "Modal Verbs",
			Description: "Modal verbs are used to express ability, possibility, permission, or obligation.",
			Examples:    "I can swim. You may leave now. She must finish her homework.",
		},
		{
			Title:       "Conditionals",
			Description: "Conditional sentences show the possible result of a certain condition.",
			Examples:    "If it rains, I will stay at home. If I had more money, I would travel the world.",
		},
	}

	for _, rule := range grammarRules {
		if err := db.DB.Create(&rule).Error; err != nil {
			log.Printf("Error creating grammar rule: %v", err)
		}
	}
}

func createSampleQuestions() {
	// Get grammar rules
	var grammarRules []models.GrammarRule
	if err := db.DB.Find(&grammarRules).Error; err != nil {
		log.Fatalf("Error fetching grammar rules: %v", err)
	}

	if len(grammarRules) < 5 {
		log.Fatalf("Not enough grammar rules found in the database")
	}

	// Sample questions
	questions := []struct {
		Text     string
		Category string
		Options  []struct {
			Label       string
			IsCorrect   bool
			Explanation string
		}
		GrammarRuleIndex int
	}{
		{
			Text:     "She ____ to work every day.",
			Category: "present_simple",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "go", IsCorrect: false, Explanation: "'go' needs an 's' with third-person singular subjects in present simple."},
				{Label: "goes", IsCorrect: true, Explanation: "Correct! In present simple, third-person singular subjects (he/she/it) use verb + s."},
				{Label: "going", IsCorrect: false, Explanation: "'going' is the present continuous form, not present simple."},
				{Label: "to go", IsCorrect: false, Explanation: "'to go' is the infinitive form, not the conjugated verb needed here."},
			},
			GrammarRuleIndex: 0, // Present Simple
		},
		{
			Text:     "I saw ____ movie that you recommended.",
			Category: "articles",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "a", IsCorrect: false, Explanation: "'a' is for non-specific nouns. Here, we're talking about a specific movie that was recommended."},
				{Label: "an", IsCorrect: false, Explanation: "'an' is used before vowel sounds for non-specific nouns."},
				{Label: "the", IsCorrect: true, Explanation: "Correct! 'The' is used for specific nouns, and here we're referring to a specific movie."},
				{Label: "No article", IsCorrect: false, Explanation: "An article is needed because we're talking about a specific, singular noun."},
			},
			GrammarRuleIndex: 1, // Articles
		},
		{
			Text:     "The keys are ____ the table.",
			Category: "prepositions",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "on", IsCorrect: true, Explanation: "Correct! 'on' is used for objects resting on a surface."},
				{Label: "in", IsCorrect: false, Explanation: "'in' indicates something is inside a container or space, not appropriate for this context."},
				{Label: "at", IsCorrect: false, Explanation: "'at' indicates a point or position, not typically used for physical objects on a surface."},
				{Label: "by", IsCorrect: false, Explanation: "'by' indicates proximity but not directly on top of something."},
			},
			GrammarRuleIndex: 2, // Prepositions of Place
		},
		{
			Text:     "You ____ smoke in the hospital.",
			Category: "modal_verbs",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "must", IsCorrect: false, Explanation: "'must' indicates obligation, which is the opposite of what's needed here."},
				{Label: "should", IsCorrect: false, Explanation: "'should' suggests advice, not a strong prohibition."},
				{Label: "can", IsCorrect: false, Explanation: "'can' indicates ability or permission, not prohibition."},
				{Label: "must not", IsCorrect: true, Explanation: "Correct! 'must not' expresses prohibition or that something is not allowed."},
			},
			GrammarRuleIndex: 3, // Modal Verbs
		},
		{
			Text:     "If it rains tomorrow, we ____ the picnic.",
			Category: "conditionals",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "cancel", IsCorrect: false, Explanation: "This should be 'will cancel' in the first conditional structure."},
				{Label: "will cancel", IsCorrect: true, Explanation: "Correct! First conditional uses 'if + present simple, will + infinitive'."},
				{Label: "would cancel", IsCorrect: false, Explanation: "'would cancel' is used in second conditional, not first conditional."},
				{Label: "cancelled", IsCorrect: false, Explanation: "This is past tense, but we need future tense in the first conditional."},
			},
			GrammarRuleIndex: 4, // Conditionals
		},
		{
			Text:     "She has lived in Paris ____ 2010.",
			Category: "prepositions",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "for", IsCorrect: false, Explanation: "'for' is used with periods of time (e.g., for 5 years), not with a specific point."},
				{Label: "since", IsCorrect: true, Explanation: "Correct! 'Since' is used with a specific point in time from which something started."},
				{Label: "from", IsCorrect: false, Explanation: "'from' would need 'to' to complete the range. 'From 2010' is incomplete."},
				{Label: "at", IsCorrect: false, Explanation: "'at' is used for specific times or points, not for continuous periods."},
			},
			GrammarRuleIndex: 2, // Prepositions of Place
		},
		{
			Text:     "I don't have ____ money.",
			Category: "articles",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "a", IsCorrect: false, Explanation: "'a' can't be used with uncountable nouns like 'money'."},
				{Label: "an", IsCorrect: false, Explanation: "'an' can't be used with uncountable nouns like 'money'."},
				{Label: "the", IsCorrect: false, Explanation: "'the' would indicate specific money, but the sentence is about money in general."},
				{Label: "any", IsCorrect: true, Explanation: "Correct! 'Any' is used in negative sentences with uncountable nouns."},
			},
			GrammarRuleIndex: 1, // Articles
		},
		{
			Text:     "He ____ tennis every Saturday.",
			Category: "present_simple",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "play", IsCorrect: false, Explanation: "'play' needs an 's' with third-person singular subjects in present simple."},
				{Label: "plays", IsCorrect: true, Explanation: "Correct! In present simple, third-person singular subjects (he/she/it) use verb + s."},
				{Label: "is playing", IsCorrect: false, Explanation: "'is playing' is present continuous, not present simple which is used for habits."},
				{Label: "to play", IsCorrect: false, Explanation: "'to play' is the infinitive form, not the conjugated verb needed here."},
			},
			GrammarRuleIndex: 0, // Present Simple
		},
		{
			Text:     "If I ____ rich, I would travel around the world.",
			Category: "conditionals",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "am", IsCorrect: false, Explanation: "Present tense 'am' doesn't work in the second conditional."},
				{Label: "will be", IsCorrect: false, Explanation: "'will be' is used in the first conditional, not the second."},
				{Label: "were", IsCorrect: true, Explanation: "Correct! The second conditional uses the past tense (or 'were' for all subjects with 'be')."},
				{Label: "would be", IsCorrect: false, Explanation: "'would be' belongs in the main clause, not the 'if' clause in conditionals."},
			},
			GrammarRuleIndex: 4, // Conditionals
		},
		{
			Text:     "You ____ better see a doctor about that cough.",
			Category: "modal_verbs",
			Options: []struct {
				Label       string
				IsCorrect   bool
				Explanation string
			}{
				{Label: "had", IsCorrect: true, Explanation: "Correct! 'had better' is a modal expression used for strong advice or warnings."},
				{Label: "would", IsCorrect: false, Explanation: "'would better' is not a correct modal expression in English."},
				{Label: "must", IsCorrect: false, Explanation: "'must' is too strong and doesn't form the expression 'must better'."},
				{Label: "should", IsCorrect: false, Explanation: "'should better' is not a correct modal expression; we use 'should' alone."},
			},
			GrammarRuleIndex: 3, // Modal Verbs
		},
	}

	// Create questions with their options
	for _, q := range questions {
		question := models.Question{
			Text:      q.Text,
			Category:  q.Category,
			GrammarID: grammarRules[q.GrammarRuleIndex].ID,
		}

		// Create the question
		if err := db.DB.Create(&question).Error; err != nil {
			log.Printf("Error creating question: %v", err)
			continue
		}

		// Create options
		for _, o := range q.Options {
			option := models.Option{
				QuestionID:  question.ID,
				Label:       o.Label,
				IsCorrect:   o.IsCorrect,
				Explanation: o.Explanation,
			}

			if err := db.DB.Create(&option).Error; err != nil {
				log.Printf("Error creating option: %v", err)
			}
		}
	}
} 