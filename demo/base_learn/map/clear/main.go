package main

type p struct {
	age int
}

func main() {
	batchSize := 2

	ps := make(map[int][]*p)
	for i := 1; i < 10; i++ {
		ps[i] = append(ps[i], &p{age: i * 10})
	}

	for k, p := range ps {
		go func(j int) { //播放
			res := make(map[int]map[int]int)
			for kk, v := range p {
				tgMap := make(map[int]int)
				tgMap[kk] = v.age

				res[j] = tgMap
				if len(res) == batchSize {
					res = nil
					res = make(map[int]map[int]int)
				}
			}
		}(k)
	}
}
