package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (reply *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{}}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{}}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{}}, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (reply *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TaglistReply, err error) {
	return &v1.TaglistReply{}, nil
}
