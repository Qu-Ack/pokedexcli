package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
		createdAt time.Time 
		Val []byte
}

var Cache = make(map[string]cacheEntry)
func InitCache() map[string]cacheEntry {
		return Cache
}


func Add(key string, value []byte, mu *sync.Mutex) error {
		mu.Lock()
		Cache[key] = cacheEntry{
				createdAt: time.Now(),
				Val: value,
		}	
		defer mu.Unlock()
		return nil
}
func Get(key string, mu *sync.Mutex) cacheEntry {
		mu.Lock()
		defer mu.Unlock()
		return Cache[key] 
}

func ReapLoop(interval time.Duration) {
		ticker := time.NewTicker(time.Second)				
		go func () {
				for {
						t:= ticker.C
						time := <-t
						for key, entry := range Cache {
								duration := time.Sub(entry.createdAt)
								if interval - duration < 0 {
										delete(Cache , key)	
								}
						}	

				}

		}()
}
