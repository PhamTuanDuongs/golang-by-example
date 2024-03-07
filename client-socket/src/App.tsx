import React from "react";
import "./App.css";
import { Chart } from "./chart/chart";
import { SocketClient } from "./socket/clientSocket";

function App() {
  return (
    <div className="App">
      <Chart message="Xin chao" tag={1} />
      <SocketClient />
    </div>
  );
}

export default App;
