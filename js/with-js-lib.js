import PeerId from "peer-id";

try {
  const stringToParse =
    "bafykbzacecesdmus2x37xyvtxdwlpx3vu5cjecxed473f66fezh6tdoccqhqq";
  console.log(`Parsing ${stringToParse}...`);
  const peerIdObj = PeerId.parse(stringToParse);
  console.log("peerId is valid!", peerIdObj);
} catch (e) {
  console.error("Error while parsing peerId", e);
}
