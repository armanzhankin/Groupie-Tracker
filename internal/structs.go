package groupie

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationDate map[string][]string
}

type Relations struct {
	Id int                 `json:"id"`
	LD map[string][]string `json:"datesLocations"`
}
