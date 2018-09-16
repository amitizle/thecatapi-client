package thecatapi

type FavouriteService struct {
	client *Client // TODO needed?
}

func NewFavouriteService(client *Client) (*FavouriteService, error) {
	return &FavouriteService{
		client: client,
	}, nil // TODO
}

func (favouriteService *FavouriteService) Get() {
}
