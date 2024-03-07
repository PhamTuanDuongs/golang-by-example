import { useEffect, useState } from "react";
import useWebSocket from "react-use-websocket";
import { TypeMessage } from "../chart/types";
import { Chart } from "../chart/chart";

const socketUrl = "ws://localhost:8000/sendData";
export const SocketClient = () => {
  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);
  const [messageHistory, setMessageHistory] = useState([]);
  const [message, setMessage] = useState<TypeMessage[]>([]);
  useEffect(() => {
    if (lastMessage != null) {
      try {
        const receivedData = lastMessage.data;
        const parsedTypeMessage: TypeMessage = JSON.parse(
          receivedData
        ) as TypeMessage;
        setMessage((prev) => prev.concat(parsedTypeMessage));
        console.log(parsedTypeMessage);
      } catch (error) {
        console.log(error);
      }
    }
  }, [lastMessage]);

  return (
    <div className="flex justify-center">
      <Chart message={message} />
      <div>
        {message.map((value: TypeMessage, index) => (
          <div key={index}>
            <span>ID: {value.id.toString()}</span>
            <span>Time: {value.time.toString()}</span>
            <span>Name: {value.data}</span>
          </div>
        ))}
      </div>
    </div>
  );
  // 1. listent for a data event and update the state
  // 2. Render the line chart using state
};
