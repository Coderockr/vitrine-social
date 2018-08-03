package graphql

import (
	"html/template"
	"net/http"
	"strings"
)

type graphiqlHandler struct {
	next http.Handler
}

func (h *graphiqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	acceptHeader := r.Header.Get("Accept")
	_, raw := r.URL.Query()["raw"]
	if !raw && !strings.Contains(acceptHeader, "application/json") && strings.Contains(acceptHeader, "text/html") {
		h.renderGraphiQL(w)
		return
	}

	h.next.ServeHTTP(w, r)
}

var graphiqlTemplate *template.Template

func init() {
	graphiqlTemplate = template.New("GraphiQL")
	graphiqlTemplate.Parse(graphiqlTemplateStr)
}

func (h *graphiqlHandler) renderGraphiQL(w http.ResponseWriter) {
	err := graphiqlTemplate.ExecuteTemplate(w, "index", map[string]interface{}{
		"GraphiqlVersion": graphiqlVersion,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// graphiqlVersion is the current version of GraphiQL
const graphiqlVersion = "0.11.3"

// tmpl is the page template to render GraphiQL
const graphiqlTemplateStr = `
{{ define "index" -}}
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <title>GraphiQL</title>
  <meta name="robots" content="noindex" />
  <style>
    html, body {
      height: 100%;
      margin: 0;
      overflow: hidden;
      width: 100%;
    }

    body {
      display: flex;
      flex-direction: column;
    }

    #root {
        flex: 1;
    }

    #tokenBox {
        background-color: #f7f7f7;
        border-bottom: 1px solid #d0d0d0;
        padding: 7px 14px 6px;

        font-family: system, -apple-system, 'San Francisco', '.SFNSDisplay-Regular', 'Segoe UI', Segoe, 'Segoe WP', 'Helvetica Neue', helvetica, 'Lucida Grande', arial, sans-serif;
        font-size: 14px;
        display: flex;
        align-items: baseline;
    }

    #tokenBox * {
        color: #555;
    }

    #tokenBox label {
        padding-right: 1rem;
    }

    #tokenBox input {
        flex: 1;
        border: 1px solid #d0d0d0;
    }
  </style>
  <link href="//cdn.jsdelivr.net/npm/graphiql@{{ .GraphiqlVersion }}/graphiql.css" rel="stylesheet" />
  <script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
  <script src="//cdn.jsdelivr.net/react/15.4.2/react.min.js"></script>
  <script src="//cdn.jsdelivr.net/react/15.4.2/react-dom.min.js"></script>
  <script src="//cdn.jsdelivr.net/npm/graphiql@{{ .GraphiqlVersion }}/graphiql.min.js"></script>
</head>
<body>
  <div id="tokenBox">
    <label for="token">Token:</label>
    <input type="text" name="token" id="token" />
  </div>
  <div id="root"></div>
  <script>
    const tokenInput = document.getElementById("token");

    tokenInput.value = localStorage.getItem("graphiqlToken") || ""

    // Defines a GraphQL fetcher using the fetch API.
    function graphQLFetcher(graphQLParams) {
      let headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      };

      tokenInput.value = tokenInput.value.trim();
      localStorage.setItem("graphiqlToken", tokenInput.value)
      if (tokenInput.value.length > 0) {
        headers.Authorization = tokenInput.value;
      }

      return fetch("?", {
        method: 'post',
        headers: headers,
        body: JSON.stringify(graphQLParams),
        credentials: 'include',
      }).then(function (response) {
        return response.text();
      }).then(function (responseBody) {
        try {
          return JSON.parse(responseBody);
        } catch (error) {
          return responseBody;
        }
      });
    }

    // Render <GraphiQL /> into the body.
    ReactDOM.render(
      React.createElement(GraphiQL, {
        fetcher: graphQLFetcher,
      }),
      document.getElementById("root")
    );
  </script>
</body>
</html>
{{- end }}
`
