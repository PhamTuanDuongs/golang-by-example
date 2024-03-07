import { maxHeaderSize } from "http";
import { AppProps, TypeMessage } from "./types";
import {
  BarChart,
  Bar,
  Line,
  LineChart,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  Area,
  AreaChart,
  ResponsiveContainer,
  Legend,
} from "recharts";
const data = [
  { name: "Page A", uv: 400, pv: 2400, amt: 3400 },
  { name: "Page B", uv: 300, pv: 1400, amt: 5400 },
  { name: "Page C", uv: 200, pv: 400, amt: 6400 },
];
export const Chart = ({ message }: { message: TypeMessage[] }) => (
  <div>
    <div>
      <LineChart
        width={1500}
        height={500}
        data={message}
        margin={{ top: 5, right: 20, bottom: 5, left: 0 }}
      >
        <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
        <XAxis dataKey="id" padding={{ left: 30, right: 30 }} />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line
          type="monotone"
          dataKey="time"
          stroke="#8884d8"
          activeDot={{ r: 8 }}
        />
        <Line type="monotone" />
      </LineChart>
    </div>
  </div>
);
