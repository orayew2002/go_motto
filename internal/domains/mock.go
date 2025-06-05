package domains

var MockQuoteRequests = []CreateQuoteRequest{
	{
		Author: "Albert Einstein",
		Quote:  "Imagination is more important than knowledge.",
	},
	{
		Author: "Steve Jobs",
		Quote:  "Stay hungry, stay foolish.",
	},
	{
		Author: "Mark Twain",
		Quote:  "The secret of getting ahead is getting started.",
	},
	{
		Author: "Confucius",
		Quote:  "It does not matter how slowly you go as long as you do not stop.",
	},
	{
		Author: "Walt Disney",
		Quote:  "The way to get started is to quit talking and begin doing.",
	},
	{
		Author: "Maya Angelou",
		Quote:  "If you don't like something, change it. If you can't change it, change your attitude.",
	},
	{
		Author: "Nelson Mandela",
		Quote:  "It always seems impossible until it's done.",
	},
	{
		Author: "Oscar Wilde",
		Quote:  "Be yourself; everyone else is already taken.",
	},
	{
		Author: "Helen Keller",
		Quote:  "Keep your face to the sunshine and you cannot see a shadow.",
	},
	{
		Author: "Benjamin Franklin",
		Quote:  "Tell me and I forget. Teach me and I remember. Involve me and I learn.",
	},
}

type MockLogger struct{}

func (m MockLogger) Error(msg string, keysAndValues ...any) {}
func (m MockLogger) Info(msg string, keysAndValues ...any)  {}
func (m MockLogger) Warn(msg string, keysAndValues ...any)  {}
