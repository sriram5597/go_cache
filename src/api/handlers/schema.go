package handlers


type errorResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	Message string `json:"message"`
}

type getKeyPayload struct {
	Key int `json:"key"`
}

type getKeyResponse struct {
	Value int `json:"value"`
}

type setKeyPayload struct {
	Key int `json:"key"`
	Value int `json:"value"` 
	Expiry int `json:"Expiry"`
}

type cacheUsage struct {
	TotalKeys int `json:"totalKeys"`
	MaxKeys int `json:"maxKeys"`
}