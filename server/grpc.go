package server

import (
	"context"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/common"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type grpcserver struct {
	api.UnimplementedKarlchencloudServer
	room cloud.Room
	auth Auth
}

func (s *grpcserver) Register(_ context.Context, req *api.RegisterRequest) (*api.RegisterReply, error) {
	id := cloud.RandomLetters(8)
	secret := cloud.RandomSecret()
	s.room.Users.Add(cloud.UserId(id), req.GetName(), secret)
	log.Printf("Registered user %v with id %v", req.GetName(), id)
	return &api.RegisterReply{Id: id, Secret: secret}, nil
}

func (s *grpcserver) CheckLogin(ctx context.Context, _ *api.Empty) (*api.OkOrNot, error) {
	user, ok := GetAuthenticatedUser(ctx)
	if !ok {
		log.Print("check login failed")
		return &api.OkOrNot{Value: false}, nil
	}
	log.Printf("user %v ok", user)
	return &api.OkOrNot{Value: true}, nil
}

func (s *grpcserver) CreateTable(ctx context.Context, _ *api.Empty) (*api.TableData, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table := s.room.CreateTable(user)
	log.Printf("user %v created new table %v", s.room.Users.GetName(user), table)
	return toTableData(table, user), nil
}

func toTableData(table *cloud.Table, user cloud.UserId) *api.TableData {
	exposedInviteCode := ""
	if table.Owner() == user {
		exposedInviteCode = table.InviteCode
	}
	return &api.TableData{TableId: table.Id, Owner: string(table.Owner()), InviteCode: exposedInviteCode}
}

func (s *grpcserver) ListTables(ctx context.Context, _ *api.Empty) (*api.TableList, error) {
	user, _ := GetAuthenticatedUser(ctx)
	tables := s.room.Tables.List()
	result := make([]*api.TableData, len(tables))
	for i, table := range tables {
		result[i] = toTableData(table, user)
	}
	return &api.TableList{Tables: result}, nil
}

func (s *grpcserver) StartTable(ctx context.Context, id *api.TableId) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := s.tryGetTable(id.Value)
	if err != nil {
		return nil, err
	}
	if table.Owner() != user {
		return nil, status.Error(codes.PermissionDenied, "you are not owner of the table")
	}
	err = table.Start()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.Empty{}, nil

}

func (s *grpcserver) tryGetTable(id string) (*cloud.Table, error) {
	table, ok := s.room.Tables.ById[id]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "table does not exist")
	}
	return table, nil
}

func (s *grpcserver) JoinTable(ctx context.Context, req *api.JoinTableRequest) (*api.Empty, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := s.tryGetTable(req.TableId)
	if err != nil {
		return nil, err
	}
	if table.InviteCode != req.InviteCode {
		return nil, status.Error(codes.PermissionDenied, "invalid invite code")
	}
	err = table.Join(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.Empty{}, nil
}

func StartServer(users cloud.Users, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	room := cloud.NewRoom(users)
	auth := NewAuth(users)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(auth.Authenticate)),
		grpc.StreamInterceptor(grpcauth.StreamServerInterceptor(auth.Authenticate)))

	serv := grpcserver{room: room, auth: auth}
	api.RegisterKarlchencloudServer(grpcServer, &serv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *grpcserver) GetMatchState(ctx context.Context, tableId *api.TableId) (*api.MyMatchState, error) {
	user, _ := GetAuthenticatedUser(ctx)
	table, err := s.tryGetTable(tableId.Value)
	if err != nil {
		return nil, err
	}
	m := table.CurrentMatch
	if m == nil {
		return nil, status.Error(codes.Internal, "no active match at table")
	}
	matchState := common.ToMatchState(m)
	for p, playerUser := range m.Players {
		if user == playerUser {
			private := &api.PlayerPrivateState{HandCards: GetHandCards(m.Match, game.Player(p))}
			return &api.MyMatchState{
				MatchState: matchState,
				Role:       &api.MyMatchState_PlayerState{PlayerState: private}}, nil
		}
	}
	return &api.MyMatchState{MatchState: matchState, Role: &api.MyMatchState_Spectator{Spectator: &api.Empty{}}}, nil
}

func GetHandCards(m *match.Match, p game.Player) []*api.Card {
	if m.Phase() != match.InGame {
		return nil
	}
	cards := m.Game.HandCards[p]
	ans := make([]*api.Card, len(cards))
	for i, card := range cards {
		ans[i] = common.ToApiCard(card)
	}
	return ans
}
