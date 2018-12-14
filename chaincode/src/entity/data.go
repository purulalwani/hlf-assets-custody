package entity

import (
    "time"
)

// DATA STRUCTURE TO CAPTURE INVESTOR DETAILS BY CUSTODIAN
// Key consists of custodianPrefix + userName
type Investor struct {
    UserName     string  `json:"user_name"`
    UserFName    string  `json:"user_fname"`
    UserLName    string  `json:"user_lname"`
    UserIdentity string  `json:"user_identity"`
    KycStatus    string  `json:"kyc_status"`
    DepositoryAC string  `json:"depository_ac"`
    BankAC       string  `json:"bank_ac"`
}

// Key consists of custodianPrefix + userName
type InvestorPortfolio struct {
    StockTicker string  `json:"stock_ticker"`
    StockQty    int32   `json:"stock_qty"`
    StockPrice  float32 `json:"stock_price"`
    StockValue  float32 `json:"stock_value"`
}

// Key consists of custodianPrefix + userName + tradeUUID
type InvestorTrade struct {
    TradeUUID   string    `json:"trade_uuid"`
    TradeDate   time.Time `json:"trade_date"`
    TradeType   string    `json:"trade_type"`
    StockTicker string    `json:"stock_ticker"`
    StockQty    int32     `json:"stock_qty"`
    StockPrice  float32   `json:"stock_price"`
    StockValue  float32   `json:"stock_value"`
}

// DATA STRUCTURE TO CAPTURE TRANSACTION DETAILS BY BANK
// Key consists of exchangePrefix only
type BankMaster struct {
    UserName    string    `json:"user_name"`
    BankAC      string    `json:"bank_ac"`
    Balance     float32   `json:"balance"`
}

// Key consists of bankPrefix + userName + transUUID
type BankTransaction struct {
    TransUUID   string    `json:"trans_uuid"`
    UserName    string    `json:"user_name"`
    BankAC      string    `json:"bank_ac"`
    TransDate   time.Time `json:"trans_date"`
    TransAmount float32   `json:"trans_amount"`
    Balance     float32   `json:"balance"`
}

// DATA STRUCTURE TO CAPTURE TRADING DETAILS BY EXCHANGE
// Key consists of exchangePrefix only
type ExchangeMaster struct {
    StockTicker string    `json:"stock_ticker"`
    StockQty    int32     `json:"stock_qty"`
    StockPrice  float32   `json:"stock_price"`
}
// Key consists of exchangePrefix + tradeUUID
type ExchangeTrades struct {
    TradeUUID   string    `json:"trade_uuid"`
    TradeDate   time.Time `json:"trade_date"`
    StockTicker string    `json:"stock_ticker"`
    StockQty    int32     `json:"stock_qty"`
    StockPrice  float32   `json:"stock_price"`
    StockValue  float32   `json:"stock_value"`
}

// DATA STRUCTURE TO CAPTURE TRANSACTION DETAILS BY DEPOSITORY
// Key consists of depositoryPrefix + userName + transUUID
type DepositoryTransaction struct {
    TransUUID   string    `json:"trans_uuid"`
    UserName    string    `json:"user_name"`
    DepositoryAC string   `json:"depository_ac"`
    TradeDate   time.Time `json:"trade_date"`
    StockTicker string    `json:"stock_ticker"`
    StockQty    int32     `json:"stock_qty"`
    StockPrice  float32   `json:"stock_price"`
    StockValue  float32   `json:"stock_value"`
}

