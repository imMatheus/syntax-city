import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { GameService } from "./gen/game/v1/game_connect";

export const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
export const grpcClient = createPromiseClient(GameService, transport);
// const data = await client.getGame({ gameId: "abc" });
