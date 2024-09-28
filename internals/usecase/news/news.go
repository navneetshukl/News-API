package news

import (
	"log"
	"news-api/internals/adapter/external/news_3"
	"news-api/internals/adapter/external/news_4"
	"news-api/internals/adapter/external/news_5"
	"news-api/internals/core/news"
)

type NewsUseCaseImpl struct {
}

func NewNewsUsecase() news.NewsUseCase {
	return &NewsUseCaseImpl{}
}

func (nu *NewsUseCaseImpl) GetFirstNews() (*news.NewsAPIResponse, error) {
	newsData := &news.NewsAPIResponse{}

	articles := []news.Article{}
	data, err := news_3.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for _, value := range data.Articles {
		articles = append(articles, news.Article{
			Author:      value.Author,
			Title:       value.Title,
			Description: value.Description,
			URL:         value.URL,
			URLToImage:  value.URLToImage,
			PublishedAt: value.PublishedAt,
			Content:     value.Content,
			Source:      news.Source(value.Source),
		})

	}

	newsData.Articles = articles
	newsData.Status = data.Status
	newsData.TotalResults = data.TotalResults

	return newsData, nil
}

func (nu *NewsUseCaseImpl) GetSecondNews() (*news.NewsAPIResponse, error) {
	newsData := &news.NewsAPIResponse{}

	articles := []news.Article{}

	data, err := news_4.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for _, value := range data.Articles {
		articles = append(articles, news.Article{
			Author:      value.Author,
			Title:       value.Title,
			Description: value.Description,
			URL:         value.URL,
			URLToImage:  value.URLToImage,
			PublishedAt: value.PublishedAt,
			Content:     value.Content,
			Source:      news.Source(value.Source),
		})

	}

	newsData.Articles = articles
	newsData.Status = data.Status
	newsData.TotalResults = data.TotalResults

	return newsData, nil
}
func (nu *NewsUseCaseImpl) GetThirdNews() (*news.NewsAPIResponse, error) {
	newsData := &news.NewsAPIResponse{}

	articles := []news.Article{}
	data, err := news_5.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for _, value := range data.Articles {
		articles = append(articles, news.Article{
			Author:      value.Author,
			Title:       value.Title,
			Description: value.Description,
			URL:         value.URL,
			URLToImage:  value.URLToImage,
			PublishedAt: value.PublishedAt,
			Content:     value.Content,
			Source:      news.Source(value.Source),
		})

	}

	newsData.Articles = articles
	newsData.Status = data.Status
	newsData.TotalResults = data.TotalResults

	return newsData, nil
}
