package main

import (
	"encoding/json"
	"testing"

	"github.com/bbutton/trello_webhook/main/Godeps/_workspace/src/github.com/bbutton/trello_webhook/trelloData"
)

type Name struct {
	First string
	Last  string
}

type Model struct {
	Action struct {
		ID string
	}
}

var jsonNameData = "{\"First\": \"Brian\", \"Last\": \"Button\"}"

var moveColumnData = "{\"model\":{\"id\":\"53fb794a72ab28b254f3f471\",\"name\":\"Feature Team: Advanced Storage\",\"desc\":\"\",\"descData\":null,\"closed\":false,\"idOrganization\":\"5005e589b7ba66fe3d4da9a6\",\"pinned\":false,\"url\":\"https://trello.com/b/fY5JIykX/feature-team-advanced-storage\",\"shortUrl\":\"https://trello.com/b/fY5JIykX\",\"prefs\":{\"permissionLevel\":\"org\",\"voting\":\"members\",\"comments\":\"members\",\"invitations\":\"members\",\"selfJoin\":true,\"cardCovers\":true,\"cardAging\":\"regular\",\"calendarFeedEnabled\":false,\"background\":\"blue\",\"backgroundColor\":\"#0079BF\",\"backgroundImage\":null,\"backgroundImageScaled\":null,\"backgroundTile\":false,\"backgroundBrightness\":\"unknown\",\"canBePublic\":false,\"canBeOrg\":true,\"canBePrivate\":true,\"canInvite\":true},\"labelNames\":{\"green\":\"1 - Trivial\",\"yellow\":\"2 - Straightforward\",\"orange\":\"3 - Interesting\",\"red\":\"Expedite\",\"purple\":\"Platform\",\"blue\":\"\",\"sky\":\"External Requests\",\"lime\":\"Maintenance\",\"pink\":\"Features\",\"black\":\"OS - Manual Bucket Migration\"}},\"action\":{\"id\":\"56577cb613a1a7a7db4e9bcb\",\"idMemberCreator\":\"52cc8e1054e1b4ee3b64bc68\",\"data\":{\"listAfter\":{\"name\":\"Feature - Update Control UI\",\"id\":\"53fb794a72ab28b254f3f473\"},\"listBefore\":{\"name\":\"Feature - Block Storage MVP\",\"id\":\"557896e125d44aa467f376d7\"},\"board\":{\"shortLink\":\"fY5JIykX\",\"name\":\"Feature Team: Advanced Storage\",\"id\":\"53fb794a72ab28b254f3f471\"},\"card\":{\"shortLink\":\"Gsd2Np6M\",\"idShort\":122,\"name\":\"Add billing info\",\"id\":\"54b7f56a06e69aa6959b4269\",\"idList\":\"53fb794a72ab28b254f3f473\"},\"old\":{\"idList\":\"557896e125d44aa467f376d7\"}},\"type\":\"updateCard\",\"date\":\"2015-11-26T21:42:14.269Z\",\"memberCreator\":{\"id\":\"52cc8e1054e1b4ee3b64bc68\",\"avatarHash\":\"c5cbaff8af55fe12ac23de147d863bdc\",\"fullName\":\"Brian Button\",\"initials\":\"BAB\",\"username\":\"brianbuttonxp\"}}}"

func Test_ReadsSimpleNameData(t *testing.T) {
	var name Name
	err := json.Unmarshal([]byte(jsonNameData), &name)

	if err != nil {
		t.Fatalf("Failed to parse json, error is %s", err.Error())
	}

	if name.First != "Brian" || name.Last != "Button" {
		t.Fatalf("First Name was %s, Last Name was %s", name.First, name.Last)
	}
}

func Test_ActionID(t *testing.T) {
	var model Model
	err := json.Unmarshal([]byte(moveColumnData), &model)

	if err != nil {
		t.Fatalf("Failed to parse json, error is %s", err.Error())
	}

	if model.Action.ID != "56577cb613a1a7a7db4e9bcb" {
		t.Fatalf("expected: 56577cb613a1a7a7db4e9bcb, actual: %s", model.Action.ID)
	}
}

func TestGeneratedModelStruct(t *testing.T) {
	var model trelloData.Model
	err := json.Unmarshal([]byte(moveColumnData), &model)

	if err != nil {
		t.Fatalf("Failed to parse json, error is %s", err.Error())
	}

	if model.Action.ID != "56577cb613a1a7a7db4e9bcb" {
		t.Fatalf("expected: 56577cb613a1a7a7db4e9bcb, actual: %s", model.Action.ID)
	}

	if model.Action.Data.ListBefore.Name != "Feature - Block Storage MVP" {
		t.Fatalf("expected: Feature - Block Storage MVP actual: %s", model.Action.Data.ListBefore.Name)
	}
}
