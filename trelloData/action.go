package trelloData

// Model is taken (borrowed) from VojtechVitek's go-trello package to allow easy parsing of incoming json
type Model struct {
	Action struct {
		ID              string `json:"id"`
		IDMemberCreator string `json:"idMemberCreator"`
		Data            struct {
			DateLastEdited string `json:"dateLastEdited"`
			ListBefore     struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"listBefore"`
			ListAfter struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"listAfter"`
			CheckItem struct {
				ID    string `json:"id"`
				State string `json:"state"`
				Name  string `json:"name"`
			} `json:"checkItem"`
			CheckList struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"checklist"`
			List struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"list"`
			TextData struct {
				Emoji struct{} `json:"emoji"`
			} `json:"textData"`
			Board struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				ShortLink string `json:"shortLink"`
			} `json:"board"`
			Card struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				ShortLink string `json:"shortLink"`
				IDShort   int    `json:"idShort"`
			} `json:"card"`
			Text string `json:"text"`
		} `json:"data"`
		Type          string `json:"type"`
		Date          string `json:"date"`
		MemberCreator struct {
			ID         string `json:"id"`
			AvatarHash string `json:"avatarHash"`
			FullName   string `json:"fullName"`
			Initials   string `json:"initials"`
			Username   string `json:"username"`
		} `json:"memberCreator"`
	}
}
