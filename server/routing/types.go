package routing

// GetDataReq is the req for GTFS
type GetDataReq struct {
	AuthPwd string
}

// PutDataReq is the req for PDTS
type PutDataReq struct {
	AuthPwd string
	Vault   []byte
}

// GetDataRes is the res for GTFS
type GetDataRes struct {
	StatusCode int
	Status     string
	Vault      []byte
}

// PutDataRes is the res for PDTS
type PutDataRes struct {
	StatusCode int
	Status     string
}
