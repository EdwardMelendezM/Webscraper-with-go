package domain

// Harasser Class: Harasser
type Harasser struct {
	Username       string
	History        []string
	Location       string
	HarassmentFreq int
}

// Victim Class: Victim
type Victim struct {
	Username       string
	HarassmentHist []string
	EmotionalState string
}

// Message Class: Message
type Message struct {
	Content        string
	SentDate       string
	HarassmentType string
	Severity       string
}

// Context Class: Context
type Context struct {
	CommunicationChannel string
	VirtualLocation      string
	SentTime             string
}
