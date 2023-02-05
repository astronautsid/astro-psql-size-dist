package info

func (info *Label) CountSize() uint64 {
	var count uint64 = 0
	for _, r := range info.Tables {
		count += r.Size
	}

	return count
}
