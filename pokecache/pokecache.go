package pokecache

import "time"


type cacheEntry struct {
		createdAt time.Time
		val []byte
}

func InitCache() map[string]cacheEntry {
		cache := make(map[string]cacheEntry)
		return cache
}


func Add() {

}

func Get() {

}

func ReapLoop() {

}
