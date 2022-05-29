package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "test", "test client")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExperimentServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.CreateExperiment(ctx, &pb.CreateExperimentReq{
	// 	Title: "试验标题",
	// 	ResearcherId: "12121212121",
	// })
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetExperimentId())
	// //查询
	// rr, err := c.QueryExperiment(ctx, &pb.QueryExperimentReq{
	// 	ExperimentId: "6a56dc56-7d3d-4ee9-b2b7-9072ffcc509b",
	// })
	//  查询试验列表
	// rrr, err := c.QueryExperimentList(ctx, &pb.QueryExperimentListReq{
	// 	ResearcherId: "12121212121",
	// })
	// if err != nil {
	// 	log.Fatalf("could not query experimentlist: %v", err)
	// }
	// log.Printf("%v", rrr)
	// 更新试验
	// r, err := c.UpdateExperiment(ctx, &pb.UpdateExperimentReq{
	// 	Title: "更新标题",
	// 	ExperimentId: "49dcf8b7-b0c9-445e-adf1-6739b58ece37",
	// 	ResearcherId: "12121212121",
	// })
	// log.Printf("resp :%v", r)
	//  创建一个新的被试记录
	// r, err := c.CreateSubjectRecord(ctx, &pb.CreateSubjectRecordReq{
	// 	ExperimentId: "49dcf8b7-b0c9-445e-adf1-6739b58ece37",
	// 	ParticipantId: "9dcf8b7-b0c9-445e-adf1-6739b58ece37",
	// })
	// r, err := c.UpdateSubjectRecord(ctx, &pb.UpdateSubjectRecordReq{
	// 	SubjectRecordId: "eed88eb3-e7f8-40a5-85ff-9c16e6ce378c",
	// 	State: 1,
	// })
	// r, err := c.QuerySubjectRecord(ctx, &pb.QuerySubjectRecordReq{
	// 	SubjectRecordId: "367702ea-e164-4a4e-a85f-a273a341a96c", 
	// })
	r, err := c.QuerySubjectRecordList(ctx, &pb.QuerySubjectRecordListReq{
		ExperimentId: "49dcf8b7-b0c9-445e-adf1-6739b58ece37",
	})
	log.Printf("resp :%v", r)
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", rr.GetExpInfo().GetTitle())
}
