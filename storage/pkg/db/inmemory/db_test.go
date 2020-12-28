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
	"reflect"
	"testing"

	"github.com/amadeusitgroup/miniplanes/storage/pkg/gen/models"
)

func TestDB_GetAirlines(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.Airline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			got, err := m.GetAirlines()
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetAirlines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetAirlines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_InsertAirline(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	type args struct {
		a *models.Airline
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Airline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			got, err := m.InsertAirline(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.InsertAirline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.InsertAirline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_InsertSchedule(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	type args struct {
		s *models.Schedule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			got, err := m.InsertSchedule(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.InsertSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.InsertSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetSchedule(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			got, err := m.GetSchedule(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateSchedule(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	type args struct {
		id       int64
		schedule *models.Schedule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			got, err := m.UpdateSchedule(tt.args.id, tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.UpdateSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_DeleteSchedule(t *testing.T) {
	type fields struct {
		airlines  []*models.Airline
		airports  []*models.Airport
		schedules []*models.Schedule
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DB{
				airlines:  tt.fields.airlines,
				airports:  tt.fields.airports,
				schedules: tt.fields.schedules,
			}
			if err := m.DeleteSchedule(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DB.DeleteSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
