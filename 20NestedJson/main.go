package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IssueListResponse struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand    string `json:"expand"`
		ID        string `json:"id"`
		Self      string `json:"self"`
		Key       string `json:"key"`
		Changelog struct {
			StartAt    int `json:"startAt"`
			MaxResults int `json:"maxResults"`
			Total      int `json:"total"`
			Histories  []struct {
				ID     string `json:"id"`
				Author struct {
					Self         string `json:"self"`
					AccountID    string `json:"accountId"`
					EmailAddress string `json:"emailAddress"`
					AvatarUrls   struct {
						Four8X48  string `json:"48x48"`
						Two4X24   string `json:"24x24"`
						One6X16   string `json:"16x16"`
						Three2X32 string `json:"32x32"`
					} `json:"avatarUrls"`
					DisplayName string `json:"displayName"`
					Active      bool   `json:"active"`
					TimeZone    string `json:"timeZone"`
					AccountType string `json:"accountType"`
				} `json:"author"`
				Created string `json:"created"`
				Items   []struct {
					Field      string `json:"field"`
					Fieldtype  string `json:"fieldtype"`
					From       string `json:"from"`
					FromString string `json:"fromString"`
					To         string `json:"to"`
				} `json:"items"`
			} `json:"histories"`
		} `json:"changelog"`
		Fields struct {
			Statuscategorychangedate string `json:"statuscategorychangedate"`
			Issuetype                struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				Subtask        bool   `json:"subtask"`
				AvatarID       int    `json:"avatarId"`
				HierarchyLevel int    `json:"hierarchyLevel"`
			} `json:"issuetype"`
			Timespent interface{} `json:"timespent"`
			Project   struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Key            string `json:"key"`
				Name           string `json:"name"`
				ProjectTypeKey string `json:"projectTypeKey"`
				Simplified     bool   `json:"simplified"`
				AvatarUrls     struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
			} `json:"project"`
			FixVersions []struct {
				Self     string `json:"self"`
				ID       string `json:"id"`
				Name     string `json:"name"`
				Archived bool   `json:"archived"`
				Released bool   `json:"released"`
			} `json:"fixVersions"`
			Aggregatetimespent interface{} `json:"aggregatetimespent"`
			Resolution         interface{} `json:"resolution"`
			Resolutiondate     interface{} `json:"resolutiondate"`
			Workratio          int         `json:"workratio"`
			Watches            struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			LastViewed       interface{} `json:"lastViewed"`
			Created          string      `json:"created"`
			Customfield10020 interface{} `json:"customfield_10020"`
			Customfield10021 interface{} `json:"customfield_10021"`
			Customfield10022 interface{} `json:"customfield_10022"`
			Customfield10023 interface{} `json:"customfield_10023"`
			Priority         struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Customfield10024 interface{}   `json:"customfield_10024"`
			Customfield10025 interface{}   `json:"customfield_10025"`
			Labels           []interface{} `json:"labels"`
			Customfield10026 interface{}   `json:"customfield_10026"`
			Customfield10016 interface{}   `json:"customfield_10016"`
			Customfield10017 interface{}   `json:"customfield_10017"`
			Customfield10018 struct {
				HasEpicLinkFieldDependency bool `json:"hasEpicLinkFieldDependency"`
				ShowField                  bool `json:"showField"`
				NonEditableReason          struct {
					Reason  string `json:"reason"`
					Message string `json:"message"`
				} `json:"nonEditableReason"`
			} `json:"customfield_10018"`
			Customfield10019              string        `json:"customfield_10019"`
			Timeestimate                  interface{}   `json:"timeestimate"`
			Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
			Versions                      []interface{} `json:"versions"`
			Issuelinks                    []interface{} `json:"issuelinks"`
			Assignee                      struct {
				Self         string `json:"self"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
				AccountType string `json:"accountType"`
			} `json:"assignee"`
			Updated string `json:"updated"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Components            []interface{} `json:"components"`
			Timeoriginalestimate  interface{}   `json:"timeoriginalestimate"`
			Description           interface{}   `json:"description"`
			Customfield10010      interface{}   `json:"customfield_10010"`
			Customfield10014      interface{}   `json:"customfield_10014"`
			Customfield10015      interface{}   `json:"customfield_10015"`
			Customfield10005      interface{}   `json:"customfield_10005"`
			Customfield10006      interface{}   `json:"customfield_10006"`
			Customfield10007      interface{}   `json:"customfield_10007"`
			Security              interface{}   `json:"security"`
			Customfield10008      interface{}   `json:"customfield_10008"`
			Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
			Customfield10009      interface{}   `json:"customfield_10009"`
			Summary               string        `json:"summary"`
			Creator               struct {
				Self         string `json:"self"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
				AccountType string `json:"accountType"`
			} `json:"creator"`
			Subtasks []interface{} `json:"subtasks"`
			Reporter struct {
				Self         string `json:"self"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
				AccountType string `json:"accountType"`
			} `json:"reporter"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Customfield10000 string      `json:"customfield_10000"`
			Customfield10001 interface{} `json:"customfield_10001"`
			Customfield10002 interface{} `json:"customfield_10002"`
			Customfield10003 interface{} `json:"customfield_10003"`
			Customfield10004 interface{} `json:"customfield_10004"`
			Environment      interface{} `json:"environment"`
			Duedate          interface{} `json:"duedate"`
			Progress         struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Votes struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
		} `json:"fields"`
	} `json:"issues"`
}

type JIRAKey struct {
	JIRAID  string
	History string
}

func main() {
	url := "https://techcody.atlassian.net/rest/api/latest/search?jql=project=%22SKP%22&maxResults=100&expand=changelog"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic dGVjaGNvZHlpbmRAZ21haWwuY29tOnkyUTB6cVlQTWZQTjFiMXc0QVZ5Q0MzOA==")
	req.Header.Add("Cookie", "atlassian.xsrf.token=842afbd8-5ae8-4d51-aa32-d03c89cdb252_abd21892780a86f45a3f7df2d1eeecb9547b771c_lin")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	u := IssueListResponse{}

	json.Unmarshal([]byte(body), &u)

	for i := 0; i < len(u.Issues); i++ {

		for j := 0; j < len(u.Issues[i].Changelog.Histories); j++ {

			fmt.Printf("Issue Key is %s and History ID is %s \n", u.Issues[i].Key, u.Issues[i].Changelog.Histories[j].ID)

			jiraJson := JIRAKey{JIRAID: u.Issues[i].Key, History: u.Issues[i].Changelog.Histories[j].ID}

			MongoDBInsert(&jiraJson)

		}

	}

}

func MongoDBInsert(body *JIRAKey) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://svcrm:svcrm@cluster0.yn4cf.mongodb.net/techcody?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("techcody")
	jiracolllection := database.Collection("training")

	insertresult, err := jiracolllection.InsertOne(ctx, body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertresult.InsertedID)

}
