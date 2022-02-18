package types

import (
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/okxtuta/go-anatha/types"
)

var (
	coinPos  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000)
	coinZero = sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)
)

// test ValidateBasic for MsgCreateValidator
func TestMsgCreateValidator(t *testing.T) {
	commission1 := NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	commission2 := NewCommissionRates(sdk.NewDec(5), sdk.NewDec(5), sdk.NewDec(5))

	tests := []struct {
		name, moniker, identity, website, securityContact, details string
		CommissionRates                                            CommissionRates
		minSelfDelegation                                          sdk.Int
		validatorAddr                                              sdk.ValAddress
		pubkey                                                     crypto.PubKey
		bond                                                       sdk.Coin
		expectPass                                                 bool
	}{
		{"basic good", "a", "b", "c", "d", "e", commission1, sdk.OneInt(), valAddr1, pk1, coinPos, true},
		{"partial description", "", "", "c", "", "", commission1, sdk.OneInt(), valAddr1, pk1, coinPos, true},
		{"empty description", "", "", "", "", "", commission2, sdk.OneInt(), valAddr1, pk1, coinPos, false},
		{"empty address", "a", "b", "c", "d", "e", commission2, sdk.OneInt(), emptyAddr, pk1, coinPos, false},
		{"empty pubkey", "a", "b", "c", "d", "e", commission1, sdk.OneInt(), valAddr1, emptyPubkey, coinPos, true},
		{"empty bond", "a", "b", "c", "d", "e", commission2, sdk.OneInt(), valAddr1, pk1, coinZero, false},
		{"zero min self delegation", "a", "b", "c", "d", "e", commission1, sdk.ZeroInt(), valAddr1, pk1, coinPos, false},
		{"negative min self delegation", "a", "b", "c", "d", "e", commission1, sdk.NewInt(-1), valAddr1, pk1, coinPos, false},
		{"delegation less than min self delegation", "a", "b", "c", "d", "e", commission1, coinPos.Amount.Add(sdk.OneInt()), valAddr1, pk1, coinPos, false},
	}

	for _, tc := range tests {
		description := NewDescription(tc.moniker, tc.identity, tc.website, tc.securityContact, tc.details)
		msg := NewMsgCreateValidator(tc.validatorAddr, tc.pubkey, tc.bond, description, tc.CommissionRates, tc.minSelfDelegation)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

// test ValidateBasic for MsgEditValidator
func TestMsgEditValidator(t *testing.T) {
	tests := []struct {
		name, moniker, identity, website, securityContact, details string
		validatorAddr                                              sdk.ValAddress
		expectPass                                                 bool
	}{
		{"basic good", "a", "b", "c", "d", "e", valAddr1, true},
		{"partial description", "", "", "c", "", "", valAddr1, true},
		{"empty description", "", "", "", "", "", valAddr1, false},
		{"empty address", "a", "b", "c", "d", "e", emptyAddr, false},
	}

	for _, tc := range tests {
		description := NewDescription(tc.moniker, tc.identity, tc.website, tc.securityContact, tc.details)
		newRate := sdk.ZeroDec()
		newMinSelfDelegation := sdk.OneInt()

		msg := NewMsgEditValidator(tc.validatorAddr, description, &newRate, &newMinSelfDelegation)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

// test ValidateBasic for MsgDelegate
func TestMsgDelegate(t *testing.T) {
	tests := []struct {
		name          string
		delegatorAddr sdk.AccAddress
		validatorAddr sdk.ValAddress
		bond          sdk.Coin
		expectPass    bool
	}{
		{"basic good", sdk.AccAddress(valAddr1), valAddr2, coinPos, true},
		{"self bond", sdk.AccAddress(valAddr1), valAddr1, coinPos, true},
		{"empty delegator", sdk.AccAddress(emptyAddr), valAddr1, coinPos, false},
		{"empty validator", sdk.AccAddress(valAddr1), emptyAddr, coinPos, false},
		{"empty bond", sdk.AccAddress(valAddr1), valAddr2, coinZero, false},
	}

	for _, tc := range tests {
		msg := NewMsgDelegate(tc.delegatorAddr, tc.validatorAddr, tc.bond)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

// test ValidateBasic for MsgUnbond
func TestMsgBeginRedelegate(t *testing.T) {
	tests := []struct {
		name             string
		delegatorAddr    sdk.AccAddress
		validatorSrcAddr sdk.ValAddress
		validatorDstAddr sdk.ValAddress
		amount           sdk.Coin
		expectPass       bool
	}{
		{"regular", sdk.AccAddress(valAddr1), valAddr2, valAddr3, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), true},
		{"zero amount", sdk.AccAddress(valAddr1), valAddr2, valAddr3, sdk.NewInt64Coin(sdk.DefaultBondDenom, 0), false},
		{"empty delegator", sdk.AccAddress(emptyAddr), valAddr1, valAddr3, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), false},
		{"empty source validator", sdk.AccAddress(valAddr1), emptyAddr, valAddr3, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), false},
		{"empty destination validator", sdk.AccAddress(valAddr1), valAddr2, emptyAddr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), false},
	}

	for _, tc := range tests {
		msg := NewMsgBeginRedelegate(tc.delegatorAddr, tc.validatorSrcAddr, tc.validatorDstAddr, tc.amount)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

// test ValidateBasic for MsgUnbond
func TestMsgUndelegate(t *testing.T) {
	tests := []struct {
		name          string
		delegatorAddr sdk.AccAddress
		validatorAddr sdk.ValAddress
		amount        sdk.Coin
		expectPass    bool
	}{
		{"regular", sdk.AccAddress(valAddr1), valAddr2, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), true},
		{"zero amount", sdk.AccAddress(valAddr1), valAddr2, sdk.NewInt64Coin(sdk.DefaultBondDenom, 0), false},
		{"empty delegator", sdk.AccAddress(emptyAddr), valAddr1, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), false},
		{"empty validator", sdk.AccAddress(valAddr1), emptyAddr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1), false},
	}

	for _, tc := range tests {
		msg := NewMsgUndelegate(tc.delegatorAddr, tc.validatorAddr, tc.amount)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

//test to validate if NewMsgCreateValidator implements yaml marshaller
func TestMsgMarshalYAML(t *testing.T) {
	commission1 := NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	tc := struct {
		name, moniker, identity, website, securityContact, details string
		CommissionRates                                            CommissionRates
		minSelfDelegation                                          sdk.Int
		validatorAddr                                              sdk.ValAddress
		pubkey                                                     crypto.PubKey
		bond                                                       sdk.Coin
		expectPass                                                 bool
	}{"basic good", "a", "b", "c", "d", "e", commission1, sdk.OneInt(), valAddr1, pk1, coinPos, true}

	description := NewDescription(tc.moniker, tc.identity, tc.website, tc.securityContact, tc.details)
	msg := NewMsgCreateValidator(tc.validatorAddr, tc.pubkey, tc.bond, description, tc.CommissionRates, tc.minSelfDelegation)
	bs, err := yaml.Marshal(msg)
	require.NoError(t, err)
	bechifiedPub, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, msg.PubKey)
	require.NoError(t, err)

	want := fmt.Sprintf(`|
  description:
    moniker: a
    identity: b
    website: c
    security_contact: d
    details: e
  commission:
    rate: "0.000000000000000000"
    max_rate: "0.000000000000000000"
    max_change_rate: "0.000000000000000000"
  minselfdelegation: "1"
  delegatoraddress: %s
  validatoraddress: %s
  pubkey: %s
  value:
    denom: stake
    amount: "1000"
`, msg.DelegatorAddress, msg.ValidatorAddress, bechifiedPub)

	require.Equal(t, want, string(bs))

}
