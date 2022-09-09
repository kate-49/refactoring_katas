package gildedrose_test

import (
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"testing"
)

func EqualValues(t *testing.T, name string, actual, expected int) bool {
	if expected != actual {
		t.Errorf("%s: Got %v but expected %v", name, actual, expected)
	}
	return true
}

func Test_UpdateQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 19)
	EqualValues(t, items[0].Name, items[0].SellIn, 9)
	EqualValues(t, items[1].Name, items[1].Quality, 1)
	EqualValues(t, items[1].Name, items[1].SellIn, 1)
	EqualValues(t, items[2].Name, items[2].Quality, 6)
	EqualValues(t, items[2].Name, items[2].SellIn, 4)
	EqualValues(t, items[3].Name, items[3].Quality, 80)
	EqualValues(t, items[3].Name, items[3].SellIn, 0)
	EqualValues(t, items[4].Name, items[4].Quality, 80)
	EqualValues(t, items[4].Name, items[4].SellIn, -1)
	EqualValues(t, items[5].Name, items[5].Quality, 21)
	EqualValues(t, items[5].Name, items[5].SellIn, 14)
	EqualValues(t, items[6].Name, items[6].Quality, 50)
	EqualValues(t, items[6].Name, items[6].SellIn, 9)
	EqualValues(t, items[7].Name, items[7].Quality, 50)
	EqualValues(t, items[7].Name, items[7].SellIn, 4)
	EqualValues(t, items[8].Name, items[8].Quality, 5)
	EqualValues(t, items[8].Name, items[8].SellIn, 2)
}

func Test_UpdateDexterityVest(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 50},
		{"+5 Dexterity Vest", 10, 49},
		{"+5 Dexterity Vest", 12, 50},
		{"+5 Dexterity Vest", 12, 49},
		{"+5 Dexterity Vest", 4, 50},
		{"+5 Dexterity Vest", 4, 49},
		{"+5 Dexterity Vest", 10, 0},
		{"+5 Dexterity Vest", 12, 0},
		{"+5 Dexterity Vest", 4, 0},
		{"+5 Dexterity Vest", -1, 10},
		{"+5 Dexterity Vest", -1, 0},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 49)
	EqualValues(t, items[0].Name, items[0].SellIn, 9)
	EqualValues(t, items[1].Name, items[1].Quality, 48)
	EqualValues(t, items[1].Name, items[1].SellIn, 9)
	EqualValues(t, items[2].Name, items[2].Quality, 49)
	EqualValues(t, items[2].Name, items[2].SellIn, 11)
	EqualValues(t, items[3].Name, items[3].Quality, 48)
	EqualValues(t, items[3].Name, items[3].SellIn, 11)
	EqualValues(t, items[4].Name, items[4].Quality, 49)
	EqualValues(t, items[4].Name, items[4].SellIn, 3)
	EqualValues(t, items[5].Name, items[5].Quality, 48)
	EqualValues(t, items[5].Name, items[5].SellIn, 3)
	EqualValues(t, items[6].Name, items[6].Quality, 0)
	EqualValues(t, items[6].Name, items[6].SellIn, 9)
	EqualValues(t, items[7].Name, items[7].Quality, 0)
	EqualValues(t, items[7].Name, items[7].SellIn, 11)
	EqualValues(t, items[8].Name, items[8].Quality, 0)
	EqualValues(t, items[8].Name, items[8].SellIn, 3)
	EqualValues(t, items[9].Name, items[9].Quality, 8)
	EqualValues(t, items[9].Name, items[9].SellIn, -2)
	EqualValues(t, items[10].Name, items[10].Quality, 0)
	EqualValues(t, items[10].Name, items[10].SellIn, -2)
}

func Test_UpdateAgedBrie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 0, 50},
		{"Aged Brie", 10, 49},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 50)
	EqualValues(t, items[0].Name, items[0].SellIn, -1)
	EqualValues(t, items[1].Name, items[1].Quality, 50)
	EqualValues(t, items[1].Name, items[1].SellIn, 9)
}

