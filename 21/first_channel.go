package main

type firstChannelNewsSvc struct {
}

func (s *firstChannelNewsSvc) GetNewsList() []PieceOfNews {
	return []PieceOfNews{
		{
			ViewsCount: 1,
			Content:    "Российские ученные открыли что Крот может вырыть 900 метровый туннель за ночь",
		},
		{
			ViewsCount: 30,
			Content:    "Завтра солнечное затмение",
		},
	}
}

func NewFirstChannelNewsService() *firstChannelNewsSvc {
	return &firstChannelNewsSvc{}
}
