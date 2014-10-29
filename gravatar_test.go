package gravatar

import (
	"testing"
)

func TestNewGravatarFromEmail(t *testing.T) {
	email := "foo@example.com"
	emailHashed := "b48def645758b95537d4424c84d1a9ff"
	g := NewGravatarFromEmail(email)
	if g.Hash != emailHashed {
		t.Error("got hash: %q; expected: %q", g.Hash, emailHashed)
	}
}

func TestGravatarGetUrl(t *testing.T) {
	expectedUrl := "https://www.gravatar.com/avatar/b48def645758b95537d4424c84d1a9ff"
	hash := "b48def645758b95537d4424c84d1a9ff"
	g := NewGravatar()
	g.Hash = hash
	if url := g.GetUrl(); url != expectedUrl {
		t.Error("got url: %q; expected: %q", url, expectedUrl)
	}
}
