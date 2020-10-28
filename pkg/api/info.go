package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"quote/pkg/constants"
	"quote/pkg/model"
)

func (s *Server) info(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	infoList, err := s.infoService.GetInfoByTitleOrInfo(searchText)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"searchText": searchText,
			"error":      err,
		}).Errorf("error searching events")
		w.WriteHeader(http.StatusInternalServerError)
	}
	displayInfo(infoList, w)
	w.WriteHeader(http.StatusOK)
}

func displayInfo(filteredInfo []model.Info, w http.ResponseWriter) {
	fmt.Fprintf(w, "<h1>Info:</h1>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>SN</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>ID</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info Added</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, info := range filteredInfo {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d</td>", info.ID))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info.Title))

		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		// Display URL in different table under td
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for i, url := range info.Links {
			var youtubeLink string
			if strings.Contains(strings.ToLower(url), "youtube") {
				youtubeLink = "click me to watch on youtube"
			}
			if !strings.Contains(strings.ToLower(url), "no-link") {
				fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Links%d. %s </a></td></tr>", url, i+1, youtubeLink))
			}
		}
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		fmt.Fprintf(w, fmt.Sprintf("<pre>%s</pre>", info.Info))
		fmt.Fprintf(w, fmt.Sprintf("</td>"))

		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", info.CreationDate.Format(constants.DATE_FORMAT_INFO)))
		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
}

// swagger:operation GET /api/quote/v1/info/{id} INFO info
// ---
// description: get INFO by id
// consumes:
// - "application/json"
// parameters:
// - name: id
//   description: id to get info
//   in: path
//   required: true
//   default: 100
//   type: string
// Responses:
//   '200':
//     description: Ok
//   '400':
//     description: Bad request
//   '404':
//     description: Not found
//   '500':
//     description: Internal server error
func (s *Server) getInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Errorf("error converting string to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	info, err := s.infoService.GetInfoByID(int64(idInt))
	if err != nil {
		logrus.WithError(err).Errorf("error getting info for id=%d", idInt)
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResult, err := info.ToJson()
	if err != nil {
		logrus.WithError(err).Errorf("error converting info=%v to json", info)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jsonResult))
	w.WriteHeader(http.StatusOK)
}

// swagger:operation PUT /api/quote/v1/info/{id} INFO info
// ---
// description: Put INFO by id
// consumes:
// - "application/json"
// parameters:
// - name: id
//   description: id to put info
//   in: path
//   required: true
//   default: 100
//   type: string
// - name: infoRequest
//   in: body
//   required: true
//   schema:
//     '$ref': '#/definitions/infoRequest'
// Responses:
//  200: noContent
func (s *Server) putInfo(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Error("error in info id conversion")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var infoReq InfoRequest
	err = json.NewDecoder(r.Body).Decode(&infoReq)

	if err != nil {
		logrus.WithError(err).Error("error decoding info body request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			logrus.WithError(err).Error("cookie-token not found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		logrus.WithError(err).Error("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	var jwtKey = []byte("my_secret_key")
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			logrus.WithError(err).Error("invalid signature")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		logrus.WithError(err).Error("bad request..")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		logrus.WithError(err).Error("invalid token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	info := model.Info{
		ID:    int64(id),
		Title: infoReq.Title,
		Info:  infoReq.Info,
		Links: infoReq.Links,
	}

	err = s.infoService.UpdateInfoByID(info)
	if err != nil {
		logrus.WithError(err).Error("error updating info")
		if strings.Contains(err.Error(), "does not exist in database") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
