package handler

// if req.Method == "POST" {
// 	reqJson := req.Body

// 	var q queryStruct

// 	json.NewDecoder(reqJson).Decode(&q)

// 	conf, err := h.Compound.GetConnection()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
// 		os.Exit(1)
// 	}

// 	if conf.Slave1.Count > conf.Slave2.Count {
// 		h.wg.Add(1)

// 		var answer [][]string

// 		go func() {
// 			row, err := conf.Slave1.Slave1Compound.Query(q.SqlQuery, q.Args)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
// 				os.Exit(1)
// 			}

// 			for row.Next() {

// 				var pars []string

// 				err = row.Scan(&pars)
// 				if err != nil {
// 					panic(err)
// 				}
// 				answer = append(answer, pars)
// 			}

// 			atomic.AddUint64(&conf.Slave1.Count, 1)

// 			h.wg.Done()
// 		}()

// 		h.wg.Wait()

// 		js, err := json.Marshal(answer)
// 		if err != nil {
// 			panic(err)
// 		}

// 		resp.Write(js)
// 	} else {
// 		h.wg.Add(1)

// 		var answer [][]string

// 		go func() {
// 			row, err := conf.Slave2.Slave2Compound.Query(q.SqlQuery, q.Args)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
// 				os.Exit(1)
// 			}

// 			for row.Next() {

// 				var pars []string

// 				err = row.Scan(&pars)
// 				if err != nil {
// 					panic(err)
// 				}
// 				answer = append(answer, pars)
// 			}

// 			atomic.AddUint64(&conf.Slave2.Count, 1)

// 			h.wg.Done()
// 		}()

// 		h.wg.Wait()

// 		js, err := json.Marshal(answer)
// 		if err != nil {
// 			panic(err)
// 		}

// 		resp.Write(js)
// 	}

// } else {
// 	resp.Write([]byte("This url only handles POST requests"))
// }

// if req.Method == "POST" {
// 	reqJson := req.Body

// 	var q queryStruct

// 	json.NewDecoder(reqJson).Decode(&q)

// 	conf, err := h.Compound.GetConnection()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
// 		os.Exit(1)
// 	}

// 	_, err = conf.Main.MainCompound.Exec(q.SqlQuery, q.Args)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
// 		os.Exit(1)
// 	}

// 	err = migration(reqJson, q, conf)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fale migration on slaves: %v\n", err)
// 		os.Exit(1)
// 	}

// } else {
// 	resp.Write([]byte("This url only handles POST requests"))
// }
// }

// func migration(body io.ReadCloser, q queryStruct, conf *mysqlconn.Compound) error {
// var wg sync.WaitGroup

// wg.Add(2)

// go func() {
// 	_, err := conf.Slave1.Slave1Compound.Exec(q.SqlQuery, q.Args)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
// 		os.Exit(1)
// 	}

// 	wg.Done()
// }()

// go func() {
// 	_, err := conf.Slave2.Slave2Compound.Exec(q.SqlQuery, q.Args)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
// 		os.Exit(1)
// 	}

// 	wg.Done()
// }()

// wg.Wait()

// return nil
