package database

import (
	"time"
)

type TTLTypeMap map[string]uint

type ExtensionExpire struct {
	ttl TTLTypeMap
}

func NewExtensionExpire() ExtensionExpire {
	return ExtensionExpire{
		ttl: make(TTLTypeMap),
	}
}

func CurrentUnixTime() uint {
	return uint(time.Now().Unix())
}

func (ext *ExtensionExpire) setTTL(key string, second uint) {
	ext.ttl[key] = CurrentUnixTime() + second
}

func (ext *ExtensionExpire) GetTTL(key string) uint {
	v, _ := ext.ttl[key]
	return v
}

func (ext *ExtensionExpire) hasTTL(key string) bool {
	_, ok := ext.ttl[key]
	return ok
}

func (ext *ExtensionExpire) Diff(key string) uint {
	return ext.GetTTL(key) - CurrentUnixTime()
}

func (ext *ExtensionExpire) isExpire(key string) bool {
	v, ok := ext.ttl[key]
	if ok && CurrentUnixTime() >= v {
		return true
	}
	return false
}

func (ext *ExtensionExpire) delTLL(key string) {
	delete(ext.ttl, key)
}

func (ext *ExtensionExpire) getExpiredKeys() []string {
	expiredKeys := []string{}
	for k := range ext.ttl {
		if ext.isExpire(k) {
			_ = append(expiredKeys, k)
		}
	}
	return expiredKeys
}
