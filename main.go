package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()





			for {

				kvs := map[string]string{"a": "CHEGOU", "b": "OK"}
				myMsg :=" "
				for k, v := range kvs {
					myMsg = fmt.Sprintf("%s -> %s\n", k, v)
					err = wsutil.WriteServerMessage(conn, ws.OpText,[]byte(myMsg))
					if err != nil {
						fmt.Println(err)
					}
				}

				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println(err)
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}))


}