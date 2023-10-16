package feed

type SearchParams struct {
	Keyword       string
	Scope         string
	ExcludeFeedId string
	Source        string
	Country       string
	SortByDate    int
	Page          int
	Size          int
}

type UserSubkeywordParams struct {
	Keyword       string
	UserId        string
	ExcludeFeedId string
	Country       string
	Source        string
}

type ItunesSearchEpisodeResult struct {
	ContentAdvisoryRating string `json:"contentAdvisoryRating"`
	TrackViewUrl          string `json:"trackViewUrl"`
	ArtworkUrl60          string `json:"artworkUrl60"`
	EpisodeFileExtension  string `json:"episodeFileExtension"`
	EpisodeContentType    string `json:"episodeContentType"`
	ArtworkUrl160         string `json:"artworkUrl160"`
	FeedUrl               string `json:"feedUrl"`
	ClosedCaptioning      string `json:"closedCaptioning"`
	CollectionId          string `json:"collectionId"`
	CollectionName        string `json:"collectionName"`
	EpisodeUrl            string `json:"episodeUrl"`
	Genres                []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"genres"`
	EpisodeGuid       string `json:"episodeGuid"`
	Description       string `json:"description"`
	TrackId           int    `json:"trackId"`
	TrackName         string `json:"trackName"`
	ShortDescription  string `json:"shortDescription"`
	ReleaseDate       string `json:"releaseDate"`
	ArtistIds         []int  `json:"artistIds"`
	PreviewUrl        string `json:"previewUrl"`
	TrackTimeMillis   int    `json:"trackTimeMillis"`
	CollectionViewUrl string `json:"collectionViewUrl"`
	ArtworkUrl600     string `json:"artworkUrl600"`
	Kind              string `json:"kind"`
	WrapperType       string `json:"wrapperType"`
	Country           string `json:"country"`
}
