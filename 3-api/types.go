package main

import "math/rand"

type Playlist struct {
    ID int `json:"id"`
    PlaylistName string `json:"playlistName"`
    Songs []string `json:"songs"`
}

func NewPlaylist(name string) *Playlist {
    return &Playlist{
        ID: rand.Intn(1000),
        PlaylistName: name,
        Songs: []string{},
    }
}

type Account struct {
    ID int `json:"id"`
    Username string `json:"username"`
    Playlists []Playlist `json:"playlist"`
}

func NewAccount(u string) *Account {
    return &Account{
        ID: rand.Intn(1000),
        Username: u,
        Playlists: []Playlist{},
    }
}
