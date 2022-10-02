// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RealWorldClient is the client API for RealWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealWorldClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UnFavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TaglistReply, error)
}

type realWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealWorldClient(cc grpc.ClientConnInterface) RealWorldClient {
	return &realWorldClient{cc}
}

func (c *realWorldClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/UnfollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/FeedArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/FavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) UnFavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/UnFavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realWorldClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TaglistReply, error) {
	out := new(TaglistReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealWorldServer is the server API for RealWorld service.
// All implementations must embed UnimplementedRealWorldServer
// for forward compatibility
type RealWorldServer interface {
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error)
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*MultipleCommentsReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error)
	UnFavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*SingleArticleReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TaglistReply, error)
	mustEmbedUnimplementedRealWorldServer()
}

// UnimplementedRealWorldServer must be embedded to have forward compatible implementations.
type UnimplementedRealWorldServer struct {
}

func (UnimplementedRealWorldServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRealWorldServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRealWorldServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedRealWorldServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRealWorldServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedRealWorldServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedRealWorldServer) UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedRealWorldServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedRealWorldServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedRealWorldServer) GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedRealWorldServer) CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedRealWorldServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedRealWorldServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedRealWorldServer) AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedRealWorldServer) GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedRealWorldServer) DeleteComment(context.Context, *DeleteCommentRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedRealWorldServer) FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (UnimplementedRealWorldServer) UnFavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
func (UnimplementedRealWorldServer) GetTags(context.Context, *GetTagsRequest) (*TaglistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedRealWorldServer) mustEmbedUnimplementedRealWorldServer() {}

// UnsafeRealWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealWorldServer will
// result in compilation errors.
type UnsafeRealWorldServer interface {
	mustEmbedUnimplementedRealWorldServer()
}

func RegisterRealWorldServer(s grpc.ServiceRegistrar, srv RealWorldServer) {
	s.RegisterService(&RealWorld_ServiceDesc, srv)
}

func _RealWorld_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/UnfollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnfollowUser(ctx, req.(*UnfollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/FeedArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_FavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).FavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/FavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_UnFavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).UnFavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/UnFavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).UnFavoriteArticle(ctx, req.(*UnfavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealWorld_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealWorld_ServiceDesc is the grpc.ServiceDesc for RealWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.RealWorld",
	HandlerType: (*RealWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RealWorld_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _RealWorld_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _RealWorld_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _RealWorld_UpdateUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _RealWorld_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _RealWorld_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _RealWorld_UnfollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _RealWorld_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _RealWorld_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _RealWorld_GetArticle_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _RealWorld_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _RealWorld_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _RealWorld_DeleteArticle_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _RealWorld_AddComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _RealWorld_GetComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _RealWorld_DeleteComment_Handler,
		},
		{
			MethodName: "FavoriteArticle",
			Handler:    _RealWorld_FavoriteArticle_Handler,
		},
		{
			MethodName: "UnFavoriteArticle",
			Handler:    _RealWorld_UnFavoriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _RealWorld_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realworld/v1/realworld.proto",
}
