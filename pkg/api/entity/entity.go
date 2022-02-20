package entity

type Tweet struct {
	ID             int64         `json:"id"`
	ConversationID string        `json:"conversation_id"`
	CreatedAt      int64         `json:"created_at"`
	Date           string        `json:"date"`
	Time           string        `json:"time"`
	Timezone       string        `json:"timezone"`
	UserID         int           `json:"user_id"`
	Username       string        `json:"username"`
	Name           string        `json:"name"`
	Place          string        `json:"place"`
	Tweet          string        `json:"tweet"`
	Mentions       []string      `json:"mentions"`
	Urls           []interface{} `json:"urls"`
	Photos         []interface{} `json:"photos"`
	RepliesCount   int           `json:"replies_count"`
	RetweetsCount  int           `json:"retweets_count"`
	LikesCount     int           `json:"likes_count"`
	Hashtags       []interface{} `json:"hashtags"`
	Cashtags       []interface{} `json:"cashtags"`
	Link           string        `json:"link"`
	Retweet        bool          `json:"retweet"`
	QuoteURL       string        `json:"quote_url"`
	Video          int           `json:"video"`
	Near           string        `json:"near"`
	Geo            string        `json:"geo"`
	Source         string        `json:"source"`
	UserRtID       string        `json:"user_rt_id"`
	UserRt         string        `json:"user_rt"`
	RetweetID      string        `json:"retweet_id"`
	ReplyTo        []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"reply_to"`
	RetweetDate string `json:"retweet_date"`
	Translate   string `json:"translate"`
	TransSrc    string `json:"trans_src"`
	TransDest   string `json:"trans_dest"`
}
