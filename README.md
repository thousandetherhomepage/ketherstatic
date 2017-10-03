# ketherstatic
Generate a static version of ketherhomepage.com


## Usage

```bash
GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json ketherstatic -jsonPath="rinkeby.json" -pngPath="rinkeby.png" -rpc="https://rinkeby.infura.io/9qx9rjhXpHyOJHGWlx8B" -bucket="storage.thousandetherhomepage.com"

GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json ./ketherstatic -jsonPath="mainnet.json" -pngPath="mainnet.png" -rpc="https://mainnet.infura.io/9qx9rjhXpHyOJHGWlx8B" --address="0xb5fe93ccfec708145d6278b0c71ce60aa75ef925" -bucket="storage.thousandetherhomepage.com"
```
