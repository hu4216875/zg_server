package template

type OreInfo struct {
	Id    uint32
	Total uint32
}

type OreInfoTemplate struct {
	data map[uint32]*OreInfo
}

func (i *OreInfoTemplate) load() {
	i.data = make(map[uint32]*OreInfo)
	rf := readRf(OreInfo{})
	for k := 0; k < rf.NumRecord(); k++ {
		oreDistrict := rf.Record(k).(*OreInfo)
		i.data[oreDistrict.Id] = oreDistrict
	}
}

// GetOreDistrict 获得指定矿洞
func (i *OreInfoTemplate) GetOreDistrict(id uint32) *OreInfo {
	if ret, ok := i.data[id]; ok {
		return ret
	}
	return nil
}

// GetAll 获取所有矿洞
func (i *OreInfoTemplate) GetAll() []*OreInfo {
	var ret []*OreInfo
	for _, data := range i.data {
		ret = append(ret, data)
	}
	return ret
}
