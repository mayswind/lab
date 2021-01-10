package exchangerates

import (
	"github.com/mayswind/lab/pkg/errs"
	"github.com/mayswind/lab/pkg/settings"
)

// ExchangeRatesDataSourceContainer contains the current exchange rates data source
type ExchangeRatesDataSourceContainer struct {
	Current ExchangeRatesDataSource
}

// Initialize a exchange rates data source container singleton instance
var (
	Container = &ExchangeRatesDataSourceContainer{}
)

// InitializeExchangeRatesDataSource initializes the current exchange rates data source according to the config
func InitializeExchangeRatesDataSource(config *settings.Config) error {
	if config.ExchangeRatesDataSource == settings.EuroCentralBankDataSource {
		Container.Current = &EuroCentralBankDataSource{}
		return nil
	}

	return errs.ErrInvalidExchangeRatesDataSource
}