package persist

type UserRepository interface {
	Save(user *User) *User
	Update(user *User) *User
	Find(id uint) *User
	FindByUsername(username string) *User
}

type PostRepository interface {
	Save(post *Post) *Post
	Update(post *Post) *Post
	Find(id uint) *Post
	FindAllByOwnerUsername(ownerUsername string) []*Post
	Delete(id uint)
}

type CommentRepository interface {
	Save(comment *Comment) *Comment
	Update(comment *Comment) *Comment
	Find(id uint) *Comment
	FindAllByOwnerUsername(ownerUsername string) []*Comment
	Delete(id uint)
}
