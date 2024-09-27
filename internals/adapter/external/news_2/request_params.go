package news_2

type NewsResponse struct {
	Status       string    `json:"status" bson:"status"`
	Copyright    string    `json:"copyright" bson:"copyright"`
	NumResults   int       `json:"num_results" bson:"num_results"`
	Results      []Article `json:"results" bson:"results"`
}

type Article struct {
	SlugName          string       `json:"slug_name" bson:"slug_name"`
	Section           string       `json:"section" bson:"section"`
	Subsection        string       `json:"subsection" bson:"subsection"`
	Title             string       `json:"title" bson:"title"`
	Abstract          string       `json:"abstract" bson:"abstract"`
	URI               string       `json:"uri" bson:"uri"`
	URL               string       `json:"url" bson:"url"`
	Byline            string       `json:"byline" bson:"byline"`
	ItemType          string       `json:"item_type" bson:"item_type"`
	Source            string       `json:"source" bson:"source"`
	UpdatedDate       string       `json:"updated_date" bson:"updated_date"`
	CreatedDate       string       `json:"created_date" bson:"created_date"`
	PublishedDate     string       `json:"published_date" bson:"published_date"`
	FirstPublishedDate string      `json:"first_published_date" bson:"first_published_date"`
	MaterialTypeFacet string       `json:"material_type_facet" bson:"material_type_facet"`
	Kicker            string       `json:"kicker" bson:"kicker"`
	Subheadline       string       `json:"subheadline" bson:"subheadline"`
	DesFacet          []string     `json:"des_facet" bson:"des_facet"`
	OrgFacet          []string     `json:"org_facet" bson:"org_facet"`
	PerFacet          []string     `json:"per_facet" bson:"per_facet"`
	GeoFacet          []string     `json:"geo_facet" bson:"geo_facet"`
	RelatedURLs       []string     `json:"related_urls" bson:"related_urls"`
	Multimedia        []Multimedia  `json:"multimedia" bson:"multimedia"`
}

type Multimedia struct {
	URL      string `json:"url" bson:"url"`
	Format   string `json:"format" bson:"format"`
	Height   int    `json:"height" bson:"height"`
	Width    int    `json:"width" bson:"width"`
	Type     string `json:"type" bson:"type"`
	Subtype  string `json:"subtype" bson:"subtype"`
	Caption  string `json:"caption" bson:"caption"`
	Copyright string `json:"copyright" bson:"copyright"`
}