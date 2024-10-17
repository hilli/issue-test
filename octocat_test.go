package issues_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/octocat"
)

// ExampleApiClient_Octocat constructs an unauthenticated API client
// and makes a simple API request.
func Test_ExampleApiClient_Octocat(t *testing.T) {
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("octokit/go-sdk.example-functions"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
	)

	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	s := "Salutations"

	// create headers that accept json back; GitHub's OpenAPI definition says
	// octet-stream but that's not actually what the API returns in this case
	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	octocatRequestConfig := &abs.RequestConfiguration[octocat.OctocatRequestBuilderGetQueryParameters]{
		QueryParameters: &octocat.OctocatRequestBuilderGetQueryParameters{
			S: &s,
		},
		Headers: headers,
	}

	cat, err := client.Octocat().Get(context.Background(), octocatRequestConfig)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))
	// Output:
	// MMM.           .MMM
	//                MMMMMMMMMMMMMMMMMMM
	//                MMMMMMMMMMMMMMMMMMM      _____________
	//               MMMMMMMMMMMMMMMMMMMMM    |             |
	//              MMMMMMMMMMMMMMMMMMMMMMM   | Salutations |
	//             MMMMMMMMMMMMMMMMMMMMMMMM   |_   _________|
	//             MMMM::- -:::::::- -::MMMM    |/
	//              MM~:~ 00~:::::~ 00~:~MM
	//         .. MMMMM::.00:::+:::.00::MMMMM ..
	//               .MM::::: ._. :::::MM.
	//                  MMMM;:::::;MMMM
	//           -MM        MMMMMMM
	//           ^  M+     MMMMMMMMM
	//               MMMMMMM MM MM MM
	//                    MM MM MM MM
	//                    MM MM MM MM
	//                 .~~MM~MM~MM~MM~~.
	//              ~~~~MM:~MM~~~MM~:MM~~~~
	//             ~~~~~~==~==~~~==~==~~~~~~
	//              ~~~~~~==~==~==~==~~~~~~
	//                  :~==~==~==~==~~
}
