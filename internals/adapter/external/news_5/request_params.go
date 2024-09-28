package news_5

type NewsAPIResponse struct {
	Status       string    `json:"status" bson:"status"`
	TotalResults int       `json:"totalResults" bson:"totalResults"`
	Articles     []Article `json:"articles" bson:"articles"`
}

type Article struct {
	Source      Source  `json:"source" bson:"source"`
	Author      string `json:"author,omitempty" bson:"author,omitempty"` // Use pointer to allow for null
	Title       string  `json:"title" bson:"title"`
	Description string `json:"description,omitempty" bson:"description,omitempty"` // Use pointer to allow for null
	URL         string  `json:"url" bson:"url"`
	URLToImage  string `json:"urlToImage,omitempty" bson:"urlToImage,omitempty"` // Use pointer to allow for null
	PublishedAt string  `json:"publishedAt" bson:"publishedAt"`
	Content     string `json:"content,omitempty" bson:"content,omitempty"` // Use pointer to allow for null
}

type Source struct {
	ID   *string `json:"id,omitempty" bson:"id,omitempty"` // Use pointer to allow for null
	Name string  `json:"name" bson:"name"`
}
