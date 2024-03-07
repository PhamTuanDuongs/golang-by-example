import { AppProps } from "./types";
import {
  BarChart,
  Bar,
  Line,
  LineChart,
  XAxis,
  YAxis,
  Tooltip,
} from "recharts";
export const Chart = ({ message, tag }: AppProps) => (
  <div>
    {message} {tag}
  </div>
);
