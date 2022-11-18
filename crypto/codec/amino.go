package codec

import (
	"github.com/tendermint/tendermint/crypto/sr25519"

	"github.com/tendermint/go-amino"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// RegisterCrypto registers all crypto dependency types with the provided Amino
// codec.
func RegisterCrypto(cdc *codec.LegacyAmino) {
	pubOptions := amino.InterfaceOptions{
		Priority: []string{
			"tendermint/PubKeySecp256k1",
			"injective/PubKeyEthSecp256k1",
		},
	}
	privOptions := amino.InterfaceOptions{
		Priority: []string{
			"tendermint/PrivKeySecp256k1",
			"injective/PrivKeyEthSecp256k1",
		},
	}
	cdc.RegisterInterface((*cryptotypes.PubKey)(nil), &pubOptions)
	cdc.RegisterConcrete(sr25519.PubKey{},
		sr25519.PubKeyName, nil)
	cdc.RegisterConcrete(&ed25519.PubKey{},
		ed25519.PubKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PubKey{},
		secp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&kmultisig.LegacyAminoPubKey{},
		kmultisig.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*cryptotypes.PrivKey)(nil), &privOptions)
	cdc.RegisterConcrete(sr25519.PrivKey{},
		sr25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&ed25519.PrivKey{}, //nolint:staticcheck
		ed25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PrivKey{},
		secp256k1.PrivKeyName, nil)
}
