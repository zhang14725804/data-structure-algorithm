package linkedlist

import "fmt"

type songDoubly struct {
	name     string
	previous *songDoubly
	next     *songDoubly
}

type playlistDoubly struct {
	name       string
	head       *songDoubly
	tail       *songDoubly
	nowPlaying *songDoubly
}

func createPlaylistDoubly(name string) *playlistDoubly {
	return &playlistDoubly{
		name: name,
	}
}

func (p *playlistDoubly) addSong(name string) error {
	s := &songDoubly{
		name: name,
	}
	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	p.tail = s
	return nil
}

func (p *playlistDoubly) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("playlistDoubly is empty")
		return nil
	}
	//
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *playlistDoubly) startPlaying() *songDoubly {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlistDoubly) nextSong() *songDoubly {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func (p *playlistDoubly) previousSong() *songDoubly {
	p.nowPlaying = p.nowPlaying.previous
	return p.nowPlaying
}

// DoublyLinkedList 测试
func DoublyLinkedList() {
	playlistName := "myPlaylist"
	myPlaylist := createPlaylistDoubly(playlistName)
	myPlaylist.addSong("Ophelia")
	myPlaylist.addSong("Shape of you")
	myPlaylist.addSong("Stubborn Love")
	myPlaylist.addSong("Feels")
	fmt.Println("Showing all songs in playlistDoubly...")
	myPlaylist.showAllSongs()

	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s\n", myPlaylist.nowPlaying.name)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("Changing next songDoubly...")
	fmt.Printf("Now playing: %s\n", myPlaylist.nowPlaying.name)
	fmt.Println()
	myPlaylist.nextSong()
	fmt.Println("Changing next songDoubly...")
	fmt.Printf("Now playing: %s\n", myPlaylist.nowPlaying.name)
	fmt.Println()

	myPlaylist.previousSong()
	fmt.Println("Changing previous songDoubly...")
	fmt.Printf("Now playing: %s\n", myPlaylist.nowPlaying.name)
	fmt.Println()
	myPlaylist.previousSong()
	fmt.Println("Changing previous songDoubly...")
	fmt.Printf("Now playing: %s\n", myPlaylist.nowPlaying.name)
}
