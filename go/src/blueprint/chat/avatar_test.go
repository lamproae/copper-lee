package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrorNoAvatarURL when no value present")
	}

	testUrl := "http://en.gravatar.com/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL shoudl return no error when value present")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURL should return correct URL")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvitar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "de0ab3407a0c7b744225a9e5911c4990"}
	url, err := gravatarAvitar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvitar.GetAvatarURL should not return an error")
	}

	if url != "//www.gravatar.com/avatar/de0ab3407a0c7b744225a9e5911c4990" {
		t.Errorf("GravatarAvitar.GetAvatarURL wrongly return %s", url)
	}
}
