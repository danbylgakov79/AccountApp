package AccountApp

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type Product struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Href  string `json:"href"`
}

type ProductList struct {
	Id        int
	ListId    int
	ProductId int
}
