// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// unnamed API: CLI Commands
//
// Command:
// $ goagen
// --design=ComputingProvider/design
// --out=$(GOPATH)\src\ComputingProvider
// --version=v1.3.1

package cli

import (
	"ComputingProvider/client"
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// AddComputingProviderCommand is the command line data structure for the add action of ComputingProvider
	AddComputingProviderCommand struct {
		// ETH private key for transaction
		ETHKey string
		// data IPFS address
		Hash        string
		PrettyPrint bool
	}

	// AgreeComputingProviderCommand is the command line data structure for the agree action of ComputingProvider
	AgreeComputingProviderCommand struct {
		// ETH private key for transaction
		ETHKey string
		// computing resourse hash
		ComputingHash string
		// smart contract hash
		ContractHash string
		// ETH public key(Wallet address)
		PublicKey   string
		PrettyPrint bool
	}

	// DelComputingProviderCommand is the command line data structure for the del action of ComputingProvider
	DelComputingProviderCommand struct {
		// ETH private key for transaction
		ETHKey string
		// data IPFS address
		Hash        string
		PrettyPrint bool
	}

	// UploadResComputingProviderCommand is the command line data structure for the uploadRes action of ComputingProvider
	UploadResComputingProviderCommand struct {
		// ETH private key for transaction
		ETHKey string
		// encrypted aes key hash
		AesHash string
		// [request_id]
		RequestID int
		// encrypted result hash
		ResHash     string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add",
		Short: `add computing resource`,
	}
	tmp1 := new(AddComputingProviderCommand)
	sub = &cobra.Command{
		Use:   `computing-provider ["/computing/add/HASH/ETH_KEY"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "agree",
		Short: `agree computing request for request[ID]`,
	}
	tmp2 := new(AgreeComputingProviderCommand)
	sub = &cobra.Command{
		Use:   `computing-provider ["/computing/agree/ETH_KEY/COMPUTING_HASH/CONTRACT_HASH/PUBLIC_KEY"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "del",
		Short: `delete computing resource`,
	}
	tmp3 := new(DelComputingProviderCommand)
	sub = &cobra.Command{
		Use:   `computing-provider ["/computing/del/HASH/ETH_KEY"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "upload-res",
		Short: `upload result hash for [request_id]`,
	}
	tmp4 := new(UploadResComputingProviderCommand)
	sub = &cobra.Command{
		Use:   `computing-provider ["/computing/upload/RES_HASH/AES_HASH/ETH_KEY/REQUEST_ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swagger-ui-dist/") {
		fnd = c.DownloadSwaggerUIDist
		rpath = rpath[17:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the AddComputingProviderCommand command.
func (cmd *AddComputingProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/computing/add/%v/%v", url.QueryEscape(cmd.Hash), url.QueryEscape(cmd.ETHKey))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddComputingProvider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddComputingProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eTHKey string
	cc.Flags().StringVar(&cmd.ETHKey, "ETH_key", eTHKey, `ETH private key for transaction`)
	var hash string
	cc.Flags().StringVar(&cmd.Hash, "hash", hash, `data IPFS address`)
}

// Run makes the HTTP request corresponding to the AgreeComputingProviderCommand command.
func (cmd *AgreeComputingProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/computing/agree/%v/%v/%v/%v", url.QueryEscape(cmd.ETHKey), url.QueryEscape(cmd.ComputingHash), url.QueryEscape(cmd.ContractHash), url.QueryEscape(cmd.PublicKey))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AgreeComputingProvider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AgreeComputingProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eTHKey string
	cc.Flags().StringVar(&cmd.ETHKey, "ETH_key", eTHKey, `ETH private key for transaction`)
	var computingHash string
	cc.Flags().StringVar(&cmd.ComputingHash, "computing_hash", computingHash, `computing resourse hash`)
	var contractHash string
	cc.Flags().StringVar(&cmd.ContractHash, "contract_hash", contractHash, `smart contract hash`)
	var publicKey string
	cc.Flags().StringVar(&cmd.PublicKey, "public_key", publicKey, `ETH public key(Wallet address)`)
}

// Run makes the HTTP request corresponding to the DelComputingProviderCommand command.
func (cmd *DelComputingProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/computing/del/%v/%v", url.QueryEscape(cmd.Hash), url.QueryEscape(cmd.ETHKey))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DelComputingProvider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DelComputingProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eTHKey string
	cc.Flags().StringVar(&cmd.ETHKey, "ETH_key", eTHKey, `ETH private key for transaction`)
	var hash string
	cc.Flags().StringVar(&cmd.Hash, "hash", hash, `data IPFS address`)
}

// Run makes the HTTP request corresponding to the UploadResComputingProviderCommand command.
func (cmd *UploadResComputingProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/computing/upload/%v/%v/%v/%v", url.QueryEscape(cmd.ResHash), url.QueryEscape(cmd.AesHash), url.QueryEscape(cmd.ETHKey), cmd.RequestID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UploadResComputingProvider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UploadResComputingProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eTHKey string
	cc.Flags().StringVar(&cmd.ETHKey, "ETH_key", eTHKey, `ETH private key for transaction`)
	var aesHash string
	cc.Flags().StringVar(&cmd.AesHash, "aes_hash", aesHash, `encrypted aes key hash`)
	var requestID int
	cc.Flags().IntVar(&cmd.RequestID, "request_id", requestID, ` [request_id]`)
	var resHash string
	cc.Flags().StringVar(&cmd.ResHash, "res_hash", resHash, `encrypted result hash`)
}
