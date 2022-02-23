# klaytn-reverseproxy

클레이튼 엔드포인트 노드를 리버스 프록시 하여 Basic Authentication 을 추가하였습니다.

## Usage

### RPC
```typescript
import { encode } from 'js-base64';

const caver = new Caver(new Caver.providers.HttpProvider('http://localhost:3000/v1/rpc/baobab', {
  headers: [
    { name: 'Authorization', value: `Basic ${encode('root:root')}` }
  ],
}));
```

### Websocket
```typescript
const caver = new Caver(new Caver.providers.WebsocketProvider('ws://root:root@localhost:3000/v1/ws/baobab'));
```

## Environment
```shell
USERNAME=
PASSWORD=
CYPRESS_RPC_URL=
CYPRESS_WS_URL=
BAOBAB_RPC_URL=
BAOBAB_WS_URL=
```
