package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type extendedItem struct {
	item *Item
	sid  specialItemData
}

type qualityRateAdjustment struct {
	whenSellInBelow int
	newRate         int
}

type specialItemData struct {
	hasQualityRate         bool
	qualityRate            int
	qualityRateAdjustments []qualityRateAdjustment
	fixedSellIn            bool
	zeroQualityAfterSellIn bool
}

var specialItems = map[string]specialItemData{
	"Sulfuras, Hand of Ragnaros": {
		hasQualityRate: true,
		qualityRate:    0,
		fixedSellIn:    true,
	},
	"Aged Brie": {
		hasQualityRate: true,
		qualityRate:    1,
	},
	"Backstage passes to a TAFKAL80ETC concert": {
		hasQualityRate: true,
		qualityRate:    1,
		qualityRateAdjustments: []qualityRateAdjustment{
			{whenSellInBelow: 11, newRate: 2},
			{whenSellInBelow: 6, newRate: 3},
		},
		zeroQualityAfterSellIn: true,
	},
}

func adjustQuality(ei extendedItem) {
	item := ei.item
	qr := getQualityRate(ei)
	if qr == 0 {
		return
	}
	item.Quality += qr
	if ei.sid.zeroQualityAfterSellIn && item.SellIn <= 0 {
		item.Quality = 0
	}
	if item.Quality > 50 {
		item.Quality = 50
	}
	if item.Quality < 0 {
		item.Quality = 0
	}
}

func getQualityRate(ei extendedItem) int {
	item := ei.item
	qualityRate := -1
	if ei.sid.hasQualityRate {
		qualityRate = ei.sid.qualityRate
	}
	if item.SellIn <= 0 {
		qualityRate = qualityRate * 2
	}
	for _, qra := range ei.sid.qualityRateAdjustments {
		if item.SellIn < qra.whenSellInBelow {
			qualityRate = qra.newRate
		}
	}
	return qualityRate
}

func adjustSellIn(ei extendedItem) {
	item := ei.item
	if ei.sid.fixedSellIn {
		return
	}
	item.SellIn -= 1
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		var extraData specialItemData
		if si, found := specialItems[items[i].Name]; found {
			extraData = si
		} else {
			extraData = specialItemData{}
		}
		exti := extendedItem{
			item: items[i],
			sid:  extraData,
		}
		adjustQuality(exti)
		adjustSellIn(exti)
	}
}
