package strategy

func Main() {
	lfu := &LFU{}
	cache := InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")
	cache.Add("g", "5")
	lru := &LRU{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d", "4")
	fifo := &FIFO{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "7")

}
