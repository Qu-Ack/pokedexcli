package pokecache

import (
	"fmt"
	"time"
)


type cacheEntry struct {
		createdAt time.Time
		val []byte
}

var cache = make(map[string]cacheEntry)
func InitCache() map[string]cacheEntry {
		return cache
}


func Add(key string, value []byte) error {
		cache[key] = cacheEntry{
				createdAt: time.Now(),
				val: value,
		}	
		return nil
}

func Get(key string) cacheEntry {
		return cache[key] 
}

func ReapLoop(interval int) {
		ticker := time.NewTicker(time.Second)				
		fmt.Println(interval)
		for {
				t:= ticker.C
				fmt.Printf("the time is : %v" , t)

		}
}
