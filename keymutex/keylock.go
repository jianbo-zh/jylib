package keymutex

import "sync"

var klock = NewKeyMutex()

func Lock(key any) {
	klock.Lock(key)
}

func Unlock(key any) {
	klock.Unlock(key)
}

type mutex struct {
	sync.Mutex
	count int
}
type KeyMutex struct {
	locker sync.Mutex
	subs   map[any]*mutex
}

func NewKeyMutex() *KeyMutex {
	return &KeyMutex{
		subs: map[any]*mutex{},
	}
}

func (l *KeyMutex) Lock(key any) {
	l.locker.Lock()
	sub := l.subs[key]
	if sub == nil {
		sub = &mutex{}
		l.subs[key] = sub
	}
	sub.count++
	l.locker.Unlock()
	sub.Lock()
}

func (l *KeyMutex) Unlock(key any) {
	l.locker.Lock()
	sub := l.subs[key]
	sub.Unlock()
	sub.count--
	if sub.count == 0 {
		delete(l.subs, key)
	}
	l.locker.Unlock()
}
