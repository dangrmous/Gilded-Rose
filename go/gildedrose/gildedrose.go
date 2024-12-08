package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func adjustQuality(item *Item) {
	if item.Name == `Sulfuras, Hand of Ragnaros` {
		return
	}
	item.Quality += getQualityRate(item)
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" && item.SellIn <= 0 {
		item.Quality = 0
	}
	if item.Quality > 50 {
		item.Quality = 50
	}
}

func getQualityRate(item *Item) int {
	qualityRate := -1
	if item.Name == `Aged Brie` {
		qualityRate = 1
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		qualityRate = 1
		if item.SellIn < 11 {
			qualityRate = 2
		}
		if item.SellIn < 6 {
			qualityRate = 3
		}
	}
	return qualityRate
}

func adjustSellIn(item *Item) {
	if item.Name == `Sulfuras, Hand of Ragnaros` {
		return
	}
	item.SellIn -= 1
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		adjustQuality(items[i])
		adjustSellIn(items[i])
	}

}
