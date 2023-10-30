package main

type firstChannelNewsService interface {
	GetNewsList() []PieceOfNews
}

type service struct {
	newsService firstChannelNewsService
}

func (s *service) GetTheMostPopular() (output PieceOfNews) {
	list := s.newsService.GetNewsList()
	for _, n := range list {
		if output.ViewsCount < n.ViewsCount {
			output = n
		}
	}
	return output
}

func NewService(newsService firstChannelNewsService) *service {
	return &service{
		newsService: newsService,
	}
}
