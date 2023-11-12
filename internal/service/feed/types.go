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

type ItunesSearchPodcastResult struct {
    WrapperType              string   `json:"wrapperType"`
    Kind                     string   `json:"kind"`
    CollectionId             int      `json:"collectionId"`
    TrackId                  int      `json:"trackId"`
    ArtistName               string   `json:"artistName"`
    CollectionName           string   `json:"collectionName"`
    TrackName                string   `json:"trackName"`
    CollectionCensoredName   string   `json:"collectionCensoredName"`
    TrackCensoredName        string   `json:"trackCensoredName"`
    CollectionViewUrl        string   `json:"collectionViewUrl"`
    FeedUrl                  string   `json:"feedUrl"`
    TrackViewUrl             string   `json:"trackViewUrl"`
    ArtworkUrl30             string   `json:"artworkUrl30"`
    ArtworkUrl60             string   `json:"artworkUrl60"`
    ArtworkUrl100            string   `json:"artworkUrl100"`
    CollectionPrice          float64  `json:"collectionPrice"`
    TrackPrice               float64  `json:"trackPrice"`
    CollectionHdPrice        int      `json:"collectionHdPrice"`
    ReleaseDate              string   `json:"releaseDate"`
    CollectionExplicitness   string   `json:"collectionExplicitness"`
    TrackExplicitness        string   `json:"trackExplicitness"`
    TrackCount               int      `json:"trackCount"`
    TrackTimeMillis          int      `json:"trackTimeMillis"`
    Country                  string   `json:"country"`
    Currency                 string   `json:"currency"`
    PrimaryGenreName         string   `json:"primaryGenreName"`
    ContentAdvisoryRating    string   `json:"contentAdvisoryRating"`
    ArtworkUrl600            string   `json:"artworkUrl600"`
    GenreIds                 []string `json:"genreIds"`
    Genres                   []string `json:"genres"`
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
