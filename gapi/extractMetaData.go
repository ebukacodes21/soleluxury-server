package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	grpcGatewayUserAgentHeader = "grpc-gateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type MetaData struct {
	UserAgent string
	ClientIp  string
}

func (s *Server) extractMetaData(ctx context.Context) *MetaData {
	metaData := &MetaData{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			metaData.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			metaData.UserAgent = userAgents[0]
		}

		if clientIp := md.Get(xForwardedForHeader); len(clientIp) > 0 {
			metaData.ClientIp = clientIp[0]
		}
	}

	if md, ok := peer.FromContext(ctx); ok {
		metaData.ClientIp = md.Addr.String()
	}

	return metaData
}
