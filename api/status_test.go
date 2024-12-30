package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inventory/main/db"
	"inventory/main/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStatus(t *testing.T) db.RefStatus {
	arg := db.RefStatus{
		Description: util.RandomStatus(),
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/status"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	status, err := requireBodyMatchStatus(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, status)
	require.Equal(t, arg.Description, status.Description)
	require.WithinDuration(t, time.Now().Local(), status.CreatedAt, time.Second)

	return status
}

func getStatus(id int64) (db.RefStatus, error) {
	status := db.RefStatus{}
	url := fmt.Sprintf("/status/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return status, err
	}

	err = addAuthorization(request)
	if err != nil {
		return status, err
	}

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusInternalServerError {
		return status, fmt.Errorf("internal testServer error: %v", err)
	}

	status, err = requireBodyMatchStatus(recorder.Body)
	return status, err
}

func requireBodyMatchStatus(body *bytes.Buffer) (db.RefStatus, error) {
	var gotStatus db.RefStatus

	data, err := io.ReadAll(body)
	if err != nil {
		return gotStatus, err
	}

	err = json.Unmarshal(data, &gotStatus)
	return gotStatus, err
}

func TestCreateStatus(t *testing.T) {
	createRandomStatus(t)
}

func TestGetStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	status2, err := getStatus(status1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, status2)

	require.Equal(t, status1.ID, status2.ID)
	require.Equal(t, status1.Description, status2.Description)
	require.WithinDuration(t, status1.CreatedAt, status2.CreatedAt, time.Second)
}

func TestListStatus(t *testing.T) {
	for range 3 {
		createRandomStatus(t)
	}

	arg := listStatusRequest{
		Size: 3,
		Page: 1,
	}

	url := fmt.Sprintf("/statuses/%d/%d", arg.Size, arg.Page)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	results, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	statuses := []db.RefStatus{}
	err = json.Unmarshal(results, &statuses)
	require.NoError(t, err)
	require.NotEmpty(t, statuses)
	require.Len(t, statuses, arg.Size)
}

func TestUpdateStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	arg := updateStatusRequest{
		ID:          status1.ID,
		Description: util.RandomStatus(),
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/status"
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	status2, err := getStatus(status1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, status2)

	require.Equal(t, status1.ID, status2.ID)
	require.Equal(t, arg.Description, status2.Description)
	require.WithinDuration(t, status1.CreatedAt, status2.CreatedAt, time.Second)
}

func TestDeleteStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	url := fmt.Sprintf("/status/%d", status1.ID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	status2, err := getStatus(status1.ID)
	require.Error(t, err)
	require.Empty(t, status2)
}
