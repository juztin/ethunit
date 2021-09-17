## Ethereum Units

Simple library to convert between Ethereum units.


_example_

```golang
import (
  ...
  "github.com/juztin/ethunit"
  ...
)

func printTransaction(w io.Writer, tx *types.Transaction) error {
  // Specify the units
	value := ethunit.ConvertInt(tx.Value(), Wei, Ether)
  // Use available helpers
	cost := ethunit.WeiToEther(tx.Cost())
	price := ethunit.WeiToGwei(tx.GasPrice())
	_, err := fmt.Fprintf(w, `  Transaction: %s
         Data: %x
        Value: %s Ether
      Tx Cost: %s Ether
    Gas Price: %s Gwei
        Nonce: %d
`,
		tx.Hash().String(),
		tx.Data(),
		value.Text('f', 9),
		cost.Text('f', 9),
		price.Text('f', 4),
		tx.Nonce())

	return err
}

```
