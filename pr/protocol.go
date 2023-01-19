package pr

import (
	"golang.org/x/net/context"
	"strangerDB/dbIndexing"
	"strangerDB/dbService"
)

type Server struct{}

func (s Server) HandleRequest(ctx context.Context, request *Request) (*Response, error) {
	//TODO implement me
	command := request.Command

	switch command.(type) {
	case *Request_Read:
		{
			p := dbIndexing.Index[request.GetRead().Key]
			data := dbService.ReadRecord("chunks/data", p.Offset, p.Size)
			return &Response{Message: string(data)}, nil
		}
	case *Request_Write:
		{
			data := dbService.KeyValue{Key: request.GetWrite().Key, Value: request.GetWrite().Value}
			dbIndexing.AddIndex("chunks/data", data)
			dbService.WriteRecord("chunks/data", data)
			return &Response{Message: "ЗАПИСАЛ"}, nil
		}
	case *Request_Delete:
		{
		}
	}

	return nil, nil
}

func (s Server) mustEmbedUnimplementedDBServer() {
	//TODO implement me
	panic("implement me")
}
