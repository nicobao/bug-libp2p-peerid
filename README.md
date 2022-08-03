## Bug libp2p

This repo aims to reproduce a bug I noticed between 	"github.com/libp2p/go-libp2p-core/peer" and js-peer-id.

A peerId isn't decoded correctly by go (as a CID), but is parse correclty by the JS lib.

## Usage

In the main directory, run:
- `./run.sh` 

And you will see the difference right away.
