package usecase

import "cat-bookmarks/domain/model"

type BookmarkCase struct {
}

func NewBookmark() *BookmarkCase {
	return &BookmarkCase{}
}

func (b *BookmarkCase) AllBookmarks() ([]model.Bookmark, error) {
	return nil, nil
}
