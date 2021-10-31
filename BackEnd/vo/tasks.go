package vo

type Tasks []Task

type Task struct {
	TransactionId      int64  `json:"transaction_id"`
	State              string `json:"state"`
	DockerImage        string `json:"docker_image"`
	Port               string `json:"port"`
	FlagMessage        string `json:"flag_message"`
	Url                string `json:"url"`
	CreationTimeStamps int64  `json:"creation_time_stamps"`
	WithDrawVisible    bool   `json:"with_draw_visible"`
}
