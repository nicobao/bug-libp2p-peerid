## Bug libp2p

This repo aims to reproduce a bug I noticed between 	"github.com/libp2p/go-libp2p-core/peer" and js-peer-id.

A peerId isn't decoded correctly by go (as a CID), but is parse correclty by the JS lib.

## Usage

In the main directory, run:
- `./run.sh` 

And you will see the difference right away.

## Result

```
=== JS VERSION ===

> bug-libp2p-peerid@1.0.0 start /home/nicolas/nicobao/bug-libp2p-peerid/js
> node with-js-lib.js

Parsing bafykbzacecesdmus2x37xyvtxdwlpx3vu5cjecxed473f66fezh6tdoccqhqq...
peerId is valid! PeerId {
  _id: Uint8Array(36) [
    160, 228,   2,  32, 137,  33, 178, 146,
    213, 247, 251, 226, 179, 184, 236, 183,
    223, 117, 167,  68, 146,  10, 228,  31,
     63, 178, 251, 197,  38,  79, 233, 141,
    194,  20,  15,   8
  ],
  _idB58String: '2DrjgbDVCNmAACH2oLHRPSg874iiE13cTnHK9dvVhXJqN2UQXR',
  _privKey: undefined,
  _pubKey: undefined
}

=== GO VERSION ===
Parsing bafykbzacecesdmus2x37xyvtxdwlpx3vu5cjecxed473f66fezh6tdoccqhqq ...
Error can't convert CID of type "dag-pb" to a peer ID
```