func Test_UpdateElixirOfTheMongoose(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Elixir of the Mongoose", 0, 50},
		{"Elixir of the Mongoose", 10, 50},
		{"Elixir of the Mongoose", 10, 49},
		{"Elixir of the Mongoose", 12, 50},
		{"Elixir of the Mongoose", 12, 49},
		{"Elixir of the Mongoose", 4, 50},
		{"Elixir of the Mongoose", 4, 49},
		{"Elixir of the Mongoose", 10, 0},
		{"Elixir of the Mongoose", 12, 0},
		{"Elixir of the Mongoose", 4, 0},
		{"Elixir of the Mongoose", -1, 10},
		{"Elixir of the Mongoose", -1, 0},
		{"Elixir of the Mongoose", 0, 10},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 48)
	EqualValues(t, items[0].Name, items[0].SellIn, -1)
	EqualValues(t, items[1].Name, items[1].Quality, 49)
	EqualValues(t, items[1].Name, items[1].SellIn, 9)
	EqualValues(t, items[2].Name, items[2].Quality, 48)
	EqualValues(t, items[2].Name, items[2].SellIn, 9)
	EqualValues(t, items[3].Name, items[3].Quality, 49)
	EqualValues(t, items[3].Name, items[3].SellIn, 11)
	EqualValues(t, items[4].Name, items[4].Quality, 48)
	EqualValues(t, items[4].Name, items[4].SellIn, 11)
	EqualValues(t, items[5].Name, items[5].Quality, 49)
	EqualValues(t, items[5].Name, items[5].SellIn, 3)
	EqualValues(t, items[6].Name, items[6].Quality, 48)
	EqualValues(t, items[6].Name, items[6].SellIn, 3)
	EqualValues(t, items[7].Name, items[7].Quality, 0)
	EqualValues(t, items[7].Name, items[7].SellIn, 9)
	EqualValues(t, items[8].Name, items[8].Quality, 0)
	EqualValues(t, items[8].Name, items[8].SellIn, 11)
	EqualValues(t, items[9].Name, items[9].Quality, 0)
	EqualValues(t, items[9].Name, items[9].SellIn, 3)
	EqualValues(t, items[10].Name, items[10].Quality, 8)
	EqualValues(t, items[10].Name, items[10].SellIn, -2)
	EqualValues(t, items[11].Name, items[11].Quality, 0)
	EqualValues(t, items[11].Name, items[11].SellIn, -2)
}

func Test_UpdateSulfurasHandOfRagnaros(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 10, 50},
		{"Sulfuras, Hand of Ragnaros", 10, 49},
		{"Sulfuras, Hand of Ragnaros", 0, 50},
		{"Sulfuras, Hand of Ragnaros", 0, 49},
		{"Sulfuras, Hand of Ragnaros", 4, 20},
		{"Sulfuras, Hand of Ragnaros", 4, 10},
		{"Sulfuras, Hand of Ragnaros", 0, 5},
		{"Sulfuras, Hand of Ragnaros", 0, 5},
		{"Sulfuras, Hand of Ragnaros", 5, 0},
		{"Sulfuras, Hand of Ragnaros", -1, 10},
		{"Sulfuras, Hand of Ragnaros", -1, 20},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 50)
	EqualValues(t, items[0].Name, items[0].SellIn, 10)
	EqualValues(t, items[1].Name, items[1].Quality, 49)
	EqualValues(t, items[1].Name, items[1].SellIn, 10)
	EqualValues(t, items[2].Name, items[2].Quality, 50)
	EqualValues(t, items[2].Name, items[2].SellIn, 0)
	EqualValues(t, items[3].Name, items[3].Quality, 49)
	EqualValues(t, items[3].Name, items[3].SellIn, 0)
	EqualValues(t, items[4].Name, items[4].Quality, 20)
	EqualValues(t, items[4].Name, items[4].SellIn, 4)
	EqualValues(t, items[5].Name, items[5].Quality, 10)
	EqualValues(t, items[5].Name, items[5].SellIn, 4)
	EqualValues(t, items[6].Name, items[6].Quality, 5)
	EqualValues(t, items[6].Name, items[6].SellIn, 0)
	EqualValues(t, items[7].Name, items[7].Quality, 5)
	EqualValues(t, items[7].Name, items[7].SellIn, 0)
	EqualValues(t, items[8].Name, items[8].Quality, 0)
	EqualValues(t, items[8].Name, items[8].SellIn, 5)
	EqualValues(t, items[9].Name, items[9].Quality, 10)
	EqualValues(t, items[9].Name, items[9].SellIn, -1)
	EqualValues(t, items[10].Name, items[10].Quality, 20)
	EqualValues(t, items[10].Name, items[10].SellIn, -1)
}

