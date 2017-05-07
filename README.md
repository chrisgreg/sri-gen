# sri-gen

### What?
A multi-threaded Golang SRI hash generator

### What's SRI?
SRI is a new [W3C specification](https://www.w3.org/TR/SRI/) that allows web developers to ensure that resources hosted on third-party servers have not been tampered with. Use of SRI is recommended as a best-practice, whenever libraries are loaded from a third-party source.

[More about SRI on MDN](https://developer.mozilla.org/docs/Web/Security/Subresource_Integrity)

### Which browsers support SRI?
At the time of writing, Chrome & Firefox.

Check out [Can I Use](http://caniuse.com/#feat=subresource-integrity) for more information

## Usage
`$ sri-gen -hash(256|384|512) [file paths]` will output a result.txt file with resultant hashes ready for SRI with their respective file names

## Example

`$ sri-gen -hash256 ./example.js ./example2.js`

will output

```
File	Hash

./example.js	sha256-e9b3814bb4959f2749876624f0a9b04ade95f9faff092d673e00f18cd0c503c2
./example2.js	sha256-a1cc3b3fe09a7aa98c6c8251d2dd7d48a2ddb9f1fec12a3040a5b66c7a6f1171
```
