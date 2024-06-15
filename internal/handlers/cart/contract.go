package cart

type ctxKey string

const (
	ctxKeyUserID ctxKey = "user_id"
)

type RequestAddToCard struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type ReqRemoveCart struct {
	ProductID uint `json:"product_id"`
}
