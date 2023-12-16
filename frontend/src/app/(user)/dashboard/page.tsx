import ApiKeyTable from "@/components/ui/apikeytable";
import { currentUser } from "@clerk/nextjs";
import { sql } from "../../../../neon/neonclient";
import { UUID } from "crypto";
import DashboardTiles from "@/components/ui/dashboardtiles";

export type ApiKeyPayload = {
  id: number;
  created_at: string;
  api_key: UUID;
};
export default async function Page() {
  const clerkUser = await currentUser();

  if (!clerkUser) {
    throw new Error("Clerk failed to authenticate.");
  }

  const res: ApiKeyPayload[] = (await sql(
    `SELECT id, created_at, api_key FROM api_keys WHERE user_id= '${clerkUser.id}';`
  )) as ApiKeyPayload[];

  return (
    <div>
      <DashboardTiles username={clerkUser.username!} />
      <ApiKeyTable userId={clerkUser.id} apiKeyData={res} />
    </div>
  );
}
