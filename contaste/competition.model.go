package contaste

type ContasteManager struct {
}
type Competition map[string]CompetitionObject

type CoupleName string

type CoupleResult struct {
	Category   string
	Award      string
	Outof      string
	CoupleName string
}

type CompetitionObject struct {
	AgeFrom            int                  `json:"ageFrom,omitempty"`
	AgeTill            int                  `json:"ageTill,omitempty"`
	AwardStatus        string               `json:"awardStatus,omitempty"`
	StoredContestTitle string               `json:"storedContestTitle,omitempty"`
	Title              Title                `json:"title,omitempty"`
	Type               string               `json:"type,omitempty"`
	DancingLevel       string               `json:"dancingLevel,omitempty"`
	Achivments         map[string]Achivment `json:"achi,omitempty"`
	Dancers            map[string]Dancer    `json:"dancers,omitempty"`
	Group              string               `json:"group,omitempty"`
}

type Title struct {
	Eng string `json:"eng,omitempty"`
}

type Achivment struct {
	Award string `json:"award,omitempty"`
	OutOf string `json:"outof,omitempty"`
}

type Dancer struct {
	Title   string `json:"title,omitempty"`
	Checkin string `json:"checkin,omitempty"`
}
