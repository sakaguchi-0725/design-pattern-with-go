package main

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type NewsService struct {
	observers []Observer
	news      string
}

func MakeNewsService() *NewsService {
	return &NewsService{
		observers: make([]Observer, 0),
	}
}

func (n *NewsService) RegisterObserver(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsService) RemoveObserver(observer Observer) {
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsService) NotifyObservers() {
	for _, observer := range n.observers {
		observer.Update(n.news)
	}
}

func (n *NewsService) AddNews(news string) {
	n.news = news
	n.NotifyObservers()
}
