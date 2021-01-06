package clientapi

func (c *clientAPI) PortsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("wrong method")
		c.JSON(w, http.StatusBadRequest, &ErrorResponse{Error: "Wrong request method"})
		return
	}

	dec := json.NewDecoder(r.Body)

	_, err := dec.Token()
	if err != nil {
		c.JSON(w, http.StatusBadRequest, ErrorResponse{Error: "Bad json"})
		return
	}

	upsert, err := c.serviceClient.Upsert(r.Context())
	if err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}
	defer upsert.CloseAndRecv()

	for dec.More() {
		portID, err := dec.Token()
		if err != nil {
			panic(err)
		}

		// TODO: this is not DDD but time is running out.
		port := &pb.Port{}
		if err = dec.Decode(port); err != nil {
			c.JSON(w, http.StatusInternalServerError, err)
			return
		}

		portIDStr, ok := portID.(string)
		if !ok {
			c.JSON(w, http.StatusBadRequest, ErrorResponse{Error: "Bad json"})
			return
		}

		port.IdStr = portIDStr

		if err := upsert.Send(port); err != nil {
			c.JSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
