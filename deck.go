package main

import ("fmt"
		"strings"
		"io/ioutil"
		"os"
		"math/rand"
		"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}


//this is a reciever function the 
//in (d deck) the d here holds the reference of the calling object
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck,handSize int) (deck,deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) toString()(string){
	//type cast deck to slice of string []string(d)
	//joins  every string in slice with the give sep and returns a single string
	return strings.Join([]string(d),",")

}


//error is built in type
func (d deck) saveToFile(fileName string) error{
	//WriteFile takes bytes as an arg hence type cast string to []byte
	//0666 is the permission to the file
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666)
}

func newDeckFromFile(fileName string)(deck){
	bs,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		fmt.Println("Error:\n",err)
		//exits the code gracefully 
		os.Exit(-1)
	}
	jointString:=string(bs)
	s:=strings.Split(jointString,",")
	return deck(s)
		 
}


func (d deck) shuffle(){
	//refer seeding in rand before this
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for index:= range d{
		newPosition:=r.Intn(len(d)-1)
		d[index],d[newPosition]=d[newPosition],d[index]
	}
} 