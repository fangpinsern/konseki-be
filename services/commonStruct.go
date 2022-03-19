package services

type MessageInfo struct {
	Id string
	ExposureDate int64
	MsgType string
	// actual message will depend on the date from exposure
}

type Event struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Attended []string `json:"attended"`
	Date int64 `json:"date"`
	CreatorId string `json:"creator_id"`
}
