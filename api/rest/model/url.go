package model

type(
	RequestURL struct {
		URL string `json:"url"`
	}

	ResponseURL struct {
		ShortURL string `json:"result"`
	}
)