package memorystorage

import (
	"context"
	"fmt"
	"testing"

	"github.com/dingowd/otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	stor := New()
	stor.Connect(context.Background(), "")
	e := storage.Event{
		Owner: 1, Title: "Check", Descr: "Check",
		StartDate: "2022-03-08", StartTime: "00:00:00", EndDate: "2022-03-17", EndTime: "00:00:00",
	}
	e2 := storage.Event{
		Owner: 1, Title: "Update", Descr: "Update",
		StartDate: "2022-03-09", StartTime: "12:00:00", EndDate: "2022-03-17", EndTime: "22:00:00",
	}

	t.Run("Create", func(t *testing.T) {
		err := stor.Create(e)
		require.NoError(t, err)
	})
	t.Run("IsEventExist", func(t *testing.T) {
		is, err := stor.IsEventExist(e)
		require.NoError(t, err)
		require.Equal(t, true, is)
	})
	fmt.Println(stor.Events)

	t.Run("Update", func(t *testing.T) {
		err := stor.Update(0, e2)
		require.NoError(t, err)
	})
	fmt.Println(stor.Events)
	t.Run("Update", func(t *testing.T) {
		err := stor.Update(0, e)
		require.NoError(t, err)
	})
	fmt.Println(stor.Events)
	t.Run("Delete", func(t *testing.T) {
		err := stor.Delete(0)
		require.NoError(t, err)
	})
	stor.Create(e)
	events := make([]storage.Event, 0)
	events = append(events, e)
	fmt.Println(events)
	t.Run("GetIntervalEvent", func(t *testing.T) {
		inter, _ := stor.GetIntervalEvent("2022-03-07", 168)
		require.Equal(t, events, inter)
	})
}
