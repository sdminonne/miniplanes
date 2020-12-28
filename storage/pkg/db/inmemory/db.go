/*

MIT License

Copyright (c) 2019 Amadeus s.a.s.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package inmemory

import (
	"fmt"
	"sync"

	"github.com/amadeusitgroup/miniplanes/storage/pkg/gen/models"
	"github.com/jinzhu/copier"
)

// DB defines db interface
type DB struct {
	airlines  []*models.Airline
	airports  []*models.Airport
	schedules []*models.Schedule
}

var (
	instance DB
	once     sync.Once
)

// TheDB return the in memory DB
func TheDB() *DB {
	once.Do(func() {
		instance = DB{}
	})
	return &instance
}

// Ping returns always OK
func (m *DB) Ping() error {
	return nil
}

// DialString returns an empty string
func (m *DB) DialString() string {
	return ""
}

// GetAirlines returns all airlines stored so fa
func (m *DB) GetAirlines() ([]*models.Airline, error) {
	return m.airlines, nil
}

// InsertAirline insert airline if is not already stored in mongo db
func (m *DB) InsertAirline(a *models.Airline) (*models.Airline, error) {
	m.airlines = append(m.airlines, a) // TODO checks whether a is already present
	return a, nil
}

// GetAirports returns the airports stored so far
func (m *DB) GetAirports() ([]*models.Airport, error) {
	return m.airports, nil
}

// InsertAirport insert airport if is not already stored in mongo db
func (m *DB) InsertAirport(a *models.Airport) (*models.Airport, error) {
	m.airports = append(m.airports, a)
	return a, nil
}

// InsertSchedule inserts schedule
func (m *DB) InsertSchedule(s *models.Schedule) (*models.Schedule, error) {
	m.schedules = append(m.schedules, s)
	return s, nil
}

// GetSchedules returns schedules stored so far
func (m *DB) GetSchedules() ([]*models.Schedule, error) {
	return m.schedules, nil
}

// GetSchedule by Id return schedule with specified id (if any)
func (m *DB) GetSchedule(id int64) (*models.Schedule, error) {
	for i := range m.schedules {
		if *m.schedules[i].ScheduleID == id {
			return m.schedules[i], nil
		}
	}
	return nil, fmt.Errorf("Couldn't find schedule with ID %d", id)
}

// UpdateSchedule updates scheduleID with supplied schedule
func (m *DB) UpdateSchedule(id int64, schedule *models.Schedule) (*models.Schedule, error) {
	for i := range m.schedules {
		if *m.schedules[i].ScheduleID == id {
			copier.Copy(m.schedules[i], schedule)
			return m.schedules[i], nil
		}
	}
	return nil, fmt.Errorf("Couldn't find schedule with ID %d", id)
}

// DeleteSchedule deletes a schedule
func (m *DB) DeleteSchedule(id int64) error {
	for i := range m.schedules {
		if *m.schedules[i].ScheduleID == id {
			m.schedules = append(m.schedules[:i], m.schedules[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find schedule with ID %d", id)
}
