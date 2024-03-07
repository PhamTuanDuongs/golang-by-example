import { AppProps } from "./types";
import {
  BarChart,
  Bar,
  Line,
  LineChart,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
} from "recharts";
const data = [
  { name: "Page A", uv: 400, pv: 2400, amt: 3400 },
  { name: "Page B", uv: 300, pv: 1400, amt: 5400 },
  { name: "Page C", uv: 200, pv: 400, amt: 6400 },
];
export const Chart = ({ message, tag }: AppProps) => (
  <div>
    {message} {tag}
    <LineChart
      width={600}
      height={300}
      data={data}
      margin={{ top: 5, right: 20, bottom: 5, left: 0 }}
    >
      <Line type="monotone" dataKey="uv" stroke="#8884d8" />
      <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
      <XAxis dataKey="name" />
      <YAxis />
      <Tooltip />
    </LineChart>
  </div>
);
