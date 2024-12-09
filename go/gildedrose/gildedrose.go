package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type qualityRateAdjustment struct {
	whenSellInBelow int
	newRate         int
}

type specialItemData struct {
	qualityRate            int
	qualityRateAdjustments []qualityRateAdjustment
	fixedSellIn            bool
	hasZeroQualitySellIn   bool
	zeroQualitySellIn      int
}

var specialItems = map[string]specialItemData{
	"Sulfuras, Hand of Ragnaros": {
		qualityRate: 0,
		fixedSellIn: true,
	},
	"Aged Brie": {
		qualityRate: 1,
	},
	"Backstage passes to a TAFKAL80ETC concert": {
		qualityRateAdjustments: []qualityRateAdjustment{
			{whenSellInBelow: 11, newRate: 2},
			{whenSellInBelow: 6, newRate: 3},
		},
		hasZeroQualitySellIn: true,
		zeroQualitySellIn:    0,
	},
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
	if item.Quality < 0 {
		item.Quality = 0
	}
}

func getQualityRate(item *Item) int {
	qualityRate := -1
	if item.Name == `Aged Brie` {
		qualityRate = 1
	}
	if item.SellIn <= 0 {
		qualityRate = qualityRate * 2
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
