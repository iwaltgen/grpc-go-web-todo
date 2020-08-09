import { TodoClient } from '../../grpc';

const grpcHostname = 'https://localhost:8443';

export const todoClient = new TodoClient(grpcHostname);

// eslint-disable-next-line @typescript-eslint/no-empty-function
const enableDevTools = globalThis.__GRPCWEB_DEVTOOLS__ || (() => {});
enableDevTools([todoClient.client]);
