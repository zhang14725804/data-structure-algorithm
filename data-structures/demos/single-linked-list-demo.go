package linkedlist

import "fmt"

/*
	Append(t) adds an Item t to the end of the linked list
	Insert(i, t) adds an Item t at position i
	RemoveAt(i) removes a node at position i
	IndexOf() returns the position of the Item t
	IsEmpty() returns true if the list is empty
	Size() returns the linked list size
	String() returns a string representation of the list
	Head() returns the first node, so we can iterate on it

	一个简易播放器
	https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9
	Create a playlist
	Add song
	Show all songs
	Start playing
	Next/Previous Song
*/

type song struct {
	Name   string
	Artist string
	Next   *song
}

type playlist struct {
	Name       string
	Head       *song
	NowPlaying *song
}

// 创建
func createPlaylist(name string) *playlist {
	return &playlist{
		Name: name,
	}
}

// 尾插入
func (p *playlist) addSong(name, artist string) error {
	s := &song{
		Name:   name,
		Artist: artist,
	}
	if p.Head == nil {
		p.Head = s
	} else {
		currendNode := p.Head
		for currendNode.Next != nil {
			currendNode = currendNode.Next
		}
		currendNode.Next = s
	}
	return nil
}

// 删除节点
func (p *playlist) removeSong(name string) error {
	currentNode := p.Head
	if currentNode.Name == name {
		h := currentNode.Next
		p.Head = h
	} else {
		for currentNode.Next != nil {
			node := currentNode
			currentNode = currentNode.Next
			if currentNode.Name == name {
				node.Next = currentNode.Next
			}
		}
	}
	return nil
}
func (p *playlist) showAllSongs() error {
	currentNode := p.Head
	if currentNode == nil {
		fmt.Println("playlist is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *playlist) startPlaying() *song {
	p.NowPlaying = p.Head
	return p.NowPlaying
}

func (p *playlist) nextSong() *song {
	p.NowPlaying = p.NowPlaying.Next
	return p.NowPlaying
}

// SingleLinkedListDemo 测试简单链表
func SingleLinkedListDemo() {
	playlistName := "我的最爱"
	myPlaylist := createPlaylist(playlistName)
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.showAllSongs()
	myPlaylist.removeSong("Shape of you")
	myPlaylist.showAllSongs()
}
