package schema

type (
	Page struct {
		Total    int64 `json:"total"`
		Page     int   `json:"page,omitempty"`
		PageSize int   `json:"pageSize,omitempty"`
	}
)
