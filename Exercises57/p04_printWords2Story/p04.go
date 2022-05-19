package p04

import "fmt"

func RunP04() {

	var strNoun, strVerb, strAdj, strAdv string

	fmt.Println("Enter a noun: ")
	fmt.Scanln(&strNoun)

	fmt.Println("Enter a verb: ")
	fmt.Scanln(&strVerb)

	fmt.Println("Enter a adjective: ")
	fmt.Scanln(&strAdj)

	fmt.Println("Enter a adverb: ")
	fmt.Scanln(&strAdv)

	r := Practice04Normal(strNoun, strVerb, strAdj, strAdv)

	fmt.Println(r)

}

func Practice04Normal(strNoun, strVerb, strAdj, strAdv string) (r string) {

	r = SentancePrefix1 + strVerb + SentancePrefix2 + strAdj + " " + strNoun + " " + strAdv + SentancePrefix3

	return r
}
