import { getAuth } from "@clerk/nextjs/server";
import { NextRequest, NextResponse } from "next/server";
import { sql } from "../../../../neon/neonclient";
export async function POST(req: NextRequest) {
  const { userId } = getAuth(req);
  if (!userId) {
    return NextResponse.json({ error: "Not authenicated" }, { status: 401 });
  }

  const data = await req.json();

  try {
    await sql(
      `INSERT INTO api_keys (api_key, user_id) VALUES ('${data.api_key}','${data.user_id}');`,
    );
    return NextResponse.json({ message: "Added api key" }, { status: 200 });
  } catch (error) {
    console.error("Error adding api key to db: ", error);
    return NextResponse.json({ message: error }, { status: 500 });
  }
}

// export const config = {
//   runtime: 'edge',
// };
