package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/commons"
	"github.com/brunopadz/amictl/providers"
	"github.com/spf13/cobra"
)

// listAmiCmd represents the listAmi command
var listUnused = &cobra.Command{
	Use:   "list-unused",
	Short: "List unused AMIs",
	Long:  `List not used AMIs for a given region and account.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("You need to specify 2 arguments: [ACCOUNT_ID] [REGION]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Creates a input filter to get AMIs
		f := &ec2.DescribeImagesInput{
			Owners: []*string{
				aws.String(args[0]),
			},
		}

		// Establishes new authenticated session to AWS
		s := providers.AwsSession(args[1])

		// Filter AMIs based on input filter
		a, err := s.DescribeImages(f)
		if err != nil {
			fmt.Println(err)
		}

		// Compare AMI list
		l, u := providers.AwsListNotUsed(a, s)

		n := commons.Compare(l, u)
		r := strings.Join(n, "\n")

		fmt.Println(r)
		fmt.Println("Total of", len(n), "not used AMIs")

	},
}

func init() {
	awsCmd.AddCommand(listUnused)
}
