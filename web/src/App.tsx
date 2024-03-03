import { TransportProvider } from "@connectrpc/connect-query";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Board } from "./views/Board";
import { transport } from "./gprc";

const queryClient = new QueryClient();

function App() {
  return (
    <TransportProvider transport={transport}>
      <QueryClientProvider client={queryClient}>
        <Board />
      </QueryClientProvider>
    </TransportProvider>
  );
}

export default App;
