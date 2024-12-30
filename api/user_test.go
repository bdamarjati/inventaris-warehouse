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

func createRandomUser() (db.User, error) {
	arg := createUserRequest{
		Username: util.RandomName(),
		Password: util.RandomText(8),
		Role:     "Admin",
	}

	data, err := json.Marshal(arg)
	if err != nil {
		return db.User{}, err
	}

	url := "/user"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return db.User{}, err
	}

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		return db.User{}, fmt.Errorf("error expected %d actual %d\n", http.StatusOK, recorder.Code)
	}

	user, err := requireBodyMatchUser(recorder.Body)
	if err != nil {
		return db.User{}, err
	}

	if arg.Username != user.Username {
		return db.User{}, fmt.Errorf("error expected %s actual %s\n", arg.Username, user.Username)
	}
	if arg.Role != user.Role {
		return db.User{}, fmt.Errorf("error expected %s actual %s\n", arg.Role, user.Role)
	}
	return db.User{
		Username:  user.Username,
		Password:  arg.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

func getAccessToken(arg loginUserRequest) (loginUserResponse, error) {
	data, err := json.Marshal(arg)
	if err != nil {
		return loginUserResponse{}, err
	}

	url := "/login"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return loginUserResponse{}, err
	}

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		return loginUserResponse{}, fmt.Errorf("error expected %d actual %d\n", http.StatusOK, recorder.Code)
	}

	user := loginUserResponse{}
	res, err := io.ReadAll(recorder.Body)
	if err != nil {
		return loginUserResponse{}, err
	}

	err = json.Unmarshal(res, &user)
	if err != nil {
		return loginUserResponse{}, err
	}
	return user, nil
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
	user, err := createRandomUser()
	require.NoError(t, err)
	require.WithinDuration(t, time.Now().Local(), user.CreatedAt, time.Second)
}

func TestLoginUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NotEmpty(t, user1)
	require.NoError(t, err)

	arg := loginUserRequest{
		Username: user1.Username,
		Password: user1.Password,
	}

	user2, err := getAccessToken(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.NotNil(t, user2.AccessToken)
	require.Equal(t, user1.Username, user2.User.Username)
	require.Equal(t, user1.Role, user2.User.Role)
}
