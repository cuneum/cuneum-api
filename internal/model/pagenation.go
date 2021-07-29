package model

import "github.com/cuneum/cuneum-api/pkg/pb"

type Pagenation struct {
	Current int32
	Pages   int32
	Limit   int32
	Total   int32
}

func (p *Pagenation) Into() *pb.PagenationResponse {
	return &pb.PagenationResponse{
		Current: p.Current,
		Pages:   p.Pages,
		Limit:   p.Limit,
		Total:   p.Total,
	}
}
