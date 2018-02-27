package todolist

import (
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"
	"github.com/ukdave/goswagger-todolist/gen/models"
)

type List struct {
	items     map[int64]*models.Item
	lastID    int64
	itemsLock *sync.Mutex
}

func NewList() *List {
	return &List{
		items:     make(map[int64]*models.Item),
		lastID:    0,
		itemsLock: &sync.Mutex{},
	}
}

func (l *List) newItemID() int64 {
	return atomic.AddInt64(&l.lastID, 1)
}

func (l *List) AddItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	l.itemsLock.Lock()
	defer l.itemsLock.Unlock()

	newID := l.newItemID()
	item.ID = newID
	l.items[newID] = item

	return nil
}

func (l *List) UpdateItem(id int64, item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	l.itemsLock.Lock()
	defer l.itemsLock.Unlock()

	_, exists := l.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	l.items[id] = item
	return nil
}

func (l *List) DeleteItem(id int64) error {
	l.itemsLock.Lock()
	defer l.itemsLock.Unlock()

	_, exists := l.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(l.items, id)
	return nil
}

func (l *List) AllItems(limit int) []*models.Item {
	if limit <= 0 {
		limit = len(l.items)
	}
	result := []*models.Item{}
	for _, item := range l.items {
		if len(result) >= int(limit) {
			break
		}
		result = append(result, item)
	}
	return result
}

func (l *List) GetItem(id int64) *models.Item {
	return l.items[id]
}
