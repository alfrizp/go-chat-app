package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar struct
type AuthAvatar struct{}

// UseAuthAvatar type is AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL function in AuthAvatar
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar struct
type GravatarAvatar struct{}

// UseGravatar types is GravatarAvatar
var UseGravatar GravatarAvatar

// GetAvatarURL function in GravatarAvatar
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userID, ok := c.userData["userID"]; ok {
		if useridStr, ok := userID.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar struct
type FileSystemAvatar struct{}

// UseFileSystemAvatar us FileSystemAvatar
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL function in FileSystemAvatar
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userID, ok := c.userData["userID"]; ok {
		if useridStr, ok := userID.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}

				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
