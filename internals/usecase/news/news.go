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
	news, err := news_3.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for idx, value := range news.Articles {
		newsData.Articles[idx].Author = value.Author
		newsData.Articles[idx].Title = value.Title
		newsData.Articles[idx].Description = value.Description
		newsData.Articles[idx].URL = value.URL
		newsData.Articles[idx].URLToImage = value.URLToImage
		newsData.Articles[idx].PublishedAt = value.PublishedAt
		newsData.Articles[idx].Content = value.Content
		newsData.Status = "ok"
		newsData.TotalResults = len(news.Articles)

	}

	return newsData, nil
}

func (nu *NewsUseCaseImpl) GetSecondNews() (*news.NewsAPIResponse, error) {
	newsData := &news.NewsAPIResponse{}
	news, err := news_4.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for idx, value := range news.Articles {
		newsData.Articles[idx].Author = value.Author
		newsData.Articles[idx].Title = value.Title
		newsData.Articles[idx].Description = value.Description
		newsData.Articles[idx].URL = value.URL
		newsData.Articles[idx].URLToImage = value.URLToImage
		newsData.Articles[idx].PublishedAt = value.PublishedAt
		newsData.Articles[idx].Content = value.Content
		newsData.Status = "ok"
		newsData.TotalResults = len(news.Articles)

	}

	return newsData, nil
}
func (nu *NewsUseCaseImpl) GetThirdNews() (*news.NewsAPIResponse, error) {
	newsData := &news.NewsAPIResponse{}
	news, err := news_5.NewsSvc()
	if err != nil {
		log.Println("error in getting the news-1", err)
		return nil, err
	}

	for idx, value := range news.Articles {
		newsData.Articles[idx].Author = value.Author
		newsData.Articles[idx].Title = value.Title
		newsData.Articles[idx].Description = value.Description
		newsData.Articles[idx].URL = value.URL
		newsData.Articles[idx].URLToImage = value.URLToImage
		newsData.Articles[idx].PublishedAt = value.PublishedAt
		newsData.Articles[idx].Content = value.Content
		newsData.Status = "ok"
		newsData.TotalResults = len(news.Articles)

	}

	return newsData, nil
}
