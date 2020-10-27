package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"

	"github.com/gorilla/mux"
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
