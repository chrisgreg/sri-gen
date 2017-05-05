# sri-gen

### What?
A Golang SRI hash generator

### What's SRI?
SRI is a new [W3C specification](https://www.w3.org/TR/SRI/) that allows web developers to ensure that resources hosted on third-party servers have not been tampered with. Use of SRI is recommended as a best-practice, whenever libraries are loaded from a third-party source.

[More about SRI on MDN](https://developer.mozilla.org/docs/Web/Security/Subresource_Integrity)

### Which browsers support SRI?
At the time of writing, Chrome & Firefox.

Check out [Can I Use](http://caniuse.com/#feat=subresource-integrity) for more information

## Usage
`sri-gen -hash(256|384|512) [file path array...]`
