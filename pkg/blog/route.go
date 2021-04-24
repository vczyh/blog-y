package blog

import (
	"blog-y/pkg/common/route"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	r := router.Group("/blog/v1")

	// posts
	posts := r.Group("/posts")
	posts.GET("", route.Handle(PostListAPI))                // 获取文章
	posts.POST("", route.Handle(CreatePostAPI))             // 添加文章 auth
	posts.GET("/:postId", route.Handle(GetPostByPostIdAPI)) // 根据ID获取文章
	posts.PUT("/:postId", route.Handle(UpdatePostAPI))      // 更新文章 auth
	posts.DELETE("/:postId", route.Handle(DeletePostAPI))
	posts.PUT("/:postId/tags", route.Handle(AddPostTagAPI))       // 添加标签 auth
	posts.DELETE("/:postId/tags", route.Handle(DeletePostTagAPI)) // 删除标签 auth

	// tags
	tags := r.Group("/tags")
	tags.GET("", route.Handle(TagListAPI))   // 获取所有tag
	tags.PUT("", route.Handle(CreateTagAPI)) // 添加tag

	// admin
	admin := r.Group("/admin")
	admin.PUT("/login", route.Handle(LoginAPI))
	admin.PUT("/auth", route.Handle(AuthAPI))
}
