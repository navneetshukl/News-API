package news_1


type Response struct {
	Status       string    `json:"status" bson:"status"`
	Copyright    string    `json:"copyright" bson:"copyright"`
	Section      string    `json:"section" bson:"section"`
	LastUpdated  string    `json:"last_updated" bson:"last_updated"`
	NumResults   int       `json:"num_results" bson:"num_results"`
	Results      []Result  `json:"results" bson:"results"`
}

type Result struct {
	Section          string       `json:"section" bson:"section"`
	Subsection       string       `json:"subsection" bson:"subsection"`
	Title            string       `json:"title" bson:"title"`
	Abstract         string       `json:"abstract" bson:"abstract"`
	URL              string       `json:"url" bson:"url"`
	URI              string       `json:"uri" bson:"uri"`
	Byline           string       `json:"byline" bson:"byline"`
	ItemType         string       `json:"item_type" bson:"item_type"`
	UpdatedDate      string       `json:"updated_date" bson:"updated_date"`
	CreatedDate      string       `json:"created_date" bson:"created_date"`
	PublishedDate    string       `json:"published_date" bson:"published_date"`
	MaterialTypeFacet string      `json:"material_type_facet" bson:"material_type_facet"`
	Kicker           string       `json:"kicker" bson:"kicker"`
	DesFacet         []string     `json:"des_facet" bson:"des_facet"`
	OrgFacet         []string     `json:"org_facet" bson:"org_facet"`
	PerFacet         []string     `json:"per_facet" bson:"per_facet"`
	GeoFacet         []string     `json:"geo_facet" bson:"geo_facet"`
	Multimedia       []Multimedia  `json:"multimedia" bson:"multimedia"`
}

type Multimedia struct {
	URL     string `json:"url" bson:"url"`
	Format  string `json:"format" bson:"format"`
	Height  int    `json:"height" bson:"height"`
	Width   int    `json:"width" bson:"width"`
	Type    string `json:"type" bson:"type"`
	Subtype string `json:"subtype" bson:"subtype"`
	Caption string `json:"caption" bson:"caption"`
	Copyright string `json:"copyright" bson:"copyright"`
}