func Test_UpdateBackstagePassesToATAFKAL80ETCConcert(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 5},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 51},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 50},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 50)
	EqualValues(t, items[0].Name, items[0].SellIn, 9)
	EqualValues(t, items[1].Name, items[1].Quality, 50)
	EqualValues(t, items[1].Name, items[1].SellIn, 9)
	EqualValues(t, items[2].Name, items[2].Quality, 50)
	EqualValues(t, items[2].Name, items[2].SellIn, 11)
	EqualValues(t, items[3].Name, items[3].Quality, 50)
	EqualValues(t, items[3].Name, items[3].SellIn, 11)
	EqualValues(t, items[4].Name, items[4].Quality, 50)
	EqualValues(t, items[4].Name, items[4].SellIn, 5)
	EqualValues(t, items[5].Name, items[5].Quality, 12)
	EqualValues(t, items[5].Name, items[5].SellIn, 5)
	EqualValues(t, items[6].Name, items[6].Quality, 50)
	EqualValues(t, items[6].Name, items[6].SellIn, 3)
	EqualValues(t, items[7].Name, items[7].Quality, 8)
	EqualValues(t, items[7].Name, items[7].SellIn, 3)
	EqualValues(t, items[8].Name, items[8].Quality, 0)
	EqualValues(t, items[8].Name, items[8].SellIn, -2)
	EqualValues(t, items[9].Name, items[9].Quality, 0)
	EqualValues(t, items[9].Name, items[9].SellIn, -2)
	EqualValues(t, items[10].Name, items[10].Quality, 0)
	EqualValues(t, items[10].Name, items[10].SellIn, -2)
}

func Test_UpdateConjuredManaCake(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 10, 50},
		{"Conjured Mana Cake", 10, 49},
		{"Conjured Mana Cake", 12, 50},
		{"Conjured Mana Cake", 12, 49},
		{"Conjured Mana Cake", 6, 50},
		{"Conjured Mana Cake", 6, 10},
		{"Conjured Mana Cake", 4, 50},
		{"Conjured Mana Cake", 4, 5},
		{"Conjured Mana Cake", -1, 0},
		{"Conjured Mana Cake", -1, 51},
		{"Conjured Mana Cake", -1, 50},
	}

	gildedrose.UpdateQuality(items)

	EqualValues(t, items[0].Name, items[0].Quality, 49)
	EqualValues(t, items[0].Name, items[0].SellIn, 9)
	EqualValues(t, items[1].Name, items[1].Quality, 48)
	EqualValues(t, items[1].Name, items[1].SellIn, 9)
	EqualValues(t, items[2].Name, items[2].Quality, 49)
	EqualValues(t, items[2].Name, items[2].SellIn, 11)
	EqualValues(t, items[3].Name, items[3].Quality, 48)
	EqualValues(t, items[3].Name, items[3].SellIn, 11)
	EqualValues(t, items[4].Name, items[4].Quality, 49)
	EqualValues(t, items[4].Name, items[4].SellIn, 5)
	EqualValues(t, items[5].Name, items[5].Quality, 9)
	EqualValues(t, items[5].Name, items[5].SellIn, 5)
	EqualValues(t, items[6].Name, items[6].Quality, 49)
	EqualValues(t, items[6].Name, items[6].SellIn, 3)
	EqualValues(t, items[7].Name, items[7].Quality, 4)
	EqualValues(t, items[7].Name, items[7].SellIn, 3)
	EqualValues(t, items[8].Name, items[8].Quality, 0)
	EqualValues(t, items[8].Name, items[8].SellIn, -2)
	EqualValues(t, items[9].Name, items[9].Quality, 49)
	EqualValues(t, items[9].Name, items[9].SellIn, -2)
	EqualValues(t, items[10].Name, items[10].Quality, 48)
	EqualValues(t, items[10].Name, items[10].SellIn, -2)
}
