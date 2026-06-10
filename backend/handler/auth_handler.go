package handler

func (shop *Marineshop) NewAuthHandler() AuthHandler {
	return &Authentication{AuthService: shop.AuthService}
}
