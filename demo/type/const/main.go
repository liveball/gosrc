package main

const (
	_ byte = iota
	//ArtView 阅读
	ArtView
	//ArtReply 评论
	ArtReply
	//ArtShare 分享
	ArtShare
	//ArtCoin 硬币
	ArtCoin
	//ArtFavTBL 收藏
	ArtFavTBL
	//ArtLikeTBL 喜欢
	ArtLikeTBL
)

var (
	artTypeMap = map[byte]struct{}{
		ArtView:    {},
		ArtReply:   {},
		ArtShare:   {},
		ArtCoin:    {},
		ArtFavTBL:  {},
		ArtLikeTBL: {},
	}
)

//CheckType check article data type.
func CheckType(ty byte) bool {
	_, ok := artTypeMap[ty]
	if ok {
		return true
	}
	return false
}

func main() {
	for i := 1; i <= 6; i++ {
		println(CheckType(byte(i)))
	}
}
