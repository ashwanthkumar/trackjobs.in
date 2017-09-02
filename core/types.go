package core

// Message represents the parsed email
type Message struct {
	From       string
	To         string
	ParsedFrom string // when forwarding emails, who initially sent the message
	Parts      []*MessagePart
}

// MessagePart represents the differnt versions of the email when they're sent as multipart/alternative
type MessagePart struct {
	Type    string // MIME Content-Type of the message part
	Content string
}
