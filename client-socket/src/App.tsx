import React from "react";
import "./App.css";
import { Chart } from "./chart/chart";
import { SocketClient } from "./socket/clientSocket";

function App() {
  return (
    <div className="App">
      <h1 className="font-bold text-5xl">Client Socket</h1>
      <SocketClient />
    </div>
  );
}

export default App;
