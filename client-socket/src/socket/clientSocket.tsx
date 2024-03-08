import { useState } from "react";
import useWebSocket from "react-use-websocket";
import { TypeMessage } from "../chart/types";
import { Chart } from "../chart/chart";

const socketUrl = process.env.REACT_APP_WEBSOCKET_URL as string;
export const SocketClient = () => {
  const [message, setMessage] = useState<TypeMessage[]>([]);
  const { lastMessage, readyState } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("Websocket connection established.");
    },
    onMessage: (message) => {
      const parsedTypeMessage: TypeMessage = JSON.parse(
        message.data
      ) as TypeMessage;
      setMessage((prev) => prev.concat(parsedTypeMessage));
    },

    shouldReconnect: (closeEvent) => false,
  });

  return (
    <div>
      <div>State {readyState}</div>
      <Chart message={message} />
      <ol>
        {message.map((value: TypeMessage, index) => (
          <li key={index}>
            {value.id.toString()}--
            {value.time.toString()}--
            {value.data}
          </li>
        ))}
      </ol>
    </div>
  );
};
