package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/providers"
	"github.com/spf13/cobra"
)

func init() {
	awsCmd.AddCommand(listAll)
}

// listAmiCmd represents the listAmi command
var listAll = &cobra.Command{
	Use:     "list-all",
	Short:   "List all AMIs",
	Long:    `List all AMIs for a given region and account.`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1`,
	RunE:    runAll,
}

func runAll(cmd *cobra.Command, args []string) error {

	// Creates a input filter to get AMIs
	f := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}

	// Establishes new authenticated session to AWS
	s := providers.AwsSession(region)

	// Filter AMIs based on input filter
	a, err := s.DescribeImages(f)
	if err != nil {
		fmt.Println(err)
	}

	l := providers.AwsListAll(a, s)

	r := strings.Join(l, "\n")
	fmt.Println(r)
	fmt.Println("Total of", len(l), "AMIs")

	return nil
}
