package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/development/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entries {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entrie, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entrie)

	require.Equal(t, arg.AccountID, entrie.AccountID)
	require.Equal(t, arg.Amount, entrie.Amount)

	require.NotZero(t, entrie.ID)
	require.NotZero(t, entrie.CreatedAt)
	return entrie
}
func TestCreateEntry(t *testing.T) {
	entrie := createRandomEntry(t)
	title := "Create Entrie ID # " + strconv.FormatInt(entrie.ID, 10)
	util.PrintStructColor(title, entrie)
}

func TestGetEntry(t *testing.T) {
	entrie1 := createRandomEntry(t)
	entrie2, err := testQueries.GetEntry(context.Background(), entrie1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entrie2)

	require.Equal(t, entrie1.ID, entrie2.ID)
	require.Equal(t, entrie1.AccountID, entrie2.AccountID)
	require.Equal(t, entrie1.Amount, entrie2.Amount)
	require.WithinDuration(t, entrie1.CreatedAt, entrie2.CreatedAt, time.Second)

	title := "Get Entrie ID # " + strconv.FormatInt(entrie2.ID, 10)
	util.PrintStructColor(title, entrie1)
}

func TestUpdateEntrie(t *testing.T) {
	entrie1 := createRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entrie1.ID,
		Amount: util.RandomMoney(),
	}

	entrie2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entrie2)

	require.Equal(t, entrie1.ID, entrie2.ID)
	require.Equal(t, entrie1.AccountID, entrie2.AccountID)
	require.Equal(t, arg.Amount, entrie2.Amount)
	require.WithinDuration(t, entrie1.CreatedAt, entrie2.CreatedAt, time.Second)

	title := "Update Entrie ID # " + strconv.FormatInt(entrie2.ID, 10)
	util.PrintStructColor(title, entrie2)
}

func TestDeleteEntrie(t *testing.T) {
	entrie1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entrie1.ID)
	require.NoError(t, err)

	entrie2, err := testQueries.GetEntry(context.Background(), entrie1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entrie2)

	title := "Delete Entrie ID # " + strconv.FormatInt(entrie1.ID, 10)
	util.PrintStructColor(title, entrie1)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 2; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  2,
		Offset: 0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 2)

	for _, entrie := range entries {
		require.NotEmpty(t, entrie)
	}

	title := "List Entries"
	util.PrintStructColor(title, entries)
}
