package domain

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	posts = []Post{
		{1, "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"},
		{2, "qui est esse", "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"},
		{3, "ea molestias quasi exercitationem repellat qui ipsa sit aut", "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut"},
	}
	lastPostID = len(posts)
)

func GetAllPosts() []Post {
	return posts
}

func FindPostByID(id int) *Post {
	var res *Post

	for _, p := range posts {
		if p.Id == id {
			res = &p
			break
		}
	}

	return res
}

func AddPost(post *Post) {
	lastPostID++

	post.Id = lastPostID

	posts = append(posts, *post)
}

func UpdatePost(post *Post) {
	postIndex := getPostIndex(*post)

	posts[postIndex] = *post
}

func RemovePost(post *Post) {
	i := getPostIndex(*post)

	posts = append(posts[:i], posts[i+1:]...)
}

func getPostIndex(post Post) int {
	for i, p := range posts {
		if p.Id == post.Id {
			return i
		}
	}

	return -1
}
