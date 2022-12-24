package lobsang

import (
	"errors"
	"fmt"
)

type WorkDay struct {
	entries []Entry
}

type EntryMap map[string]Entry

func (d WorkDay) Entries() EntryMap {
	m := make(EntryMap)
	for index, entry := range d.entries {
		m[fmt.Sprintf("%v-%v", index, entry.Hash())] = entry
	}
	return m
}

func (m EntryMap) ToWorkDay() WorkDay {
	entries := make([]Entry, 0)
	for _, entry := range m {
		entries = append(entries, entry)
	}
	return WorkDay{entries: entries}
}

func (d *WorkDay) Update(id string, e Entry) error {
	entries := d.Entries()
	if _, ok := entries[id]; !ok {
		return errors.New("unknown id: " + id)
	}
	entries[id] = e

	// TODO persist workday
	entries.ToWorkDay()

	return nil
}
