package ethunit

import "math/big"

func WeiToGwei(i *big.Int) *big.Float { return ConvertInt(i, Wei, Gwei) }

func WeiToEther(i *big.Int) *big.Float { return ConvertInt(i, Wei, Ether) }

func GweiToWei(i *big.Int) *big.Float { return ConvertInt(i, Gwei, Wei) }

func GweiToEth(i *big.Int) *big.Float { return ConvertInt(i, Gwei, Ether) }

func EtherToWei(i *big.Int) *big.Float { return ConvertInt(i, Ether, Wei) }

func EtherToGwei(i *big.Int) *big.Float { return ConvertInt(i, Ether, Gwei) }
