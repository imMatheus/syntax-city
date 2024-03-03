import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { GameService } from "./gen/game/v1/game_connect";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(GameService, transport);
const data = await client.getGame({ gameId: "abc" });

console.log(data);
function App() {
  return <div className="bg-blue-400">Syntax city</div>;
}

export default App;
