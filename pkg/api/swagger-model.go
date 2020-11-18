package api

// swagger:model infoRequest
type InfoRequest struct {
	//example: Bhajan: O Kanha ab to murli ki madhur suna do taan
	Title string `json:"title"`
	//example: Heart touching bhajan
	Info string `json:"info"`
	//example: ["https://www.youtube.com/watch?v=XP9rlhzJoxc"]
	Links []string `json:"links"`
}

// swagger:model infoPutResponse
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"message"`
}
