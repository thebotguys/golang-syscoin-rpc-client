# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. 

Please note we have a code of conduct, please follow it in all your interactions with the project.

## Implementing a new RPC function

1. All functions must belong to their group as described in the [Syscoin RPC API Readme](https://syscoin.readme.io/v3.2.0/reference).
2. Every function must have his test function(s)

   For Example, if you add the `getaddressbalance` function to the `addressindex.go` file (representing `addressindex` functions group) you have to create the
   following test functions in `addressindex_test.go`:

   ``` go
   func TestGetAddressBalanceInvalid(t *testing.T) {
       // instantiate the client
       cl, err := syscoinrpc.NewClient(InvalidURL, "", "")
       // use the github.com/testify/require library
       require.NoError(t, err, "Must have no error on creation, even with invalid URL")

       // Then you call the function and expect error.
       _, err = cl.AddressIndex.GetAddressBalance(nil, false)
	   require.Error(t, err, "Must error on any method with invalid URL")
   }

   // Remember to align your functions with the other test functions
   func TestGetAddressBalanceOK(t *testing.T) {
       testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}
       testSeparatedOutput := false

       bal, err := cl.AddressIndex.GetAddressBalance(testAddresses, testSeparatedOutput)
       require.NoError(t, err, "Must not error on valid URL, check if the node is running")

       balJSON, _ := json.Marshal(bal)
       t.Log(string(balJSON))
   }
   ```

### Testing Process

You must test your code with a local node before pushing any code.

``` bash
go test -v
```

After that, if tests pass, you can create a branch called `feature/your-function-name`, one per function, if possible. Use vocative names if you implement multiple functions at once.

[Travis CI](https://travis-ci.org/thebotguys/golang-syscoin-rpc-client) will than judge your code before a reviewer will see and evaluate again.