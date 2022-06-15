package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Provider string             `bson:"provider"`
	Token    string             `bson:"token"`
}

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubUserData struct {
	Login          string `json:"login"`
	Id             string `json:"id"`
	NodeId         string `json:"node_id"`
	AvatarUrl      string `json:"avatar_url"`
	GravatarId     string `json:"gravatar_id"`
	Url            string `json:"url"`
	HtmlUrl        string `json:"html_url"`
	FollowerUrl    string `json:"followers_url"`
	FollowingUrl   string `json:"following_url"`
	Gists          string `json:"gists_url"`
	Starred        string `json:"starred_url"`
	Subscriptions  string `json:"subscriptions_url"`
	Organizations  string `json:"organizations_url"`
	Repos          string `json:"repos_url"`
	Events         string `json:"events_url"`
	ReceivedEvents string `json:"received_events_url"`
	Type           string `json:"type"`
	SiteAdmin      string `json:"site_admin"`
	Name           string `json:"name"`
	Company        string `json:"company"`
	Blog           string `json:"blog"`
	Location       string `json:"location"`
	Email          string `json:"email"`
	HireAble       string `json:"hireable"`
	Bio            string `json:"bio"`
	Twitter        string `json:"twitter_username"`
	PublicRepos    string `json:"public_repos"`
	PublicGists    string `json:"public_gists"`
	Follower       string `json:"followers"`
	Following      string `json:"following"`
	Created        string `json:"created_at"`
	Updated        string `json:"updated_at"`
}
