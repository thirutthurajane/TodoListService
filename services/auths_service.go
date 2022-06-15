package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"todoservice/configs"
	"todoservice/models"
)

var gitHubEndPoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

func exchangeGithubToken(code string) models.GithubAccessTokenResponse {

	cid, secret := "", ""

	switch configs.Configuration.Server.Env {
	case "Develop":
		cid, secret = configs.Configuration.Server.Develop.Github.Cid, configs.Configuration.Server.Develop.Github.Secret
	case "Production":
		cid, secret = configs.Configuration.Server.Production.Github.Cid, configs.Configuration.Server.Production.Github.Secret
	default:
		cid, secret = configs.Configuration.Server.Develop.Github.Cid, configs.Configuration.Server.Develop.Github.Secret
	}

	reqBody := map[string]string{
		"client_id":     cid,
		"client_secret": secret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(reqBody)

	req, reqerr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)

	if reqerr != nil {
		errors.Wrap(reqerr, reqerr.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		errors.Wrap(resperr, resperr.Error())
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	tokenResponse := models.GithubAccessTokenResponse{}
	err := json.Unmarshal(respbody, &tokenResponse)
	if err != nil {
		errors.Wrap(err, err.Error())
		return models.GithubAccessTokenResponse{}
	}
	return tokenResponse
}

func getGithubData(accessToken string) models.GithubUserData {
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		errors.Wrap(resperr, resperr.Error())
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	dataResponse := models.GithubUserData{}
	err := json.Unmarshal(respbody, &dataResponse)
	if err != nil {
		errors.Wrap(err, err.Error())
	}

	return dataResponse
}

func HandleCallback(provider string, code string) (*mongo.UpdateResult, error) {
	accessToken := exchangeGithubToken(code)
	gitData := getGithubData(accessToken.AccessToken)
	ctx, _ := configs.GetContext()
	client, err := configs.GetConn()
	if err != nil {
		return nil, errors.Wrap(err, "Configuration error")
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Connection error")
	}
	defer client.Disconnect(ctx)
	database := client.Database(configs.Configuration.Server.DbName)
	userCollection := database.Collection("Users")
	result, err := userCollection.UpdateOne(ctx,
		bson.D{{"email", gitData.Email}, {"provider", provider}},
		bson.D{{"$set", bson.D{{"token", accessToken.AccessToken}}}},
		options.Update().SetUpsert(true))
	if err != nil {
		return nil, errors.Wrap(err, "Login error")
	}

	return result, nil
}
