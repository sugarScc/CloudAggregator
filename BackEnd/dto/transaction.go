package dto

type CommitRequest struct {
	DockerImage       string `json:"docker_image"`
	Port              string `json:"port"`
	FlagMessage       string `json:"flag_message"`
	Url               string `json:"url"`
}

type WithdrawRequest struct {
	TransactionId int64 `json:"transaction_id"`
	//ReturnAddress int64 `json:"return_address"`
}