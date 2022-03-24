package codeinformation

//CodeInformationRequest is a struct to request
type CodeInformationRequest struct {
	UserID   string `json:"userId"`
	Customer struct {
		ID       string `json:"id"`
		SyncFlag bool   `json:"syncFlag"`
	} `json:"customer"`
}
