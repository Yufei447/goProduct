package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" goProduct:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" goProduct:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" goProduct:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" goProduct:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" goProduct:"ProductUrl"`
}
