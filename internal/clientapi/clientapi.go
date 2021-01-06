package clientapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pb "github.com/toncek345/port_manager/internal/portdomainsvc/grpc"
	"google.golang.org/grpc"
)

type clientAPI struct {
	router *http.ServeMux
	port   int

	serviceClient pb.PortDomainClient

	server *http.Server
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(port int, grpcServiceAddr string) (*clientAPI, error) {
	conn, err := grpc.Dial(grpcServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to grpc server: %w", err)
	}

	api := &clientAPI{
		router:        http.NewServeMux(),
		port:          port,
		serviceClient: pb.NewPortDomainClient(conn),
	}

	api.router.HandleFunc("/ports", api.PortsHandler)

	return api, nil
}

// Start runs http server.
func (c *clientAPI) Start() error {
	c.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", c.port),
		Handler: c.router,
	}

	if err := c.server.ListenAndServe(); err != nil {
		return fmt.Errorf("listening and serving: %w", err)
	}

	return nil
}

// Stop stops http server.
func (c *clientAPI) Stop() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	if err := c.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutting down server: %w", err)
	}

	return nil
}

func (c *clientAPI) JSON(w http.ResponseWriter, status int, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("unable to marshall json: %w", err)
	}

	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	w.Write(data)
	return nil
}

func (c *clientAPI) PortsHandler(w http.ResponseWriter, r *http.Request) {
	upsert, err := c.serviceClient.Upsert(r.Context())
	if err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}

	if err := upsert.Send(&pb.Port{}); err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte("heeeeelo"))
	return

	// if r.Method != http.MethodPost {
	// 	log.Println("wrong method")
	// 	a.JSON(w, http.StatusBadRequest, &ErrorResponse{Error: "Wrong request method"})
	// 	return
	// }

	// 	jsonmsg := `
	// {
	//   "AEAJM": {
	//     "name": "Ajman",
	//     "city": "Ajman",
	//     "country": "United Arab Emirates",
	//     "alias": [],
	//     "regions": [],
	//     "coordinates": [
	//       55.5136433,
	//       25.4052165
	//     ],
	//     "province": "Ajman",
	//     "timezone": "Asia/Dubai",
	//     "unlocs": [
	//       "AEAJM"
	//     ],
	//     "code": "52000"
	//   },
	//   "AEAUH": {
	//     "name": "Abu Dhabi",
	//     "coordinates": [
	//       54.37,
	//       24.47
	//     ],
	//     "city": "Abu Dhabi",
	//     "province": "Abu Z¸aby [Abu Dhabi]",
	//     "country": "United Arab Emirates",
	//     "alias": [],
	//     "regions": [],
	//     "timezone": "Asia/Dubai",
	//     "unlocs": [
	//       "AEAUH"
	//     ],
	//     "code": "52001"
	//   },
	//   "AEDXB": {
	//     "name": "Dubai",
	//     "coordinates": [
	//       55.27,
	//       25.25
	//     ],
	//     "city": "Dubai",
	//     "province": "Dubayy [Dubai]",
	//     "country": "United Arab Emirates",
	//     "alias": [],
	//     "regions": [],
	//     "timezone": "Asia/Dubai",
	//     "unlocs": [
	//       "AEDXB"
	//     ],
	//     "code": "52005"
	//   }}`

	// dec := json.NewDecoder(strings.NewReader(jsonmsg))

	// // read open bracket
	// t, err := dec.Token()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%T: %v\n", t, t)

	// type Message struct {
	// 	A struct {
	// 		Name string `json:"name"`
	// 	} `json:"AEAJM"`
	// 	Name string `json:"name"`
	// }

	// while the array contains values
	// for dec.More() {
	// 	fmt.Println("REAAAAAAAAAAAAAAAD")
	// 	data := make([]byte, 0, 20)
	// 	s, err := dec.Buffered().Read(p []byte)

	// 	// var m Message
	// 	m := make(map[string]interface{})
	// 	// decode an array value (Message)
	// 	err := dec.Decode(&m)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%#v\n\n", m)
	// }
}
