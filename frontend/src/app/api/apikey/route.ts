import { getAuth } from "@clerk/nextjs/server";
import { NextRequest, NextResponse } from "next/server";
import { sql } from "../../../../neon/neonclient";

type countItem = {
  count: string;
};

export async function POST(req: NextRequest) {
  const { userId } = getAuth(req);
  if (!userId) {
    return NextResponse.json({ error: "Not authenicated" }, { status: 401 });
  }

  // throw Error();
  // TODO: check data types
  // let count: number = 0
  let countPayload: countItem[];

  const data = await req.json();
  try {
    countPayload = (await sql(
      "SELECT COUNT(api_key) FROM api_keys WHERE user_id = $1;",
      [userId],
    )) as countItem[];
  } catch (error) {
    console.error("Error validating api key count: ", error);
    return NextResponse.json({ message: error }, { status: 500 });
  }

  if (Number(countPayload[0]?.count) >= 3) {
    console.log("IN HERE");
    return NextResponse.json(
      { message: "Unable to generate an API key. Limited to 3." },
      { status: 429 },
    );
  }

  try {
    await sql("INSERT INTO api_keys (api_key, user_id) VALUES ($1, $2)", [
      data.api_key,
      data.user_id,
    ]);
    return NextResponse.json({ message: "Added api key" }, { status: 200 });
  } catch (error) {
    console.error("Error adding api key to db: ", error);
    return NextResponse.json({ message: error }, { status: 500 });
  }
}

// Delete by user id & api key for api_keys table
export async function DELETE(req: NextRequest) {
  const { userId } = getAuth(req);
  if (!userId) {
    return NextResponse.json({ error: "Not authenicated" }, { status: 401 });
  }

  const data = await req.json();

  try {
    await sql("DELETE FROM api_keys WHERE api_key = $1 AND user_id = $2;", [
      data.api_key,
      data.user_id,
    ]);
    return NextResponse.json({ message: "Deleted api key" }, { status: 200 });
  } catch (error) {
    console.error("Error deleting api key: ", error);
    return NextResponse.json({ message: error }, { status: 500 });
  }
}

// export const config = {
//   runtime: 'edge',
// };
