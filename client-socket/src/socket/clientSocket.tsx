import { useEffect } from "react";
import { io } from "socket.io-client";

var socket = io("http://localhost:8000", {
  transports: ["websocket", "polling"],
});
export const SocketClient = () => {
  useEffect(() => {
    socket.on("bye", (msg) => {
      console.log(msg);
    });
  });

  return <h1>Client Socket</h1>;
  // 1. listent for a data event and update the state
  // 2. Render the line chart using state
};
