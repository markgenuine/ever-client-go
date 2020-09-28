# go-ton-sdk

FreeTON SDK Golang language based itself on [TON-SDK](https://github.com/tonlabs/TON-SDK).

## Install
```sh
$ go get -u github.com/move-ton/ton-client-go
```
## Usage
```go
import goton "github.com/move-ton/ton-client-go"
```
## Supported methods
### Based
- [+]	version
- [+]	setup

### Contracts
- [+]	contracts.run.local.msg
- [+]	contracts.run.local
- [+]	contracts.run.fee
- [+]	tvm.get
- [+]	contracts.deploy.address
- [+]	contracts.run.unknown.input
- [+]	contracts.run.unknown.output
- [+]	contracts.deploy.encode_unsigned_message
- [+]	contracts.resolve.error
- [+]	contracts.parse.message
- [+]	contracts.run.fee.msg
- [+]	contracts.deploy.data
- [+]	contracts.deploy.message
- [+]	contracts.encode_message_with_sign
- [+]	contracts.send.message
- [+]	contracts.run.message
- [+]	contracts.function.id
- [+]	contracts.process.message
- [+]	contracts.run
- [+]	contracts.process.transaction
- [+]	contracts.run.output
- [+]	contracts.run.encode_unsigned_message
- [+]	contracts.run.body
- [+]	contracts.address.convert
- [+]	contracts.deploy
- [+]	contracts.wait.transaction
- [+]	contracts.load
- [+]	contracts.find.shard

### Crypto
- [+]	crypto.mnemonic.derive.sign.keys
- [+]	crypto.hdkey.xprv.secret
- [+]	crypto.hdkey.xprv.derive.path
- [+]	crypto.sha256
- [+]	crypto.hdkey.xprv.derive
- [+]	crypto.nacl.sign.open
- [+]	crypto.nacl.box.open
- [+]	crypto.nacl.secret.box
- [+]	crypto.nacl.sign.detached
- [+]	crypto.nacl.sign.keypair.fromSecretKey
- [+]	crypto.ton_public_key_string
- [+]	crypto.nacl.secret.box.open
- [+]	crypto.nacl.sign.keypair
- [+]	crypto.nacl.sign
- [+]	crypto.math.modularPower
- [+]	crypto.ed25519.keypair
- [+]	crypto.random.generateBytes
- [+]	crypto.nacl.box.keypair
- [+]	crypto.sha512
- [+]	crypto.mnemonic.from.random
- [+]	crypto.math.factorize
- [+]	crypto.nacl.box.keypair.fromSecretKey
- [+]	crypto.mnemonic.verify
- [+]	crypto.ton_crc16
- [+]	crypto.mnemonic.from.entropy
- [+]	crypto.hdkey.xprv.from.mnemonic
- [+]	crypto.nacl.box
- [+]	crypto.hdkey.xprv.public
- [+]	crypto.scrypt
- [+]	crypto.mnemonic.words

### Queries
- [+]	queries.unsubscribe
- [+]	queries.subscribe
- [+]	queries.query
- [+]	queries.get.next
- [+]	queries.wait.for
