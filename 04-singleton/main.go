package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

// GetSingletonDatabase returns a singelton pointer to an instance of the db
func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("capitals.txt")
		db := singletonDatabase{caps}
		if e == nil {
			db.capitals = caps
		} else {
			log.Fatal(e)
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("seoul")
	fmt.Println("population of Seoul = ", pop)
}
