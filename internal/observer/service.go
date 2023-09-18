package observer

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/nakiner/eth-parser/internal/models"
)

const (
	maxWorkersPerAddr = 2
	maxBlockCount     = 100
)

type Observer struct {
	ethService          ethService
	ethProvider         ethProvider
	subscriberJobs      chan jobDescription
	wg                  sync.WaitGroup
	jobsCountByAddress  map[string]int // how many jobs actually working for particular address
	mu                  sync.Mutex
	expectedBlockPerJob map[string]int64
	currEthBlockNumber  int64
}

type jobDescription struct {
	address   string
	blockFrom int64
	blockTo   int64
}

type ethService interface {
	GetSubscribers() []string
	SetBlockNumber(block int64)
	GetBlockNumber() int64
	AddTransactions(address string, transactions []models.Transaction)
}

type ethProvider interface {
	GetBlockNumber(ctx context.Context) (int64, error)
	GetLogs(ctx context.Context, fromBlock int64, toBlock int64, address string) ([]models.Transaction, error)
}

func NewObserver(ethService ethService, ethProvider ethProvider) *Observer {
	return &Observer{
		ethService:          ethService,
		ethProvider:         ethProvider,
		subscriberJobs:      make(chan jobDescription),
		wg:                  sync.WaitGroup{},
		jobsCountByAddress:  make(map[string]int),
		mu:                  sync.Mutex{},
		expectedBlockPerJob: make(map[string]int64),
	}
}

func (o *Observer) Serve(ctx context.Context, count int) (func() error, func(error)) {
	observerCtx, observerCancel := context.WithCancel(ctx)

	o.setCurrentBlockByAddress(strings.ToLower("0x0B50800f891C1966C4263e90E8081ee0e6cE55cF"), 18156440)

	execute := func() error {
		o.startWorkers(observerCtx, count)
		ethBlockNumUpdateTick := time.NewTicker(time.Second * 30)
		addJobTicker := time.NewTicker(time.Second * 15)
		defer ethBlockNumUpdateTick.Stop()
		defer addJobTicker.Stop()
		logger.InfoKV(ctx, "started main worker")
		for {
			select {
			case <-observerCtx.Done():
				logger.Info(ctx, "stopped main worker")
				return nil
			case <-ethBlockNumUpdateTick.C:
				currentEthBlock, err := o.ethProvider.GetBlockNumber(ctx)
				if err != nil {
					logger.ErrorKV(ctx, "could not fetch ethProvider.GetBlockNumber", "err", err)
					continue
				}
				o.currEthBlockNumber = currentEthBlock
			case <-addJobTicker.C:
				if o.currEthBlockNumber < 1 {
					continue
				}
				subs := o.ethService.GetSubscribers()
				parsedAddresses := make([]string, 0)
				subsToParse := make([]string, 0)

				for _, sub := range subs {
					currBlock := o.getCurrentBlockByAddress(sub)
					if currBlock == o.currEthBlockNumber {
						parsedAddresses = append(parsedAddresses, sub)
						continue
					}
					subsToParse = append(subsToParse, sub)
				}

				if len(parsedAddresses) == len(subs) { // all addresses have current block, we can refresh it now
					o.ethService.SetBlockNumber(o.currEthBlockNumber)
					continue
				}

				for _, sub := range subsToParse {
					blockFrom, blockTo := o.calculateBlockSizeForStart(o.currEthBlockNumber, sub)
					j := jobDescription{
						address:   sub,
						blockFrom: blockFrom,
						blockTo:   blockTo,
					}
					if o.currEthBlockNumber != blockTo {
						blockTo++
					}
					o.setCurrentBlockByAddress(sub, blockTo)
					fmt.Println(j)
					o.subscriberJobs <- j
				}
			}
		}
	}

	interrupt := func(err error) {
		observerCancel()
		logger.Warn(ctx, "observer is stopped")
	}

	return execute, interrupt
}

func (o *Observer) startWorkers(ctx context.Context, count int) {
	for i := 0; i < count; i++ {
		go func(ctx context.Context) {
			logger.Info(ctx, "started worker")
			for {
				select {
				case <-ctx.Done():
					logger.Info(ctx, "worker stopped")
					return
				case job := <-o.subscriberJobs:
					logger.WarnKV(ctx, "got new job to parse", "addr", job.address)
					if o.getRunningJobCountByAddress(job.address) >= maxWorkersPerAddr {
						logger.Info(ctx, "job rejected due to max capacity per address")
						o.subscriberJobs <- job
						continue
					}
					o.runningJobCountByAddressInc(job.address)
					transactions, err := o.ethProvider.GetLogs(ctx, job.blockFrom, job.blockTo, job.address)
					if err != nil {
						logger.ErrorKV(ctx, "could not fetch ethProvider.GetLogs", "err", err)
						continue
					}
					o.ethService.AddTransactions(job.address, transactions)
					o.runningJobCountByAddressDec(job.address)
					logger.Info(ctx, "complete")
				}
			}
		}(ctx)
	}
}

func (o *Observer) getCurrentBlockByAddress(address string) int64 {
	num, ok := o.expectedBlockPerJob[address]
	if !ok {
		return 0
	}

	return num
}

func (o *Observer) setCurrentBlockByAddress(address string, block int64) {
	o.expectedBlockPerJob[address] = block
}

func (o *Observer) getRunningJobCountByAddress(address string) int {
	o.mu.Lock()
	defer o.mu.Unlock()
	num, ok := o.jobsCountByAddress[address]
	if !ok {
		return 0
	}

	return num
}

func (o *Observer) runningJobCountByAddressInc(address string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.jobsCountByAddress[address]++
}

func (o *Observer) runningJobCountByAddressDec(address string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.jobsCountByAddress[address]--
}

func (o *Observer) calculateBlockSizeForStart(currentBlock int64, address string) (int64, int64) {
	oldBlock := o.getCurrentBlockByAddress(address)
	newBlockTo := oldBlock
	if currentBlock-oldBlock >= maxBlockCount {
		newBlockTo += maxBlockCount
	} else {
		newBlockTo = currentBlock
	}

	return oldBlock, newBlockTo
}
