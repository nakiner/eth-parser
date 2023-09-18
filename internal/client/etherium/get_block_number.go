package etherium

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/pkg/errors"
)

type GetBlockNumberRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}

type GetBlockNumberResponse struct {
	ID      int
	JsonRpc string
	Result  string
}

func (c *Client) GetBlockNumber(ctx context.Context) (int64, error) {
	req := GetBlockNumberRequest{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
		ID:      83,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not marshal json")
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://cloudflare-eth.com", bytes.NewReader(data))
	if err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not make http request")
	}

	resp, err := c.cl.Do(httpReq)
	defer func() {
		errD := resp.Body.Close()
		if errD != nil {
			logger.WarnKV(ctx, "GetBlockNumber could not close http body", "err", errD)
		}
	}()
	if err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not Do http req")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not read response body")
	}

	if resp.StatusCode != http.StatusOK {
		errResp := errors.New("GetBlockNumber status not OK")
		logger.ErrorKV(ctx, "GetBlockNumber status not OK", "resp", string(body))
		return 0, errors.Wrap(errResp, string(body))
	}

	var respData GetBlockNumberResponse

	if err = json.Unmarshal(body, &respData); err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not unmarshal json response")
	}

	hexVal, err := fromHex(respData.Result)
	if err != nil {
		return 0, errors.Wrap(err, "GetBlockNumber could not decode hex")
	}

	return hexVal, nil
}
