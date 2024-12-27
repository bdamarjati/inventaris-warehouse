package api

import (
	"bytes"
	"encoding/json"
	"inventory/main/db"
	"inventory/main/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username: util.RandomName(),
		Password: util.RandomText(8),
		Role:     "Admin",
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/user"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	user, err := requireBodyMatchUser(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Role, user.Role)
	require.WithinDuration(t, time.Now().Local(), user.CreatedAt, time.Second)
	return user
}

func requireBodyMatchUser(body *bytes.Buffer) (db.User, error) {
	var gotUser db.User

	data, err := io.ReadAll(body)
	if err != nil {
		return gotUser, err
	}

	err = json.Unmarshal(data, &gotUser)
	return gotUser, err
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
