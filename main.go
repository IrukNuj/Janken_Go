package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Please request to '/battle/:your_hand' like this '/battle/ぱー'")
}

func Battle(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	hand := []string{"ぐー", "ちょき", "ぱー"}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(hand))

	user_hand := p.ByName("hand")
	goper_hand := hand[i]
	
	judge := Judge(user_hand, goper_hand)

	fmt.Fprintln(w, "Your hand     : ", p.ByName("hand"))
	fmt.Fprintln(w, "Gopher's hand : ", hand[i])
	fmt.Fprintln(w, judge)

}

func Judge(u string, g string) string {
	var u_hand, g_hand int

	if u == "ぐー" {
		u_hand = 0
	} else if u == "ちょき" {
		u_hand = 1
	} else  if u == "ぱー" {
		u_hand = 2
	} else {
		return "ぐー、ちょき、ぱーの中でリクエストしてね〜！"
	}

	if g == "ぐー" {
		g_hand = 0
	} else if g == "ちょき" {
		g_hand = 1
	} else  if g == "ぱー" {
		g_hand = 2
	}

	result_num := u_hand - g_hand + 3

	if result_num % 3 == 0 {
		return "あいこでしたー！"
	} else if result_num % 3 == 1 {
		return "うーん、負け！ｗｗｗ"
	} else if result_num % 3 == 2 {
		return "勝ちです。"
	}

	return "何かしらの障害が起きてます！"
}


func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/battle/:hand", Battle)
	log.Fatal(http.ListenAndServe(":8080", router))
}
