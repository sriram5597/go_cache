package cache

import (
	"time"
	"log"
	"sync"
)

type CacheElement struct {
	Value  int   `json:"value"`
	Expiry int64 `json:"expiry"`
}

var goCache map[int]CacheElement = nil
var KeysCount int = 0
const MaxKeys int = 2
var mutexLock sync.RWMutex

func NewCache() {
	goCache = make(map[int]CacheElement)
}


type CacheError struct {
	Message string 
}

func (err *CacheError) Error() string {
	return err.Message
}

func GetKey(key int) int {
	if goCache == nil {
		panic("cache not initialized..")
	}

	mutexLock.RLock()
	log.Println("Acquired Read Lock")
	element, ok := goCache[key]
	log.Println("Releasing Read Lock")
	mutexLock.RUnlock()
	if !ok {
		return -1
	}

	if element.Expiry < time.Now().Unix() {
		mutexLock.Lock()
		log.Println("Acquired RW Lock")
		delete(goCache, key)
		KeysCount = KeysCount - 1
		log.Println("Releasing RW Lock")
		mutexLock.Unlock()
		return -1
	}

	return element.Value
}

func SetKey(key, value, expiry_in_seconds int) error {
	if goCache == nil {
		panic("cache not initialized...")
	} 

	mutexLock.Lock()
	log.Println("acquired RW lock")
	if KeysCount + 1 == MaxKeys {
		cacheError := CacheError{Message: "Cache memory is full"}
		log.Println("Releasing RW lock")
		mutexLock.Unlock()
		return &cacheError
	}
	_, ok := goCache[key]
	goCache[key] = CacheElement{Value: value, Expiry: time.Now().Unix() + int64(expiry_in_seconds)}
	if !ok {
		KeysCount = KeysCount + 1
	}
	log.Println("releasing RW lock")
	mutexLock.Unlock()
	log.Println(goCache)
	return nil
}
