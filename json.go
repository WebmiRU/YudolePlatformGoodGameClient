package main

//type Smile struct {
//	Key    string          `json:"key"`
//	Images ChatSmileImages `json:"images"`
//}

type ChatSmile struct {
	Id         string          `json:"id"`
	Key        string          `json:"key"`
	Level      int             `json:"level"`
	Paid       string          `json:"paid"`
	Bind       string          `json:"bind"`
	InternalId int             `json:"internal_id"`
	ChannelId  int             `json:"channel_id"`
	Channel    interface{}     `json:"channel"`
	Nickname   string          `json:"nickname"`
	Donat      int             `json:"donat"`
	Premium    int             `json:"premium"`
	Animated   int             `json:"animated"`
	Images     ChatSmileImages `json:"images"`
	Tags       string          `json:"tags"`
	TagsArray  []string        `json:"tagsArray"`
}

type ChatSmileImages struct {
	Small string `json:"small"`
	Big   string `json:"big"`
	Gif   string `json:"gif"`
}

type ChatMessage struct {
	Type string `json:"type"`
	//Data any    `json:"data"`
}

type ChatMessageWelcome struct {
}

type ChatMessageJoin struct {
	Type string              `json:"type"`
	Data ChatMessageJoinData `json:"data"`
}

type ChatMessageJoinData struct {
	ChannelId string `json:"channel_id"`
	Hidden    int    `json:"hidden"`
	Mobile    bool   `json:"mobile"`
	Reload    bool   `json:"reload"`
}

type ChatMessageMessage struct {
	Type string                 `json:"type"`
	Data ChatMessageMessageData `json:"data"`
}

type ChatMessageMessageData struct {
	ChannelId  string   `json:"channel_id"`
	UserId     int      `json:"user_id"`
	UserName   string   `json:"user_name"`
	UserRights int      `json:"user_rights"`
	Premium    int      `json:"premium"`
	Premiums   []string `json:"premiums"`
	Resubs     struct {
		Field1 int `json:"15365"`
		Field2 int `json:"40756"`
	} `json:"resubs"`
	Staff       int    `json:"staff"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	Role        string `json:"role"`
	Mobile      int    `json:"mobile"`
	Payments    int    `json:"payments"`
	PaymentsAll struct {
		Field1 int `json:"5"`
		Field2 int `json:"6515"`
		Field3 int `json:"15365"`
		Field4 int `json:"183946"`
	} `json:"paymentsAll"`
	GgPlusTier int    `json:"gg_plus_tier"`
	IsStatus   int    `json:"isStatus"`
	MessageId  int64  `json:"message_id"`
	Timestamp  int    `json:"timestamp"`
	Text       string `json:"text"`
	Regtime    int    `json:"regtime"`
}
