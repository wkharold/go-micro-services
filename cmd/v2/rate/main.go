package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/harlow/go-micro-services/pb/v1/rate"
	"github.com/harlow/go-micro-services/pb/v2/profile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type stay struct {
	HotelID string
	InDate  string
	OutDate string
}

type server struct {
	profileClient profile.ProfileClient
	rateTable     map[stay]*rate.RatePlan
}

// GetRates gets rates for hotels for specific date range.
func (s *server) GetRates(ctx context.Context, req *rate.Request) (*rate.Result, error) {
	log.Printf("fetching rates")
	res := new(rate.Result)
	for _, hotelID := range req.HotelIds {
		stay := stay{
			HotelID: hotelID,
			InDate:  req.InDate,
			OutDate: req.OutDate,
		}
		if s.rateTable[stay] != nil {
			log.Printf("handling hotel %s", hotelID)
			res.RatePlans = append(res.RatePlans, s.rateTable[stay])
			if s.profileClient != nil {
				log.Printf("Rating hotel %s", hotelID)
				s.profileClient.Rate(ctx, &profile.HotelID{Id: hotelID})
			}
		}
	}
	return res, nil
}

// loadRates loads rate codes from JSON file.
func loadRateTable(path string) map[stay]*rate.RatePlan {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot read: %s [%v]", path, err)
		panic(err)
	}

	rates := []*rate.RatePlan{}
	if err := json.Unmarshal(file, &rates); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	rateTable := make(map[stay]*rate.RatePlan)
	for _, ratePlan := range rates {
		stay := stay{
			HotelID: ratePlan.HotelId,
			InDate:  ratePlan.InDate,
			OutDate: ratePlan.OutDate,
		}
		rateTable[stay] = ratePlan
	}

	log.Printf("rate table loaded")
	return rateTable
}

func NewProfileClient(addr string) profile.ProfileClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to dial profile: %v", err)
		return nil
	}

	log.Printf("got a profile client")
	return profile.NewProfileClient(conn)
}

func main() {
	var (
		port        = flag.Int("port", 8080, "The server port")
		profileAddr = flag.String("profileaddr", "profile:8080", "Profile service addr")
	)
	flag.Parse()

	// tcp listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpc server
	srv := grpc.NewServer()
	rate.RegisterRateServer(srv, &server{
		profileClient: NewProfileClient(*profileAddr),
		rateTable:     loadRateTable("/data/inventory.json"),
	})

	log.Printf("starting server")
	srv.Serve(lis)
}
