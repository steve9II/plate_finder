package search

func SearchByPart(part int) []int {
	feeds, err := RetrieveFeeds()

	var hitPlate []int
	if err != nil {
		return []int{}
	}
	for _, feed := range feeds {
		plateList, err := GetPlateListFromRangeString(feed.PlateRange)
		if err != nil {
			// 返回空切片
			return []int{}
		}
		for _, plateNumber := range plateList {
			if PartMatcher(plateNumber, part) {
				hitPlate = append(hitPlate, plateNumber)
				println(plateNumber)
			}
		}
	}
	return hitPlate
}
