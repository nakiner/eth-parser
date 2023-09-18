package etherium

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/pkg/errors"
)

type Params struct {
	FromBlock string        `json:"fromBlock"`
	ToBlock   string        `json:"toBlock"`
	Topics    []interface{} `json:"topics"`
}

type ResultRow struct {
	Address          string   `json:"address,omitempty"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

type GetLogsRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []Params `json:"params"`
	ID      int      `json:"id"`
}

type GetLogsResponse struct {
	ID      int         `json:"id,omitempty"`
	JsonRpc string      `json:"jsonRpc,omitempty"`
	Result  []ResultRow `json:"result,omitempty"`
}

func (c *Client) GetLogs(ctx context.Context, fromBlock int64, toBlock int64, address string) (GetLogsResponse, error) {
	req := GetLogsRequest{
		JsonRpc: "2.0",
		Method:  "eth_getLogs",
		Params: []Params{{
			FromBlock: toHex(fromBlock),
			ToBlock:   toHex(toBlock),
			Topics: []interface{}{
				"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
				[]any{
					nil,
					fmt.Sprintf("0x000000000000000000000000%s", address[2:]),
				},
				[]any{
					fmt.Sprintf("0x000000000000000000000000%s", address[2:]),
					nil,
				},
			},
		}},
		ID: 0,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return GetLogsResponse{}, errors.Wrap(err, "GetLogs could not marshal json")
	}

	fmt.Println(string(data))

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://cloudflare-eth.com", bytes.NewReader(data))
	if err != nil {
		return GetLogsResponse{}, errors.Wrap(err, "GetLogs could not make http request")
	}

	resp, err := c.cl.Do(httpReq)
	if err != nil {
		return GetLogsResponse{}, errors.Wrap(err, "GetLogs could not Do http req")
	}
	defer func() {
		errD := resp.Body.Close()
		if errD != nil {
			logger.WarnKV(ctx, "GetLogs could not close http body", "err", errD)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetLogsResponse{}, errors.Wrap(err, "GetLogs could not read response body")
	}

	if resp.StatusCode != http.StatusOK {
		errResp := errors.New("GetLogs status not OK")
		logger.ErrorKV(ctx, "GetLogs status not OK", "resp", string(body))
		return GetLogsResponse{}, errors.Wrap(errResp, string(body))
	}

	var respData GetLogsResponse

	if err = json.Unmarshal(body, &respData); err != nil {
		return GetLogsResponse{}, errors.Wrap(err, "GetLogs could not unmarshal json response")
	}

	return respData, nil
}
