package grpc

type Config struct {
	Host        string `yaml:"host"`
	GrpcPort    int    `yaml:"grpc_port"`
	HttpPort    int    `yaml:"http_port"`
	GraphqlPort int    `yaml:"graphql_port"`
}
