package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func adjustQuality(item *Item) {
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
	if item.Name == `Sulfuras, Hand of Ragnaros` {
		qualityRate = 0
	}
	if item.Name == `Aged Brie` {
		qualityRate = 1
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.SellIn < 11 {
			qualityRate = 2
		}
		if item.SellIn < 6 {
			qualityRate = 3
		}
	}
	return qualityRate
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						adjustQuality(items[i])
					}
					if items[i].SellIn < 6 {
						adjustQuality(items[i])
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}

}
