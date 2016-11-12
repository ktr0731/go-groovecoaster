package groovecoaster

// Statistics is the structure that represent stage statistics of all musics in Groove Coaster
// type Statistics struct {
// 	AverageScore int All          int
// 	Clear        int
// 	Fullchain    int
// 	Nomiss       int
// 	Perfect      int
// 	S            int
// 	Ss           int
// 	Sss          int
// }
//
// // Personal is the structure that represent personal configration in Groove Coaster
// type Personal struct {
// 	Statistics        *Statistics
// 	Avatar            string
// 	FriendApplication bool
// 	Level             int
// 	Name              string
// 	Rank              int
// 	Title             string
// 	TotalMusic        int
// 	TotalPlayMusic    int
// 	TotalScore        int
// 	TotalTrophy       int
// 	TrophyRank        int
// }

// FetchPersonal fetch player.json from mypage in GrooveCoaster
// func FetchPersonal() *Personal {
// 	return &Personal{}
// }

// Personal hgoe
func (p *APIClient) Personal() string {
	return "hoge"
}
