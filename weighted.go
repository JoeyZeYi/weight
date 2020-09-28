package weight

type Weighted interface {
	GetId() uint32
	GetWeight() int
}

//初始化一个池子
func NewPool(servers []Weighted) *Load {
	new := &Load{}
	new.UpdateServers(servers)
	return new
}

type Training struct {
	Server          Weighted `json:"server"`
	Weight          int      `json:"weight"`
	CurrentWeight   int      `json:"current_weight"`
	EffectiveWeight int      `json:"effective_weight"`
}

type Load struct {
	Weighted []Weighted  `json:"servers"`
	Training []*Training `json:"weighted"`
}

func (l *Load) UpdateServers(servers []Weighted) {
	weighted := make([]*Training, 0)
	for _, v := range servers {
		w := &Training{
			Server:          v,
			Weight:          v.GetWeight(),
			CurrentWeight:   0,
			EffectiveWeight: v.GetWeight(),
		}
		weighted = append(weighted, w)
	}
	l.Training = weighted
	l.Weighted = servers
}

//remove为需要屏蔽的ID，没有的话传nil
func (l *Load) Draw(remove []uint) Weighted {
	if len(l.Training) == 0 {
		return nil
	}
	w := l.nextWeighted(l.Training, remove)
	if w == nil {
		return nil
	}
	return w.Server
}
func (l *Load) nextWeighted(servers []*Training, remove []uint) (best *Training) {
	total := 0
	for i := 0; i < len(servers); i++ {
		w := servers[i]
		if w == nil {
			continue
		}
		isFind := false
		for _, v := range remove {
			if v == uint(w.Server.GetId()) {
				isFind = true
			}
		}
		if isFind {
			continue
		}

		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight
		if w.EffectiveWeight < w.Weight {
			w.EffectiveWeight++
		}

		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}
	if best == nil {
		return best
	}
	best.CurrentWeight -= total
	return best
}
