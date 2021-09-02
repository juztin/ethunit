## Ethereum Units

Simple library to convert between Ethereum units.


_example_

```golang
func printTransaction(w io.Writer, tx *types.Transaction) error {
	value := ConvertInt(tx.Value(), Wei, Ether)
	cost := ConvertInt(tx.Cost(), Wei, Ether)
	price := ConvertInt(tx.GasPrice(), Wei, Gwei)
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
