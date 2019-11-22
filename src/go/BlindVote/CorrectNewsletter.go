package BlindVote

import "fmt"

type CorrectNewsletters struct {
	list []*Newsletter
}

func (cn *CorrectNewsletters) Add(newsletter *Newsletter) {
	cn.list = append(cn.list, newsletter)
}

func (cn *CorrectNewsletters) PrintAll() {
	for i, el := range cn.list {
		fmt.Printf("[%d] N:[%v] S:[%v]\n", i, el.N, el.S)
	}
}
