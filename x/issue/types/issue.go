package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
	"math/big"
	"strings"
)

var (
	CoinMaxTotalSupply, _        = sdk.NewIntFromString("1000000000000000000000000000000000000")
	CoinIssueMaxId        uint64 = 999999999999
	CoinIssueMinId        uint64 = 1
)

const (
	IDPreStr = "coin"
	Custom   = "custom"
)

type IIssue interface {
	GetIssueId() string
	SetIssueId(string)

	GetIssuer() sdk.AccAddress
	SetIssuer(sdk.AccAddress)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetSymbol() string
	SetSymbol(string)

	GetDecimals() uint
	SetDecimals(uint)

	ToCoin() sdk.Coin
}

type CoinIssues []CoinIssue

//nolint
func (coinIssues CoinIssues) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-10s|%-6s|%-18s|%-8s|%s\n",
		"IssueID", "Owner", "Name", "Symbol", "TotalSupply", "Decimals", "IssueTime")
	for _, issue := range coinIssues {
		out += fmt.Sprintf("%-17s|%-44s|%-10s|%-6s|%-18s|%-8d|%d\n",
			issue.IssueId, issue.GetOwner().String(), issue.Name, issue.Symbol, issue.TotalSupply.String(), issue.Decimals, issue.IssueTime)
	}
	return strings.TrimSpace(out)
}

type CoinIssue struct {
	IssueId            string         `json:"issue_id"`
	Issuer             sdk.AccAddress `json:"issuer"`
	Owner              sdk.AccAddress `json:"owner"`
	Name               string         `json:"name"`
	Symbol             string         `json:"symbol"`
	TotalSupply        sdk.Int        `json:"total_supply"`
	Decimals           uint           `json:"decimals"`
	Description        string         `json:"description"`
	IssueTime          int64          `json:"issue_time"`
	BurnOwnerDisabled  bool           `json:"burn_owner_disabled"`
	BurnHolderDisabled bool           `json:"burn_holder_disabled"`
	BurnFromDisabled   bool           `json:"burn_from_disabled"`
	FreezeDisabled     bool           `json:"freeze_disabled"`
	MintingFinished    bool           `json:"minting_finished"`
}

func NewCoinIssue(owner, issuer sdk.AccAddress, params *IssueParams) *CoinIssue {
	var ci CoinIssue
	_ = mapstructure.Decode(params, &ci)
	ci.Owner = owner
	ci.Issuer = issuer
	ci.Symbol = strings.ToUpper(ci.Symbol)
	ci.TotalSupply = params.TotalSupply

	return &ci
}

func (ci *CoinIssue) GetIssueId() string {
	return ci.IssueId
}

func (ci *CoinIssue) SetIssueId(issueId string) {
	ci.IssueId = issueId
}

func (ci *CoinIssue) GetIssuer() sdk.AccAddress {
	return ci.Issuer
}

func (ci *CoinIssue) SetIssuer(issuer sdk.AccAddress) {
	ci.Issuer = issuer
}

func (ci *CoinIssue) GetOwner() sdk.AccAddress {
	return ci.Owner
}

func (ci *CoinIssue) SetOwner(owner sdk.AccAddress) {
	ci.Owner = owner
}

func (ci *CoinIssue) GetSymbol() string {
	return ci.Symbol
}

func (ci *CoinIssue) SetSymbol(symbol string) {
	ci.Symbol = symbol
}

func (ci *CoinIssue) GetDecimals() uint {
	return ci.Decimals
}
func (ci *CoinIssue) SetDecimals(decimals uint) {
	ci.Decimals = decimals
}

func (ci *CoinIssue) ToCoin() sdk.Coin {
	return sdk.NewCoin(ci.IssueId, ci.TotalSupply)
}

func (ci *CoinIssue) ToCoins() sdk.Coins {
	return sdk.NewCoins(ci.ToCoin())
}

func getDecimalsInt(decimals uint) sdk.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return sdk.NewIntFromBigInt(exp)
}

func (ci *CoinIssue) QuoDecimals(amount sdk.Int) sdk.Int {
	return amount.Quo(getDecimalsInt(ci.GetDecimals()))
}

func IsIssueId(issueID string) bool {
	return strings.HasPrefix(issueID, IDPreStr)
}

func CheckIssueId(issueID string) sdk.Error {
	if !IsIssueId(issueID) {
		return ErrIssueID(issueID)
	}
	return nil
}