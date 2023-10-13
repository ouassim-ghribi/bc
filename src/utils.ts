import { createHash } from "crypto";

export enum Direction {
  LEFT = "left",
  RIGHT = "right",
}

export const sha256 = (data: string) => {
  return createHash("sha256").update(data).digest().toString("hex");
};
