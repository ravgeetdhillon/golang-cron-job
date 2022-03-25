import { nanoid } from "nanoid";

export default async function (params) {
  const randomId = nanoid();
  console.log("Random ID", randomId);
  return { randomId };
}
