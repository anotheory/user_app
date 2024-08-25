package healthcheck

type HealthcheckResponse struct {
	Version string `json:"version"`
	Message string `json:"message"`
}
