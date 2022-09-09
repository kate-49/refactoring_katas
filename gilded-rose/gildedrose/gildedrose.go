package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func updateItemQualityWhenItemQualityUnder50(item *Item) *Item {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	return item
}

func CalculateAgedBrie(item *Item) *Item {
	updateItemQualityWhenItemQualityUnder50(item)
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		updateItemQualityWhenItemQualityUnder50(item)
	}
	return item
}

func CalculateBackstagePasses(item *Item) *Item {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
		if item.SellIn < 11 {
			updateItemQualityWhenItemQualityUnder50(item)
		}
		if item.SellIn < 6 {
			updateItemQualityWhenItemQualityUnder50(item)
		}
	}
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		item.Quality = item.Quality - item.Quality
	}
	return item
}

func calculateQuality(item *Item) *Item {
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}
	}
	return item
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" && items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
			calculateQuality(items[i])
		} else {
			if items[i].Name == "Aged Brie" {
				CalculateAgedBrie(items[i])
			}
			if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
				CalculateBackstagePasses(items[i])
			}
		}
	}
}
