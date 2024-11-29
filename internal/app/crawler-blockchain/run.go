package crawlerblockchain

func (service *CrawlerBlockchain) Run() {
	service.wg.Add(2)
	service.deps.Logger.Out.Info("Started crawling")
	go service.produce()
	go service.consume()
}
