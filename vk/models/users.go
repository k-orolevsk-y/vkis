package models

import "vkIntership/vk/helpers"

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Domain   string `json:"domain"`
	Bdate    string `json:"bdate"`
	City     struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Country struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	Timezone               int    `json:"timezone"`
	Photo200               string `json:"photo_200"`
	PhotoMax               string `json:"photo_max"`
	Photo200Orig           string `json:"photo_200_orig"`
	Photo400Orig           string `json:"photo_400_orig"`
	PhotoMaxOrig           string `json:"photo_max_orig"`
	PhotoId                string `json:"photo_id"`
	HasPhoto               int    `json:"has_photo"`
	HasMobile              int    `json:"has_mobile"`
	IsFriend               int    `json:"is_friend"`
	CanPost                int    `json:"can_post"`
	CanSeeAllPosts         int    `json:"can_see_all_posts"`
	CanSeeAudio            int    `json:"can_see_audio"`
	Interests              string `json:"interests"`
	Books                  string `json:"books"`
	Tv                     string `json:"tv"`
	Quotes                 string `json:"quotes"`
	About                  string `json:"about"`
	Games                  string `json:"games"`
	Movies                 string `json:"movies"`
	Activities             string `json:"activities"`
	Music                  string `json:"music"`
	CanWritePrivateMessage int    `json:"can_write_private_message"`
	CanSendFriendRequest   int    `json:"can_send_friend_request"`
	CanBeInvitedGroup      bool   `json:"can_be_invited_group"`
	MobilePhone            string `json:"mobile_phone"`
	HomePhone              string `json:"home_phone"`
	Site                   string `json:"site"`
	Status                 string `json:"status"`
	LastSeen               struct {
		Platform int `json:"platform"`
		Time     int `json:"time"`
	} `json:"last_seen"`
	Exports   []interface{} `json:"exports"`
	CropPhoto struct {
		Photo Photo `json:"photo"`
		Crop  struct {
			X  float64 `json:"x"`
			Y  int     `json:"y"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"crop"`
		Rect struct {
			X  int `json:"x"`
			Y  int `json:"y"`
			X2 int `json:"x2"`
			Y2 int `json:"y2"`
		} `json:"rect"`
	} `json:"crop_photo"`
	FollowersCount   int `json:"followers_count"`
	Blacklisted      int `json:"blacklisted"`
	BlacklistedByMe  int `json:"blacklisted_by_me"`
	IsFavorite       int `json:"is_favorite"`
	IsHiddenFromFeed int `json:"is_hidden_from_feed"`
	CommonCount      int `json:"common_count"`
	Occupation       struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"occupation"`
	Career []struct {
		CountryId int    `json:"country_id"`
		From      int    `json:"from"`
		GroupId   int    `json:"group_id"`
		Position  string `json:"position"`
	} `json:"career"`
	Military       []interface{} `json:"military"`
	University     int           `json:"university"`
	UniversityName string        `json:"university_name"`
	Faculty        int           `json:"faculty"`
	FacultyName    string        `json:"faculty_name"`
	Graduation     int           `json:"graduation"`
	HomeTown       string        `json:"home_town"`
	Relation       int           `json:"relation"`
	Personal       struct {
		Alcohol    int      `json:"alcohol"`
		InspiredBy string   `json:"inspired_by"`
		Langs      []string `json:"langs"`
		LangsFull  []struct {
			Id         int    `json:"id"`
			NativeName string `json:"native_name"`
		} `json:"langs_full"`
		LifeMain   int `json:"life_main"`
		PeopleMain int `json:"people_main"`
		Political  int `json:"political"`
		Smoking    int `json:"smoking"`
	} `json:"personal"`
	Universities []struct {
		City        int    `json:"city"`
		Country     int    `json:"country"`
		Faculty     int    `json:"faculty"`
		FacultyName string `json:"faculty_name"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"universities"`
	Schools   []interface{} `json:"schools"`
	Relatives []struct {
		Type string `json:"type"`
		Id   int    `json:"id"`
	} `json:"relatives"`
	Sex             int    `json:"sex"`
	ScreenName      string `json:"screen_name"`
	Photo50         string `json:"photo_50"`
	Photo100        string `json:"photo_100"`
	Online          int    `json:"online"`
	Verified        int    `json:"verified"`
	FriendStatus    int    `json:"friend_status"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	CanAccessClosed bool   `json:"can_access_closed"`
	IsClosed        bool   `json:"is_closed"`
}

type ResultUsersGet struct {
	Error    helpers.Error `json:"error"`
	Response []User        `json:"response"`
}
