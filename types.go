package klout

type Identity struct {
	Id, Network string
}

type Deltas struct {
	Day, Week, Month float32
}

type UserScore struct {
	Score  float32
	Bucket string
}

type User struct {
	Id    string     `json:"kloutId"`
	Nick  string     `json:"nick"`
	Score *UserScore `json:"score"`
	Delta *Deltas    `json:"scoreDeltas"`
}

type Score struct {
	S      float32 `json:"score"`
	Delta  *Deltas `json:"scoreDelta"`
	Bucket string  `json:"bucket"`
}

type Topic struct {
	Id      string `json:"id"`
	Display string `json:"displayName"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Image   string `json:"imageUrl"`
	Type    string `json:"displayType"`
}

type Topics []*Topic

type EntityObject struct {
	Id      string `json:"id"`
	Payload *User  `json:"payload"`
}

type Entity struct {
	E *EntityObject `json:"entity"`
}

type Influence struct {
	Ers      []*Entity `json:"myInfluencers"`
	Ees      []*Entity `json:"myInfluencees"`
	ErsCount int       `json:"myInfluencersCount"`
	EesCount int       `json:"myInfluenceesCount"`
}
