package transactions

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	UserId      int       `json:"user_id"`
	TypeTrans   string    `json:"type"`
	NameProduct string    `json:"name_product"`
	Quantity    int64     `json:"quantity"`
	Rate        int       `json:"rate"`
	SumMoney    int64     `json:"sum_money"`
	DateTime    time.Time `json:"date_time"`
}
